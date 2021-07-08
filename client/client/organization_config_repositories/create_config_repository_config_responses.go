// Code generated by go-swagger; DO NOT EDIT.

package organization_config_repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/cycloid-cli/client/models"
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

/*CreateConfigRepositoryConfigNoContent handles this case with default header values.

SC config files have been created successfully
*/
type CreateConfigRepositoryConfigNoContent struct {
}

func (o *CreateConfigRepositoryConfigNoContent) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfigNoContent ", 204)
}

func (o *CreateConfigRepositoryConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateConfigRepositoryConfigForbidden creates a CreateConfigRepositoryConfigForbidden with default headers values
func NewCreateConfigRepositoryConfigForbidden() *CreateConfigRepositoryConfigForbidden {
	return &CreateConfigRepositoryConfigForbidden{}
}

/*CreateConfigRepositoryConfigForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type CreateConfigRepositoryConfigForbidden struct {
	Payload *models.ErrorPayload
}

func (o *CreateConfigRepositoryConfigForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfigForbidden  %+v", 403, o.Payload)
}

func (o *CreateConfigRepositoryConfigForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateConfigRepositoryConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*CreateConfigRepositoryConfigNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type CreateConfigRepositoryConfigNotFound struct {
	Payload *models.ErrorPayload
}

func (o *CreateConfigRepositoryConfigNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfigNotFound  %+v", 404, o.Payload)
}

func (o *CreateConfigRepositoryConfigNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateConfigRepositoryConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*CreateConfigRepositoryConfigLengthRequired handles this case with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type CreateConfigRepositoryConfigLengthRequired struct {
}

func (o *CreateConfigRepositoryConfigLengthRequired) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfigLengthRequired ", 411)
}

func (o *CreateConfigRepositoryConfigLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateConfigRepositoryConfigUnprocessableEntity creates a CreateConfigRepositoryConfigUnprocessableEntity with default headers values
func NewCreateConfigRepositoryConfigUnprocessableEntity() *CreateConfigRepositoryConfigUnprocessableEntity {
	return &CreateConfigRepositoryConfigUnprocessableEntity{}
}

/*CreateConfigRepositoryConfigUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type CreateConfigRepositoryConfigUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *CreateConfigRepositoryConfigUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfigUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateConfigRepositoryConfigUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateConfigRepositoryConfigUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*CreateConfigRepositoryConfigDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreateConfigRepositoryConfigDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the create config repository config default response
func (o *CreateConfigRepositoryConfigDefault) Code() int {
	return o._statusCode
}

func (o *CreateConfigRepositoryConfigDefault) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories/{config_repository_canonical}/config][%d] createConfigRepositoryConfig default  %+v", o._statusCode, o.Payload)
}

func (o *CreateConfigRepositoryConfigDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateConfigRepositoryConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
