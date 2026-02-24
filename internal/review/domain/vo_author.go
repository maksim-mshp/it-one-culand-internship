package domain

import "culand-internship/internal/shared/validation"

type Author struct {
	value string
}

func NewAuthor(v string) (Author, error) {
	val, err := validation.ValidateStringLength("author", v, 2, 100)
	if err != nil {
		if err.CurrentLength == 0 {
			return Author{}, validation.RequiredFieldMissingError{Field: "author"}
		}
		return Author{}, *err
	}
	return Author{value: val}, nil
}

func (a Author) Value() string {
	return a.value
}

func ReconstituteAuthor(v string) Author {
	return Author{value: v}
}
