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

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// GetServiceCatalogSourceReader is a Reader for the GetServiceCatalogSource structure.
type GetServiceCatalogSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceCatalogSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceCatalogSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetServiceCatalogSourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetServiceCatalogSourceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServiceCatalogSourceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceCatalogSourceOK creates a GetServiceCatalogSourceOK with default headers values
func NewGetServiceCatalogSourceOK() *GetServiceCatalogSourceOK {
	return &GetServiceCatalogSourceOK{}
}

/*GetServiceCatalogSourceOK handles this case with default header values.

Organization Service Catalog Sources.
*/
type GetServiceCatalogSourceOK struct {
	Payload *GetServiceCatalogSourceOKBody
}

func (o *GetServiceCatalogSourceOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_id}][%d] getServiceCatalogSourceOK  %+v", 200, o.Payload)
}

func (o *GetServiceCatalogSourceOK) GetPayload() *GetServiceCatalogSourceOKBody {
	return o.Payload
}

func (o *GetServiceCatalogSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetServiceCatalogSourceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceCatalogSourceForbidden creates a GetServiceCatalogSourceForbidden with default headers values
func NewGetServiceCatalogSourceForbidden() *GetServiceCatalogSourceForbidden {
	return &GetServiceCatalogSourceForbidden{}
}

/*GetServiceCatalogSourceForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetServiceCatalogSourceForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetServiceCatalogSourceForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_id}][%d] getServiceCatalogSourceForbidden  %+v", 403, o.Payload)
}

func (o *GetServiceCatalogSourceForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogSourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceCatalogSourceUnprocessableEntity creates a GetServiceCatalogSourceUnprocessableEntity with default headers values
func NewGetServiceCatalogSourceUnprocessableEntity() *GetServiceCatalogSourceUnprocessableEntity {
	return &GetServiceCatalogSourceUnprocessableEntity{}
}

/*GetServiceCatalogSourceUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetServiceCatalogSourceUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *GetServiceCatalogSourceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_id}][%d] getServiceCatalogSourceUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetServiceCatalogSourceUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogSourceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceCatalogSourceDefault creates a GetServiceCatalogSourceDefault with default headers values
func NewGetServiceCatalogSourceDefault(code int) *GetServiceCatalogSourceDefault {
	return &GetServiceCatalogSourceDefault{
		_statusCode: code,
	}
}

/*GetServiceCatalogSourceDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetServiceCatalogSourceDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get service catalog source default response
func (o *GetServiceCatalogSourceDefault) Code() int {
	return o._statusCode
}

func (o *GetServiceCatalogSourceDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_id}][%d] getServiceCatalogSource default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceCatalogSourceDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogSourceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetServiceCatalogSourceOKBody get service catalog source o k body
swagger:model GetServiceCatalogSourceOKBody
*/
type GetServiceCatalogSourceOKBody struct {

	// data
	// Required: true
	Data *models.ServiceCatalogSource `json:"data"`
}

// Validate validates this get service catalog source o k body
func (o *GetServiceCatalogSourceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceCatalogSourceOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getServiceCatalogSourceOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceCatalogSourceOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceCatalogSourceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceCatalogSourceOKBody) UnmarshalBinary(b []byte) error {
	var res GetServiceCatalogSourceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}