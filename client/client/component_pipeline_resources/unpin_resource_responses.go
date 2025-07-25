// Code generated by go-swagger; DO NOT EDIT.

package component_pipeline_resources

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

// UnpinResourceReader is a Reader for the UnpinResource structure.
type UnpinResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnpinResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUnpinResourceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUnpinResourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUnpinResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUnpinResourceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnpinResourceNoContent creates a UnpinResourceNoContent with default headers values
func NewUnpinResourceNoContent() *UnpinResourceNoContent {
	return &UnpinResourceNoContent{}
}

/*
UnpinResourceNoContent describes a response with status code 204, with default header values.

Resource has been unpinned
*/
type UnpinResourceNoContent struct {
}

// IsSuccess returns true when this unpin resource no content response has a 2xx status code
func (o *UnpinResourceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unpin resource no content response has a 3xx status code
func (o *UnpinResourceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unpin resource no content response has a 4xx status code
func (o *UnpinResourceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this unpin resource no content response has a 5xx status code
func (o *UnpinResourceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this unpin resource no content response a status code equal to that given
func (o *UnpinResourceNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the unpin resource no content response
func (o *UnpinResourceNoContent) Code() int {
	return 204
}

func (o *UnpinResourceNoContent) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/unpin][%d] unpinResourceNoContent", 204)
}

func (o *UnpinResourceNoContent) String() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/unpin][%d] unpinResourceNoContent", 204)
}

func (o *UnpinResourceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnpinResourceForbidden creates a UnpinResourceForbidden with default headers values
func NewUnpinResourceForbidden() *UnpinResourceForbidden {
	return &UnpinResourceForbidden{}
}

/*
UnpinResourceForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type UnpinResourceForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this unpin resource forbidden response has a 2xx status code
func (o *UnpinResourceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unpin resource forbidden response has a 3xx status code
func (o *UnpinResourceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unpin resource forbidden response has a 4xx status code
func (o *UnpinResourceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this unpin resource forbidden response has a 5xx status code
func (o *UnpinResourceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this unpin resource forbidden response a status code equal to that given
func (o *UnpinResourceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the unpin resource forbidden response
func (o *UnpinResourceForbidden) Code() int {
	return 403
}

func (o *UnpinResourceForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/unpin][%d] unpinResourceForbidden %s", 403, payload)
}

func (o *UnpinResourceForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/unpin][%d] unpinResourceForbidden %s", 403, payload)
}

func (o *UnpinResourceForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UnpinResourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUnpinResourceNotFound creates a UnpinResourceNotFound with default headers values
func NewUnpinResourceNotFound() *UnpinResourceNotFound {
	return &UnpinResourceNotFound{}
}

/*
UnpinResourceNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type UnpinResourceNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this unpin resource not found response has a 2xx status code
func (o *UnpinResourceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unpin resource not found response has a 3xx status code
func (o *UnpinResourceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unpin resource not found response has a 4xx status code
func (o *UnpinResourceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this unpin resource not found response has a 5xx status code
func (o *UnpinResourceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this unpin resource not found response a status code equal to that given
func (o *UnpinResourceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the unpin resource not found response
func (o *UnpinResourceNotFound) Code() int {
	return 404
}

func (o *UnpinResourceNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/unpin][%d] unpinResourceNotFound %s", 404, payload)
}

func (o *UnpinResourceNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/unpin][%d] unpinResourceNotFound %s", 404, payload)
}

func (o *UnpinResourceNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UnpinResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUnpinResourceDefault creates a UnpinResourceDefault with default headers values
func NewUnpinResourceDefault(code int) *UnpinResourceDefault {
	return &UnpinResourceDefault{
		_statusCode: code,
	}
}

/*
UnpinResourceDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UnpinResourceDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this unpin resource default response has a 2xx status code
func (o *UnpinResourceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unpin resource default response has a 3xx status code
func (o *UnpinResourceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unpin resource default response has a 4xx status code
func (o *UnpinResourceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unpin resource default response has a 5xx status code
func (o *UnpinResourceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unpin resource default response a status code equal to that given
func (o *UnpinResourceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the unpin resource default response
func (o *UnpinResourceDefault) Code() int {
	return o._statusCode
}

func (o *UnpinResourceDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/unpin][%d] unpinResource default %s", o._statusCode, payload)
}

func (o *UnpinResourceDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/unpin][%d] unpinResource default %s", o._statusCode, payload)
}

func (o *UnpinResourceDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UnpinResourceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
