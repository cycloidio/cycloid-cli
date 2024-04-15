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
