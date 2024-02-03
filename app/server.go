//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ffc1e12/kopo-kopo-transactionmanagementbackend/internal/handlers"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handler)
	r.HandleFunc("/transactions", handlers.CreateTransactionHandler).Methods("POST")
	r.HandleFunc("/transactions/{transaction_id}", handlers.GetTransactionByIDHandler).Methods("GET")
	r.HandleFunc("/accounts/{account_id}", handlers.GetAccountHandler).Methods("GET")
	r.HandleFunc("/transactions", handlers.GetTransactionsHandler).Methods("GET")

	port := 8080
	log.Printf("Server running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
