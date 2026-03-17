package printer

import "strings"

// ErrHTTPResponse is implemented by errors that attach an API HTTP response
// for diagnostics (e.g. JSON printer fallback when marshaling fails).
type ErrHTTPResponse interface {
	error
	HTTPStatusCode() int
	HTTPResponseBody() []byte
}

// RequestPather is optionally implemented when the error includes the request URI.
type RequestPather interface {
	HTTPRequestPath() string
}

// FirstNLinesFromBytes returns the first n lines of b (split on '\n'),
// rejoined with '\n'. Trailing whitespace on the joined string is trimmed.
func FirstNLinesFromBytes(b []byte, n int) string {
	if n <= 0 {
		return ""
	}
	s := string(b)
	lines := strings.Split(s, "\n")
	if len(lines) > n {
		lines = lines[:n]
	}
	return strings.TrimSpace(strings.Join(lines, "\n"))
}
