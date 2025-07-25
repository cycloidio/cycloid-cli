// Code generated by go-swagger; DO NOT EDIT.

package component_pipeline_resources_versions

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

// EnableResourceVersionReader is a Reader for the EnableResourceVersion structure.
type EnableResourceVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableResourceVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEnableResourceVersionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewEnableResourceVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEnableResourceVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEnableResourceVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnableResourceVersionNoContent creates a EnableResourceVersionNoContent with default headers values
func NewEnableResourceVersionNoContent() *EnableResourceVersionNoContent {
	return &EnableResourceVersionNoContent{}
}

/*
EnableResourceVersionNoContent describes a response with status code 204, with default header values.

Resource version have been enabled
*/
type EnableResourceVersionNoContent struct {
}

// IsSuccess returns true when this enable resource version no content response has a 2xx status code
func (o *EnableResourceVersionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enable resource version no content response has a 3xx status code
func (o *EnableResourceVersionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable resource version no content response has a 4xx status code
func (o *EnableResourceVersionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this enable resource version no content response has a 5xx status code
func (o *EnableResourceVersionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this enable resource version no content response a status code equal to that given
func (o *EnableResourceVersionNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the enable resource version no content response
func (o *EnableResourceVersionNoContent) Code() int {
	return 204
}

func (o *EnableResourceVersionNoContent) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/versions/{resource_version_id}/enable][%d] enableResourceVersionNoContent", 204)
}

func (o *EnableResourceVersionNoContent) String() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/versions/{resource_version_id}/enable][%d] enableResourceVersionNoContent", 204)
}

func (o *EnableResourceVersionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableResourceVersionForbidden creates a EnableResourceVersionForbidden with default headers values
func NewEnableResourceVersionForbidden() *EnableResourceVersionForbidden {
	return &EnableResourceVersionForbidden{}
}

/*
EnableResourceVersionForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type EnableResourceVersionForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this enable resource version forbidden response has a 2xx status code
func (o *EnableResourceVersionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enable resource version forbidden response has a 3xx status code
func (o *EnableResourceVersionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable resource version forbidden response has a 4xx status code
func (o *EnableResourceVersionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this enable resource version forbidden response has a 5xx status code
func (o *EnableResourceVersionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this enable resource version forbidden response a status code equal to that given
func (o *EnableResourceVersionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the enable resource version forbidden response
func (o *EnableResourceVersionForbidden) Code() int {
	return 403
}

func (o *EnableResourceVersionForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/versions/{resource_version_id}/enable][%d] enableResourceVersionForbidden %s", 403, payload)
}

func (o *EnableResourceVersionForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/versions/{resource_version_id}/enable][%d] enableResourceVersionForbidden %s", 403, payload)
}

func (o *EnableResourceVersionForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *EnableResourceVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewEnableResourceVersionNotFound creates a EnableResourceVersionNotFound with default headers values
func NewEnableResourceVersionNotFound() *EnableResourceVersionNotFound {
	return &EnableResourceVersionNotFound{}
}

/*
EnableResourceVersionNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type EnableResourceVersionNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this enable resource version not found response has a 2xx status code
func (o *EnableResourceVersionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enable resource version not found response has a 3xx status code
func (o *EnableResourceVersionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable resource version not found response has a 4xx status code
func (o *EnableResourceVersionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this enable resource version not found response has a 5xx status code
func (o *EnableResourceVersionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this enable resource version not found response a status code equal to that given
func (o *EnableResourceVersionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the enable resource version not found response
func (o *EnableResourceVersionNotFound) Code() int {
	return 404
}

func (o *EnableResourceVersionNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/versions/{resource_version_id}/enable][%d] enableResourceVersionNotFound %s", 404, payload)
}

func (o *EnableResourceVersionNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/versions/{resource_version_id}/enable][%d] enableResourceVersionNotFound %s", 404, payload)
}

func (o *EnableResourceVersionNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *EnableResourceVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewEnableResourceVersionDefault creates a EnableResourceVersionDefault with default headers values
func NewEnableResourceVersionDefault(code int) *EnableResourceVersionDefault {
	return &EnableResourceVersionDefault{
		_statusCode: code,
	}
}

/*
EnableResourceVersionDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type EnableResourceVersionDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this enable resource version default response has a 2xx status code
func (o *EnableResourceVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this enable resource version default response has a 3xx status code
func (o *EnableResourceVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this enable resource version default response has a 4xx status code
func (o *EnableResourceVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this enable resource version default response has a 5xx status code
func (o *EnableResourceVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this enable resource version default response a status code equal to that given
func (o *EnableResourceVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the enable resource version default response
func (o *EnableResourceVersionDefault) Code() int {
	return o._statusCode
}

func (o *EnableResourceVersionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/versions/{resource_version_id}/enable][%d] enableResourceVersion default %s", o._statusCode, payload)
}

func (o *EnableResourceVersionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/versions/{resource_version_id}/enable][%d] enableResourceVersion default %s", o._statusCode, payload)
}

func (o *EnableResourceVersionDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *EnableResourceVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
