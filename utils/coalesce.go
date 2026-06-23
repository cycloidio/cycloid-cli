package utils

// Coalesce returns the first non-nil pointer value
func Coalesce[T any](values ...*T) *T {
	for _, v := range values {
		if v != nil {
			return v
		}
	}

	return nil
}

// CoalesceNonZero returns the first non-zero value
func CoalesceNonZero[T comparable](values ...T) T {
	var zero T
	for _, v := range values {
		if v != zero {
			return v
		}
	}
	return zero
}

// CoalesceNonZeroPtr returns a pointer to the first non-zero value, or nil
func CoalesceNonZeroPtr[T comparable](values ...T) *T {
	var zero T
	for _, v := range values {
		if v != zero {
			return &v
		}
	}
	return nil
}
