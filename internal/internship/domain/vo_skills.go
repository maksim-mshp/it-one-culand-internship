package domain

import (
	"strings"
	"unicode/utf8"
)

type Skills struct {
	value []string
}

func NewSkills(v *[]string) (Skills, error) {
	if v == nil {
		return Skills{}, nil
	}
	val := *v
	for i := range val {
		val[i] = strings.TrimSpace(val[i])
		n := utf8.RuneCountInString(val[i])
		if n < 5 || n > 255 {
			return Skills{}, NewInvalidFieldLengthError("skills", 5, 255, n)
		}
	}
	return Skills{value: val}, nil
}

func (s Skills) Value() []string {
	return s.value
}

func ReconstituteSkills(v []string) Skills {
	return Skills{value: v}
}
