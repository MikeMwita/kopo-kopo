package router

import (
	"github.com/ffc1e12/kopo-kopo-transactionmanagementbackend/internal/handlers"
	"github.com/gorilla/mux"
)

// NewRouter creates a new router with defined routes
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handlers.HealthCheckHandler).Methods("GET")
	r.HandleFunc("/transactions", handlers.CreateTransactionHandler).Methods("POST")
	r.HandleFunc("/transactions", handlers.GetTransactionsHandler).Methods("GET")
	r.HandleFunc("/transactions/{transaction_id}", handlers.GetTransactionByIDHandler).Methods("GET")
	r.HandleFunc("/accounts/{account_id}", handlers.GetAccountHandler).Methods("GET")

	return r
}
