// Code generated by go-swagger; DO NOT EDIT.

package organization_service_catalog_sources

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

// RefreshServiceCatalogSourceReader is a Reader for the RefreshServiceCatalogSource structure.
type RefreshServiceCatalogSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshServiceCatalogSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRefreshServiceCatalogSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRefreshServiceCatalogSourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 411:
		result := NewRefreshServiceCatalogSourceLengthRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRefreshServiceCatalogSourceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRefreshServiceCatalogSourceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRefreshServiceCatalogSourceOK creates a RefreshServiceCatalogSourceOK with default headers values
func NewRefreshServiceCatalogSourceOK() *RefreshServiceCatalogSourceOK {
	return &RefreshServiceCatalogSourceOK{}
}

/*RefreshServiceCatalogSourceOK handles this case with default header values.

Success refresh
*/
type RefreshServiceCatalogSourceOK struct {
	Payload *RefreshServiceCatalogSourceOKBody
}

func (o *RefreshServiceCatalogSourceOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_canonical}/refresh][%d] refreshServiceCatalogSourceOK  %+v", 200, o.Payload)
}

func (o *RefreshServiceCatalogSourceOK) GetPayload() *RefreshServiceCatalogSourceOKBody {
	return o.Payload
}

func (o *RefreshServiceCatalogSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RefreshServiceCatalogSourceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshServiceCatalogSourceNotFound creates a RefreshServiceCatalogSourceNotFound with default headers values
func NewRefreshServiceCatalogSourceNotFound() *RefreshServiceCatalogSourceNotFound {
	return &RefreshServiceCatalogSourceNotFound{}
}

/*RefreshServiceCatalogSourceNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type RefreshServiceCatalogSourceNotFound struct {
	Payload *models.ErrorPayload
}

func (o *RefreshServiceCatalogSourceNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_canonical}/refresh][%d] refreshServiceCatalogSourceNotFound  %+v", 404, o.Payload)
}

func (o *RefreshServiceCatalogSourceNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RefreshServiceCatalogSourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshServiceCatalogSourceLengthRequired creates a RefreshServiceCatalogSourceLengthRequired with default headers values
func NewRefreshServiceCatalogSourceLengthRequired() *RefreshServiceCatalogSourceLengthRequired {
	return &RefreshServiceCatalogSourceLengthRequired{}
}

/*RefreshServiceCatalogSourceLengthRequired handles this case with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type RefreshServiceCatalogSourceLengthRequired struct {
}

func (o *RefreshServiceCatalogSourceLengthRequired) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_canonical}/refresh][%d] refreshServiceCatalogSourceLengthRequired ", 411)
}

func (o *RefreshServiceCatalogSourceLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRefreshServiceCatalogSourceUnprocessableEntity creates a RefreshServiceCatalogSourceUnprocessableEntity with default headers values
func NewRefreshServiceCatalogSourceUnprocessableEntity() *RefreshServiceCatalogSourceUnprocessableEntity {
	return &RefreshServiceCatalogSourceUnprocessableEntity{}
}

/*RefreshServiceCatalogSourceUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type RefreshServiceCatalogSourceUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *RefreshServiceCatalogSourceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_canonical}/refresh][%d] refreshServiceCatalogSourceUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *RefreshServiceCatalogSourceUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RefreshServiceCatalogSourceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshServiceCatalogSourceDefault creates a RefreshServiceCatalogSourceDefault with default headers values
func NewRefreshServiceCatalogSourceDefault(code int) *RefreshServiceCatalogSourceDefault {
	return &RefreshServiceCatalogSourceDefault{
		_statusCode: code,
	}
}

/*RefreshServiceCatalogSourceDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type RefreshServiceCatalogSourceDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the refresh service catalog source default response
func (o *RefreshServiceCatalogSourceDefault) Code() int {
	return o._statusCode
}

func (o *RefreshServiceCatalogSourceDefault) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_canonical}/refresh][%d] refreshServiceCatalogSource default  %+v", o._statusCode, o.Payload)
}

func (o *RefreshServiceCatalogSourceDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RefreshServiceCatalogSourceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RefreshServiceCatalogSourceOKBody refresh service catalog source o k body
swagger:model RefreshServiceCatalogSourceOKBody
*/
type RefreshServiceCatalogSourceOKBody struct {

	// data
	// Required: true
	Data *models.ServiceCatalogSource `json:"data"`
}

// Validate validates this refresh service catalog source o k body
func (o *RefreshServiceCatalogSourceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RefreshServiceCatalogSourceOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("refreshServiceCatalogSourceOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("refreshServiceCatalogSourceOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RefreshServiceCatalogSourceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RefreshServiceCatalogSourceOKBody) UnmarshalBinary(b []byte) error {
	var res RefreshServiceCatalogSourceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
