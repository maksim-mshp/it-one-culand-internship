package domain

import (
	"strings"
	"unicode/utf8"
)

type Label struct {
	value string
}

func NewLabel(v *string) (Label, error) {
	if v == nil {
		return Label{}, nil
	}
	val := strings.TrimSpace(*v)
	n := utf8.RuneCountInString(val)
	if n < 1 || n > 50 {
		return Label{}, NewInvalidFieldLengthError("label", 1, 50, n)
	}
	return Label{value: val}, nil
}

func (l Label) Value() string {
	return l.value
}

func ReconstituteLabel(v string) Label {
	return Label{value: v}
}
