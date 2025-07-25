// Code generated by go-swagger; DO NOT EDIT.

package organization_team_members

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

// UnassignMemberFromTeamReader is a Reader for the UnassignMemberFromTeam structure.
type UnassignMemberFromTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnassignMemberFromTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUnassignMemberFromTeamNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUnassignMemberFromTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUnassignMemberFromTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUnassignMemberFromTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnassignMemberFromTeamNoContent creates a UnassignMemberFromTeamNoContent with default headers values
func NewUnassignMemberFromTeamNoContent() *UnassignMemberFromTeamNoContent {
	return &UnassignMemberFromTeamNoContent{}
}

/*
UnassignMemberFromTeamNoContent describes a response with status code 204, with default header values.

The user has been unassigned of the team.
*/
type UnassignMemberFromTeamNoContent struct {
}

// IsSuccess returns true when this unassign member from team no content response has a 2xx status code
func (o *UnassignMemberFromTeamNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unassign member from team no content response has a 3xx status code
func (o *UnassignMemberFromTeamNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unassign member from team no content response has a 4xx status code
func (o *UnassignMemberFromTeamNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this unassign member from team no content response has a 5xx status code
func (o *UnassignMemberFromTeamNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this unassign member from team no content response a status code equal to that given
func (o *UnassignMemberFromTeamNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the unassign member from team no content response
func (o *UnassignMemberFromTeamNoContent) Code() int {
	return 204
}

func (o *UnassignMemberFromTeamNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/teams/{team_canonical}/members/{member_id}][%d] unassignMemberFromTeamNoContent", 204)
}

func (o *UnassignMemberFromTeamNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/teams/{team_canonical}/members/{member_id}][%d] unassignMemberFromTeamNoContent", 204)
}

func (o *UnassignMemberFromTeamNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnassignMemberFromTeamForbidden creates a UnassignMemberFromTeamForbidden with default headers values
func NewUnassignMemberFromTeamForbidden() *UnassignMemberFromTeamForbidden {
	return &UnassignMemberFromTeamForbidden{}
}

/*
UnassignMemberFromTeamForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type UnassignMemberFromTeamForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this unassign member from team forbidden response has a 2xx status code
func (o *UnassignMemberFromTeamForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unassign member from team forbidden response has a 3xx status code
func (o *UnassignMemberFromTeamForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unassign member from team forbidden response has a 4xx status code
func (o *UnassignMemberFromTeamForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this unassign member from team forbidden response has a 5xx status code
func (o *UnassignMemberFromTeamForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this unassign member from team forbidden response a status code equal to that given
func (o *UnassignMemberFromTeamForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the unassign member from team forbidden response
func (o *UnassignMemberFromTeamForbidden) Code() int {
	return 403
}

func (o *UnassignMemberFromTeamForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/teams/{team_canonical}/members/{member_id}][%d] unassignMemberFromTeamForbidden %s", 403, payload)
}

func (o *UnassignMemberFromTeamForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/teams/{team_canonical}/members/{member_id}][%d] unassignMemberFromTeamForbidden %s", 403, payload)
}

func (o *UnassignMemberFromTeamForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UnassignMemberFromTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUnassignMemberFromTeamNotFound creates a UnassignMemberFromTeamNotFound with default headers values
func NewUnassignMemberFromTeamNotFound() *UnassignMemberFromTeamNotFound {
	return &UnassignMemberFromTeamNotFound{}
}

/*
UnassignMemberFromTeamNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type UnassignMemberFromTeamNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this unassign member from team not found response has a 2xx status code
func (o *UnassignMemberFromTeamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unassign member from team not found response has a 3xx status code
func (o *UnassignMemberFromTeamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unassign member from team not found response has a 4xx status code
func (o *UnassignMemberFromTeamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this unassign member from team not found response has a 5xx status code
func (o *UnassignMemberFromTeamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this unassign member from team not found response a status code equal to that given
func (o *UnassignMemberFromTeamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the unassign member from team not found response
func (o *UnassignMemberFromTeamNotFound) Code() int {
	return 404
}

func (o *UnassignMemberFromTeamNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/teams/{team_canonical}/members/{member_id}][%d] unassignMemberFromTeamNotFound %s", 404, payload)
}

func (o *UnassignMemberFromTeamNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/teams/{team_canonical}/members/{member_id}][%d] unassignMemberFromTeamNotFound %s", 404, payload)
}

func (o *UnassignMemberFromTeamNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UnassignMemberFromTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUnassignMemberFromTeamDefault creates a UnassignMemberFromTeamDefault with default headers values
func NewUnassignMemberFromTeamDefault(code int) *UnassignMemberFromTeamDefault {
	return &UnassignMemberFromTeamDefault{
		_statusCode: code,
	}
}

/*
UnassignMemberFromTeamDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UnassignMemberFromTeamDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this unassign member from team default response has a 2xx status code
func (o *UnassignMemberFromTeamDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unassign member from team default response has a 3xx status code
func (o *UnassignMemberFromTeamDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unassign member from team default response has a 4xx status code
func (o *UnassignMemberFromTeamDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unassign member from team default response has a 5xx status code
func (o *UnassignMemberFromTeamDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unassign member from team default response a status code equal to that given
func (o *UnassignMemberFromTeamDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the unassign member from team default response
func (o *UnassignMemberFromTeamDefault) Code() int {
	return o._statusCode
}

func (o *UnassignMemberFromTeamDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/teams/{team_canonical}/members/{member_id}][%d] unassignMemberFromTeam default %s", o._statusCode, payload)
}

func (o *UnassignMemberFromTeamDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/teams/{team_canonical}/members/{member_id}][%d] unassignMemberFromTeam default %s", o._statusCode, payload)
}

func (o *UnassignMemberFromTeamDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UnassignMemberFromTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
