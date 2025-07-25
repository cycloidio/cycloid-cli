// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// GetEventsCountReader is a Reader for the GetEventsCount structure.
type GetEventsCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventsCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventsCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetEventsCountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEventsCountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetEventsCountUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetEventsCountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEventsCountOK creates a GetEventsCountOK with default headers values
func NewGetEventsCountOK() *GetEventsCountOK {
	return &GetEventsCountOK{}
}

/*
GetEventsCountOK describes a response with status code 200, with default header values.

The count of events which fulfills the query parameters filter
*/
type GetEventsCountOK struct {
	Payload *GetEventsCountOKBody
}

// IsSuccess returns true when this get events count o k response has a 2xx status code
func (o *GetEventsCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get events count o k response has a 3xx status code
func (o *GetEventsCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get events count o k response has a 4xx status code
func (o *GetEventsCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get events count o k response has a 5xx status code
func (o *GetEventsCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get events count o k response a status code equal to that given
func (o *GetEventsCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get events count o k response
func (o *GetEventsCountOK) Code() int {
	return 200
}

func (o *GetEventsCountOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/count][%d] getEventsCountOK %s", 200, payload)
}

func (o *GetEventsCountOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/count][%d] getEventsCountOK %s", 200, payload)
}

func (o *GetEventsCountOK) GetPayload() *GetEventsCountOKBody {
	return o.Payload
}

func (o *GetEventsCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetEventsCountOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsCountForbidden creates a GetEventsCountForbidden with default headers values
func NewGetEventsCountForbidden() *GetEventsCountForbidden {
	return &GetEventsCountForbidden{}
}

/*
GetEventsCountForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetEventsCountForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get events count forbidden response has a 2xx status code
func (o *GetEventsCountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get events count forbidden response has a 3xx status code
func (o *GetEventsCountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get events count forbidden response has a 4xx status code
func (o *GetEventsCountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get events count forbidden response has a 5xx status code
func (o *GetEventsCountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get events count forbidden response a status code equal to that given
func (o *GetEventsCountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get events count forbidden response
func (o *GetEventsCountForbidden) Code() int {
	return 403
}

func (o *GetEventsCountForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/count][%d] getEventsCountForbidden %s", 403, payload)
}

func (o *GetEventsCountForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/count][%d] getEventsCountForbidden %s", 403, payload)
}

func (o *GetEventsCountForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetEventsCountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEventsCountNotFound creates a GetEventsCountNotFound with default headers values
func NewGetEventsCountNotFound() *GetEventsCountNotFound {
	return &GetEventsCountNotFound{}
}

/*
GetEventsCountNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetEventsCountNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get events count not found response has a 2xx status code
func (o *GetEventsCountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get events count not found response has a 3xx status code
func (o *GetEventsCountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get events count not found response has a 4xx status code
func (o *GetEventsCountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get events count not found response has a 5xx status code
func (o *GetEventsCountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get events count not found response a status code equal to that given
func (o *GetEventsCountNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get events count not found response
func (o *GetEventsCountNotFound) Code() int {
	return 404
}

func (o *GetEventsCountNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/count][%d] getEventsCountNotFound %s", 404, payload)
}

func (o *GetEventsCountNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/count][%d] getEventsCountNotFound %s", 404, payload)
}

func (o *GetEventsCountNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetEventsCountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEventsCountUnprocessableEntity creates a GetEventsCountUnprocessableEntity with default headers values
func NewGetEventsCountUnprocessableEntity() *GetEventsCountUnprocessableEntity {
	return &GetEventsCountUnprocessableEntity{}
}

/*
GetEventsCountUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetEventsCountUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get events count unprocessable entity response has a 2xx status code
func (o *GetEventsCountUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get events count unprocessable entity response has a 3xx status code
func (o *GetEventsCountUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get events count unprocessable entity response has a 4xx status code
func (o *GetEventsCountUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get events count unprocessable entity response has a 5xx status code
func (o *GetEventsCountUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get events count unprocessable entity response a status code equal to that given
func (o *GetEventsCountUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get events count unprocessable entity response
func (o *GetEventsCountUnprocessableEntity) Code() int {
	return 422
}

func (o *GetEventsCountUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/count][%d] getEventsCountUnprocessableEntity %s", 422, payload)
}

func (o *GetEventsCountUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/count][%d] getEventsCountUnprocessableEntity %s", 422, payload)
}

func (o *GetEventsCountUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetEventsCountUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEventsCountDefault creates a GetEventsCountDefault with default headers values
func NewGetEventsCountDefault(code int) *GetEventsCountDefault {
	return &GetEventsCountDefault{
		_statusCode: code,
	}
}

/*
GetEventsCountDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetEventsCountDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get events count default response has a 2xx status code
func (o *GetEventsCountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get events count default response has a 3xx status code
func (o *GetEventsCountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get events count default response has a 4xx status code
func (o *GetEventsCountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get events count default response has a 5xx status code
func (o *GetEventsCountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get events count default response a status code equal to that given
func (o *GetEventsCountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get events count default response
func (o *GetEventsCountDefault) Code() int {
	return o._statusCode
}

func (o *GetEventsCountDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/count][%d] getEventsCount default %s", o._statusCode, payload)
}

func (o *GetEventsCountDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/count][%d] getEventsCount default %s", o._statusCode, payload)
}

func (o *GetEventsCountDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetEventsCountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
GetEventsCountOKBody get events count o k body
swagger:model GetEventsCountOKBody
*/
type GetEventsCountOKBody struct {

	// data
	// Required: true
	Data []*models.EventsCount `json:"data"`
}

// Validate validates this get events count o k body
func (o *GetEventsCountOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEventsCountOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getEventsCountOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getEventsCountOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getEventsCountOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get events count o k body based on the context it is used
func (o *GetEventsCountOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEventsCountOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getEventsCountOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getEventsCountOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEventsCountOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEventsCountOKBody) UnmarshalBinary(b []byte) error {
	var res GetEventsCountOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
