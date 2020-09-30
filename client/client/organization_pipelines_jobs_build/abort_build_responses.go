// Code generated by go-swagger; DO NOT EDIT.

package organization_pipelines_jobs_build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// AbortBuildReader is a Reader for the AbortBuild structure.
type AbortBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AbortBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAbortBuildNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAbortBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAbortBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAbortBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAbortBuildNoContent creates a AbortBuildNoContent with default headers values
func NewAbortBuildNoContent() *AbortBuildNoContent {
	return &AbortBuildNoContent{}
}

/*AbortBuildNoContent handles this case with default header values.

The build has been aborted.
*/
type AbortBuildNoContent struct {
}

func (o *AbortBuildNoContent) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/abort][%d] abortBuildNoContent ", 204)
}

func (o *AbortBuildNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAbortBuildForbidden creates a AbortBuildForbidden with default headers values
func NewAbortBuildForbidden() *AbortBuildForbidden {
	return &AbortBuildForbidden{}
}

/*AbortBuildForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type AbortBuildForbidden struct {
	Payload *models.ErrorPayload
}

func (o *AbortBuildForbidden) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/abort][%d] abortBuildForbidden  %+v", 403, o.Payload)
}

func (o *AbortBuildForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *AbortBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAbortBuildNotFound creates a AbortBuildNotFound with default headers values
func NewAbortBuildNotFound() *AbortBuildNotFound {
	return &AbortBuildNotFound{}
}

/*AbortBuildNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type AbortBuildNotFound struct {
	Payload *models.ErrorPayload
}

func (o *AbortBuildNotFound) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/abort][%d] abortBuildNotFound  %+v", 404, o.Payload)
}

func (o *AbortBuildNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *AbortBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAbortBuildDefault creates a AbortBuildDefault with default headers values
func NewAbortBuildDefault(code int) *AbortBuildDefault {
	return &AbortBuildDefault{
		_statusCode: code,
	}
}

/*AbortBuildDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type AbortBuildDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the abort build default response
func (o *AbortBuildDefault) Code() int {
	return o._statusCode
}

func (o *AbortBuildDefault) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/abort][%d] abortBuild default  %+v", o._statusCode, o.Payload)
}

func (o *AbortBuildDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *AbortBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
