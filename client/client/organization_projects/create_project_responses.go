// Code generated by go-swagger; DO NOT EDIT.

package organization_projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// CreateProjectReader is a Reader for the CreateProject structure.
type CreateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCreateProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 411:
		result := NewCreateProjectLengthRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateProjectUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateProjectOK creates a CreateProjectOK with default headers values
func NewCreateProjectOK() *CreateProjectOK {
	return &CreateProjectOK{}
}

/*CreateProjectOK handles this case with default header values.

Project created. The body contains the information of the new project of the organization.
*/
type CreateProjectOK struct {
	Payload *CreateProjectOKBody
}

func (o *CreateProjectOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects][%d] createProjectOK  %+v", 200, o.Payload)
}

func (o *CreateProjectOK) GetPayload() *CreateProjectOKBody {
	return o.Payload
}

func (o *CreateProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateProjectOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectNotFound creates a CreateProjectNotFound with default headers values
func NewCreateProjectNotFound() *CreateProjectNotFound {
	return &CreateProjectNotFound{}
}

/*CreateProjectNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type CreateProjectNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *CreateProjectNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects][%d] createProjectNotFound  %+v", 404, o.Payload)
}

func (o *CreateProjectNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertUint64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "uint64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectLengthRequired creates a CreateProjectLengthRequired with default headers values
func NewCreateProjectLengthRequired() *CreateProjectLengthRequired {
	return &CreateProjectLengthRequired{}
}

/*CreateProjectLengthRequired handles this case with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type CreateProjectLengthRequired struct {
}

func (o *CreateProjectLengthRequired) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects][%d] createProjectLengthRequired ", 411)
}

func (o *CreateProjectLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProjectUnprocessableEntity creates a CreateProjectUnprocessableEntity with default headers values
func NewCreateProjectUnprocessableEntity() *CreateProjectUnprocessableEntity {
	return &CreateProjectUnprocessableEntity{}
}

/*CreateProjectUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type CreateProjectUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *CreateProjectUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects][%d] createProjectUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateProjectUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateProjectUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertUint64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "uint64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectDefault creates a CreateProjectDefault with default headers values
func NewCreateProjectDefault(code int) *CreateProjectDefault {
	return &CreateProjectDefault{
		_statusCode: code,
	}
}

/*CreateProjectDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreateProjectDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the create project default response
func (o *CreateProjectDefault) Code() int {
	return o._statusCode
}

func (o *CreateProjectDefault) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects][%d] createProject default  %+v", o._statusCode, o.Payload)
}

func (o *CreateProjectDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertUint64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "uint64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateProjectOKBody create project o k body
swagger:model CreateProjectOKBody
*/
type CreateProjectOKBody struct {

	// data
	// Required: true
	Data *models.Project `json:"data"`
}

// Validate validates this create project o k body
func (o *CreateProjectOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateProjectOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("createProjectOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createProjectOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateProjectOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateProjectOKBody) UnmarshalBinary(b []byte) error {
	var res CreateProjectOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
