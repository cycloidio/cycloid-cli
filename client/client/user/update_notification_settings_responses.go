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

// UpdateNotificationSettingsReader is a Reader for the UpdateNotificationSettings structure.
type UpdateNotificationSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNotificationSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateNotificationSettingsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewUpdateNotificationSettingsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateNotificationSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateNotificationSettingsNoContent creates a UpdateNotificationSettingsNoContent with default headers values
func NewUpdateNotificationSettingsNoContent() *UpdateNotificationSettingsNoContent {
	return &UpdateNotificationSettingsNoContent{}
}

/*
UpdateNotificationSettingsNoContent describes a response with status code 204, with default header values.

The notification settings were updated
*/
type UpdateNotificationSettingsNoContent struct {
}

// IsSuccess returns true when this update notification settings no content response has a 2xx status code
func (o *UpdateNotificationSettingsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update notification settings no content response has a 3xx status code
func (o *UpdateNotificationSettingsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update notification settings no content response has a 4xx status code
func (o *UpdateNotificationSettingsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update notification settings no content response has a 5xx status code
func (o *UpdateNotificationSettingsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update notification settings no content response a status code equal to that given
func (o *UpdateNotificationSettingsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update notification settings no content response
func (o *UpdateNotificationSettingsNoContent) Code() int {
	return 204
}

func (o *UpdateNotificationSettingsNoContent) Error() string {
	return fmt.Sprintf("[POST /user/notification_settings][%d] updateNotificationSettingsNoContent", 204)
}

func (o *UpdateNotificationSettingsNoContent) String() string {
	return fmt.Sprintf("[POST /user/notification_settings][%d] updateNotificationSettingsNoContent", 204)
}

func (o *UpdateNotificationSettingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNotificationSettingsUnprocessableEntity creates a UpdateNotificationSettingsUnprocessableEntity with default headers values
func NewUpdateNotificationSettingsUnprocessableEntity() *UpdateNotificationSettingsUnprocessableEntity {
	return &UpdateNotificationSettingsUnprocessableEntity{}
}

/*
UpdateNotificationSettingsUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateNotificationSettingsUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update notification settings unprocessable entity response has a 2xx status code
func (o *UpdateNotificationSettingsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update notification settings unprocessable entity response has a 3xx status code
func (o *UpdateNotificationSettingsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update notification settings unprocessable entity response has a 4xx status code
func (o *UpdateNotificationSettingsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update notification settings unprocessable entity response has a 5xx status code
func (o *UpdateNotificationSettingsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update notification settings unprocessable entity response a status code equal to that given
func (o *UpdateNotificationSettingsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update notification settings unprocessable entity response
func (o *UpdateNotificationSettingsUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateNotificationSettingsUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/notification_settings][%d] updateNotificationSettingsUnprocessableEntity %s", 422, payload)
}

func (o *UpdateNotificationSettingsUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/notification_settings][%d] updateNotificationSettingsUnprocessableEntity %s", 422, payload)
}

func (o *UpdateNotificationSettingsUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateNotificationSettingsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateNotificationSettingsDefault creates a UpdateNotificationSettingsDefault with default headers values
func NewUpdateNotificationSettingsDefault(code int) *UpdateNotificationSettingsDefault {
	return &UpdateNotificationSettingsDefault{
		_statusCode: code,
	}
}

/*
UpdateNotificationSettingsDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateNotificationSettingsDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update notification settings default response has a 2xx status code
func (o *UpdateNotificationSettingsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update notification settings default response has a 3xx status code
func (o *UpdateNotificationSettingsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update notification settings default response has a 4xx status code
func (o *UpdateNotificationSettingsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update notification settings default response has a 5xx status code
func (o *UpdateNotificationSettingsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update notification settings default response a status code equal to that given
func (o *UpdateNotificationSettingsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update notification settings default response
func (o *UpdateNotificationSettingsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateNotificationSettingsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/notification_settings][%d] updateNotificationSettings default %s", o._statusCode, payload)
}

func (o *UpdateNotificationSettingsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/notification_settings][%d] updateNotificationSettings default %s", o._statusCode, payload)
}

func (o *UpdateNotificationSettingsDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateNotificationSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
