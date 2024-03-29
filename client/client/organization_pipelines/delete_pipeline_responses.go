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

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// DeletePipelineReader is a Reader for the DeletePipeline structure.
type DeletePipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePipelineNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeletePipelineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePipelineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeletePipelineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePipelineNoContent creates a DeletePipelineNoContent with default headers values
func NewDeletePipelineNoContent() *DeletePipelineNoContent {
	return &DeletePipelineNoContent{}
}

/*DeletePipelineNoContent handles this case with default header values.

Pipeline has been deleted.
*/
type DeletePipelineNoContent struct {
}

func (o *DeletePipelineNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}][%d] deletePipelineNoContent ", 204)
}

func (o *DeletePipelineNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePipelineForbidden creates a DeletePipelineForbidden with default headers values
func NewDeletePipelineForbidden() *DeletePipelineForbidden {
	return &DeletePipelineForbidden{}
}

/*DeletePipelineForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type DeletePipelineForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *DeletePipelineForbidden) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}][%d] deletePipelineForbidden  %+v", 403, o.Payload)
}

func (o *DeletePipelineForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeletePipelineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeletePipelineNotFound creates a DeletePipelineNotFound with default headers values
func NewDeletePipelineNotFound() *DeletePipelineNotFound {
	return &DeletePipelineNotFound{}
}

/*DeletePipelineNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type DeletePipelineNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *DeletePipelineNotFound) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}][%d] deletePipelineNotFound  %+v", 404, o.Payload)
}

func (o *DeletePipelineNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeletePipelineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeletePipelineDefault creates a DeletePipelineDefault with default headers values
func NewDeletePipelineDefault(code int) *DeletePipelineDefault {
	return &DeletePipelineDefault{
		_statusCode: code,
	}
}

/*DeletePipelineDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DeletePipelineDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the delete pipeline default response
func (o *DeletePipelineDefault) Code() int {
	return o._statusCode
}

func (o *DeletePipelineDefault) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}][%d] deletePipeline default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePipelineDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeletePipelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
