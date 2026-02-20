package domain

import (
	"strings"
	"unicode/utf8"
)

type Description struct {
	value string
}

func NewDescription(v *string) (Description, error) {
	if v == nil {
		return Description{}, nil
	}
	val := strings.TrimSpace(*v)
	n := utf8.RuneCountInString(val)
	if n < 10 || n > 1000 {
		return Description{}, NewInvalidFieldLengthError("description", 5, 500, n)
	}
	return Description{value: val}, nil
}

func (d Description) Value() string {
	return d.value
}

func ReconstituteDescription(v string) Description {
	return Description{value: v}
}
