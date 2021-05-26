// Code generated by go-swagger; DO NOT EDIT.

package organization_config_repositories

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

// GetConfigRepositoriesReader is a Reader for the GetConfigRepositories structure.
type GetConfigRepositoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigRepositoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigRepositoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetConfigRepositoriesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetConfigRepositoriesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetConfigRepositoriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetConfigRepositoriesOK creates a GetConfigRepositoriesOK with default headers values
func NewGetConfigRepositoriesOK() *GetConfigRepositoriesOK {
	return &GetConfigRepositoriesOK{}
}

/*GetConfigRepositoriesOK handles this case with default header values.

List of the config repositories.
*/
type GetConfigRepositoriesOK struct {
	Payload *GetConfigRepositoriesOKBody
}

func (o *GetConfigRepositoriesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/config_repositories][%d] getConfigRepositoriesOK  %+v", 200, o.Payload)
}

func (o *GetConfigRepositoriesOK) GetPayload() *GetConfigRepositoriesOKBody {
	return o.Payload
}

func (o *GetConfigRepositoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetConfigRepositoriesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigRepositoriesForbidden creates a GetConfigRepositoriesForbidden with default headers values
func NewGetConfigRepositoriesForbidden() *GetConfigRepositoriesForbidden {
	return &GetConfigRepositoriesForbidden{}
}

/*GetConfigRepositoriesForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetConfigRepositoriesForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetConfigRepositoriesForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/config_repositories][%d] getConfigRepositoriesForbidden  %+v", 403, o.Payload)
}

func (o *GetConfigRepositoriesForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetConfigRepositoriesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetConfigRepositoriesUnprocessableEntity creates a GetConfigRepositoriesUnprocessableEntity with default headers values
func NewGetConfigRepositoriesUnprocessableEntity() *GetConfigRepositoriesUnprocessableEntity {
	return &GetConfigRepositoriesUnprocessableEntity{}
}

/*GetConfigRepositoriesUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetConfigRepositoriesUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetConfigRepositoriesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/config_repositories][%d] getConfigRepositoriesUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetConfigRepositoriesUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetConfigRepositoriesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetConfigRepositoriesDefault creates a GetConfigRepositoriesDefault with default headers values
func NewGetConfigRepositoriesDefault(code int) *GetConfigRepositoriesDefault {
	return &GetConfigRepositoriesDefault{
		_statusCode: code,
	}
}

/*GetConfigRepositoriesDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetConfigRepositoriesDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get config repositories default response
func (o *GetConfigRepositoriesDefault) Code() int {
	return o._statusCode
}

func (o *GetConfigRepositoriesDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/config_repositories][%d] getConfigRepositories default  %+v", o._statusCode, o.Payload)
}

func (o *GetConfigRepositoriesDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetConfigRepositoriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetConfigRepositoriesOKBody get config repositories o k body
swagger:model GetConfigRepositoriesOKBody
*/
type GetConfigRepositoriesOKBody struct {

	// data
	// Required: true
	Data []*models.ConfigRepository `json:"data"`
}

// Validate validates this get config repositories o k body
func (o *GetConfigRepositoriesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigRepositoriesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getConfigRepositoriesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigRepositoriesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetConfigRepositoriesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetConfigRepositoriesOKBody) UnmarshalBinary(b []byte) error {
	var res GetConfigRepositoriesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
