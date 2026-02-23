package http

import (
	"encoding/json"
	"errors"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"io"
	"log"
	"mime"
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

	if ct := r.Header.Get("Content-Type"); ct != "" {
		mediaType, _, err := mime.ParseMediaType(ct)
		if err != nil || mediaType != "application/json" {
			return &APIError{
				StatusCode: http.StatusUnsupportedMediaType,
				Error:      "UNSUPPORTED_MEDIA_TYPE",
			}
		}
	}

	const maxBodyBytes = 1024 * 1024
	r.Body = http.MaxBytesReader(nil, r.Body, maxBodyBytes)

	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(data); err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			return &APIError{
				StatusCode: http.StatusRequestEntityTooLarge,
				Error:      "REQUEST_ENTITY_TOO_LARGE",
			}
		}
		return &ErrInvalidBody
	}

	if err := dec.Decode(&struct{}{}); err != io.EOF {
		log.Printf("failed to decode json: %v", err)
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			return &APIError{
				StatusCode: http.StatusRequestEntityTooLarge,
				Error:      "REQUEST_ENTITY_TOO_LARGE",
			}
		}
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

func Respond(w http.ResponseWriter, statusCode int, data any) {
	respondJSON(w, statusCode, data)
}

// @Title						Internships API
// @Servers.Url					/api/v1
// @SecurityDefinitions.APIKey	Bearer
// @In							header
// @Name						Authorization
// @Description					Формат: `Bearer {token}`
func NewServer(port int, handler http.Handler) (*http.Server, error) {
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}, nil
}

func RegisterSwagger(mux *http.ServeMux) error {
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
