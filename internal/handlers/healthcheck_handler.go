package handlers

import "net/http"

// HealthCheckHandler handles health check requests

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("The service is up and running."))
}
