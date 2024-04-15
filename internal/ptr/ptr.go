package ptr

// Ptr returns a pointer address of given value
func Ptr[V any](v V) *V { return &v }

// Value returns a value of given pointer
func Value[V any](v *V) V {
	if v == nil {
		var noop V
		return noop
	}
	return *v
}
