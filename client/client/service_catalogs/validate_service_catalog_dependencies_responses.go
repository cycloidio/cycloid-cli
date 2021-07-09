// Code generated by go-swagger; DO NOT EDIT.

package service_catalogs

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

// ValidateServiceCatalogDependenciesReader is a Reader for the ValidateServiceCatalogDependencies structure.
type ValidateServiceCatalogDependenciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateServiceCatalogDependenciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateServiceCatalogDependenciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewValidateServiceCatalogDependenciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewValidateServiceCatalogDependenciesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewValidateServiceCatalogDependenciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewValidateServiceCatalogDependenciesOK creates a ValidateServiceCatalogDependenciesOK with default headers values
func NewValidateServiceCatalogDependenciesOK() *ValidateServiceCatalogDependenciesOK {
	return &ValidateServiceCatalogDependenciesOK{}
}

/*ValidateServiceCatalogDependenciesOK handles this case with default header values.

The result of the service catalog's dependencies validation
*/
type ValidateServiceCatalogDependenciesOK struct {
	Payload *ValidateServiceCatalogDependenciesOKBody
}

func (o *ValidateServiceCatalogDependenciesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/dependencies/validate][%d] validateServiceCatalogDependenciesOK  %+v", 200, o.Payload)
}

func (o *ValidateServiceCatalogDependenciesOK) GetPayload() *ValidateServiceCatalogDependenciesOKBody {
	return o.Payload
}

func (o *ValidateServiceCatalogDependenciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ValidateServiceCatalogDependenciesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateServiceCatalogDependenciesForbidden creates a ValidateServiceCatalogDependenciesForbidden with default headers values
func NewValidateServiceCatalogDependenciesForbidden() *ValidateServiceCatalogDependenciesForbidden {
	return &ValidateServiceCatalogDependenciesForbidden{}
}

/*ValidateServiceCatalogDependenciesForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type ValidateServiceCatalogDependenciesForbidden struct {
	Payload *models.ErrorPayload
}

func (o *ValidateServiceCatalogDependenciesForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/dependencies/validate][%d] validateServiceCatalogDependenciesForbidden  %+v", 403, o.Payload)
}

func (o *ValidateServiceCatalogDependenciesForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ValidateServiceCatalogDependenciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateServiceCatalogDependenciesUnprocessableEntity creates a ValidateServiceCatalogDependenciesUnprocessableEntity with default headers values
func NewValidateServiceCatalogDependenciesUnprocessableEntity() *ValidateServiceCatalogDependenciesUnprocessableEntity {
	return &ValidateServiceCatalogDependenciesUnprocessableEntity{}
}

/*ValidateServiceCatalogDependenciesUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type ValidateServiceCatalogDependenciesUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *ValidateServiceCatalogDependenciesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/dependencies/validate][%d] validateServiceCatalogDependenciesUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ValidateServiceCatalogDependenciesUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ValidateServiceCatalogDependenciesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateServiceCatalogDependenciesDefault creates a ValidateServiceCatalogDependenciesDefault with default headers values
func NewValidateServiceCatalogDependenciesDefault(code int) *ValidateServiceCatalogDependenciesDefault {
	return &ValidateServiceCatalogDependenciesDefault{
		_statusCode: code,
	}
}

/*ValidateServiceCatalogDependenciesDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type ValidateServiceCatalogDependenciesDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the validate service catalog dependencies default response
func (o *ValidateServiceCatalogDependenciesDefault) Code() int {
	return o._statusCode
}

func (o *ValidateServiceCatalogDependenciesDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/dependencies/validate][%d] validateServiceCatalogDependencies default  %+v", o._statusCode, o.Payload)
}

func (o *ValidateServiceCatalogDependenciesDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ValidateServiceCatalogDependenciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ValidateServiceCatalogDependenciesOKBody validate service catalog dependencies o k body
swagger:model ValidateServiceCatalogDependenciesOKBody
*/
type ValidateServiceCatalogDependenciesOKBody struct {

	// data
	// Required: true
	Data *models.ServiceCatalogDependenciesValidationResult `json:"data"`
}

// Validate validates this validate service catalog dependencies o k body
func (o *ValidateServiceCatalogDependenciesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ValidateServiceCatalogDependenciesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("validateServiceCatalogDependenciesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validateServiceCatalogDependenciesOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ValidateServiceCatalogDependenciesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidateServiceCatalogDependenciesOKBody) UnmarshalBinary(b []byte) error {
	var res ValidateServiceCatalogDependenciesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
