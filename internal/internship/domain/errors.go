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

type TitleTooShortError struct {
	ErrorCode     string
	CurrentLength int
	MinLength     int
}

func NewTitleTooShortError(currentLength, minLength int) TitleTooShortError {
	return TitleTooShortError{
		"INTERNSHIP_TITLE_TOO_SHORT",
		currentLength,
		minLength,
	}
}
func (e TitleTooShortError) Error() string {
	return fmt.Sprintf("internship title too short (current: %d, min: %d)", e.CurrentLength, e.MinLength)
}

type TitleTooLongError struct {
	ErrorCode     string
	CurrentLength int
	MaxLength     int
}

func NewTitleTooLongError(currentLength, maxLength int) TitleTooLongError {
	return TitleTooLongError{
		"INTERNSHIP_TITLE_TOO_SHORT",
		currentLength,
		maxLength,
	}
}
func (e TitleTooLongError) Error() string {
	return fmt.Sprintf("internship title too long (current: %d, max: %d)", e.CurrentLength, e.MaxLength)
}

type MissingFieldError struct {
	ErrorCode string
	Field     string
}

func NewMissingFieldError(field string) MissingFieldError {
	return MissingFieldError{"INTERNSHIP_MISSING_FIELD", field}
}
func (e MissingFieldError) Error() string {
	return fmt.Sprintf("filed %s is missing", e.Field)
}
