// Code generated by go-swagger; DO NOT EDIT.

package organization_pipelines

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

/*
PausePipelineNoContent describes a response with status code 204, with default header values.

Pipeline has been paused.
*/
type PausePipelineNoContent struct {
}

// IsSuccess returns true when this pause pipeline no content response has a 2xx status code
func (o *PausePipelineNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pause pipeline no content response has a 3xx status code
func (o *PausePipelineNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pause pipeline no content response has a 4xx status code
func (o *PausePipelineNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this pause pipeline no content response has a 5xx status code
func (o *PausePipelineNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this pause pipeline no content response a status code equal to that given
func (o *PausePipelineNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the pause pipeline no content response
func (o *PausePipelineNoContent) Code() int {
	return 204
}

func (o *PausePipelineNoContent) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/pause][%d] pausePipelineNoContent", 204)
}

func (o *PausePipelineNoContent) String() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/pause][%d] pausePipelineNoContent", 204)
}

func (o *PausePipelineNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPausePipelineForbidden creates a PausePipelineForbidden with default headers values
func NewPausePipelineForbidden() *PausePipelineForbidden {
	return &PausePipelineForbidden{}
}

/*
PausePipelineForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type PausePipelineForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this pause pipeline forbidden response has a 2xx status code
func (o *PausePipelineForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pause pipeline forbidden response has a 3xx status code
func (o *PausePipelineForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pause pipeline forbidden response has a 4xx status code
func (o *PausePipelineForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pause pipeline forbidden response has a 5xx status code
func (o *PausePipelineForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pause pipeline forbidden response a status code equal to that given
func (o *PausePipelineForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pause pipeline forbidden response
func (o *PausePipelineForbidden) Code() int {
	return 403
}

func (o *PausePipelineForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/pause][%d] pausePipelineForbidden %s", 403, payload)
}

func (o *PausePipelineForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/pause][%d] pausePipelineForbidden %s", 403, payload)
}

func (o *PausePipelineForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *PausePipelineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPausePipelineNotFound creates a PausePipelineNotFound with default headers values
func NewPausePipelineNotFound() *PausePipelineNotFound {
	return &PausePipelineNotFound{}
}

/*
PausePipelineNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type PausePipelineNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this pause pipeline not found response has a 2xx status code
func (o *PausePipelineNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pause pipeline not found response has a 3xx status code
func (o *PausePipelineNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pause pipeline not found response has a 4xx status code
func (o *PausePipelineNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pause pipeline not found response has a 5xx status code
func (o *PausePipelineNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pause pipeline not found response a status code equal to that given
func (o *PausePipelineNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pause pipeline not found response
func (o *PausePipelineNotFound) Code() int {
	return 404
}

func (o *PausePipelineNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/pause][%d] pausePipelineNotFound %s", 404, payload)
}

func (o *PausePipelineNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/pause][%d] pausePipelineNotFound %s", 404, payload)
}

func (o *PausePipelineNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *PausePipelineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPausePipelineDefault creates a PausePipelineDefault with default headers values
func NewPausePipelineDefault(code int) *PausePipelineDefault {
	return &PausePipelineDefault{
		_statusCode: code,
	}
}

/*
PausePipelineDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type PausePipelineDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this pause pipeline default response has a 2xx status code
func (o *PausePipelineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pause pipeline default response has a 3xx status code
func (o *PausePipelineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pause pipeline default response has a 4xx status code
func (o *PausePipelineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pause pipeline default response has a 5xx status code
func (o *PausePipelineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pause pipeline default response a status code equal to that given
func (o *PausePipelineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the pause pipeline default response
func (o *PausePipelineDefault) Code() int {
	return o._statusCode
}

func (o *PausePipelineDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/pause][%d] pausePipeline default %s", o._statusCode, payload)
}

func (o *PausePipelineDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/pause][%d] pausePipeline default %s", o._statusCode, payload)
}

func (o *PausePipelineDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *PausePipelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
