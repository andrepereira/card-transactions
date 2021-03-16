package models

import "time"

// Account Account data
type Account struct {
	AccountID            int     `json:"account_id,omitempty"`
	DocumentNumber       string  `json:"document_number,omitempty"`
	AvaliableCreditLimit float32 `json:"avaliable_credit_limit,omitempty"`
}

// OperationType Any operation has an ID and a Description
type OperationType struct {
	OperationTypeID int16  `json:"operation_type_id,omitempty"`
	Description     string `json:"description,omitempty"`
}

// Transaction Any transaction has a ID, is linked to an Account ID,
// an Operation type ID, have an amount and a Time and Date of event
type Transaction struct {
	TransactionID   int       `json:"transaction_id,omitempty"`
	AccountID       int       `json:"account_id,omitempty"`
	OperationTypeID int16     `json:"operation_type_id,omitempty"`
	Amount          float64   `json:"amount,omitempty"`
	EventDate       time.Time `json:"event_date,omitempty"`
}

// Accounts is a slice of Account to simulate a table of accounts
var Accounts []Account

// OperationTypes is a slice of OperationType to simulate a table of operation types
var OperationTypes []OperationType

// Transactions is a slice of Transaction to simulate a table of transactions
var Transactions []Transaction
