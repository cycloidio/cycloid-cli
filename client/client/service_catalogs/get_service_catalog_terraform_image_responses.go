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

// GetServiceCatalogTerraformImageReader is a Reader for the GetServiceCatalogTerraformImage structure.
type GetServiceCatalogTerraformImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceCatalogTerraformImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceCatalogTerraformImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetServiceCatalogTerraformImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServiceCatalogTerraformImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServiceCatalogTerraformImageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceCatalogTerraformImageOK creates a GetServiceCatalogTerraformImageOK with default headers values
func NewGetServiceCatalogTerraformImageOK() *GetServiceCatalogTerraformImageOK {
	return &GetServiceCatalogTerraformImageOK{}
}

/*GetServiceCatalogTerraformImageOK handles this case with default header values.

The SC TF Image
*/
type GetServiceCatalogTerraformImageOK struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *GetServiceCatalogTerraformImageOKBody
}

func (o *GetServiceCatalogTerraformImageOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram/image][%d] getServiceCatalogTerraformImageOK  %+v", 200, o.Payload)
}

func (o *GetServiceCatalogTerraformImageOK) GetPayload() *GetServiceCatalogTerraformImageOKBody {
	return o.Payload
}

func (o *GetServiceCatalogTerraformImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(GetServiceCatalogTerraformImageOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceCatalogTerraformImageForbidden creates a GetServiceCatalogTerraformImageForbidden with default headers values
func NewGetServiceCatalogTerraformImageForbidden() *GetServiceCatalogTerraformImageForbidden {
	return &GetServiceCatalogTerraformImageForbidden{}
}

/*GetServiceCatalogTerraformImageForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetServiceCatalogTerraformImageForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetServiceCatalogTerraformImageForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram/image][%d] getServiceCatalogTerraformImageForbidden  %+v", 403, o.Payload)
}

func (o *GetServiceCatalogTerraformImageForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogTerraformImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServiceCatalogTerraformImageNotFound creates a GetServiceCatalogTerraformImageNotFound with default headers values
func NewGetServiceCatalogTerraformImageNotFound() *GetServiceCatalogTerraformImageNotFound {
	return &GetServiceCatalogTerraformImageNotFound{}
}

/*GetServiceCatalogTerraformImageNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetServiceCatalogTerraformImageNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetServiceCatalogTerraformImageNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram/image][%d] getServiceCatalogTerraformImageNotFound  %+v", 404, o.Payload)
}

func (o *GetServiceCatalogTerraformImageNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogTerraformImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServiceCatalogTerraformImageDefault creates a GetServiceCatalogTerraformImageDefault with default headers values
func NewGetServiceCatalogTerraformImageDefault(code int) *GetServiceCatalogTerraformImageDefault {
	return &GetServiceCatalogTerraformImageDefault{
		_statusCode: code,
	}
}

/*GetServiceCatalogTerraformImageDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetServiceCatalogTerraformImageDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get service catalog terraform image default response
func (o *GetServiceCatalogTerraformImageDefault) Code() int {
	return o._statusCode
}

func (o *GetServiceCatalogTerraformImageDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram/image][%d] getServiceCatalogTerraformImage default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceCatalogTerraformImageDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogTerraformImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetServiceCatalogTerraformImageOKBody get service catalog terraform image o k body
swagger:model GetServiceCatalogTerraformImageOKBody
*/
type GetServiceCatalogTerraformImageOKBody struct {

	// data
	// Required: true
	Data *models.TerraformImage `json:"data"`
}

// Validate validates this get service catalog terraform image o k body
func (o *GetServiceCatalogTerraformImageOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceCatalogTerraformImageOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getServiceCatalogTerraformImageOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceCatalogTerraformImageOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceCatalogTerraformImageOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceCatalogTerraformImageOKBody) UnmarshalBinary(b []byte) error {
	var res GetServiceCatalogTerraformImageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
