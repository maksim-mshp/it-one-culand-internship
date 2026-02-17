package http

import (
	"bytes"
	"culand-internship/internal/core/security"
	"log"
	"mime"
	"net/http"
	"strings"
)

type responseWriter struct {
	dst    http.ResponseWriter
	header http.Header
	body   bytes.Buffer
	status int
}

func (w *responseWriter) Header() http.Header {
	return w.header
}

func (w *responseWriter) WriteHeader(code int) {
	w.status = code
	if w.dst != nil {
		w.dst.WriteHeader(code)
	}
}

func (w *responseWriter) Write(p []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}

	if w.dst != nil {
		return w.dst.Write(p)
	}
	return w.body.Write(p)
}

func (w *responseWriter) StatusCode() int {
	if w.status == 0 {
		return http.StatusOK
	}
	return w.status
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{
			dst:    w,
			header: w.Header(),
		}

		next.ServeHTTP(rw, r)

		log.Printf("%s %s | %d | %s | %s",
			r.Method, r.URL.Path, rw.StatusCode(), r.RemoteAddr, r.UserAgent(),
		)
	})
}

func HTTPErrorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{
			header: make(http.Header),
		}

		next.ServeHTTP(rw, r)

		contentType := rw.Header().Get("Content-Type")
		status := rw.StatusCode()

		if status == http.StatusNotFound && !isJSON(contentType) {
			RespondError(w, APIError{
				StatusCode: http.StatusNotFound,
				Error:      "NOT_FOUND",
				Details: map[string]any{
					"path": r.URL.Path,
				},
			})
			return
		}

		if status == http.StatusMethodNotAllowed && !isJSON(contentType) {
			allow := strings.Split(rw.Header().Get("Allow"), ", ")

			RespondError(w, APIError{
				StatusCode: http.StatusMethodNotAllowed,
				Error:      "METHOD_NOT_ALLOWED",
				Details: map[string]any{
					"method": r.Method,
					"path":   r.URL.Path,
					"allow":  allow,
				},
			})
			return
		}

		for k, vals := range rw.Header() {
			for _, v := range vals {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(status)
		_, _ = w.Write(rw.body.Bytes())
	})
}

func isJSON(contentType string) bool {
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return false
	}
	return mediaType == "application/json"
}

func RequireAdminMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			const prefix = "Bearer "

			apiErr := ErrUnauthorized
			apiErr.Details = make(map[string]any)

			if !strings.HasPrefix(auth, prefix) {
				apiErr.Details["reason"] = "INVALID_TOKEN_FORMAT"
				RespondError(w, apiErr)
				return
			}

			token := strings.TrimPrefix(auth, prefix)
			isAdmin, err := security.IsAdminRole(token)

			if err != nil {
				apiErr.Details["reason"] = "INVALID_TOKEN_FORMAT"
				RespondError(w, apiErr)
				return
			}

			if !isAdmin {
				apiErr.Details["reason"] = "NOT_ADMIN"
				RespondError(w, apiErr)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
