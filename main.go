//go:build !codeanalysis
// +build !codeanalysis

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// HealthCheck handles health check requests.
func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func runAPIServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", HealthCheck)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	fmt.Println("Starting server on :8080")

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server failed:", err)
	}
}

func main() {
	runAPIServer()
}
