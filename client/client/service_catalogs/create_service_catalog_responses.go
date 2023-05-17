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

// CreateServiceCatalogReader is a Reader for the CreateServiceCatalog structure.
type CreateServiceCatalogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServiceCatalogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateServiceCatalogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateServiceCatalogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateServiceCatalogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateServiceCatalogUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateServiceCatalogDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateServiceCatalogOK creates a CreateServiceCatalogOK with default headers values
func NewCreateServiceCatalogOK() *CreateServiceCatalogOK {
	return &CreateServiceCatalogOK{}
}

/*CreateServiceCatalogOK handles this case with default header values.

The information of the service catalog.
*/
type CreateServiceCatalogOK struct {
	Payload *CreateServiceCatalogOKBody
}

func (o *CreateServiceCatalogOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalogs][%d] createServiceCatalogOK  %+v", 200, o.Payload)
}

func (o *CreateServiceCatalogOK) GetPayload() *CreateServiceCatalogOKBody {
	return o.Payload
}

func (o *CreateServiceCatalogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateServiceCatalogOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceCatalogForbidden creates a CreateServiceCatalogForbidden with default headers values
func NewCreateServiceCatalogForbidden() *CreateServiceCatalogForbidden {
	return &CreateServiceCatalogForbidden{}
}

/*CreateServiceCatalogForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type CreateServiceCatalogForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *CreateServiceCatalogForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalogs][%d] createServiceCatalogForbidden  %+v", 403, o.Payload)
}

func (o *CreateServiceCatalogForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateServiceCatalogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateServiceCatalogNotFound creates a CreateServiceCatalogNotFound with default headers values
func NewCreateServiceCatalogNotFound() *CreateServiceCatalogNotFound {
	return &CreateServiceCatalogNotFound{}
}

/*CreateServiceCatalogNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type CreateServiceCatalogNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *CreateServiceCatalogNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalogs][%d] createServiceCatalogNotFound  %+v", 404, o.Payload)
}

func (o *CreateServiceCatalogNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateServiceCatalogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateServiceCatalogUnprocessableEntity creates a CreateServiceCatalogUnprocessableEntity with default headers values
func NewCreateServiceCatalogUnprocessableEntity() *CreateServiceCatalogUnprocessableEntity {
	return &CreateServiceCatalogUnprocessableEntity{}
}

/*CreateServiceCatalogUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type CreateServiceCatalogUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *CreateServiceCatalogUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalogs][%d] createServiceCatalogUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateServiceCatalogUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateServiceCatalogUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateServiceCatalogDefault creates a CreateServiceCatalogDefault with default headers values
func NewCreateServiceCatalogDefault(code int) *CreateServiceCatalogDefault {
	return &CreateServiceCatalogDefault{
		_statusCode: code,
	}
}

/*CreateServiceCatalogDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreateServiceCatalogDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the create service catalog default response
func (o *CreateServiceCatalogDefault) Code() int {
	return o._statusCode
}

func (o *CreateServiceCatalogDefault) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalogs][%d] createServiceCatalog default  %+v", o._statusCode, o.Payload)
}

func (o *CreateServiceCatalogDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateServiceCatalogDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*CreateServiceCatalogOKBody create service catalog o k body
swagger:model CreateServiceCatalogOKBody
*/
type CreateServiceCatalogOKBody struct {

	// data
	// Required: true
	Data *models.ServiceCatalog `json:"data"`
}

// Validate validates this create service catalog o k body
func (o *CreateServiceCatalogOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateServiceCatalogOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("createServiceCatalogOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createServiceCatalogOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateServiceCatalogOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateServiceCatalogOKBody) UnmarshalBinary(b []byte) error {
	var res CreateServiceCatalogOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
