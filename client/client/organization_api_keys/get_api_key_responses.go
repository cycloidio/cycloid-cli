// Code generated by go-swagger; DO NOT EDIT.

package organization_api_keys

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

// GetAPIKeyReader is a Reader for the GetAPIKey structure.
type GetAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetAPIKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAPIKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAPIKeyOK creates a GetAPIKeyOK with default headers values
func NewGetAPIKeyOK() *GetAPIKeyOK {
	return &GetAPIKeyOK{}
}

/*
GetAPIKeyOK describes a response with status code 200, with default header values.

The information of the API key of the organization which has the specified canonical.
*/
type GetAPIKeyOK struct {
	Payload *GetAPIKeyOKBody
}

// IsSuccess returns true when this get Api key o k response has a 2xx status code
func (o *GetAPIKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api key o k response has a 3xx status code
func (o *GetAPIKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api key o k response has a 4xx status code
func (o *GetAPIKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api key o k response has a 5xx status code
func (o *GetAPIKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api key o k response a status code equal to that given
func (o *GetAPIKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api key o k response
func (o *GetAPIKeyOK) Code() int {
	return 200
}

func (o *GetAPIKeyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] getApiKeyOK %s", 200, payload)
}

func (o *GetAPIKeyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] getApiKeyOK %s", 200, payload)
}

func (o *GetAPIKeyOK) GetPayload() *GetAPIKeyOKBody {
	return o.Payload
}

func (o *GetAPIKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAPIKeyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIKeyForbidden creates a GetAPIKeyForbidden with default headers values
func NewGetAPIKeyForbidden() *GetAPIKeyForbidden {
	return &GetAPIKeyForbidden{}
}

/*
GetAPIKeyForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetAPIKeyForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get Api key forbidden response has a 2xx status code
func (o *GetAPIKeyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api key forbidden response has a 3xx status code
func (o *GetAPIKeyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api key forbidden response has a 4xx status code
func (o *GetAPIKeyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api key forbidden response has a 5xx status code
func (o *GetAPIKeyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api key forbidden response a status code equal to that given
func (o *GetAPIKeyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get Api key forbidden response
func (o *GetAPIKeyForbidden) Code() int {
	return 403
}

func (o *GetAPIKeyForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] getApiKeyForbidden %s", 403, payload)
}

func (o *GetAPIKeyForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] getApiKeyForbidden %s", 403, payload)
}

func (o *GetAPIKeyForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetAPIKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAPIKeyNotFound creates a GetAPIKeyNotFound with default headers values
func NewGetAPIKeyNotFound() *GetAPIKeyNotFound {
	return &GetAPIKeyNotFound{}
}

/*
GetAPIKeyNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetAPIKeyNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get Api key not found response has a 2xx status code
func (o *GetAPIKeyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Api key not found response has a 3xx status code
func (o *GetAPIKeyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api key not found response has a 4xx status code
func (o *GetAPIKeyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Api key not found response has a 5xx status code
func (o *GetAPIKeyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api key not found response a status code equal to that given
func (o *GetAPIKeyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get Api key not found response
func (o *GetAPIKeyNotFound) Code() int {
	return 404
}

func (o *GetAPIKeyNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] getApiKeyNotFound %s", 404, payload)
}

func (o *GetAPIKeyNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] getApiKeyNotFound %s", 404, payload)
}

func (o *GetAPIKeyNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetAPIKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAPIKeyDefault creates a GetAPIKeyDefault with default headers values
func NewGetAPIKeyDefault(code int) *GetAPIKeyDefault {
	return &GetAPIKeyDefault{
		_statusCode: code,
	}
}

/*
GetAPIKeyDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetAPIKeyDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get API key default response has a 2xx status code
func (o *GetAPIKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get API key default response has a 3xx status code
func (o *GetAPIKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get API key default response has a 4xx status code
func (o *GetAPIKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get API key default response has a 5xx status code
func (o *GetAPIKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get API key default response a status code equal to that given
func (o *GetAPIKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get API key default response
func (o *GetAPIKeyDefault) Code() int {
	return o._statusCode
}

func (o *GetAPIKeyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] getAPIKey default %s", o._statusCode, payload)
}

func (o *GetAPIKeyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] getAPIKey default %s", o._statusCode, payload)
}

func (o *GetAPIKeyDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetAPIKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetAPIKeyOKBody get API key o k body
swagger:model GetAPIKeyOKBody
*/
type GetAPIKeyOKBody struct {

	// data
	// Required: true
	Data *models.APIKey `json:"data"`
}

// Validate validates this get API key o k body
func (o *GetAPIKeyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAPIKeyOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getApiKeyOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getApiKeyOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getApiKeyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get API key o k body based on the context it is used
func (o *GetAPIKeyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAPIKeyOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getApiKeyOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getApiKeyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAPIKeyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAPIKeyOKBody) UnmarshalBinary(b []byte) error {
	var res GetAPIKeyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
