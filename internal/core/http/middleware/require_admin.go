package middleware

import (
	corehttp "culand-internship/internal/core/http"
	"culand-internship/internal/core/security"
	"net/http"
	"strings"
)

func RequireAdminMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			const prefix = "Bearer "

			if !strings.HasPrefix(auth, prefix) {
				corehttp.RespondError(w, corehttp.APIError{
					StatusCode: http.StatusUnauthorized,
					Error:      "INVALID_TOKEN_FORMAT",
				})
				return
			}

			token := strings.TrimPrefix(auth, prefix)
			isAdmin, err := security.IsAdminRole(token)

			if err != nil {
				corehttp.RespondError(w, corehttp.APIError{
					StatusCode: http.StatusUnauthorized,
					Error:      "INVALID_TOKEN_FORMAT",
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
