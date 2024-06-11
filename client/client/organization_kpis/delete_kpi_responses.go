// Code generated by go-swagger; DO NOT EDIT.

package organization_kpis

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

// DeleteKpiReader is a Reader for the DeleteKpi structure.
type DeleteKpiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteKpiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteKpiNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteKpiForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteKpiNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteKpiDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteKpiNoContent creates a DeleteKpiNoContent with default headers values
func NewDeleteKpiNoContent() *DeleteKpiNoContent {
	return &DeleteKpiNoContent{}
}

/*
DeleteKpiNoContent describes a response with status code 204, with default header values.

Organization's KPI has been deleted
*/
type DeleteKpiNoContent struct {
}

// IsSuccess returns true when this delete kpi no content response has a 2xx status code
func (o *DeleteKpiNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete kpi no content response has a 3xx status code
func (o *DeleteKpiNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete kpi no content response has a 4xx status code
func (o *DeleteKpiNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete kpi no content response has a 5xx status code
func (o *DeleteKpiNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete kpi no content response a status code equal to that given
func (o *DeleteKpiNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete kpi no content response
func (o *DeleteKpiNoContent) Code() int {
	return 204
}

func (o *DeleteKpiNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] deleteKpiNoContent", 204)
}

func (o *DeleteKpiNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] deleteKpiNoContent", 204)
}

func (o *DeleteKpiNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteKpiForbidden creates a DeleteKpiForbidden with default headers values
func NewDeleteKpiForbidden() *DeleteKpiForbidden {
	return &DeleteKpiForbidden{}
}

/*
DeleteKpiForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type DeleteKpiForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete kpi forbidden response has a 2xx status code
func (o *DeleteKpiForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete kpi forbidden response has a 3xx status code
func (o *DeleteKpiForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete kpi forbidden response has a 4xx status code
func (o *DeleteKpiForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete kpi forbidden response has a 5xx status code
func (o *DeleteKpiForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete kpi forbidden response a status code equal to that given
func (o *DeleteKpiForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete kpi forbidden response
func (o *DeleteKpiForbidden) Code() int {
	return 403
}

func (o *DeleteKpiForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] deleteKpiForbidden %s", 403, payload)
}

func (o *DeleteKpiForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] deleteKpiForbidden %s", 403, payload)
}

func (o *DeleteKpiForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteKpiForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteKpiNotFound creates a DeleteKpiNotFound with default headers values
func NewDeleteKpiNotFound() *DeleteKpiNotFound {
	return &DeleteKpiNotFound{}
}

/*
DeleteKpiNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type DeleteKpiNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete kpi not found response has a 2xx status code
func (o *DeleteKpiNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete kpi not found response has a 3xx status code
func (o *DeleteKpiNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete kpi not found response has a 4xx status code
func (o *DeleteKpiNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete kpi not found response has a 5xx status code
func (o *DeleteKpiNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete kpi not found response a status code equal to that given
func (o *DeleteKpiNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete kpi not found response
func (o *DeleteKpiNotFound) Code() int {
	return 404
}

func (o *DeleteKpiNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] deleteKpiNotFound %s", 404, payload)
}

func (o *DeleteKpiNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] deleteKpiNotFound %s", 404, payload)
}

func (o *DeleteKpiNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteKpiNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteKpiDefault creates a DeleteKpiDefault with default headers values
func NewDeleteKpiDefault(code int) *DeleteKpiDefault {
	return &DeleteKpiDefault{
		_statusCode: code,
	}
}

/*
DeleteKpiDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DeleteKpiDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete kpi default response has a 2xx status code
func (o *DeleteKpiDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete kpi default response has a 3xx status code
func (o *DeleteKpiDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete kpi default response has a 4xx status code
func (o *DeleteKpiDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete kpi default response has a 5xx status code
func (o *DeleteKpiDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete kpi default response a status code equal to that given
func (o *DeleteKpiDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete kpi default response
func (o *DeleteKpiDefault) Code() int {
	return o._statusCode
}

func (o *DeleteKpiDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] deleteKpi default %s", o._statusCode, payload)
}

func (o *DeleteKpiDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] deleteKpi default %s", o._statusCode, payload)
}

func (o *DeleteKpiDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteKpiDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
