package v1

import (
	corehttp "culand-internship/internal/core/http"
	"culand-internship/internal/internship/domain"
	"errors"
	"net/http"
)

func mapError(err error) corehttp.APIError {
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
