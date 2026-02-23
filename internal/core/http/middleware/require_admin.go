package middleware

import (
	corehttp "culand-internship/internal/core/http"
	"culand-internship/internal/core/security"
	"net/http"
	"strings"
)

func RequireAdminMiddleware(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			parts := strings.Fields(r.Header.Get("Authorization"))

			if len(parts) < 2 || !strings.EqualFold(parts[0], "Bearer") {
				w.Header().Set("WWW-Authenticate", "Bearer realm=\"API\"")
				corehttp.RespondError(w, corehttp.APIError{
					StatusCode: http.StatusUnauthorized,
					Error:      "INVALID_TOKEN",
				})
				return
			}

			isAdmin, err := security.IsAdminRole(parts[1], secretKey)

			if err != nil {
				w.Header().Set("WWW-Authenticate", "Bearer realm=\"API\"")
				corehttp.RespondError(w, corehttp.APIError{
					StatusCode: http.StatusUnauthorized,
					Error:      "INVALID_TOKEN",
				})
				return
			}

			if !isAdmin {
				corehttp.RespondError(w, corehttp.APIError{
					StatusCode: http.StatusForbidden,
					Error:      "FORBIDDEN",
					Details: map[string]any{
						"reason": "NOT_ADMIN",
					},
				})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
