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

type MissingFieldError struct {
	ErrorCode string
	Field     string
}

func NewMissingFieldError(field string) MissingFieldError {
	return MissingFieldError{"INTERNSHIP_MISSING_FIELD", field}
}
func (e MissingFieldError) Error() string {
	return fmt.Sprintf("field %s is missing", e.Field)
}

type InvalidFieldLengthError struct {
	ErrorCode     string
	Field         string
	MinLength     int
	MaxLength     int
	CurrentLength int
}

func NewInvalidFieldLengthError(field string, min, max, cur int) InvalidFieldLengthError {
	return InvalidFieldLengthError{
		"INVALID_FIELD_LENGTH",
		field,
		min,
		max,
		cur,
	}
}
func (e InvalidFieldLengthError) Error() string {
	return fmt.Sprintf("field %s has invalid length (cur: %d, min: %d, max: %d)", e.Field, e.CurrentLength, e.MinLength, e.MaxLength)
}
