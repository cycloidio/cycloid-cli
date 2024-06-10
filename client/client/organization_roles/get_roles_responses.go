// Code generated by go-swagger; DO NOT EDIT.

package organization_roles

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

// GetRolesReader is a Reader for the GetRoles structure.
type GetRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRolesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRolesOK creates a GetRolesOK with default headers values
func NewGetRolesOK() *GetRolesOK {
	return &GetRolesOK{}
}

/*
GetRolesOK describes a response with status code 200, with default header values.

List of roles which are available in the organization.
*/
type GetRolesOK struct {
	Payload *GetRolesOKBody
}

// IsSuccess returns true when this get roles o k response has a 2xx status code
func (o *GetRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get roles o k response has a 3xx status code
func (o *GetRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles o k response has a 4xx status code
func (o *GetRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get roles o k response has a 5xx status code
func (o *GetRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles o k response a status code equal to that given
func (o *GetRolesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get roles o k response
func (o *GetRolesOK) Code() int {
	return 200
}

func (o *GetRolesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/roles][%d] getRolesOK %s", 200, payload)
}

func (o *GetRolesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/roles][%d] getRolesOK %s", 200, payload)
}

func (o *GetRolesOK) GetPayload() *GetRolesOKBody {
	return o.Payload
}

func (o *GetRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetRolesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesForbidden creates a GetRolesForbidden with default headers values
func NewGetRolesForbidden() *GetRolesForbidden {
	return &GetRolesForbidden{}
}

/*
GetRolesForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetRolesForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get roles forbidden response has a 2xx status code
func (o *GetRolesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles forbidden response has a 3xx status code
func (o *GetRolesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles forbidden response has a 4xx status code
func (o *GetRolesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles forbidden response has a 5xx status code
func (o *GetRolesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles forbidden response a status code equal to that given
func (o *GetRolesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get roles forbidden response
func (o *GetRolesForbidden) Code() int {
	return 403
}

func (o *GetRolesForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/roles][%d] getRolesForbidden %s", 403, payload)
}

func (o *GetRolesForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/roles][%d] getRolesForbidden %s", 403, payload)
}

func (o *GetRolesForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRolesNotFound creates a GetRolesNotFound with default headers values
func NewGetRolesNotFound() *GetRolesNotFound {
	return &GetRolesNotFound{}
}

/*
GetRolesNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetRolesNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get roles not found response has a 2xx status code
func (o *GetRolesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles not found response has a 3xx status code
func (o *GetRolesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles not found response has a 4xx status code
func (o *GetRolesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles not found response has a 5xx status code
func (o *GetRolesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles not found response a status code equal to that given
func (o *GetRolesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get roles not found response
func (o *GetRolesNotFound) Code() int {
	return 404
}

func (o *GetRolesNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/roles][%d] getRolesNotFound %s", 404, payload)
}

func (o *GetRolesNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/roles][%d] getRolesNotFound %s", 404, payload)
}

func (o *GetRolesNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRolesDefault creates a GetRolesDefault with default headers values
func NewGetRolesDefault(code int) *GetRolesDefault {
	return &GetRolesDefault{
		_statusCode: code,
	}
}

/*
GetRolesDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetRolesDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get roles default response has a 2xx status code
func (o *GetRolesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get roles default response has a 3xx status code
func (o *GetRolesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get roles default response has a 4xx status code
func (o *GetRolesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get roles default response has a 5xx status code
func (o *GetRolesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get roles default response a status code equal to that given
func (o *GetRolesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get roles default response
func (o *GetRolesDefault) Code() int {
	return o._statusCode
}

func (o *GetRolesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/roles][%d] getRoles default %s", o._statusCode, payload)
}

func (o *GetRolesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/roles][%d] getRoles default %s", o._statusCode, payload)
}

func (o *GetRolesDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetRolesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetRolesOKBody get roles o k body
swagger:model GetRolesOKBody
*/
type GetRolesOKBody struct {

	// data
	// Required: true
	Data []*models.Role `json:"data"`

	// pagination
	// Required: true
	Pagination *models.Pagination `json:"pagination"`
}

// Validate validates this get roles o k body
func (o *GetRolesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRolesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getRolesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getRolesOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getRolesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetRolesOKBody) validatePagination(formats strfmt.Registry) error {

	if err := validate.Required("getRolesOK"+"."+"pagination", "body", o.Pagination); err != nil {
		return err
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getRolesOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getRolesOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get roles o k body based on the context it is used
func (o *GetRolesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRolesOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getRolesOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getRolesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetRolesOKBody) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if o.Pagination != nil {

		if err := o.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getRolesOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getRolesOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRolesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRolesOKBody) UnmarshalBinary(b []byte) error {
	var res GetRolesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
