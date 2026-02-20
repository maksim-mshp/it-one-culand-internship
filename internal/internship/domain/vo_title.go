package domain

import (
	"strings"
	"unicode/utf8"
)

type Title struct {
	value string
}

func NewTitle(v *string) (Title, error) {
	if v == nil {
		return Title{}, NewMissingFieldError("title")
	}
	val := strings.TrimSpace(*v)
	n := utf8.RuneCountInString(val)
	if n < 5 || n > 100 {
		return Title{}, NewInvalidFieldLengthError("title", 5, 100, n)
	}
	return Title{value: val}, nil
}

func (t Title) Value() string {
	return t.value
}

func ReconstituteTitle(v string) Title {
	return Title{value: v}
}
