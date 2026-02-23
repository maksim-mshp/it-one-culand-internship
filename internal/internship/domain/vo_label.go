package domain

import (
	"culand-internship/internal/core/validation"
)

type Label struct {
	value string
}

func NewLabel(v string) (Label, error) {
	val, err := validation.ValidateStringLength("label", v, 1, 50)
	if err != nil && err.CurrentLength != 0 {
		return Label{}, *err
	}
	return Label{value: val}, nil
}

func (l Label) Value() string {
	return l.value
}

func ReconstituteLabel(v string) Label {
	return Label{value: v}
}
