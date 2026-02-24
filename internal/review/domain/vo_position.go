package domain

import (
	"culand-internship/internal/shared/validation"
)

type Position struct {
	value string
}

func NewPosition(v string) (Position, error) {
	val, err := validation.ValidateStringLength("position", v, 5, 100)
	if err != nil {
		if err.CurrentLength == 0 {
			return Position{}, validation.RequiredFieldMissingError{Field: "position"}
		}
		return Position{}, *err
	}
	return Position{value: val}, nil
}

func (p Position) Value() string {
	return p.value
}

func ReconstitutePosition(v string) Position {
	return Position{value: v}
}
