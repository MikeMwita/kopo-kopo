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
	http.HandleFunc("/ping", handler)
	http.HandleFunc("/transactions", handlers.CreateTransactionHandler)
	http.HandleFunc("/transactions/{transaction_id}", handlers.GetTransactionByIDHandler)
	http.HandleFunc("/accounts/{account_id}", handlers.GetAccountHandler)
	http.HandleFunc("/transactions", handlers.GetTransactionsHandler)

	port := 8080
	log.Printf("Server running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
