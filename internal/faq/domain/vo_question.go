package domain

import "culand-internship/internal/shared/validation"

type Question struct {
	value string
}

func NewQuestion(v string) (Question, error) {
	val, err := validation.ValidateStringLength("question", v, 5, 500)
	if err != nil {
		if err.CurrentLength == 0 {
			return Question{}, validation.RequiredFieldMissingError{Field: "question"}
		}
		return Question{}, *err
	}
	return Question{value: val}, nil
}

func (t Question) Value() string {
	return t.value
}

func ReconstituteQuestion(v string) Question {
	return Question{value: v}
}
