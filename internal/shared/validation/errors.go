package validation

import "fmt"

type RequiredFieldMissingError struct {
	Field string
}

func (e RequiredFieldMissingError) Error() string {
	return fmt.Sprintf("required field %s is missing", e.Field)
}

type InvalidFieldLengthError struct {
	Field         string
	Value         string
	MinLength     int
	MaxLength     int
	CurrentLength int
}

func (e InvalidFieldLengthError) Error() string {
	return fmt.Sprintf("field %s=%s has invalid length (cur: %d, min: %d, max: %d)",
		e.Field,
		e.Value,
		e.CurrentLength,
		e.MinLength,
		e.MaxLength,
	)
}

type InvalidArrayItemLengthError struct {
	Field         string
	Index         int
	Value         string
	MinLength     int
	MaxLength     int
	CurrentLength int
}

func (e InvalidArrayItemLengthError) Error() string {
	return fmt.Sprintf("field %s[%d]=%s has invalid length (cur: %d, min: %d, max: %d)",
		e.Field,
		e.Index,
		e.Value,
		e.CurrentLength,
		e.MinLength,
		e.MaxLength,
	)
}
