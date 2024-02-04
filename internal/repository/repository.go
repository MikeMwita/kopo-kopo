package repository

import (
	"database/sql"
	"fmt"
	_ "fmt"
	"github.com/ffc1e12/kopo-kopo-transactionmanagementbackend/internal/models"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./transactions.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create tables if not exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS transactions (
			transaction_id TEXT PRIMARY KEY,
			account_id TEXT,
			amount INTEGER,
			created_at DATETIME
		);

		CREATE TABLE IF NOT EXISTS accounts (
			account_id TEXT PRIMARY KEY,
			balance INTEGER
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateTransaction creates a new transaction in the database
func CreateTransaction(transactionID, accountID string, amount int) error {
	_, err := db.Exec(`
		INSERT INTO transactions (transaction_id, account_id, amount, created_at)
		VALUES (?, ?, ?, ?)
	`, transactionID, accountID, amount, time.Now().UTC())

	if err != nil {
		return fmt.Errorf("error creating transaction: %v", err)
	}

	// Update account balance

	_, err = db.Exec(`
		INSERT INTO accounts (account_id, balance)
		VALUES (?, ?)
		ON CONFLICT(account_id) DO UPDATE SET balance = balance + ?
	`, accountID, amount, amount)

	if err != nil {
		return fmt.Errorf("error updating account balance: %v", err)
	}

	return nil
}

// GetTransactions retrieves all transactions from the database

func GetTransactions() ([]models.Transaction, error) {
	rows, err := db.Query(`
		SELECT transaction_id, account_id, amount, created_at
		FROM transactions
	`)

	if err != nil {
		return nil, fmt.Errorf("error retrieving transactions: %v", err)
	}
	defer rows.Close()

	var transactions []models.Transaction

	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(&t.TransactionID, &t.AccountID, &t.Amount, &t.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning transaction rows: %v", err)
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}

// GetTransactionByID retrieves a transaction by ID from the databas
func GetTransactionByID(transactionID string) (*models.Transaction, error) {
	var t models.Transaction
	err := db.QueryRow(`
		SELECT transaction_id, account_id, amount, created_at
		FROM transactions
		WHERE transaction_id = ?
	`, transactionID).Scan(&t.TransactionID, &t.AccountID, &t.Amount, &t.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil // Transaction not found
	} else if err != nil {
		return nil, fmt.Errorf("error retrieving transaction by ID: %v", err)
	}

	return &t, nil
}

// GetAccountByID retrieves an account by ID from the database

func GetAccountByID(accountID string) (*models.Account, error) {
	var a models.Account
	err := db.QueryRow(`
		SELECT account_id, balance
		FROM accounts
		WHERE account_id = ?
	`, accountID).Scan(&a.AccountID, &a.Balance)

	if err == sql.ErrNoRows {
		return nil, nil // Account not found
	} else if err != nil {
		return nil, fmt.Errorf("error retrieving account by ID: %v", err)
	}

	return &a, nil
}
