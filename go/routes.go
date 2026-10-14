package main

import (
	_ "embed"
	"encoding/json"
	"html/template"
	"net/http"
	"runtime"
	"strings"
)

//go:embed welcome.html
var welcomeHTML string

var welcome = template.Must(template.New("welcome").Parse(welcomeHTML))

// routes registers every route and returns the app's handler.
func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /api", api)

	return mux
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	welcome.Execute(w, map[string]string{"Version": strings.TrimPrefix(runtime.Version(), "go")})
}

func api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello from Laravel Cloud"})
}
