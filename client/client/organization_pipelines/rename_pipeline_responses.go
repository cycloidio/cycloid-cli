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

// RenamePipelineReader is a Reader for the RenamePipeline structure.
type RenamePipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenamePipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRenamePipelineNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRenamePipelineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRenamePipelineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRenamePipelineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRenamePipelineNoContent creates a RenamePipelineNoContent with default headers values
func NewRenamePipelineNoContent() *RenamePipelineNoContent {
	return &RenamePipelineNoContent{}
}

/*RenamePipelineNoContent handles this case with default header values.

Pipeline has been renamed.
*/
type RenamePipelineNoContent struct {
}

func (o *RenamePipelineNoContent) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/rename][%d] renamePipelineNoContent ", 204)
}

func (o *RenamePipelineNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRenamePipelineForbidden creates a RenamePipelineForbidden with default headers values
func NewRenamePipelineForbidden() *RenamePipelineForbidden {
	return &RenamePipelineForbidden{}
}

/*RenamePipelineForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type RenamePipelineForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *RenamePipelineForbidden) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/rename][%d] renamePipelineForbidden  %+v", 403, o.Payload)
}

func (o *RenamePipelineForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RenamePipelineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRenamePipelineNotFound creates a RenamePipelineNotFound with default headers values
func NewRenamePipelineNotFound() *RenamePipelineNotFound {
	return &RenamePipelineNotFound{}
}

/*RenamePipelineNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type RenamePipelineNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *RenamePipelineNotFound) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/rename][%d] renamePipelineNotFound  %+v", 404, o.Payload)
}

func (o *RenamePipelineNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RenamePipelineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRenamePipelineDefault creates a RenamePipelineDefault with default headers values
func NewRenamePipelineDefault(code int) *RenamePipelineDefault {
	return &RenamePipelineDefault{
		_statusCode: code,
	}
}

/*RenamePipelineDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type RenamePipelineDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the rename pipeline default response
func (o *RenamePipelineDefault) Code() int {
	return o._statusCode
}

func (o *RenamePipelineDefault) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/rename][%d] renamePipeline default  %+v", o._statusCode, o.Payload)
}

func (o *RenamePipelineDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RenamePipelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}