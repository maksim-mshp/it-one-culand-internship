package domain

import (
	"culand-internship/internal/core/validation"
)

type Title struct {
	value string
}

func NewTitle(v string) (Title, error) {
	val, err := validation.ValidateStringLength("title", v, 5, 100)
	if err != nil {
		if err.CurrentLength == 0 {
			return Title{}, validation.RequiredFieldMissingError{Field: "title"}
		}
		return Title{}, *err
	}
	return Title{value: val}, nil
}

func (t Title) Value() string {
	return t.value
}

func ReconstituteTitle(v string) Title {
	return Title{value: v}
}
