// Code generated by go-swagger; DO NOT EDIT.

package organization_team_members

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

// GetTeamMemberReader is a Reader for the GetTeamMember structure.
type GetTeamMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetTeamMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTeamMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetTeamMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTeamMemberOK creates a GetTeamMemberOK with default headers values
func NewGetTeamMemberOK() *GetTeamMemberOK {
	return &GetTeamMemberOK{}
}

/*
GetTeamMemberOK describes a response with status code 200, with default header values.

The information of the member of the team.
*/
type GetTeamMemberOK struct {
	Payload *GetTeamMemberOKBody
}

// IsSuccess returns true when this get team member o k response has a 2xx status code
func (o *GetTeamMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team member o k response has a 3xx status code
func (o *GetTeamMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team member o k response has a 4xx status code
func (o *GetTeamMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team member o k response has a 5xx status code
func (o *GetTeamMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get team member o k response a status code equal to that given
func (o *GetTeamMemberOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get team member o k response
func (o *GetTeamMemberOK) Code() int {
	return 200
}

func (o *GetTeamMemberOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/teams/{team_canonical}/members/{member_id}][%d] getTeamMemberOK %s", 200, payload)
}

func (o *GetTeamMemberOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/teams/{team_canonical}/members/{member_id}][%d] getTeamMemberOK %s", 200, payload)
}

func (o *GetTeamMemberOK) GetPayload() *GetTeamMemberOKBody {
	return o.Payload
}

func (o *GetTeamMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetTeamMemberOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMemberForbidden creates a GetTeamMemberForbidden with default headers values
func NewGetTeamMemberForbidden() *GetTeamMemberForbidden {
	return &GetTeamMemberForbidden{}
}

/*
GetTeamMemberForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetTeamMemberForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get team member forbidden response has a 2xx status code
func (o *GetTeamMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team member forbidden response has a 3xx status code
func (o *GetTeamMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team member forbidden response has a 4xx status code
func (o *GetTeamMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team member forbidden response has a 5xx status code
func (o *GetTeamMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get team member forbidden response a status code equal to that given
func (o *GetTeamMemberForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get team member forbidden response
func (o *GetTeamMemberForbidden) Code() int {
	return 403
}

func (o *GetTeamMemberForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/teams/{team_canonical}/members/{member_id}][%d] getTeamMemberForbidden %s", 403, payload)
}

func (o *GetTeamMemberForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/teams/{team_canonical}/members/{member_id}][%d] getTeamMemberForbidden %s", 403, payload)
}

func (o *GetTeamMemberForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetTeamMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetTeamMemberNotFound creates a GetTeamMemberNotFound with default headers values
func NewGetTeamMemberNotFound() *GetTeamMemberNotFound {
	return &GetTeamMemberNotFound{}
}

/*
GetTeamMemberNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetTeamMemberNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get team member not found response has a 2xx status code
func (o *GetTeamMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team member not found response has a 3xx status code
func (o *GetTeamMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team member not found response has a 4xx status code
func (o *GetTeamMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team member not found response has a 5xx status code
func (o *GetTeamMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get team member not found response a status code equal to that given
func (o *GetTeamMemberNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get team member not found response
func (o *GetTeamMemberNotFound) Code() int {
	return 404
}

func (o *GetTeamMemberNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/teams/{team_canonical}/members/{member_id}][%d] getTeamMemberNotFound %s", 404, payload)
}

func (o *GetTeamMemberNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/teams/{team_canonical}/members/{member_id}][%d] getTeamMemberNotFound %s", 404, payload)
}

func (o *GetTeamMemberNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetTeamMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetTeamMemberDefault creates a GetTeamMemberDefault with default headers values
func NewGetTeamMemberDefault(code int) *GetTeamMemberDefault {
	return &GetTeamMemberDefault{
		_statusCode: code,
	}
}

/*
GetTeamMemberDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetTeamMemberDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get team member default response has a 2xx status code
func (o *GetTeamMemberDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get team member default response has a 3xx status code
func (o *GetTeamMemberDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get team member default response has a 4xx status code
func (o *GetTeamMemberDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get team member default response has a 5xx status code
func (o *GetTeamMemberDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get team member default response a status code equal to that given
func (o *GetTeamMemberDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get team member default response
func (o *GetTeamMemberDefault) Code() int {
	return o._statusCode
}

func (o *GetTeamMemberDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/teams/{team_canonical}/members/{member_id}][%d] getTeamMember default %s", o._statusCode, payload)
}

func (o *GetTeamMemberDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/teams/{team_canonical}/members/{member_id}][%d] getTeamMember default %s", o._statusCode, payload)
}

func (o *GetTeamMemberDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetTeamMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetTeamMemberOKBody get team member o k body
swagger:model GetTeamMemberOKBody
*/
type GetTeamMemberOKBody struct {

	// data
	// Required: true
	Data *models.MemberTeam `json:"data"`
}

// Validate validates this get team member o k body
func (o *GetTeamMemberOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTeamMemberOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getTeamMemberOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTeamMemberOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getTeamMemberOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get team member o k body based on the context it is used
func (o *GetTeamMemberOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTeamMemberOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTeamMemberOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getTeamMemberOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTeamMemberOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTeamMemberOKBody) UnmarshalBinary(b []byte) error {
	var res GetTeamMemberOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
