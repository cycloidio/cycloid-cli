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

// GetServiceCatalogTerraformDiagramReader is a Reader for the GetServiceCatalogTerraformDiagram structure.
type GetServiceCatalogTerraformDiagramReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceCatalogTerraformDiagramReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceCatalogTerraformDiagramOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetServiceCatalogTerraformDiagramForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServiceCatalogTerraformDiagramNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServiceCatalogTerraformDiagramDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceCatalogTerraformDiagramOK creates a GetServiceCatalogTerraformDiagramOK with default headers values
func NewGetServiceCatalogTerraformDiagramOK() *GetServiceCatalogTerraformDiagramOK {
	return &GetServiceCatalogTerraformDiagramOK{}
}

/*GetServiceCatalogTerraformDiagramOK handles this case with default header values.

The information of Terraform Diagram
*/
type GetServiceCatalogTerraformDiagramOK struct {
	Payload *GetServiceCatalogTerraformDiagramOKBody
}

func (o *GetServiceCatalogTerraformDiagramOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] getServiceCatalogTerraformDiagramOK  %+v", 200, o.Payload)
}

func (o *GetServiceCatalogTerraformDiagramOK) GetPayload() *GetServiceCatalogTerraformDiagramOKBody {
	return o.Payload
}

func (o *GetServiceCatalogTerraformDiagramOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetServiceCatalogTerraformDiagramOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceCatalogTerraformDiagramForbidden creates a GetServiceCatalogTerraformDiagramForbidden with default headers values
func NewGetServiceCatalogTerraformDiagramForbidden() *GetServiceCatalogTerraformDiagramForbidden {
	return &GetServiceCatalogTerraformDiagramForbidden{}
}

/*GetServiceCatalogTerraformDiagramForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetServiceCatalogTerraformDiagramForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetServiceCatalogTerraformDiagramForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] getServiceCatalogTerraformDiagramForbidden  %+v", 403, o.Payload)
}

func (o *GetServiceCatalogTerraformDiagramForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogTerraformDiagramForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServiceCatalogTerraformDiagramNotFound creates a GetServiceCatalogTerraformDiagramNotFound with default headers values
func NewGetServiceCatalogTerraformDiagramNotFound() *GetServiceCatalogTerraformDiagramNotFound {
	return &GetServiceCatalogTerraformDiagramNotFound{}
}

/*GetServiceCatalogTerraformDiagramNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetServiceCatalogTerraformDiagramNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetServiceCatalogTerraformDiagramNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] getServiceCatalogTerraformDiagramNotFound  %+v", 404, o.Payload)
}

func (o *GetServiceCatalogTerraformDiagramNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogTerraformDiagramNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServiceCatalogTerraformDiagramDefault creates a GetServiceCatalogTerraformDiagramDefault with default headers values
func NewGetServiceCatalogTerraformDiagramDefault(code int) *GetServiceCatalogTerraformDiagramDefault {
	return &GetServiceCatalogTerraformDiagramDefault{
		_statusCode: code,
	}
}

/*GetServiceCatalogTerraformDiagramDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetServiceCatalogTerraformDiagramDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get service catalog terraform diagram default response
func (o *GetServiceCatalogTerraformDiagramDefault) Code() int {
	return o._statusCode
}

func (o *GetServiceCatalogTerraformDiagramDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] getServiceCatalogTerraformDiagram default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceCatalogTerraformDiagramDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogTerraformDiagramDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetServiceCatalogTerraformDiagramOKBody get service catalog terraform diagram o k body
swagger:model GetServiceCatalogTerraformDiagramOKBody
*/
type GetServiceCatalogTerraformDiagramOKBody struct {

	// created at
	// Minimum: 0
	CreatedAt *uint64 `json:"created_at,omitempty"`

	// data
	// Required: true
	Data models.TerraformJSONDiagram `json:"data"`

	// updated at
	// Minimum: 0
	UpdatedAt *uint64 `json:"updated_at,omitempty"`
}

// Validate validates this get service catalog terraform diagram o k body
func (o *GetServiceCatalogTerraformDiagramOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceCatalogTerraformDiagramOKBody) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("getServiceCatalogTerraformDiagramOK"+"."+"created_at", "body", int64(*o.CreatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (o *GetServiceCatalogTerraformDiagramOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getServiceCatalogTerraformDiagramOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

func (o *GetServiceCatalogTerraformDiagramOKBody) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.UpdatedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("getServiceCatalogTerraformDiagramOK"+"."+"updated_at", "body", int64(*o.UpdatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceCatalogTerraformDiagramOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceCatalogTerraformDiagramOKBody) UnmarshalBinary(b []byte) error {
	var res GetServiceCatalogTerraformDiagramOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
