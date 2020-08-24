// Code generated by go-swagger; DO NOT EDIT.

package service_catalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// GetServiceCatalogsReader is a Reader for the GetServiceCatalogs structure.
type GetServiceCatalogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceCatalogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceCatalogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetServiceCatalogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServiceCatalogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceCatalogsOK creates a GetServiceCatalogsOK with default headers values
func NewGetServiceCatalogsOK() *GetServiceCatalogsOK {
	return &GetServiceCatalogsOK{}
}

/*GetServiceCatalogsOK handles this case with default header values.

List of the service catalogs.
*/
type GetServiceCatalogsOK struct {
	Payload *GetServiceCatalogsOKBody
}

func (o *GetServiceCatalogsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs][%d] getServiceCatalogsOK  %+v", 200, o.Payload)
}

func (o *GetServiceCatalogsOK) GetPayload() *GetServiceCatalogsOKBody {
	return o.Payload
}

func (o *GetServiceCatalogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetServiceCatalogsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceCatalogsForbidden creates a GetServiceCatalogsForbidden with default headers values
func NewGetServiceCatalogsForbidden() *GetServiceCatalogsForbidden {
	return &GetServiceCatalogsForbidden{}
}

/*GetServiceCatalogsForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetServiceCatalogsForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetServiceCatalogsForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs][%d] getServiceCatalogsForbidden  %+v", 403, o.Payload)
}

func (o *GetServiceCatalogsForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServiceCatalogsDefault creates a GetServiceCatalogsDefault with default headers values
func NewGetServiceCatalogsDefault(code int) *GetServiceCatalogsDefault {
	return &GetServiceCatalogsDefault{
		_statusCode: code,
	}
}

/*GetServiceCatalogsDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetServiceCatalogsDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get service catalogs default response
func (o *GetServiceCatalogsDefault) Code() int {
	return o._statusCode
}

func (o *GetServiceCatalogsDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs][%d] getServiceCatalogs default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceCatalogsDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetServiceCatalogsOKBody get service catalogs o k body
swagger:model GetServiceCatalogsOKBody
*/
type GetServiceCatalogsOKBody struct {

	// data
	// Required: true
	Data []*models.ServiceCatalog `json:"data"`

	// pagination
	// Required: true
	Pagination *models.Pagination `json:"pagination"`
}

// Validate validates this get service catalogs o k body
func (o *GetServiceCatalogsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceCatalogsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getServiceCatalogsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getServiceCatalogsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetServiceCatalogsOKBody) validatePagination(formats strfmt.Registry) error {

	if err := validate.Required("getServiceCatalogsOK"+"."+"pagination", "body", o.Pagination); err != nil {
		return err
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceCatalogsOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceCatalogsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceCatalogsOKBody) UnmarshalBinary(b []byte) error {
	var res GetServiceCatalogsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
