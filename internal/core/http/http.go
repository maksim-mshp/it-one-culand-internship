package http

import (
	"encoding/json"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
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
// @version 1.0
// @description API для управления стажировками
// @host localhost:8080
// @BasePath /
func NewServer(port int, mux *http.ServeMux) *http.Server {
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("api"))))
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/swagger.json"), // путь к спецификации
	))

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		RespondError(w, http.StatusNotFound, "not found")
	})

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
}
