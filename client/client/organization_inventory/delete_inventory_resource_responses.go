// Code generated by go-swagger; DO NOT EDIT.

package organization_inventory

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

// DeleteInventoryResourceReader is a Reader for the DeleteInventoryResource structure.
type DeleteInventoryResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInventoryResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteInventoryResourceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteInventoryResourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteInventoryResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteInventoryResourceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteInventoryResourceNoContent creates a DeleteInventoryResourceNoContent with default headers values
func NewDeleteInventoryResourceNoContent() *DeleteInventoryResourceNoContent {
	return &DeleteInventoryResourceNoContent{}
}

/*
DeleteInventoryResourceNoContent describes a response with status code 204, with default header values.

Inventory Resource deleted
*/
type DeleteInventoryResourceNoContent struct {
}

// IsSuccess returns true when this delete inventory resource no content response has a 2xx status code
func (o *DeleteInventoryResourceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete inventory resource no content response has a 3xx status code
func (o *DeleteInventoryResourceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete inventory resource no content response has a 4xx status code
func (o *DeleteInventoryResourceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete inventory resource no content response has a 5xx status code
func (o *DeleteInventoryResourceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete inventory resource no content response a status code equal to that given
func (o *DeleteInventoryResourceNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete inventory resource no content response
func (o *DeleteInventoryResourceNoContent) Code() int {
	return 204
}

func (o *DeleteInventoryResourceNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/inventory/resources/{inventory_resource_id}][%d] deleteInventoryResourceNoContent", 204)
}

func (o *DeleteInventoryResourceNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/inventory/resources/{inventory_resource_id}][%d] deleteInventoryResourceNoContent", 204)
}

func (o *DeleteInventoryResourceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInventoryResourceForbidden creates a DeleteInventoryResourceForbidden with default headers values
func NewDeleteInventoryResourceForbidden() *DeleteInventoryResourceForbidden {
	return &DeleteInventoryResourceForbidden{}
}

/*
DeleteInventoryResourceForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type DeleteInventoryResourceForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete inventory resource forbidden response has a 2xx status code
func (o *DeleteInventoryResourceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete inventory resource forbidden response has a 3xx status code
func (o *DeleteInventoryResourceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete inventory resource forbidden response has a 4xx status code
func (o *DeleteInventoryResourceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete inventory resource forbidden response has a 5xx status code
func (o *DeleteInventoryResourceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete inventory resource forbidden response a status code equal to that given
func (o *DeleteInventoryResourceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete inventory resource forbidden response
func (o *DeleteInventoryResourceForbidden) Code() int {
	return 403
}

func (o *DeleteInventoryResourceForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/inventory/resources/{inventory_resource_id}][%d] deleteInventoryResourceForbidden %s", 403, payload)
}

func (o *DeleteInventoryResourceForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/inventory/resources/{inventory_resource_id}][%d] deleteInventoryResourceForbidden %s", 403, payload)
}

func (o *DeleteInventoryResourceForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteInventoryResourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteInventoryResourceNotFound creates a DeleteInventoryResourceNotFound with default headers values
func NewDeleteInventoryResourceNotFound() *DeleteInventoryResourceNotFound {
	return &DeleteInventoryResourceNotFound{}
}

/*
DeleteInventoryResourceNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type DeleteInventoryResourceNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete inventory resource not found response has a 2xx status code
func (o *DeleteInventoryResourceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete inventory resource not found response has a 3xx status code
func (o *DeleteInventoryResourceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete inventory resource not found response has a 4xx status code
func (o *DeleteInventoryResourceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete inventory resource not found response has a 5xx status code
func (o *DeleteInventoryResourceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete inventory resource not found response a status code equal to that given
func (o *DeleteInventoryResourceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete inventory resource not found response
func (o *DeleteInventoryResourceNotFound) Code() int {
	return 404
}

func (o *DeleteInventoryResourceNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/inventory/resources/{inventory_resource_id}][%d] deleteInventoryResourceNotFound %s", 404, payload)
}

func (o *DeleteInventoryResourceNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/inventory/resources/{inventory_resource_id}][%d] deleteInventoryResourceNotFound %s", 404, payload)
}

func (o *DeleteInventoryResourceNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteInventoryResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteInventoryResourceDefault creates a DeleteInventoryResourceDefault with default headers values
func NewDeleteInventoryResourceDefault(code int) *DeleteInventoryResourceDefault {
	return &DeleteInventoryResourceDefault{
		_statusCode: code,
	}
}

/*
DeleteInventoryResourceDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DeleteInventoryResourceDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete inventory resource default response has a 2xx status code
func (o *DeleteInventoryResourceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete inventory resource default response has a 3xx status code
func (o *DeleteInventoryResourceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete inventory resource default response has a 4xx status code
func (o *DeleteInventoryResourceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete inventory resource default response has a 5xx status code
func (o *DeleteInventoryResourceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete inventory resource default response a status code equal to that given
func (o *DeleteInventoryResourceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete inventory resource default response
func (o *DeleteInventoryResourceDefault) Code() int {
	return o._statusCode
}

func (o *DeleteInventoryResourceDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/inventory/resources/{inventory_resource_id}][%d] deleteInventoryResource default %s", o._statusCode, payload)
}

func (o *DeleteInventoryResourceDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/inventory/resources/{inventory_resource_id}][%d] deleteInventoryResource default %s", o._statusCode, payload)
}

func (o *DeleteInventoryResourceDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteInventoryResourceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
