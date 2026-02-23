package http

import (
	"culand-internship/internal/shared/validation"
	"errors"
	"net/http"
)

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

var ErrInvalidBody = APIError{
	StatusCode: http.StatusBadRequest,
	Error:      "INVALID_BODY",
}

func MapErrors(err error) *APIError {
	var requiredFieldMissing validation.RequiredFieldMissingError
	if errors.As(err, &requiredFieldMissing) {
		return &APIError{
			StatusCode: http.StatusUnprocessableEntity,
			Error:      "REQUIRED_FIELD_MISSING",
			Details: map[string]any{
				"field": requiredFieldMissing.Field,
			},
		}
	}

	var invalidFieldLength validation.InvalidFieldLengthError
	if errors.As(err, &invalidFieldLength) {
		return &APIError{
			StatusCode: http.StatusUnprocessableEntity,
			Error:      "INVALID_FIELD_LENGTH",
			Details: map[string]any{
				"field":         invalidFieldLength.Field,
				"value":         invalidFieldLength.Value,
				"minLength":     invalidFieldLength.MinLength,
				"maxLength":     invalidFieldLength.MaxLength,
				"currentLength": invalidFieldLength.CurrentLength,
			},
		}
	}

	var invalidArrayItemLengthLength validation.InvalidArrayItemLengthError
	if errors.As(err, &invalidArrayItemLengthLength) {
		return &APIError{
			StatusCode: http.StatusUnprocessableEntity,
			Error:      "INVALID_ARRAY_ITEM_LENGTH",
			Details: map[string]any{
				"field":         invalidArrayItemLengthLength.Field,
				"index":         invalidArrayItemLengthLength.Index,
				"value":         invalidArrayItemLengthLength.Value,
				"minLength":     invalidArrayItemLengthLength.MinLength,
				"maxLength":     invalidArrayItemLengthLength.MaxLength,
				"currentLength": invalidArrayItemLengthLength.CurrentLength,
			},
		}
	}

	return nil
}
