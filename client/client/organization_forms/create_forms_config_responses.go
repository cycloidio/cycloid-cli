// Code generated by go-swagger; DO NOT EDIT.

package organization_forms

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

// CreateFormsConfigReader is a Reader for the CreateFormsConfig structure.
type CreateFormsConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFormsConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateFormsConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateFormsConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateFormsConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateFormsConfigUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateFormsConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateFormsConfigOK creates a CreateFormsConfigOK with default headers values
func NewCreateFormsConfigOK() *CreateFormsConfigOK {
	return &CreateFormsConfigOK{}
}

/*CreateFormsConfigOK handles this case with default header values.

Set of config to create the project / push onto repositories
*/
type CreateFormsConfigOK struct {
	Payload *CreateFormsConfigOKBody
}

func (o *CreateFormsConfigOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/forms/config][%d] createFormsConfigOK  %+v", 200, o.Payload)
}

func (o *CreateFormsConfigOK) GetPayload() *CreateFormsConfigOKBody {
	return o.Payload
}

func (o *CreateFormsConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateFormsConfigOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFormsConfigForbidden creates a CreateFormsConfigForbidden with default headers values
func NewCreateFormsConfigForbidden() *CreateFormsConfigForbidden {
	return &CreateFormsConfigForbidden{}
}

/*CreateFormsConfigForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type CreateFormsConfigForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *CreateFormsConfigForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/forms/config][%d] createFormsConfigForbidden  %+v", 403, o.Payload)
}

func (o *CreateFormsConfigForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateFormsConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFormsConfigNotFound creates a CreateFormsConfigNotFound with default headers values
func NewCreateFormsConfigNotFound() *CreateFormsConfigNotFound {
	return &CreateFormsConfigNotFound{}
}

/*CreateFormsConfigNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type CreateFormsConfigNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *CreateFormsConfigNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/forms/config][%d] createFormsConfigNotFound  %+v", 404, o.Payload)
}

func (o *CreateFormsConfigNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateFormsConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFormsConfigUnprocessableEntity creates a CreateFormsConfigUnprocessableEntity with default headers values
func NewCreateFormsConfigUnprocessableEntity() *CreateFormsConfigUnprocessableEntity {
	return &CreateFormsConfigUnprocessableEntity{}
}

/*CreateFormsConfigUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type CreateFormsConfigUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *CreateFormsConfigUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/forms/config][%d] createFormsConfigUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateFormsConfigUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateFormsConfigUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFormsConfigDefault creates a CreateFormsConfigDefault with default headers values
func NewCreateFormsConfigDefault(code int) *CreateFormsConfigDefault {
	return &CreateFormsConfigDefault{
		_statusCode: code,
	}
}

/*CreateFormsConfigDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreateFormsConfigDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the create forms config default response
func (o *CreateFormsConfigDefault) Code() int {
	return o._statusCode
}

func (o *CreateFormsConfigDefault) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/forms/config][%d] createFormsConfig default  %+v", o._statusCode, o.Payload)
}

func (o *CreateFormsConfigDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateFormsConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*CreateFormsConfigOKBody create forms config o k body
swagger:model CreateFormsConfigOKBody
*/
type CreateFormsConfigOKBody struct {

	// data
	// Required: true
	Data interface{} `json:"data"`
}

// Validate validates this create forms config o k body
func (o *CreateFormsConfigOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateFormsConfigOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("createFormsConfigOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateFormsConfigOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateFormsConfigOKBody) UnmarshalBinary(b []byte) error {
	var res CreateFormsConfigOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
