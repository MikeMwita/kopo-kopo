package handlers

import (
	"encoding/json"
	"github.com/ffc1e12/kopo-kopo-transactionmanagementbackend/internal/repository"
	"github.com/gorilla/mux"
	"net/http"
)

// GetTransactionByIDHandler handles getting a transaction by ID
func GetTransactionByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transactionID := vars["transaction_id"]

	// Call repository.GetTransactionByID to retrieve the transaction by ID
	transactionDetails, err := repository.GetTransactionByID(transactionID)
	if err != nil {
		http.Error(w, "Error retrieving transaction details", http.StatusInternalServerError)
		return
	}

	// Respond with the transaction details
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactionDetails)
}
