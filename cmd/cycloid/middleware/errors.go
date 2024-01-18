package middleware

import (
	"fmt"
	"regexp"

	"github.com/cycloidio/cycloid-cli/client/models"
)

type ErrorPayloader interface {
	Error() string
	GetPayload() *models.ErrorPayload
}

type ApiError struct {
	HttpMethod string
	URL        string
	HttpCode   string
	ApiAction  string

	Payload *models.ErrorPayload
}

var reApiError = regexp.MustCompile(`\[(?P<httpmethod>\w+)\s(?P<url>.*)\]\[(?P<httpcode>\d{3})\]\s(?P<apiaction>\w+)\s`)

// NewApiError will try to convert the err to a more standard one if possible,
// if the err does not implement ErrorPayloader and not match the reApiError
// then nothing will be done and the same err will be returned
func NewApiError(err error) error {
	ep, ok := err.(ErrorPayloader)
	// If it's not implementing the interface then we just return
	// the old error
	if !ok {
		return err
	}

	match := reApiError.FindStringSubmatch(ep.Error())
	if match == nil {
		// If even being the ErrorPayload we cannot match the string
		// we just still return the old err
		return err
	}

	apierr := ApiError{
		HttpMethod: match[1],
		URL:        match[2],
		HttpCode:   match[3],
		ApiAction:  match[4],

		Payload: ep.GetPayload(),
	}

	return &apierr
}

func (a *ApiError) Error() string {
	var msg string

	if a.Payload != nil && len(a.Payload.Errors) != 0 && a.Payload.Errors[0].Message != nil {
		msg = *a.Payload.Errors[0].Message
	}
	return fmt.Sprintf("A %s error was returned on %q call with message: %s", a.HttpCode, a.ApiAction, msg)
}
