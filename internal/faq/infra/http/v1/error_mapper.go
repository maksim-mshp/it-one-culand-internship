package v1

import (
	corehttp "culand-internship/internal/core/http"
	"culand-internship/internal/faq/domain"
	"errors"
	"net/http"
)

func mapError(err error) corehttp.APIError {
	if apiErr := corehttp.MapErrors(err); apiErr != nil {
		return *apiErr
	}

	var notFound domain.NotFoundError
	if errors.As(err, &notFound) {
		return corehttp.APIError{
			StatusCode: http.StatusNotFound,
			Error:      "FAQ_NOT_FOUND",
			Details: map[string]any{
				"id": notFound.ID,
			},
		}
	}

	return corehttp.ErrInternal
}
