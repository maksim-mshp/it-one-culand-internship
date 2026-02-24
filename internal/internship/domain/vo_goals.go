package domain

import "culand-internship/internal/shared/validation"

type Goals struct {
	value []string
}

func NewGoals(v []string) (Goals, error) {
	val, err := validation.ValidateSliceItemsLength("goals", v, 5, 200)
	if err != nil {
		return Goals{}, *err
	}
	return Goals{value: val}, nil
}

func (g Goals) Value() []string {
	return g.value
}

func ReconstituteGoals(v []string) Goals {
	return Goals{value: v}
}
