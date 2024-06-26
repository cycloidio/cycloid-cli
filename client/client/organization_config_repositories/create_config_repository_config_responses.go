// Code generated by go-swagger; DO NOT EDIT.

package organization_config_repositories

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

// CreateConfigRepositoryConfigReader is a Reader for the CreateConfigRepositoryConfig structure.
type CreateConfigRepositoryConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateConfigRepositoryConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCreateConfigRepositoryConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateConfigRepositoryConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateConfigRepositoryConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 411:
		result := NewCreateConfigRepositoryConfigLengthRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateConfigRepositoryConfigUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateConfigRepositoryConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateConfigRepositoryConfigNoContent creates a CreateConfigRepositoryConfigNoContent with default headers values
func NewCreateConfigRepositoryConfigNoContent() *CreateConfigRepositoryConfigNoContent {
	return &CreateConfigRepositoryConfigNoContent{}
}

/*
CreateConfigRepositoryConfigNoContent describes a response with status code 204, with default header values.

SC config files have been created successfully
*/
type CreateConfigRepositoryConfigNoContent struct {
}

// IsSuccess returns true when this create config repository config no content response has a 2xx status code
func (o *CreateConfigRepositoryConfigNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create config repository config no content response has a 3xx status code
func (o *CreateConfigRepositoryConfigNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create config repository config no content response has a 4xx status code
func (o *CreateConfigRepositoryConfigNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this create config repository config no content response has a 5xx status code
func (o *CreateConfigRepositoryConfigNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this create config repository config no content response a status code equal to that given
func (o *CreateConfigRepositoryConfigNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the create config repository config no content response
func (o *CreateConfigRepositoryConfigNoContent) Code() int {
	return 204
}

func (o *CreateConfigRepositoryConfigNoContent) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfigNoContent", 204)
}

func (o *CreateConfigRepositoryConfigNoContent) String() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfigNoContent", 204)
}

func (o *CreateConfigRepositoryConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateConfigRepositoryConfigForbidden creates a CreateConfigRepositoryConfigForbidden with default headers values
func NewCreateConfigRepositoryConfigForbidden() *CreateConfigRepositoryConfigForbidden {
	return &CreateConfigRepositoryConfigForbidden{}
}

/*
CreateConfigRepositoryConfigForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type CreateConfigRepositoryConfigForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create config repository config forbidden response has a 2xx status code
func (o *CreateConfigRepositoryConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create config repository config forbidden response has a 3xx status code
func (o *CreateConfigRepositoryConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create config repository config forbidden response has a 4xx status code
func (o *CreateConfigRepositoryConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create config repository config forbidden response has a 5xx status code
func (o *CreateConfigRepositoryConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create config repository config forbidden response a status code equal to that given
func (o *CreateConfigRepositoryConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create config repository config forbidden response
func (o *CreateConfigRepositoryConfigForbidden) Code() int {
	return 403
}

func (o *CreateConfigRepositoryConfigForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfigForbidden %s", 403, payload)
}

func (o *CreateConfigRepositoryConfigForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfigForbidden %s", 403, payload)
}

func (o *CreateConfigRepositoryConfigForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateConfigRepositoryConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateConfigRepositoryConfigNotFound creates a CreateConfigRepositoryConfigNotFound with default headers values
func NewCreateConfigRepositoryConfigNotFound() *CreateConfigRepositoryConfigNotFound {
	return &CreateConfigRepositoryConfigNotFound{}
}

/*
CreateConfigRepositoryConfigNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type CreateConfigRepositoryConfigNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create config repository config not found response has a 2xx status code
func (o *CreateConfigRepositoryConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create config repository config not found response has a 3xx status code
func (o *CreateConfigRepositoryConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create config repository config not found response has a 4xx status code
func (o *CreateConfigRepositoryConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create config repository config not found response has a 5xx status code
func (o *CreateConfigRepositoryConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create config repository config not found response a status code equal to that given
func (o *CreateConfigRepositoryConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create config repository config not found response
func (o *CreateConfigRepositoryConfigNotFound) Code() int {
	return 404
}

func (o *CreateConfigRepositoryConfigNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfigNotFound %s", 404, payload)
}

func (o *CreateConfigRepositoryConfigNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfigNotFound %s", 404, payload)
}

func (o *CreateConfigRepositoryConfigNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateConfigRepositoryConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateConfigRepositoryConfigLengthRequired creates a CreateConfigRepositoryConfigLengthRequired with default headers values
func NewCreateConfigRepositoryConfigLengthRequired() *CreateConfigRepositoryConfigLengthRequired {
	return &CreateConfigRepositoryConfigLengthRequired{}
}

/*
CreateConfigRepositoryConfigLengthRequired describes a response with status code 411, with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type CreateConfigRepositoryConfigLengthRequired struct {
}

// IsSuccess returns true when this create config repository config length required response has a 2xx status code
func (o *CreateConfigRepositoryConfigLengthRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create config repository config length required response has a 3xx status code
func (o *CreateConfigRepositoryConfigLengthRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create config repository config length required response has a 4xx status code
func (o *CreateConfigRepositoryConfigLengthRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this create config repository config length required response has a 5xx status code
func (o *CreateConfigRepositoryConfigLengthRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this create config repository config length required response a status code equal to that given
func (o *CreateConfigRepositoryConfigLengthRequired) IsCode(code int) bool {
	return code == 411
}

// Code gets the status code for the create config repository config length required response
func (o *CreateConfigRepositoryConfigLengthRequired) Code() int {
	return 411
}

func (o *CreateConfigRepositoryConfigLengthRequired) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfigLengthRequired", 411)
}

func (o *CreateConfigRepositoryConfigLengthRequired) String() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfigLengthRequired", 411)
}

func (o *CreateConfigRepositoryConfigLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateConfigRepositoryConfigUnprocessableEntity creates a CreateConfigRepositoryConfigUnprocessableEntity with default headers values
func NewCreateConfigRepositoryConfigUnprocessableEntity() *CreateConfigRepositoryConfigUnprocessableEntity {
	return &CreateConfigRepositoryConfigUnprocessableEntity{}
}

/*
CreateConfigRepositoryConfigUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type CreateConfigRepositoryConfigUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create config repository config unprocessable entity response has a 2xx status code
func (o *CreateConfigRepositoryConfigUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create config repository config unprocessable entity response has a 3xx status code
func (o *CreateConfigRepositoryConfigUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create config repository config unprocessable entity response has a 4xx status code
func (o *CreateConfigRepositoryConfigUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create config repository config unprocessable entity response has a 5xx status code
func (o *CreateConfigRepositoryConfigUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create config repository config unprocessable entity response a status code equal to that given
func (o *CreateConfigRepositoryConfigUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create config repository config unprocessable entity response
func (o *CreateConfigRepositoryConfigUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateConfigRepositoryConfigUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfigUnprocessableEntity %s", 422, payload)
}

func (o *CreateConfigRepositoryConfigUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfigUnprocessableEntity %s", 422, payload)
}

func (o *CreateConfigRepositoryConfigUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateConfigRepositoryConfigUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateConfigRepositoryConfigDefault creates a CreateConfigRepositoryConfigDefault with default headers values
func NewCreateConfigRepositoryConfigDefault(code int) *CreateConfigRepositoryConfigDefault {
	return &CreateConfigRepositoryConfigDefault{
		_statusCode: code,
	}
}

/*
CreateConfigRepositoryConfigDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreateConfigRepositoryConfigDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create config repository config default response has a 2xx status code
func (o *CreateConfigRepositoryConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create config repository config default response has a 3xx status code
func (o *CreateConfigRepositoryConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create config repository config default response has a 4xx status code
func (o *CreateConfigRepositoryConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create config repository config default response has a 5xx status code
func (o *CreateConfigRepositoryConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create config repository config default response a status code equal to that given
func (o *CreateConfigRepositoryConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create config repository config default response
func (o *CreateConfigRepositoryConfigDefault) Code() int {
	return o._statusCode
}

func (o *CreateConfigRepositoryConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfig default %s", o._statusCode, payload)
}

func (o *CreateConfigRepositoryConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfig default %s", o._statusCode, payload)
}

func (o *CreateConfigRepositoryConfigDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateConfigRepositoryConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
