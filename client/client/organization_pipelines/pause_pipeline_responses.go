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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// PausePipelineReader is a Reader for the PausePipeline structure.
type PausePipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PausePipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPausePipelineNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPausePipelineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPausePipelineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPausePipelineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPausePipelineNoContent creates a PausePipelineNoContent with default headers values
func NewPausePipelineNoContent() *PausePipelineNoContent {
	return &PausePipelineNoContent{}
}

/*PausePipelineNoContent handles this case with default header values.

Pipeline has been paused.
*/
type PausePipelineNoContent struct {
}

func (o *PausePipelineNoContent) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/pause][%d] pausePipelineNoContent ", 204)
}

func (o *PausePipelineNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPausePipelineForbidden creates a PausePipelineForbidden with default headers values
func NewPausePipelineForbidden() *PausePipelineForbidden {
	return &PausePipelineForbidden{}
}

/*PausePipelineForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type PausePipelineForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *PausePipelineForbidden) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/pause][%d] pausePipelineForbidden  %+v", 403, o.Payload)
}

func (o *PausePipelineForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *PausePipelineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPausePipelineNotFound creates a PausePipelineNotFound with default headers values
func NewPausePipelineNotFound() *PausePipelineNotFound {
	return &PausePipelineNotFound{}
}

/*PausePipelineNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type PausePipelineNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *PausePipelineNotFound) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/pause][%d] pausePipelineNotFound  %+v", 404, o.Payload)
}

func (o *PausePipelineNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *PausePipelineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPausePipelineDefault creates a PausePipelineDefault with default headers values
func NewPausePipelineDefault(code int) *PausePipelineDefault {
	return &PausePipelineDefault{
		_statusCode: code,
	}
}

/*PausePipelineDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type PausePipelineDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the pause pipeline default response
func (o *PausePipelineDefault) Code() int {
	return o._statusCode
}

func (o *PausePipelineDefault) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/pause][%d] pausePipeline default  %+v", o._statusCode, o.Payload)
}

func (o *PausePipelineDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *PausePipelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}