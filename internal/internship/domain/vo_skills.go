package domain

import (
	"culand-internship/internal/shared/validation"
)

type Skills struct {
	value []string
}

func NewSkills(v []string) (Skills, error) {
	val, err := validation.ValidateSliceItemsLength("skills", v, 5, 200)
	if err != nil {
		return Skills{}, *err
	}
	return Skills{value: val}, nil
}

func (s Skills) Value() []string {
	return s.value
}

func ReconstituteSkills(v []string) Skills {
	return Skills{value: v}
}
