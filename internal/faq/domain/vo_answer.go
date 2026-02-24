package domain

import "culand-internship/internal/shared/validation"

type Answer struct {
	value string
}

func NewAnswer(v string) (Answer, error) {
	val, err := validation.ValidateStringLength("answer", v, 5, 500)
	if err != nil {
		if err.CurrentLength == 0 {
			return Answer{}, validation.RequiredFieldMissingError{Field: "answer"}
		}
		return Answer{}, *err
	}
	return Answer{value: val}, nil
}

func (t Answer) Value() string {
	return t.value
}

func ReconstituteAnswer(v string) Answer {
	return Answer{value: v}
}
