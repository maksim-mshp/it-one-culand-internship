package domain

import (
	"fmt"
)

type NotFoundError struct {
	ErrorCode string
	ID        int
}

func NewNotFoundError(id int) NotFoundError {
	return NotFoundError{"INTERNSHIP_NOT_FOUND", id}
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("internship %d not found", e.ID)
}

type InvalidStatusError struct {
	ErrorCode      string
	Value          string
	PossibleValues []string
}

func NewInvalidStatusError(value string, possible []string) InvalidStatusError {
	return InvalidStatusError{
		"INTERNSHIP_INVALID_STATUS",
		value,
		possible,
	}
}

func (e InvalidStatusError) Error() string {
	return fmt.Sprintf("status %s is invalid (possible: %v)", e.Value, e.PossibleValues)
}
