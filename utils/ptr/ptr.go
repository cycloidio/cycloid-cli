package ptr

import "github.com/go-openapi/strfmt"

// Email returns the address of a literal strfmt.Email
func Email(s string) *strfmt.Email { e := strfmt.Email(s); return &e }

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

// PtrOrNil returns a pointer to the value or nil if it's an empty value.
// Can be used as a quick check if a struct is empty.
// Careful! Default values like boolean false and integer 0 are treated as empty values.
// For them too the function will return nil.
func PtrOrNil[V comparable](v V) *V {
	var empty V
	if v == empty {
		return nil
	}
	return &v
}
