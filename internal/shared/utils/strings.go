package utils

func EmptyStringToNil(s string) *string {
	if len(s) == 0 {
		return nil
	}
	return &s
}
