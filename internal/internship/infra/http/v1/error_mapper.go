package v1

import (
	corehttp "culand-internship/internal/core/http"
	"culand-internship/internal/internship/domain"
	"errors"
	"net/http"
)

func mapError(err error) corehttp.APIError {
	var fieldMissing domain.MissingFieldError
	if errors.As(err, &fieldMissing) {
		return corehttp.APIError{
			StatusCode: http.StatusUnprocessableEntity,
			Error:      fieldMissing.ErrorCode,
			Details: map[string]any{
				"field": fieldMissing.Field,
			},
		}
	}

	var invalidFieldLength domain.InvalidFieldLengthError
	if errors.As(err, &invalidFieldLength) {
		return corehttp.APIError{
			StatusCode: http.StatusUnprocessableEntity,
			Error:      invalidFieldLength.ErrorCode,
			Details: map[string]any{
				"field":         invalidFieldLength.Field,
				"minLength":     invalidFieldLength.MinLength,
				"maxLength":     invalidFieldLength.MaxLength,
				"currentLength": invalidFieldLength.CurrentLength,
			},
		}
	}

	var notFound domain.NotFoundError
	if errors.As(err, &notFound) {
		return corehttp.APIError{
			StatusCode: http.StatusNotFound,
			Error:      notFound.ErrorCode,
			Details: map[string]any{
				"id": notFound.ID,
			},
		}
	}

	return corehttp.ErrInternal
}
