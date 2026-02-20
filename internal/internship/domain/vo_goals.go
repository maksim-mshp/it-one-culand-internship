package domain

import (
	"strings"
	"unicode/utf8"
)

type Goals struct {
	value []string
}

func NewGoals(v *[]string) (Goals, error) {
	if v == nil {
		return Goals{}, nil
	}
	val := *v
	for i := range val {
		val[i] = strings.TrimSpace(val[i])
		n := utf8.RuneCountInString(val[i])
		if n < 5 || n > 255 {
			return Goals{}, NewInvalidFieldLengthError("goals", 5, 255, n)
		}
	}
	return Goals{value: val}, nil
}

func (g Goals) Value() []string {
	return g.value
}

func ReconstituteGoals(v []string) Goals {
	return Goals{value: v}
}
