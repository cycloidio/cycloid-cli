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

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// GetServiceCatalogReader is a Reader for the GetServiceCatalog structure.
type GetServiceCatalogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceCatalogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceCatalogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetServiceCatalogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServiceCatalogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServiceCatalogDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceCatalogOK creates a GetServiceCatalogOK with default headers values
func NewGetServiceCatalogOK() *GetServiceCatalogOK {
	return &GetServiceCatalogOK{}
}

/*GetServiceCatalogOK handles this case with default header values.

The information of the service catalog.
*/
type GetServiceCatalogOK struct {
	Payload *GetServiceCatalogOKBody
}

func (o *GetServiceCatalogOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}][%d] getServiceCatalogOK  %+v", 200, o.Payload)
}

func (o *GetServiceCatalogOK) GetPayload() *GetServiceCatalogOKBody {
	return o.Payload
}

func (o *GetServiceCatalogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetServiceCatalogOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceCatalogForbidden creates a GetServiceCatalogForbidden with default headers values
func NewGetServiceCatalogForbidden() *GetServiceCatalogForbidden {
	return &GetServiceCatalogForbidden{}
}

/*GetServiceCatalogForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetServiceCatalogForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetServiceCatalogForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}][%d] getServiceCatalogForbidden  %+v", 403, o.Payload)
}

func (o *GetServiceCatalogForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServiceCatalogNotFound creates a GetServiceCatalogNotFound with default headers values
func NewGetServiceCatalogNotFound() *GetServiceCatalogNotFound {
	return &GetServiceCatalogNotFound{}
}

/*GetServiceCatalogNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetServiceCatalogNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetServiceCatalogNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}][%d] getServiceCatalogNotFound  %+v", 404, o.Payload)
}

func (o *GetServiceCatalogNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServiceCatalogDefault creates a GetServiceCatalogDefault with default headers values
func NewGetServiceCatalogDefault(code int) *GetServiceCatalogDefault {
	return &GetServiceCatalogDefault{
		_statusCode: code,
	}
}

/*GetServiceCatalogDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetServiceCatalogDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get service catalog default response
func (o *GetServiceCatalogDefault) Code() int {
	return o._statusCode
}

func (o *GetServiceCatalogDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}][%d] getServiceCatalog default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceCatalogDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetServiceCatalogOKBody get service catalog o k body
swagger:model GetServiceCatalogOKBody
*/
type GetServiceCatalogOKBody struct {

	// data
	// Required: true
	Data *models.ServiceCatalog `json:"data"`
}

// Validate validates this get service catalog o k body
func (o *GetServiceCatalogOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceCatalogOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getServiceCatalogOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceCatalogOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceCatalogOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceCatalogOKBody) UnmarshalBinary(b []byte) error {
	var res GetServiceCatalogOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
