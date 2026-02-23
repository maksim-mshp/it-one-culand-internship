package middleware

import (
	corehttp "culand-internship/internal/core/http"
	"log"
	"net/http"
	"runtime/debug"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recovered := recover(); recovered != nil {
				log.Printf("panic recovered: %v\n%s", recovered, debug.Stack())
				corehttp.RespondError(w, corehttp.ErrInternal)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
