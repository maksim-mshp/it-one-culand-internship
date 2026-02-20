package http

import (
	"culand-internship/internal/core/security"
	"log"
	"mime"
	"net/http"
	"strings"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (w *responseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *responseWriter) StatusCode() int {
	if w.status == 0 {
		return http.StatusOK
	}
	return w.status
}

func (w *responseWriter) Flush() {
	if f, ok := w.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{
			ResponseWriter: w,
		}

		next.ServeHTTP(rw, r)

		log.Printf("%s %s | %d | %s | %s",
			r.Method, r.URL.Path, rw.StatusCode(), r.RemoteAddr, r.UserAgent(),
		)
	})
}

type errorInterceptor struct {
	http.ResponseWriter
	req         *http.Request
	wroteHeader bool
	hijacked    bool
}

func (w *errorInterceptor) WriteHeader(code int) {
	if w.wroteHeader {
		return
	}
	w.wroteHeader = true

	ct := w.Header().Get("Content-Type")
	if (code == http.StatusNotFound || code == http.StatusMethodNotAllowed) && !isJSON(ct) {
		w.hijacked = true

		if code == http.StatusNotFound {
			RespondError(w.ResponseWriter, APIError{
				StatusCode: http.StatusNotFound,
				Error:      "NOT_FOUND",
				Details: map[string]any{
					"path": w.req.URL.Path,
				},
			})
			return
		}

		allow := w.Header().Get("Allow")
		w.ResponseWriter.Header().Set("Allow", allow)
		RespondError(w.ResponseWriter, APIError{
			StatusCode: http.StatusMethodNotAllowed,
			Error:      "METHOD_NOT_ALLOWED",
			Details: map[string]any{
				"method": w.req.Method,
				"path":   w.req.URL.Path,
				"allow":  strings.Split(allow, ", "),
			},
		})
		return
	}

	w.ResponseWriter.WriteHeader(code)
}

func (w *errorInterceptor) Write(p []byte) (int, error) {
	if w.hijacked {
		return len(p), nil
	}
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(p)
}

func (w *errorInterceptor) Flush() {
	if f, ok := w.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

func HTTPErrorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		interceptor := &errorInterceptor{
			ResponseWriter: w,
			req:            r,
		}
		next.ServeHTTP(interceptor, r)
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
