// Code generated by go-swagger; DO NOT EDIT.

package organization_members

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

// GetOrgMemberReader is a Reader for the GetOrgMember structure.
type GetOrgMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetOrgMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOrgMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrgMemberOK creates a GetOrgMemberOK with default headers values
func NewGetOrgMemberOK() *GetOrgMemberOK {
	return &GetOrgMemberOK{}
}

/*
GetOrgMemberOK describes a response with status code 200, with default header values.

The information of the member of the organization.
*/
type GetOrgMemberOK struct {
	Payload *GetOrgMemberOKBody
}

// IsSuccess returns true when this get org member o k response has a 2xx status code
func (o *GetOrgMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get org member o k response has a 3xx status code
func (o *GetOrgMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org member o k response has a 4xx status code
func (o *GetOrgMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org member o k response has a 5xx status code
func (o *GetOrgMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get org member o k response a status code equal to that given
func (o *GetOrgMemberOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get org member o k response
func (o *GetOrgMemberOK) Code() int {
	return 200
}

func (o *GetOrgMemberOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members/{username}][%d] getOrgMemberOK %s", 200, payload)
}

func (o *GetOrgMemberOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members/{username}][%d] getOrgMemberOK %s", 200, payload)
}

func (o *GetOrgMemberOK) GetPayload() *GetOrgMemberOKBody {
	return o.Payload
}

func (o *GetOrgMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrgMemberOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgMemberForbidden creates a GetOrgMemberForbidden with default headers values
func NewGetOrgMemberForbidden() *GetOrgMemberForbidden {
	return &GetOrgMemberForbidden{}
}

/*
GetOrgMemberForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetOrgMemberForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get org member forbidden response has a 2xx status code
func (o *GetOrgMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org member forbidden response has a 3xx status code
func (o *GetOrgMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org member forbidden response has a 4xx status code
func (o *GetOrgMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org member forbidden response has a 5xx status code
func (o *GetOrgMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get org member forbidden response a status code equal to that given
func (o *GetOrgMemberForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get org member forbidden response
func (o *GetOrgMemberForbidden) Code() int {
	return 403
}

func (o *GetOrgMemberForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members/{username}][%d] getOrgMemberForbidden %s", 403, payload)
}

func (o *GetOrgMemberForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members/{username}][%d] getOrgMemberForbidden %s", 403, payload)
}

func (o *GetOrgMemberForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOrgMemberNotFound creates a GetOrgMemberNotFound with default headers values
func NewGetOrgMemberNotFound() *GetOrgMemberNotFound {
	return &GetOrgMemberNotFound{}
}

/*
GetOrgMemberNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetOrgMemberNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get org member not found response has a 2xx status code
func (o *GetOrgMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org member not found response has a 3xx status code
func (o *GetOrgMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org member not found response has a 4xx status code
func (o *GetOrgMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org member not found response has a 5xx status code
func (o *GetOrgMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get org member not found response a status code equal to that given
func (o *GetOrgMemberNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get org member not found response
func (o *GetOrgMemberNotFound) Code() int {
	return 404
}

func (o *GetOrgMemberNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members/{username}][%d] getOrgMemberNotFound %s", 404, payload)
}

func (o *GetOrgMemberNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members/{username}][%d] getOrgMemberNotFound %s", 404, payload)
}

func (o *GetOrgMemberNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOrgMemberDefault creates a GetOrgMemberDefault with default headers values
func NewGetOrgMemberDefault(code int) *GetOrgMemberDefault {
	return &GetOrgMemberDefault{
		_statusCode: code,
	}
}

/*
GetOrgMemberDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetOrgMemberDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get org member default response has a 2xx status code
func (o *GetOrgMemberDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get org member default response has a 3xx status code
func (o *GetOrgMemberDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get org member default response has a 4xx status code
func (o *GetOrgMemberDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get org member default response has a 5xx status code
func (o *GetOrgMemberDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get org member default response a status code equal to that given
func (o *GetOrgMemberDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get org member default response
func (o *GetOrgMemberDefault) Code() int {
	return o._statusCode
}

func (o *GetOrgMemberDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members/{username}][%d] getOrgMember default %s", o._statusCode, payload)
}

func (o *GetOrgMemberDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members/{username}][%d] getOrgMember default %s", o._statusCode, payload)
}

func (o *GetOrgMemberDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetOrgMemberOKBody get org member o k body
swagger:model GetOrgMemberOKBody
*/
type GetOrgMemberOKBody struct {

	// data
	// Required: true
	Data *models.MemberOrg `json:"data"`
}

// Validate validates this get org member o k body
func (o *GetOrgMemberOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrgMemberOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getOrgMemberOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrgMemberOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrgMemberOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get org member o k body based on the context it is used
func (o *GetOrgMemberOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrgMemberOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrgMemberOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrgMemberOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrgMemberOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrgMemberOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrgMemberOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
