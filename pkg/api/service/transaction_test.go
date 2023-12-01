package service

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/eddie023/wex-tag/pkg/test"
	"github.com/eddie023/wex-tag/pkg/types"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gotest.tools/assert"
)

func TestCreatePurchase(t *testing.T) {
	type testcase struct {
		name    string
		payload types.CreateNewPurchaseTransaction
		wantErr error
		want    string
	}

	testcases := []testcase{
		{
			name: "Should correctly create new transaction for valid positive amounts",
			payload: types.CreateNewPurchaseTransaction{
				Amount:      "123.16",
				Description: "Positive amount",
			},
			wantErr: nil,
			want:    "123.16",
		},
		{
			name: "Should correctly create new transaction for valid positive amounts",
			payload: types.CreateNewPurchaseTransaction{
				Amount:      "0",
				Description: "Positive amount",
			},
			wantErr: nil,
			want:    "0.00",
		},
		{
			name: "Should faild for negative amount value",
			payload: types.CreateNewPurchaseTransaction{
				Amount:      "-123.123",
				Description: "This is negative amount",
			},
			wantErr: errors.New("amount cannot be negative number"),
			want:    "",
		},
		{
			name: "Should fail for invalid amount value",
			payload: types.CreateNewPurchaseTransaction{
				Amount:      "-123.123abcd",
				Description: "Invalid amount",
			},
			wantErr: errors.New("can't convert"),
			want:    "",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			ent := test.NewDatabase(t)
			defer ent.Close()

			s := Service{
				Ent: ent,
			}

			newTransaction, err := s.CreateNewPurchaseTransaction(context.TODO(), tc.payload)
			if err != nil {
				if strings.Contains(err.Error(), tc.wantErr.Error()) {
					return
				}

				t.Errorf("want = %s got = %s", tc.wantErr.Error(), err.Error())

				return
			}

			assert.Equal(t, newTransaction.AmountInUSD, tc.want)
		})
	}

}

// Amount should be valid positive number rounded to the nearest cent
func TestRoundToNearestCent(t *testing.T) {

	tests := []struct {
		name  string
		given decimal.Decimal
		want  string
	}{
		{
			name:  "Should not round",
			given: decimal.NewFromFloat(12.6544),
			want:  "12.65",
		},
		{
			name:  "Should not round",
			given: decimal.NewFromFloat(0),
			want:  "0",
		},
		{
			name:  "Should not round",
			given: decimal.NewFromFloat(1.5),
			want:  "1.5",
		},
		{
			name:  "should not round",
			given: decimal.NewFromFloat(12.65466),
			want:  "12.65",
		},
		{
			name:  "should round",
			given: decimal.NewFromFloat(12.65766),
			want:  "12.66",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RoundToNearestCent(tt.given)

			assert.Equal(t, tt.want, got.String())
		})
	}
}

func TestParseStringToUUID(t *testing.T) {

	tests := []struct {
		name    string
		given   string
		want    uuid.UUID
		wantErr bool
	}{
		{
			name:    "valid UUID string",
			given:   "8d8b30af-5b77-4fa0-9270-a85bec6600dd",
			wantErr: false,
		},
		{
			name:    "valid version 1 UUID string",
			given:   "114dec4e-8f91-11ee-b9d1-0242ac120002",
			wantErr: false,
		},
		{
			name:    "valid version 4 UUID string",
			given:   "69802563-9655-42c3-94ff-41537d1f8332",
			wantErr: false,
		},
		{
			name:    "invalid UUID string with invalid length",
			given:   "69802563-9655-42c3-94ff-invalid",
			wantErr: true,
		},
		{
			name:    "invalid UUID string with valid length",
			given:   "69802563-9655-42c3-94ff-41537d1f()^^",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseStringToUUID(tt.given)
			if err != nil {
				if tt.wantErr {
					return
				}

				t.Fatal()
			}
			assert.Equal(t, tt.given, got.String())
		})
	}
}