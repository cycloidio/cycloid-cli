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

// DiffPipelineReader is a Reader for the DiffPipeline structure.
type DiffPipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DiffPipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDiffPipelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDiffPipelineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDiffPipelineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 411:
		result := NewDiffPipelineLengthRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDiffPipelineUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDiffPipelineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDiffPipelineOK creates a DiffPipelineOK with default headers values
func NewDiffPipelineOK() *DiffPipelineOK {
	return &DiffPipelineOK{}
}

/*DiffPipelineOK handles this case with default header values.

The diff between the provided pipeline configuration and the existing pipeline has been done.
*/
type DiffPipelineOK struct {
	Payload *DiffPipelineOKBody
}

func (o *DiffPipelineOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/pipelines/{inpath_pipeline_name}/diff][%d] diffPipelineOK  %+v", 200, o.Payload)
}

func (o *DiffPipelineOK) GetPayload() *DiffPipelineOKBody {
	return o.Payload
}

func (o *DiffPipelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DiffPipelineOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiffPipelineForbidden creates a DiffPipelineForbidden with default headers values
func NewDiffPipelineForbidden() *DiffPipelineForbidden {
	return &DiffPipelineForbidden{}
}

/*DiffPipelineForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type DiffPipelineForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *DiffPipelineForbidden) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/pipelines/{inpath_pipeline_name}/diff][%d] diffPipelineForbidden  %+v", 403, o.Payload)
}

func (o *DiffPipelineForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DiffPipelineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDiffPipelineNotFound creates a DiffPipelineNotFound with default headers values
func NewDiffPipelineNotFound() *DiffPipelineNotFound {
	return &DiffPipelineNotFound{}
}

/*DiffPipelineNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type DiffPipelineNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *DiffPipelineNotFound) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/pipelines/{inpath_pipeline_name}/diff][%d] diffPipelineNotFound  %+v", 404, o.Payload)
}

func (o *DiffPipelineNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DiffPipelineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDiffPipelineLengthRequired creates a DiffPipelineLengthRequired with default headers values
func NewDiffPipelineLengthRequired() *DiffPipelineLengthRequired {
	return &DiffPipelineLengthRequired{}
}

/*DiffPipelineLengthRequired handles this case with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type DiffPipelineLengthRequired struct {
}

func (o *DiffPipelineLengthRequired) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/pipelines/{inpath_pipeline_name}/diff][%d] diffPipelineLengthRequired ", 411)
}

func (o *DiffPipelineLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDiffPipelineUnprocessableEntity creates a DiffPipelineUnprocessableEntity with default headers values
func NewDiffPipelineUnprocessableEntity() *DiffPipelineUnprocessableEntity {
	return &DiffPipelineUnprocessableEntity{}
}

/*DiffPipelineUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type DiffPipelineUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *DiffPipelineUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/pipelines/{inpath_pipeline_name}/diff][%d] diffPipelineUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DiffPipelineUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DiffPipelineUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDiffPipelineDefault creates a DiffPipelineDefault with default headers values
func NewDiffPipelineDefault(code int) *DiffPipelineDefault {
	return &DiffPipelineDefault{
		_statusCode: code,
	}
}

/*DiffPipelineDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DiffPipelineDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the diff pipeline default response
func (o *DiffPipelineDefault) Code() int {
	return o._statusCode
}

func (o *DiffPipelineDefault) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/pipelines/{inpath_pipeline_name}/diff][%d] diffPipeline default  %+v", o._statusCode, o.Payload)
}

func (o *DiffPipelineDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DiffPipelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*DiffPipelineOKBody diff pipeline o k body
swagger:model DiffPipelineOKBody
*/
type DiffPipelineOKBody struct {

	// data
	// Required: true
	Data *models.PipelineDiffs `json:"data"`
}

// Validate validates this diff pipeline o k body
func (o *DiffPipelineOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DiffPipelineOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("diffPipelineOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("diffPipelineOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DiffPipelineOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DiffPipelineOKBody) UnmarshalBinary(b []byte) error {
	var res DiffPipelineOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
