// Code generated by go-swagger; DO NOT EDIT.

package stacks

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

// GetStackTerraformImageReader is a Reader for the GetStackTerraformImage structure.
type GetStackTerraformImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStackTerraformImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStackTerraformImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetStackTerraformImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStackTerraformImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStackTerraformImageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStackTerraformImageOK creates a GetStackTerraformImageOK with default headers values
func NewGetStackTerraformImageOK() *GetStackTerraformImageOK {
	return &GetStackTerraformImageOK{}
}

/*GetStackTerraformImageOK handles this case with default header values.

The Stack TF Image
*/
type GetStackTerraformImageOK struct {
	Payload *GetStackTerraformImageOKBody
}

func (o *GetStackTerraformImageOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/stacks/{stack_ref}/terraform/diagram/image][%d] getStackTerraformImageOK  %+v", 200, o.Payload)
}

func (o *GetStackTerraformImageOK) GetPayload() *GetStackTerraformImageOKBody {
	return o.Payload
}

func (o *GetStackTerraformImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetStackTerraformImageOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStackTerraformImageForbidden creates a GetStackTerraformImageForbidden with default headers values
func NewGetStackTerraformImageForbidden() *GetStackTerraformImageForbidden {
	return &GetStackTerraformImageForbidden{}
}

/*GetStackTerraformImageForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetStackTerraformImageForbidden struct {
	Payload *models.ErrorPayload
}

func (o *GetStackTerraformImageForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/stacks/{stack_ref}/terraform/diagram/image][%d] getStackTerraformImageForbidden  %+v", 403, o.Payload)
}

func (o *GetStackTerraformImageForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetStackTerraformImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStackTerraformImageNotFound creates a GetStackTerraformImageNotFound with default headers values
func NewGetStackTerraformImageNotFound() *GetStackTerraformImageNotFound {
	return &GetStackTerraformImageNotFound{}
}

/*GetStackTerraformImageNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetStackTerraformImageNotFound struct {
	Payload *models.ErrorPayload
}

func (o *GetStackTerraformImageNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/stacks/{stack_ref}/terraform/diagram/image][%d] getStackTerraformImageNotFound  %+v", 404, o.Payload)
}

func (o *GetStackTerraformImageNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetStackTerraformImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStackTerraformImageDefault creates a GetStackTerraformImageDefault with default headers values
func NewGetStackTerraformImageDefault(code int) *GetStackTerraformImageDefault {
	return &GetStackTerraformImageDefault{
		_statusCode: code,
	}
}

/*GetStackTerraformImageDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetStackTerraformImageDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get stack terraform image default response
func (o *GetStackTerraformImageDefault) Code() int {
	return o._statusCode
}

func (o *GetStackTerraformImageDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/stacks/{stack_ref}/terraform/diagram/image][%d] getStackTerraformImage default  %+v", o._statusCode, o.Payload)
}

func (o *GetStackTerraformImageDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetStackTerraformImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetStackTerraformImageOKBody get stack terraform image o k body
swagger:model GetStackTerraformImageOKBody
*/
type GetStackTerraformImageOKBody struct {

	// data
	// Required: true
	Data *models.TerraformImage `json:"data"`
}

// Validate validates this get stack terraform image o k body
func (o *GetStackTerraformImageOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStackTerraformImageOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getStackTerraformImageOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStackTerraformImageOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStackTerraformImageOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStackTerraformImageOKBody) UnmarshalBinary(b []byte) error {
	var res GetStackTerraformImageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
