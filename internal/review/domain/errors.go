package domain

import "fmt"

type NotFoundError struct {
	ID int
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("review %d not found", e.ID)
}
