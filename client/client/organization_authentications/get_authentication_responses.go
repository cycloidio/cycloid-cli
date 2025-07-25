// Code generated by go-swagger; DO NOT EDIT.

package organization_authentications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// GetAuthenticationReader is a Reader for the GetAuthentication structure.
type GetAuthenticationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthenticationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthenticationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetAuthenticationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthenticationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAuthenticationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAuthenticationOK creates a GetAuthenticationOK with default headers values
func NewGetAuthenticationOK() *GetAuthenticationOK {
	return &GetAuthenticationOK{}
}

/*
GetAuthenticationOK describes a response with status code 200, with default header values.

Authentication available in the organization with such canonical.
*/
type GetAuthenticationOK struct {
	Payload *GetAuthenticationOKBody
}

// IsSuccess returns true when this get authentication o k response has a 2xx status code
func (o *GetAuthenticationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get authentication o k response has a 3xx status code
func (o *GetAuthenticationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authentication o k response has a 4xx status code
func (o *GetAuthenticationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authentication o k response has a 5xx status code
func (o *GetAuthenticationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get authentication o k response a status code equal to that given
func (o *GetAuthenticationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get authentication o k response
func (o *GetAuthenticationOK) Code() int {
	return 200
}

func (o *GetAuthenticationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/authentications/{authentication_type}][%d] getAuthenticationOK %s", 200, payload)
}

func (o *GetAuthenticationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/authentications/{authentication_type}][%d] getAuthenticationOK %s", 200, payload)
}

func (o *GetAuthenticationOK) GetPayload() *GetAuthenticationOKBody {
	return o.Payload
}

func (o *GetAuthenticationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAuthenticationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthenticationForbidden creates a GetAuthenticationForbidden with default headers values
func NewGetAuthenticationForbidden() *GetAuthenticationForbidden {
	return &GetAuthenticationForbidden{}
}

/*
GetAuthenticationForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetAuthenticationForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get authentication forbidden response has a 2xx status code
func (o *GetAuthenticationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authentication forbidden response has a 3xx status code
func (o *GetAuthenticationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authentication forbidden response has a 4xx status code
func (o *GetAuthenticationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authentication forbidden response has a 5xx status code
func (o *GetAuthenticationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get authentication forbidden response a status code equal to that given
func (o *GetAuthenticationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get authentication forbidden response
func (o *GetAuthenticationForbidden) Code() int {
	return 403
}

func (o *GetAuthenticationForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/authentications/{authentication_type}][%d] getAuthenticationForbidden %s", 403, payload)
}

func (o *GetAuthenticationForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/authentications/{authentication_type}][%d] getAuthenticationForbidden %s", 403, payload)
}

func (o *GetAuthenticationForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetAuthenticationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAuthenticationNotFound creates a GetAuthenticationNotFound with default headers values
func NewGetAuthenticationNotFound() *GetAuthenticationNotFound {
	return &GetAuthenticationNotFound{}
}

/*
GetAuthenticationNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetAuthenticationNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get authentication not found response has a 2xx status code
func (o *GetAuthenticationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authentication not found response has a 3xx status code
func (o *GetAuthenticationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authentication not found response has a 4xx status code
func (o *GetAuthenticationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authentication not found response has a 5xx status code
func (o *GetAuthenticationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get authentication not found response a status code equal to that given
func (o *GetAuthenticationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get authentication not found response
func (o *GetAuthenticationNotFound) Code() int {
	return 404
}

func (o *GetAuthenticationNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/authentications/{authentication_type}][%d] getAuthenticationNotFound %s", 404, payload)
}

func (o *GetAuthenticationNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/authentications/{authentication_type}][%d] getAuthenticationNotFound %s", 404, payload)
}

func (o *GetAuthenticationNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetAuthenticationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAuthenticationDefault creates a GetAuthenticationDefault with default headers values
func NewGetAuthenticationDefault(code int) *GetAuthenticationDefault {
	return &GetAuthenticationDefault{
		_statusCode: code,
	}
}

/*
GetAuthenticationDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetAuthenticationDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get authentication default response has a 2xx status code
func (o *GetAuthenticationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get authentication default response has a 3xx status code
func (o *GetAuthenticationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get authentication default response has a 4xx status code
func (o *GetAuthenticationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get authentication default response has a 5xx status code
func (o *GetAuthenticationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get authentication default response a status code equal to that given
func (o *GetAuthenticationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get authentication default response
func (o *GetAuthenticationDefault) Code() int {
	return o._statusCode
}

func (o *GetAuthenticationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/authentications/{authentication_type}][%d] getAuthentication default %s", o._statusCode, payload)
}

func (o *GetAuthenticationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/authentications/{authentication_type}][%d] getAuthentication default %s", o._statusCode, payload)
}

func (o *GetAuthenticationDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetAuthenticationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetAuthenticationOKBody get authentication o k body
swagger:model GetAuthenticationOKBody
*/
type GetAuthenticationOKBody struct {

	// data
	// Required: true
	Data *models.Authentication `json:"data"`
}

// Validate validates this get authentication o k body
func (o *GetAuthenticationOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAuthenticationOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getAuthenticationOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAuthenticationOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getAuthenticationOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get authentication o k body based on the context it is used
func (o *GetAuthenticationOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAuthenticationOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAuthenticationOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getAuthenticationOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAuthenticationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAuthenticationOKBody) UnmarshalBinary(b []byte) error {
	var res GetAuthenticationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
