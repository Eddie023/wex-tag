// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/eddie023/wex-tag/ent/predicate"
	"github.com/eddie023/wex-tag/ent/transaction"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTransaction = "Transaction"
)

// TransactionMutation represents an operation that mutates the Transaction nodes in the graph.
type TransactionMutation struct {
	config
	op               Op
	typ              string
	id               *uuid.UUID
	date             *time.Time
	amount_in_usd    *decimal.Decimal
	addamount_in_usd *decimal.Decimal
	description      *string
	clearedFields    map[string]struct{}
	done             bool
	oldValue         func(context.Context) (*Transaction, error)
	predicates       []predicate.Transaction
}

var _ ent.Mutation = (*TransactionMutation)(nil)

// transactionOption allows management of the mutation configuration using functional options.
type transactionOption func(*TransactionMutation)

// newTransactionMutation creates new mutation for the Transaction entity.
func newTransactionMutation(c config, op Op, opts ...transactionOption) *TransactionMutation {
	m := &TransactionMutation{
		config:        c,
		op:            op,
		typ:           TypeTransaction,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTransactionID sets the ID field of the mutation.
func withTransactionID(id uuid.UUID) transactionOption {
	return func(m *TransactionMutation) {
		var (
			err   error
			once  sync.Once
			value *Transaction
		)
		m.oldValue = func(ctx context.Context) (*Transaction, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Transaction.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTransaction sets the old Transaction of the mutation.
func withTransaction(node *Transaction) transactionOption {
	return func(m *TransactionMutation) {
		m.oldValue = func(context.Context) (*Transaction, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TransactionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TransactionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Transaction entities.
func (m *TransactionMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TransactionMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TransactionMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Transaction.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetDate sets the "date" field.
func (m *TransactionMutation) SetDate(t time.Time) {
	m.date = &t
}

// Date returns the value of the "date" field in the mutation.
func (m *TransactionMutation) Date() (r time.Time, exists bool) {
	v := m.date
	if v == nil {
		return
	}
	return *v, true
}

// OldDate returns the old "date" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDate: %w", err)
	}
	return oldValue.Date, nil
}

// ResetDate resets all changes to the "date" field.
func (m *TransactionMutation) ResetDate() {
	m.date = nil
}

// SetAmountInUsd sets the "amount_in_usd" field.
func (m *TransactionMutation) SetAmountInUsd(d decimal.Decimal) {
	m.amount_in_usd = &d
	m.addamount_in_usd = nil
}

// AmountInUsd returns the value of the "amount_in_usd" field in the mutation.
func (m *TransactionMutation) AmountInUsd() (r decimal.Decimal, exists bool) {
	v := m.amount_in_usd
	if v == nil {
		return
	}
	return *v, true
}

// OldAmountInUsd returns the old "amount_in_usd" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldAmountInUsd(ctx context.Context) (v decimal.Decimal, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAmountInUsd is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAmountInUsd requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAmountInUsd: %w", err)
	}
	return oldValue.AmountInUsd, nil
}

// AddAmountInUsd adds d to the "amount_in_usd" field.
func (m *TransactionMutation) AddAmountInUsd(d decimal.Decimal) {
	if m.addamount_in_usd != nil {
		*m.addamount_in_usd = m.addamount_in_usd.Add(d)
	} else {
		m.addamount_in_usd = &d
	}
}

// AddedAmountInUsd returns the value that was added to the "amount_in_usd" field in this mutation.
func (m *TransactionMutation) AddedAmountInUsd() (r decimal.Decimal, exists bool) {
	v := m.addamount_in_usd
	if v == nil {
		return
	}
	return *v, true
}

// ResetAmountInUsd resets all changes to the "amount_in_usd" field.
func (m *TransactionMutation) ResetAmountInUsd() {
	m.amount_in_usd = nil
	m.addamount_in_usd = nil
}

// SetDescription sets the "description" field.
func (m *TransactionMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *TransactionMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Transaction entity.
// If the Transaction object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TransactionMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *TransactionMutation) ResetDescription() {
	m.description = nil
}

// Where appends a list predicates to the TransactionMutation builder.
func (m *TransactionMutation) Where(ps ...predicate.Transaction) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TransactionMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TransactionMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Transaction, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TransactionMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TransactionMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Transaction).
func (m *TransactionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TransactionMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.date != nil {
		fields = append(fields, transaction.FieldDate)
	}
	if m.amount_in_usd != nil {
		fields = append(fields, transaction.FieldAmountInUsd)
	}
	if m.description != nil {
		fields = append(fields, transaction.FieldDescription)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TransactionMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case transaction.FieldDate:
		return m.Date()
	case transaction.FieldAmountInUsd:
		return m.AmountInUsd()
	case transaction.FieldDescription:
		return m.Description()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TransactionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case transaction.FieldDate:
		return m.OldDate(ctx)
	case transaction.FieldAmountInUsd:
		return m.OldAmountInUsd(ctx)
	case transaction.FieldDescription:
		return m.OldDescription(ctx)
	}
	return nil, fmt.Errorf("unknown Transaction field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TransactionMutation) SetField(name string, value ent.Value) error {
	switch name {
	case transaction.FieldDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDate(v)
		return nil
	case transaction.FieldAmountInUsd:
		v, ok := value.(decimal.Decimal)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAmountInUsd(v)
		return nil
	case transaction.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	}
	return fmt.Errorf("unknown Transaction field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TransactionMutation) AddedFields() []string {
	var fields []string
	if m.addamount_in_usd != nil {
		fields = append(fields, transaction.FieldAmountInUsd)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TransactionMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case transaction.FieldAmountInUsd:
		return m.AddedAmountInUsd()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TransactionMutation) AddField(name string, value ent.Value) error {
	switch name {
	case transaction.FieldAmountInUsd:
		v, ok := value.(decimal.Decimal)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAmountInUsd(v)
		return nil
	}
	return fmt.Errorf("unknown Transaction numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TransactionMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TransactionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TransactionMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Transaction nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TransactionMutation) ResetField(name string) error {
	switch name {
	case transaction.FieldDate:
		m.ResetDate()
		return nil
	case transaction.FieldAmountInUsd:
		m.ResetAmountInUsd()
		return nil
	case transaction.FieldDescription:
		m.ResetDescription()
		return nil
	}
	return fmt.Errorf("unknown Transaction field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TransactionMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TransactionMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TransactionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TransactionMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TransactionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TransactionMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TransactionMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Transaction unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TransactionMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Transaction edge %s", name)
}
