package domain

import "culand-internship/internal/shared/validation"

type Description struct {
	value string
}

func NewDescription(v string) (Description, error) {
	val, err := validation.ValidateStringLength("description", v, 5, 100)
	if err != nil && err.CurrentLength != 0 {
		return Description{}, *err
	}
	return Description{value: val}, nil
}

func (d Description) Value() string {
	return d.value
}

func ReconstituteDescription(v string) Description {
	return Description{value: v}
}
