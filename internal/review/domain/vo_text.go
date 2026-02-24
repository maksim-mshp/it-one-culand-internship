package domain

import "culand-internship/internal/shared/validation"

type Text struct {
	value string
}

func NewText(v string) (Text, error) {
	val, err := validation.ValidateStringLength("text", v, 5, 500)
	if err != nil {
		if err.CurrentLength == 0 {
			return Text{}, validation.RequiredFieldMissingError{Field: "text"}
		}
		return Text{}, *err
	}
	return Text{value: val}, nil
}

func (t Text) Value() string {
	return t.value
}

func ReconstituteText(v string) Text {
	return Text{value: v}
}
