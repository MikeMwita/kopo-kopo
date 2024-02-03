package handlers

import (
	"encoding/json"
	"github.com/ffc1e12/kopo-kopo-transactionmanagementbackend/internal/repository"
	"net/http"
)

// GetTransactionsHandler handles getting all transactions

func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	// Call repository.GetTransactions to retrieve all transactions
	arrayOfTransactions, err := repository.GetTransactions()
	if err != nil {
		http.Error(w, "Error retrieving transactions", http.StatusInternalServerError)
		return
	}

	// Respond with the array of transactions
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(arrayOfTransactions)
}
