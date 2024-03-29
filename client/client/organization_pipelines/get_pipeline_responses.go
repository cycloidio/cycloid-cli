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

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// GetPipelineReader is a Reader for the GetPipeline structure.
type GetPipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPipelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetPipelineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPipelineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPipelineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPipelineOK creates a GetPipelineOK with default headers values
func NewGetPipelineOK() *GetPipelineOK {
	return &GetPipelineOK{}
}

/*GetPipelineOK handles this case with default header values.

The information of the pipeline which has the specified name.
*/
type GetPipelineOK struct {
	Payload *GetPipelineOKBody
}

func (o *GetPipelineOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}][%d] getPipelineOK  %+v", 200, o.Payload)
}

func (o *GetPipelineOK) GetPayload() *GetPipelineOKBody {
	return o.Payload
}

func (o *GetPipelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPipelineOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPipelineForbidden creates a GetPipelineForbidden with default headers values
func NewGetPipelineForbidden() *GetPipelineForbidden {
	return &GetPipelineForbidden{}
}

/*GetPipelineForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetPipelineForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetPipelineForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}][%d] getPipelineForbidden  %+v", 403, o.Payload)
}

func (o *GetPipelineForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetPipelineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPipelineNotFound creates a GetPipelineNotFound with default headers values
func NewGetPipelineNotFound() *GetPipelineNotFound {
	return &GetPipelineNotFound{}
}

/*GetPipelineNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetPipelineNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetPipelineNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}][%d] getPipelineNotFound  %+v", 404, o.Payload)
}

func (o *GetPipelineNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetPipelineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPipelineDefault creates a GetPipelineDefault with default headers values
func NewGetPipelineDefault(code int) *GetPipelineDefault {
	return &GetPipelineDefault{
		_statusCode: code,
	}
}

/*GetPipelineDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetPipelineDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get pipeline default response
func (o *GetPipelineDefault) Code() int {
	return o._statusCode
}

func (o *GetPipelineDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}][%d] getPipeline default  %+v", o._statusCode, o.Payload)
}

func (o *GetPipelineDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetPipelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetPipelineOKBody get pipeline o k body
swagger:model GetPipelineOKBody
*/
type GetPipelineOKBody struct {

	// data
	// Required: true
	Data *models.Pipeline `json:"data"`
}

// Validate validates this get pipeline o k body
func (o *GetPipelineOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPipelineOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getPipelineOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPipelineOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPipelineOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPipelineOKBody) UnmarshalBinary(b []byte) error {
	var res GetPipelineOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
