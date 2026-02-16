package http

import (
	"culand-internship/internal/core/http/middleware"
	"encoding/json"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func ParseJSONBody(r *http.Request, data any) *APIError {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("failed to close body reader: %v", err)
		}
	}()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return &ErrInvalidBody
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		log.Printf("failed to parse json: %v", err)
		return &ErrInvalidBody
	}

	return nil
}

func respondJSON(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if body != nil {
		if err := json.NewEncoder(w).Encode(body); err != nil {
			log.Printf("failed to encode json: %v", err)
		}
	}
}

func RespondSuccess(w http.ResponseWriter, statusCode int, data any) {
	respondJSON(w, statusCode, data)
}

// @title Internships API
// @BasePath /api/v1
// @OpenAPIVersion 3.0.1
func NewServer(port int, mux *http.ServeMux) (*http.Server, error) {
	if err := registerSwagger(mux); err != nil {
		return nil, err
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RespondError(w, APIError{
			StatusCode: http.StatusNotFound,
			Error:      "NOT_FOUND",
			Details: map[string]any{
				"path": r.URL.Path,
			},
		})
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
