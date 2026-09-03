package app

import (
	"encoding/json"
	"net/http"
)

const ServiceName = "cicd-test"

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func Hello(name string) string {
	if name == "" {
		return "Hello, World"
	}
	return "Hello, " + name
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(HealthResponse{
		Status:  "ok",
		Service: ServiceName,
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte(Hello(name)))
}

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", HealthHandler)
	mux.HandleFunc("GET /", HelloHandler)
	return mux
}
