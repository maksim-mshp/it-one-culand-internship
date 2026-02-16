package http

import "net/http"

type APIError struct {
	StatusCode int            `json:"-"`
	Error      string         `json:"error"`
	Details    map[string]any `json:"details,omitempty"`
} // @name APIError

func RespondError(w http.ResponseWriter, err APIError) {
	respondJSON(w, err.StatusCode, err)
}

var ErrInternal = APIError{
	StatusCode: http.StatusInternalServerError,
	Error:      "INTERNAL_ERROR",
}
