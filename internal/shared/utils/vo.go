package utils

func BuildVO[T any, VO any](ptr *T, factory func(T) (VO, error)) (VO, error) {
	var val T
	if ptr != nil {
		val = *ptr
	}
	return factory(val)
}

func UpdateVO[T any, VO any](ptr *T, factory func(T) (VO, error), cur VO) (VO, error) {
	if ptr == nil {
		return cur, nil
	}
	return factory(*ptr)
}

func ReconstituteVO[T any, VO any](ptr *T, factory func(T) VO) VO {
	var val T
	if ptr != nil {
		val = *ptr
	}
	return factory(val)
}
