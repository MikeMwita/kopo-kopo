package handlers

import (
	"encoding/json"
	"github.com/ffc1e12/kopo-kopo-transactionmanagementbackend/internal/models"
	"github.com/ffc1e12/kopo-kopo-transactionmanagementbackend/internal/repository"
	"net/http"
)

// CreateTransactionHandler handles creating a new transaction
func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var transactionRequest models.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&transactionRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate transactionRequest fields

	err = repository.CreateTransaction(transactionRequest.TransactionID, transactionRequest.AccountID, transactionRequest.Amount)
	if err != nil {
		http.Error(w, "Error creating transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nil)
}
