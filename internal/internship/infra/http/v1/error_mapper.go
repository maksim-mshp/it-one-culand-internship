package v1

import (
	corehttp "culand-internship/internal/core/http"
	"culand-internship/internal/internship/domain"
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
			Error:      notFound.ErrorCode,
			Details: map[string]any{
				"id": notFound.ID,
			},
		}
	}

	var invalidStatus domain.InvalidStatusError
	if errors.As(err, &invalidStatus) {
		return corehttp.APIError{
			StatusCode: http.StatusUnprocessableEntity,
			Error:      invalidStatus.ErrorCode,
			Details: map[string]any{
				"value":          invalidStatus.Value,
				"possibleValues": invalidStatus.PossibleValues,
			},
		}
	}

	return corehttp.ErrInternal
}
