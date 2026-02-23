package middleware

import (
	corehttp "culand-internship/internal/core/http"
	"culand-internship/internal/core/security"
	"net/http"
	"strings"
)

func RequireInternalMiddleware(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			const prefix = "Bearer "

			if !strings.HasPrefix(auth, prefix) {
				w.Header().Set("WWW-Authenticate", "Bearer realm=\"API\"")
				corehttp.RespondError(w, corehttp.APIError{
					StatusCode: http.StatusUnauthorized,
					Error:      "INVALID_TOKEN_FORMAT",
				})
				return
			}

			token := strings.TrimPrefix(auth, prefix)
			isInternal := security.IsValidInternal(token, secretKey)

			if !isInternal {
				w.Header().Set("WWW-Authenticate", "Bearer realm=\"API\"")
				corehttp.RespondError(w, corehttp.APIError{
					StatusCode: http.StatusUnauthorized,
					Error:      "INVALID_TOKEN",
				})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
