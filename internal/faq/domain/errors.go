package domain

import "fmt"

type NotFoundError struct {
	ID int
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("faq %d not found", e.ID)
}
