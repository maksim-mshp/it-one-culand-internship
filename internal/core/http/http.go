package http

import (
	"culand-internship/internal/core/http/middleware"
	"encoding/json"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func respondJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if body != nil {
		if err := json.NewEncoder(w).Encode(body); err != nil {
			log.Printf("failed to encode json: %v", err)
		}
	}
}

func RespondError(w http.ResponseWriter, statusCode int, msg string) {
	respondJSON(w, statusCode, ErrorResponse{Error: msg})
}

func RespondSuccess(w http.ResponseWriter, statusCode int, data interface{}) {
	respondJSON(w, statusCode, data)
}

// @title Internship API
// @BasePath /api/v1
func NewServer(port int, mux *http.ServeMux) (*http.Server, error) {
	if err := registerSwagger(mux); err != nil {
		return nil, err
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		RespondError(w, http.StatusNotFound, "not found")
	})

	handler := middleware.Logging(mux)

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}, nil
}

func registerSwagger(mux *http.ServeMux) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	apiDir := filepath.Join(wd, "api")
	mux.HandleFunc("/swagger/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(apiDir, "openapi.json"))
	})
	mux.HandleFunc("/swagger/openapi.yml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(apiDir, "openapi.yml"))
	})
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/openapi.json"),
	))
	return nil
}
