// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/eddie023/wex-tag/ent/schema"
	"github.com/eddie023/wex-tag/ent/transaction"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	transactionFields := schema.Transaction{}.Fields()
	_ = transactionFields
	// transactionDescDate is the schema descriptor for date field.
	transactionDescDate := transactionFields[1].Descriptor()
	// transaction.DefaultDate holds the default value on creation for the date field.
	transaction.DefaultDate = transactionDescDate.Default.(func() time.Time)
	// transactionDescDescription is the schema descriptor for description field.
	transactionDescDescription := transactionFields[3].Descriptor()
	// transaction.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	transaction.DescriptionValidator = transactionDescDescription.Validators[0].(func(string) error)
	// transactionDescID is the schema descriptor for id field.
	transactionDescID := transactionFields[0].Descriptor()
	// transaction.DefaultID holds the default value on creation for the id field.
	transaction.DefaultID = transactionDescID.Default.(func() uuid.UUID)
}
