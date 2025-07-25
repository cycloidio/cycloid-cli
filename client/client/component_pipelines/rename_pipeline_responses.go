// Code generated by go-swagger; DO NOT EDIT.

package component_pipelines

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

/*
RenamePipelineNoContent describes a response with status code 204, with default header values.

Pipeline has been renamed.
*/
type RenamePipelineNoContent struct {
}

// IsSuccess returns true when this rename pipeline no content response has a 2xx status code
func (o *RenamePipelineNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rename pipeline no content response has a 3xx status code
func (o *RenamePipelineNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rename pipeline no content response has a 4xx status code
func (o *RenamePipelineNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this rename pipeline no content response has a 5xx status code
func (o *RenamePipelineNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this rename pipeline no content response a status code equal to that given
func (o *RenamePipelineNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the rename pipeline no content response
func (o *RenamePipelineNoContent) Code() int {
	return 204
}

func (o *RenamePipelineNoContent) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/rename][%d] renamePipelineNoContent", 204)
}

func (o *RenamePipelineNoContent) String() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/rename][%d] renamePipelineNoContent", 204)
}

func (o *RenamePipelineNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRenamePipelineForbidden creates a RenamePipelineForbidden with default headers values
func NewRenamePipelineForbidden() *RenamePipelineForbidden {
	return &RenamePipelineForbidden{}
}

/*
RenamePipelineForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type RenamePipelineForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this rename pipeline forbidden response has a 2xx status code
func (o *RenamePipelineForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rename pipeline forbidden response has a 3xx status code
func (o *RenamePipelineForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rename pipeline forbidden response has a 4xx status code
func (o *RenamePipelineForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this rename pipeline forbidden response has a 5xx status code
func (o *RenamePipelineForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this rename pipeline forbidden response a status code equal to that given
func (o *RenamePipelineForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the rename pipeline forbidden response
func (o *RenamePipelineForbidden) Code() int {
	return 403
}

func (o *RenamePipelineForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/rename][%d] renamePipelineForbidden %s", 403, payload)
}

func (o *RenamePipelineForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/rename][%d] renamePipelineForbidden %s", 403, payload)
}

func (o *RenamePipelineForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RenamePipelineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRenamePipelineNotFound creates a RenamePipelineNotFound with default headers values
func NewRenamePipelineNotFound() *RenamePipelineNotFound {
	return &RenamePipelineNotFound{}
}

/*
RenamePipelineNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type RenamePipelineNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this rename pipeline not found response has a 2xx status code
func (o *RenamePipelineNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rename pipeline not found response has a 3xx status code
func (o *RenamePipelineNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rename pipeline not found response has a 4xx status code
func (o *RenamePipelineNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this rename pipeline not found response has a 5xx status code
func (o *RenamePipelineNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this rename pipeline not found response a status code equal to that given
func (o *RenamePipelineNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the rename pipeline not found response
func (o *RenamePipelineNotFound) Code() int {
	return 404
}

func (o *RenamePipelineNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/rename][%d] renamePipelineNotFound %s", 404, payload)
}

func (o *RenamePipelineNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/rename][%d] renamePipelineNotFound %s", 404, payload)
}

func (o *RenamePipelineNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RenamePipelineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRenamePipelineDefault creates a RenamePipelineDefault with default headers values
func NewRenamePipelineDefault(code int) *RenamePipelineDefault {
	return &RenamePipelineDefault{
		_statusCode: code,
	}
}

/*
RenamePipelineDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type RenamePipelineDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this rename pipeline default response has a 2xx status code
func (o *RenamePipelineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this rename pipeline default response has a 3xx status code
func (o *RenamePipelineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this rename pipeline default response has a 4xx status code
func (o *RenamePipelineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this rename pipeline default response has a 5xx status code
func (o *RenamePipelineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this rename pipeline default response a status code equal to that given
func (o *RenamePipelineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the rename pipeline default response
func (o *RenamePipelineDefault) Code() int {
	return o._statusCode
}

func (o *RenamePipelineDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/rename][%d] renamePipeline default %s", o._statusCode, payload)
}

func (o *RenamePipelineDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/rename][%d] renamePipeline default %s", o._statusCode, payload)
}

func (o *RenamePipelineDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RenamePipelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
