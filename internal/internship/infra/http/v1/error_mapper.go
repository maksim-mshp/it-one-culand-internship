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
			StatusCode: http.StatusNotFound,
			Error:      fieldMissing.ErrorCode,
			Details: map[string]any{
				"field": fieldMissing.Field,
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

	var titleTooShort domain.TitleTooShortError
	if errors.As(err, &titleTooShort) {
		return corehttp.APIError{
			StatusCode: http.StatusUnprocessableEntity,
			Error:      titleTooShort.ErrorCode,
			Details: map[string]any{
				"currentLength": titleTooShort.CurrentLength,
				"minLength":     titleTooShort.MinLength,
			},
		}
	}

	var titleTooLong domain.TitleTooLongError
	if errors.As(err, &titleTooLong) {
		return corehttp.APIError{
			StatusCode: http.StatusUnprocessableEntity,
			Error:      titleTooLong.ErrorCode,
			Details: map[string]any{
				"currentLength": titleTooLong.CurrentLength,
				"maxLength":     titleTooLong.MaxLength,
			},
		}
	}

	return corehttp.ErrInternal
}
