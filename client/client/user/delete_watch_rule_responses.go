// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// DeleteWatchRuleReader is a Reader for the DeleteWatchRule structure.
type DeleteWatchRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWatchRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteWatchRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteWatchRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteWatchRuleOK creates a DeleteWatchRuleOK with default headers values
func NewDeleteWatchRuleOK() *DeleteWatchRuleOK {
	return &DeleteWatchRuleOK{}
}

/*
DeleteWatchRuleOK describes a response with status code 200, with default header values.

The watch rule was deleted
*/
type DeleteWatchRuleOK struct {
}

// IsSuccess returns true when this delete watch rule o k response has a 2xx status code
func (o *DeleteWatchRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete watch rule o k response has a 3xx status code
func (o *DeleteWatchRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete watch rule o k response has a 4xx status code
func (o *DeleteWatchRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete watch rule o k response has a 5xx status code
func (o *DeleteWatchRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete watch rule o k response a status code equal to that given
func (o *DeleteWatchRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete watch rule o k response
func (o *DeleteWatchRuleOK) Code() int {
	return 200
}

func (o *DeleteWatchRuleOK) Error() string {
	return fmt.Sprintf("[DELETE /user/watch_rules/{watch_rule_canonical}][%d] deleteWatchRuleOK", 200)
}

func (o *DeleteWatchRuleOK) String() string {
	return fmt.Sprintf("[DELETE /user/watch_rules/{watch_rule_canonical}][%d] deleteWatchRuleOK", 200)
}

func (o *DeleteWatchRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWatchRuleDefault creates a DeleteWatchRuleDefault with default headers values
func NewDeleteWatchRuleDefault(code int) *DeleteWatchRuleDefault {
	return &DeleteWatchRuleDefault{
		_statusCode: code,
	}
}

/*
DeleteWatchRuleDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DeleteWatchRuleDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete watch rule default response has a 2xx status code
func (o *DeleteWatchRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete watch rule default response has a 3xx status code
func (o *DeleteWatchRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete watch rule default response has a 4xx status code
func (o *DeleteWatchRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete watch rule default response has a 5xx status code
func (o *DeleteWatchRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete watch rule default response a status code equal to that given
func (o *DeleteWatchRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete watch rule default response
func (o *DeleteWatchRuleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteWatchRuleDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /user/watch_rules/{watch_rule_canonical}][%d] deleteWatchRule default %s", o._statusCode, payload)
}

func (o *DeleteWatchRuleDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /user/watch_rules/{watch_rule_canonical}][%d] deleteWatchRule default %s", o._statusCode, payload)
}

func (o *DeleteWatchRuleDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteWatchRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Length
	hdrContentLength := response.GetHeader("Content-Length")

	if hdrContentLength != "" {
		valcontentLength, err := swag.ConvertUint64(hdrContentLength)
		if err != nil {
			return errors.InvalidType("Content-Length", "header", "uint64", hdrContentLength)
		}
		o.ContentLength = valcontentLength
	}

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
