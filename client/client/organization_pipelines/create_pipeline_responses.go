// Code generated by go-swagger; DO NOT EDIT.

package organization_pipelines

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

// CreatePipelineReader is a Reader for the CreatePipeline structure.
type CreatePipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePipelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreatePipelineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreatePipelineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 411:
		result := NewCreatePipelineLengthRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreatePipelineUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreatePipelineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePipelineOK creates a CreatePipelineOK with default headers values
func NewCreatePipelineOK() *CreatePipelineOK {
	return &CreatePipelineOK{}
}

/*CreatePipelineOK handles this case with default header values.

The information of the pipeline which has been created.
*/
type CreatePipelineOK struct {
	Payload *CreatePipelineOKBody
}

func (o *CreatePipelineOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/pipelines][%d] createPipelineOK  %+v", 200, o.Payload)
}

func (o *CreatePipelineOK) GetPayload() *CreatePipelineOKBody {
	return o.Payload
}

func (o *CreatePipelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreatePipelineOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePipelineForbidden creates a CreatePipelineForbidden with default headers values
func NewCreatePipelineForbidden() *CreatePipelineForbidden {
	return &CreatePipelineForbidden{}
}

/*CreatePipelineForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type CreatePipelineForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *CreatePipelineForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/pipelines][%d] createPipelineForbidden  %+v", 403, o.Payload)
}

func (o *CreatePipelineForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreatePipelineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreatePipelineNotFound creates a CreatePipelineNotFound with default headers values
func NewCreatePipelineNotFound() *CreatePipelineNotFound {
	return &CreatePipelineNotFound{}
}

/*CreatePipelineNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type CreatePipelineNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *CreatePipelineNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/pipelines][%d] createPipelineNotFound  %+v", 404, o.Payload)
}

func (o *CreatePipelineNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreatePipelineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreatePipelineLengthRequired creates a CreatePipelineLengthRequired with default headers values
func NewCreatePipelineLengthRequired() *CreatePipelineLengthRequired {
	return &CreatePipelineLengthRequired{}
}

/*CreatePipelineLengthRequired handles this case with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type CreatePipelineLengthRequired struct {
}

func (o *CreatePipelineLengthRequired) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/pipelines][%d] createPipelineLengthRequired ", 411)
}

func (o *CreatePipelineLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePipelineUnprocessableEntity creates a CreatePipelineUnprocessableEntity with default headers values
func NewCreatePipelineUnprocessableEntity() *CreatePipelineUnprocessableEntity {
	return &CreatePipelineUnprocessableEntity{}
}

/*CreatePipelineUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type CreatePipelineUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *CreatePipelineUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/pipelines][%d] createPipelineUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreatePipelineUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreatePipelineUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreatePipelineDefault creates a CreatePipelineDefault with default headers values
func NewCreatePipelineDefault(code int) *CreatePipelineDefault {
	return &CreatePipelineDefault{
		_statusCode: code,
	}
}

/*CreatePipelineDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreatePipelineDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

// Code gets the status code for the create pipeline default response
func (o *CreatePipelineDefault) Code() int {
	return o._statusCode
}

func (o *CreatePipelineDefault) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/pipelines][%d] createPipeline default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePipelineDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreatePipelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*CreatePipelineOKBody create pipeline o k body
swagger:model CreatePipelineOKBody
*/
type CreatePipelineOKBody struct {

	// data
	// Required: true
	Data *models.Pipeline `json:"data"`
}

// Validate validates this create pipeline o k body
func (o *CreatePipelineOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreatePipelineOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("createPipelineOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createPipelineOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreatePipelineOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePipelineOKBody) UnmarshalBinary(b []byte) error {
	var res CreatePipelineOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
