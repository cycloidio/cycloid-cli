// Code generated by go-swagger; DO NOT EDIT.

package component_pipelines_jobs_builds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cycloidio/cycloid-cli/client/models"
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

/*
AbortBuildNoContent describes a response with status code 204, with default header values.

The build has been aborted.
*/
type AbortBuildNoContent struct {
}

// IsSuccess returns true when this abort build no content response has a 2xx status code
func (o *AbortBuildNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this abort build no content response has a 3xx status code
func (o *AbortBuildNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this abort build no content response has a 4xx status code
func (o *AbortBuildNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this abort build no content response has a 5xx status code
func (o *AbortBuildNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this abort build no content response a status code equal to that given
func (o *AbortBuildNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the abort build no content response
func (o *AbortBuildNoContent) Code() int {
	return 204
}

func (o *AbortBuildNoContent) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/abort][%d] abortBuildNoContent", 204)
}

func (o *AbortBuildNoContent) String() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/abort][%d] abortBuildNoContent", 204)
}

func (o *AbortBuildNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAbortBuildForbidden creates a AbortBuildForbidden with default headers values
func NewAbortBuildForbidden() *AbortBuildForbidden {
	return &AbortBuildForbidden{}
}

/*
AbortBuildForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type AbortBuildForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this abort build forbidden response has a 2xx status code
func (o *AbortBuildForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this abort build forbidden response has a 3xx status code
func (o *AbortBuildForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this abort build forbidden response has a 4xx status code
func (o *AbortBuildForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this abort build forbidden response has a 5xx status code
func (o *AbortBuildForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this abort build forbidden response a status code equal to that given
func (o *AbortBuildForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the abort build forbidden response
func (o *AbortBuildForbidden) Code() int {
	return 403
}

func (o *AbortBuildForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/abort][%d] abortBuildForbidden %s", 403, payload)
}

func (o *AbortBuildForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/abort][%d] abortBuildForbidden %s", 403, payload)
}

func (o *AbortBuildForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *AbortBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Length
	hdrContentLength := response.GetHeader("Content-Length")

	if hdrContentLength != "" {
		valcontentLength, err := swag.ConvertUint64(hdrContentLength)
		if err != nil {
			return errors.InvalidType("Content-Length", "header", "uint64", hdrContentLength)
		}
		o.ContentLength = valcontentLength
	}

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

/*
AbortBuildNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type AbortBuildNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this abort build not found response has a 2xx status code
func (o *AbortBuildNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this abort build not found response has a 3xx status code
func (o *AbortBuildNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this abort build not found response has a 4xx status code
func (o *AbortBuildNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this abort build not found response has a 5xx status code
func (o *AbortBuildNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this abort build not found response a status code equal to that given
func (o *AbortBuildNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the abort build not found response
func (o *AbortBuildNotFound) Code() int {
	return 404
}

func (o *AbortBuildNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/abort][%d] abortBuildNotFound %s", 404, payload)
}

func (o *AbortBuildNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/abort][%d] abortBuildNotFound %s", 404, payload)
}

func (o *AbortBuildNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *AbortBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Length
	hdrContentLength := response.GetHeader("Content-Length")

	if hdrContentLength != "" {
		valcontentLength, err := swag.ConvertUint64(hdrContentLength)
		if err != nil {
			return errors.InvalidType("Content-Length", "header", "uint64", hdrContentLength)
		}
		o.ContentLength = valcontentLength
	}

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

/*
AbortBuildDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type AbortBuildDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this abort build default response has a 2xx status code
func (o *AbortBuildDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this abort build default response has a 3xx status code
func (o *AbortBuildDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this abort build default response has a 4xx status code
func (o *AbortBuildDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this abort build default response has a 5xx status code
func (o *AbortBuildDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this abort build default response a status code equal to that given
func (o *AbortBuildDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the abort build default response
func (o *AbortBuildDefault) Code() int {
	return o._statusCode
}

func (o *AbortBuildDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/abort][%d] abortBuild default %s", o._statusCode, payload)
}

func (o *AbortBuildDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/abort][%d] abortBuild default %s", o._statusCode, payload)
}

func (o *AbortBuildDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *AbortBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Length
	hdrContentLength := response.GetHeader("Content-Length")

	if hdrContentLength != "" {
		valcontentLength, err := swag.ConvertUint64(hdrContentLength)
		if err != nil {
			return errors.InvalidType("Content-Length", "header", "uint64", hdrContentLength)
		}
		o.ContentLength = valcontentLength
	}

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
