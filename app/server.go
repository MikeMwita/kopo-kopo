//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func main() {
	http.HandleFunc("/ping", handler)
	http.HandleFunc("/ping", handlers.PingHandler)
	http.HandleFunc("/transactions", handlers.CreateTransactionHandler)
	http.HandleFunc("/transactions/{transaction_id}", handlers.GetTransactionByIDHandler)
	http.HandleFunc("/accounts/{account_id}", handlers.GetAccountHandler)
	http.HandleFunc("/transactions", handlers.GetTransactionsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
