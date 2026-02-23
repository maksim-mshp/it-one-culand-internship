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
			parts := strings.Fields(r.Header.Get("Authorization"))

			if !strings.EqualFold(parts[0], "Bearer") {
				w.Header().Set("WWW-Authenticate", "Bearer realm=\"API\"")
				corehttp.RespondError(w, corehttp.APIError{
					StatusCode: http.StatusUnauthorized,
					Error:      "INVALID_TOKEN_FORMAT",
				})
				return
			}

			isInternal := security.IsValidInternal(parts[1], secretKey)

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
