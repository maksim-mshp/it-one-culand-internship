package validation

import (
	"strings"
	"unicode/utf8"
)

func ValidateStringLength(field string, v string, min, max int) (string, *InvalidFieldLengthError) {
	val := strings.TrimSpace(v)
	n := utf8.RuneCountInString(val)
	if n < min || n > max {
		return "", &InvalidFieldLengthError{field, val, min, max, n}
	}
	return val, nil
}

func ValidateSliceItemsLength(field string, arr []string, min, max int) ([]string, *InvalidArrayItemLengthError) {
	var result []string
	for i, v := range arr {
		val := strings.TrimSpace(v)
		n := utf8.RuneCountInString(val)
		if n < min || n > max {
			return nil, &InvalidArrayItemLengthError{field, i, val, min, max, n}
		}
		result = append(result, val)
	}
	return result, nil
}
