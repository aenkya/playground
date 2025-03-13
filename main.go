//go:build !codeanalysis
// +build !codeanalysis

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HealthCheck handles health check requests.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func main() {
	r := mux.NewRouter()

	// Health check route
	r.HandleFunc("/health", HealthCheck).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
