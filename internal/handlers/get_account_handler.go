package handlers

import (
	"encoding/json"
	"github.com/ffc1e12/kopo-kopo-transactionmanagementbackend/internal/repository"
	"github.com/gorilla/mux"
	"net/http"
)

// GetAccountHandler handles getting an account by ID
func GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["account_id"]

	// Call repository.GetAccountByID to retrieve the account by ID
	accountDetails, err := repository.GetAccountByID(accountID)
	if err != nil {
		http.Error(w, "Error retrieving account details", http.StatusInternalServerError)
		return
	}

	// Respond with the account details
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accountDetails)
}
