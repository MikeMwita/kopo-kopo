package main

import (
	"github.com/ffc1e12/kopo-kopo-transactionmanagementbackend/internal/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()

	port := 8080
	log.Printf("Server running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(":8080", r))
}
