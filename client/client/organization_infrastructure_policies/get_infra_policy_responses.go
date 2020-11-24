// Code generated by go-swagger; DO NOT EDIT.

package organization_infrastructure_policies

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

// GetInfraPolicyReader is a Reader for the GetInfraPolicy structure.
type GetInfraPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInfraPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInfraPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetInfraPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInfraPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetInfraPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetInfraPolicyOK creates a GetInfraPolicyOK with default headers values
func NewGetInfraPolicyOK() *GetInfraPolicyOK {
	return &GetInfraPolicyOK{}
}

/*GetInfraPolicyOK handles this case with default header values.

The information of the InfraPolicy which has the specified canonical.
*/
type GetInfraPolicyOK struct {
	Payload *GetInfraPolicyOKBody
}

func (o *GetInfraPolicyOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] getInfraPolicyOK  %+v", 200, o.Payload)
}

func (o *GetInfraPolicyOK) GetPayload() *GetInfraPolicyOKBody {
	return o.Payload
}

func (o *GetInfraPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetInfraPolicyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInfraPolicyForbidden creates a GetInfraPolicyForbidden with default headers values
func NewGetInfraPolicyForbidden() *GetInfraPolicyForbidden {
	return &GetInfraPolicyForbidden{}
}

/*GetInfraPolicyForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetInfraPolicyForbidden struct {
	Payload *models.ErrorPayload
}

func (o *GetInfraPolicyForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] getInfraPolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetInfraPolicyForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetInfraPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInfraPolicyNotFound creates a GetInfraPolicyNotFound with default headers values
func NewGetInfraPolicyNotFound() *GetInfraPolicyNotFound {
	return &GetInfraPolicyNotFound{}
}

/*GetInfraPolicyNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetInfraPolicyNotFound struct {
	Payload *models.ErrorPayload
}

func (o *GetInfraPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] getInfraPolicyNotFound  %+v", 404, o.Payload)
}

func (o *GetInfraPolicyNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetInfraPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInfraPolicyDefault creates a GetInfraPolicyDefault with default headers values
func NewGetInfraPolicyDefault(code int) *GetInfraPolicyDefault {
	return &GetInfraPolicyDefault{
		_statusCode: code,
	}
}

/*GetInfraPolicyDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetInfraPolicyDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get infra policy default response
func (o *GetInfraPolicyDefault) Code() int {
	return o._statusCode
}

func (o *GetInfraPolicyDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] getInfraPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *GetInfraPolicyDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetInfraPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetInfraPolicyOKBody get infra policy o k body
swagger:model GetInfraPolicyOKBody
*/
type GetInfraPolicyOKBody struct {

	// data
	// Required: true
	Data *models.InfraPolicy `json:"data"`
}

// Validate validates this get infra policy o k body
func (o *GetInfraPolicyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInfraPolicyOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getInfraPolicyOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getInfraPolicyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetInfraPolicyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInfraPolicyOKBody) UnmarshalBinary(b []byte) error {
	var res GetInfraPolicyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}