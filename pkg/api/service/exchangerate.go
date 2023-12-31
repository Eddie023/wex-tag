package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/eddie023/wex-tag/ent"
	"github.com/eddie023/wex-tag/pkg/apiout"
	"github.com/eddie023/wex-tag/pkg/types"
	"github.com/shopspring/decimal"
)

type ExchangeRateGetter struct{}

type ExchangeRatePayload struct {
	CountryName string
	Currency    string
	RecordDate  time.Time
}

type ExchangeRateResponse struct {
	CountryCurrencyDesc string `json:"country_currency_desc"`
	ExchangeRate        string `json:"exchange_rate"`
	RecordDate          string `json:"record_date"`
}

type ExchangeRateAPIResponse struct {
	Data []ExchangeRateResponse `json:"data"`
}

const TREASURY_RATES_OF_EXCHANGE_API_URL = "https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange"

func (e *ExchangeRateGetter) GetExchangeRate(ctx context.Context, payload ExchangeRatePayload) (ExchangeRateResponse, error) {

	req, err := http.NewRequest("GET", TREASURY_RATES_OF_EXCHANGE_API_URL, nil)
	if err != nil {
		return ExchangeRateResponse{}, err
	}

	req.URL.RawQuery = getURLWithRawQueryParms(payload)

	client := &http.Client{}

	var resp *http.Response

	slog.Debug("generated exchange rate API", "url", req.URL)

	// since, this API have a rate limiting, we will try to expotentially backoff and retry if we get too many request error from the API.
	operation := func() error {
		resp, err = client.Do(req)
		if err != nil {
			return &backoff.PermanentError{}
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			return fmt.Errorf("too many request error")
		}

		return nil
	}

	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.MaxElapsedTime = 1 * time.Minute

	err = backoff.Retry(operation, expBackoff)
	if err != nil {
		return ExchangeRateResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		slog.Debug("exchange request failed", "status_code", resp.StatusCode)
		return ExchangeRateResponse{}, apiout.NewRequestError(fmt.Errorf("the exchange rate service failed with status code %v", resp.StatusCode), http.StatusInternalServerError)
	}

	var response ExchangeRateAPIResponse

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return ExchangeRateResponse{}, err
	}

	// for invalid country or currency, API will still return 200 with empty list
	if len(response.Data) == 0 {
		return ExchangeRateResponse{}, apiout.NewRequestError(errors.New("the purchase cannot be converted to the target currency, exchange rate API returned empty result"), http.StatusBadRequest)
	}

	if resp != nil {
		resp.Body.Close()
	}

	// parse string to Date
	latestedRecordDate, err := time.Parse(time.DateOnly, response.Data[0].RecordDate)
	if err != nil {
		return ExchangeRateResponse{}, apiout.NewRequestError(errors.New("unable to parse returned record date"), http.StatusInternalServerError)
	}

	// currency conversion rate can be less than or equal to purchase date from within the last 6 months
	sixMonthBeforePurchaseDate := getSixMonthBeforePurchaseDate(payload.RecordDate)

	if latestedRecordDate.Before(sixMonthBeforePurchaseDate) {
		slog.Debug("unable to find currency conversion rate within last 6 months", "latest_date", latestedRecordDate)
		return ExchangeRateResponse{}, apiout.NewRequestError(errors.New("the purchase cannot be converted to the target currency, unable to find currency converson rate within last 6 months"), http.StatusBadRequest)
	}

	// we can return the first item since we have already sorted our API response to our need.
	return response.Data[0], nil
}

// ConvertCurrency will return the converted purchase amount for the provided purchase price in USD to exchange rate information.
func (e *ExchangeRateGetter) ConvertCurrency(payload ExchangeRatePayload, trans *ent.Transaction, er ExchangeRateResponse) (types.GetPurchaseTransaction, error) {
	exchangeRate, err := decimal.NewFromString(er.ExchangeRate)
	if err != nil {
		return types.GetPurchaseTransaction{}, err
	}

	country := strings.Trim(payload.CountryName, "\"")
	currency := strings.Trim(payload.Currency, "\"")

	convertedAmount := convertAmount(trans.AmountInUsd, exchangeRate)

	response := types.GetPurchaseTransaction{
		TransactionDetails: types.Transaction{
			AmountInUSD: trans.AmountInUsd.String(),
			Date:        trans.Date,
			Description: trans.Description,
			Id:          trans.ID.String(),
		},
		ConvertedDetails: types.ConvertedPurchasePrice{
			Amount:           RoundToNearestCent(convertedAmount).String(),
			Country:          country,
			Currency:         currency,
			ExchangeRateUsed: er.ExchangeRate,
			ExchangeRateDate: er.RecordDate,
		},
	}

	return response, nil
}

func getSixMonthBeforePurchaseDate(d time.Time) time.Time {
	return d.AddDate(0, -6, 0)
}

func convertAmount(original decimal.Decimal, exchangeRate decimal.Decimal) decimal.Decimal {
	return original.Mul(exchangeRate)
}

// getURLWithRawQueryParams will generate required query param for our exchange rate API call.
func getURLWithRawQueryParms(payload ExchangeRatePayload) string {
	// remove " quotes from our query params
	country := strings.Trim(payload.CountryName, "\"")
	country = strings.Trim(country, "'")
	currency := strings.Trim(payload.Currency, "\"")
	currency = strings.Trim(currency, "'")

	filter := fmt.Sprintf("record_date:lte:%s,country_currency_desc:eq:%s-%s", payload.RecordDate.Format(time.DateOnly), url.QueryEscape(country), url.QueryEscape(currency))
	fields := "country_currency_desc,exchange_rate,record_date"
	// sort by record_date in descending order such that we will get the first item which is closest to our purchase date within last six months
	sort := "-record_date"

	output := fmt.Sprintf("filter=%s&fields=%s&sort=%s&page[size]=1", filter, fields, sort)

	return output
}
