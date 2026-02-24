package middleware

import (
	corehttp "culand-internship/internal/core/http"
	"mime"
	"net/http"
	"strings"
)

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
			corehttp.RespondError(w.ResponseWriter, corehttp.APIError{
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
		corehttp.RespondError(w.ResponseWriter, corehttp.APIError{
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
