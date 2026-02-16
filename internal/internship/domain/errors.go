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
