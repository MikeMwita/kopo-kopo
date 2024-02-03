package models

import "time"

type Transaction struct {
	TransactionID string    `json:"transaction_id"`
	AccountID     string    `json:"account_id"`
	Amount        int       `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

type ArrayOfTransactions []Transaction

type Account struct {
	AccountID string `json:"account_id"`
	Balance   int    `json:"balance"`
}

type TransactionRequest struct {
	TransactionID string `json:"transaction_id"`
	AccountID     string `json:"account_id"`
	Amount        int    `json:"amount"`
}
