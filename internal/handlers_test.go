package handlers

import (
	"bytes"
	"github.com/ffc1e12/kopo-kopo-transactionmanagementbackend/internal/handlers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTransactionHandler(t *testing.T) {
	testCases := []struct {
		Name         string
		RequestBody  string
		ExpectedCode int
	}{
		{
			Name:         "ValidRequest",
			RequestBody:  `{"account_id": "some_account_id", "amount": 10}`,
			ExpectedCode: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/transactions", bytes.NewBuffer([]byte(tc.RequestBody)))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(handlers.CreateTransactionHandler)

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tc.ExpectedCode, rr.Code)

		})
	}
}

func TestGetTransactionsHandler(t *testing.T) {
	testCases := []struct {
		Name string
	}{
		{
			Name: "ValidRequest",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/transactions", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(handlers.GetTransactionsHandler)

			handler.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusOK, rr.Code)

		})
	}
}
