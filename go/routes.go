package main

import (
	"encoding/json"
	"net/http"
)

// routes registers every route and returns the app's handler.
func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)

	return mux
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello from Laravel Cloud"})
}
