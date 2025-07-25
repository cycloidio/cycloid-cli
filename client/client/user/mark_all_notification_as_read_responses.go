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

// MarkAllNotificationAsReadReader is a Reader for the MarkAllNotificationAsRead structure.
type MarkAllNotificationAsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MarkAllNotificationAsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMarkAllNotificationAsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMarkAllNotificationAsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMarkAllNotificationAsReadOK creates a MarkAllNotificationAsReadOK with default headers values
func NewMarkAllNotificationAsReadOK() *MarkAllNotificationAsReadOK {
	return &MarkAllNotificationAsReadOK{}
}

/*
MarkAllNotificationAsReadOK describes a response with status code 200, with default header values.

The notifications was marked as read
*/
type MarkAllNotificationAsReadOK struct {
}

// IsSuccess returns true when this mark all notification as read o k response has a 2xx status code
func (o *MarkAllNotificationAsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this mark all notification as read o k response has a 3xx status code
func (o *MarkAllNotificationAsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this mark all notification as read o k response has a 4xx status code
func (o *MarkAllNotificationAsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this mark all notification as read o k response has a 5xx status code
func (o *MarkAllNotificationAsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this mark all notification as read o k response a status code equal to that given
func (o *MarkAllNotificationAsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the mark all notification as read o k response
func (o *MarkAllNotificationAsReadOK) Code() int {
	return 200
}

func (o *MarkAllNotificationAsReadOK) Error() string {
	return fmt.Sprintf("[POST /user/notifications/read_all][%d] markAllNotificationAsReadOK", 200)
}

func (o *MarkAllNotificationAsReadOK) String() string {
	return fmt.Sprintf("[POST /user/notifications/read_all][%d] markAllNotificationAsReadOK", 200)
}

func (o *MarkAllNotificationAsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMarkAllNotificationAsReadDefault creates a MarkAllNotificationAsReadDefault with default headers values
func NewMarkAllNotificationAsReadDefault(code int) *MarkAllNotificationAsReadDefault {
	return &MarkAllNotificationAsReadDefault{
		_statusCode: code,
	}
}

/*
MarkAllNotificationAsReadDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type MarkAllNotificationAsReadDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this mark all notification as read default response has a 2xx status code
func (o *MarkAllNotificationAsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this mark all notification as read default response has a 3xx status code
func (o *MarkAllNotificationAsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this mark all notification as read default response has a 4xx status code
func (o *MarkAllNotificationAsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this mark all notification as read default response has a 5xx status code
func (o *MarkAllNotificationAsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this mark all notification as read default response a status code equal to that given
func (o *MarkAllNotificationAsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the mark all notification as read default response
func (o *MarkAllNotificationAsReadDefault) Code() int {
	return o._statusCode
}

func (o *MarkAllNotificationAsReadDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/notifications/read_all][%d] markAllNotificationAsRead default %s", o._statusCode, payload)
}

func (o *MarkAllNotificationAsReadDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/notifications/read_all][%d] markAllNotificationAsRead default %s", o._statusCode, payload)
}

func (o *MarkAllNotificationAsReadDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *MarkAllNotificationAsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
