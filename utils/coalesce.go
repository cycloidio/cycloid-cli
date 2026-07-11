package utils

// Coalesce will return the first non-empty value
func Coalesce[T any](values ...*T) *T {
	for _, v := range values {
		if v != nil {
			return v
		}
	}

	return nil
}

func CoalesceNonZero[T comparable](values ...T) T {
	var zero T
	for _, v := range values {
		if v != zero {
			return v
		}
	}
	return zero
}

func CoalesceNonZeroPtr[T comparable](values ...T) *T {
	var zero T
	for _, v := range values {
		if v != zero {
			return &v
		}
	}
	return nil
}
