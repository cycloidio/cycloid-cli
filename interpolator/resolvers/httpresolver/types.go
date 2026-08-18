package httpresolver

import (
	"fmt"
	"strings"
)

type APIResponse struct {
	Data      any          `json:"data,omitempty"`
	Errors    ErrorDetails `json:"errors,omitempty"`
	RequestID string       `json:"request_id,omitempty"`
}

type ErrorDetails []ErrorDetail

func (es ErrorDetails) Error() string {
	outErr := make([]string, len(es))
	for i, e := range es {
		outErr[i] = e.String()
	}
	return strings.Join(outErr, ": ")
}

type ErrorDetail struct {
	Message string   `json:"message"`
	Code    string   `json:"code"`
	Details []string `json:"details"`
}

func (e *ErrorDetail) String() string {
	return strings.Trim(
		fmt.Sprintf(`%s: %s %s`, e.Code, e.Message, strings.Join(e.Details, ", ")), " ",
	)
}
