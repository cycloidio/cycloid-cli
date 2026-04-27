package middleware

import "github.com/go-openapi/strfmt"

func uriPtr(s string) *strfmt.URI {
	u := strfmt.URI(s)
	return &u
}
