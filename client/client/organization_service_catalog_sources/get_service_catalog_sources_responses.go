// Code generated by go-swagger; DO NOT EDIT.

package organization_service_catalog_sources

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

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// GetServiceCatalogSourcesReader is a Reader for the GetServiceCatalogSources structure.
type GetServiceCatalogSourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceCatalogSourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceCatalogSourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetServiceCatalogSourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetServiceCatalogSourcesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServiceCatalogSourcesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceCatalogSourcesOK creates a GetServiceCatalogSourcesOK with default headers values
func NewGetServiceCatalogSourcesOK() *GetServiceCatalogSourcesOK {
	return &GetServiceCatalogSourcesOK{}
}

/*GetServiceCatalogSourcesOK handles this case with default header values.

List of the private service catalogs.
*/
type GetServiceCatalogSourcesOK struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *GetServiceCatalogSourcesOKBody
}

func (o *GetServiceCatalogSourcesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalog_sources][%d] getServiceCatalogSourcesOK  %+v", 200, o.Payload)
}

func (o *GetServiceCatalogSourcesOK) GetPayload() *GetServiceCatalogSourcesOKBody {
	return o.Payload
}

func (o *GetServiceCatalogSourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertUint64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "uint64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(GetServiceCatalogSourcesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceCatalogSourcesForbidden creates a GetServiceCatalogSourcesForbidden with default headers values
func NewGetServiceCatalogSourcesForbidden() *GetServiceCatalogSourcesForbidden {
	return &GetServiceCatalogSourcesForbidden{}
}

/*GetServiceCatalogSourcesForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetServiceCatalogSourcesForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetServiceCatalogSourcesForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalog_sources][%d] getServiceCatalogSourcesForbidden  %+v", 403, o.Payload)
}

func (o *GetServiceCatalogSourcesForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogSourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServiceCatalogSourcesUnprocessableEntity creates a GetServiceCatalogSourcesUnprocessableEntity with default headers values
func NewGetServiceCatalogSourcesUnprocessableEntity() *GetServiceCatalogSourcesUnprocessableEntity {
	return &GetServiceCatalogSourcesUnprocessableEntity{}
}

/*GetServiceCatalogSourcesUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetServiceCatalogSourcesUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetServiceCatalogSourcesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalog_sources][%d] getServiceCatalogSourcesUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetServiceCatalogSourcesUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogSourcesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServiceCatalogSourcesDefault creates a GetServiceCatalogSourcesDefault with default headers values
func NewGetServiceCatalogSourcesDefault(code int) *GetServiceCatalogSourcesDefault {
	return &GetServiceCatalogSourcesDefault{
		_statusCode: code,
	}
}

/*GetServiceCatalogSourcesDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetServiceCatalogSourcesDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get service catalog sources default response
func (o *GetServiceCatalogSourcesDefault) Code() int {
	return o._statusCode
}

func (o *GetServiceCatalogSourcesDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalog_sources][%d] getServiceCatalogSources default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceCatalogSourcesDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogSourcesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetServiceCatalogSourcesOKBody get service catalog sources o k body
swagger:model GetServiceCatalogSourcesOKBody
*/
type GetServiceCatalogSourcesOKBody struct {

	// data
	// Required: true
	Data []*models.ServiceCatalogSource `json:"data"`
}

// Validate validates this get service catalog sources o k body
func (o *GetServiceCatalogSourcesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceCatalogSourcesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getServiceCatalogSourcesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getServiceCatalogSourcesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceCatalogSourcesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceCatalogSourcesOKBody) UnmarshalBinary(b []byte) error {
	var res GetServiceCatalogSourcesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
