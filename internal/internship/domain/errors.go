package domain

import "fmt"

type NotFoundError struct {
	ID int
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("internship %d not found", e.ID)
}

type InvalidStatusError struct {
	Value          string
	PossibleValues []string
}

func (e InvalidStatusError) Error() string {
	return fmt.Sprintf("status %s is invalid (possible: %v)", e.Value, e.PossibleValues)
}
