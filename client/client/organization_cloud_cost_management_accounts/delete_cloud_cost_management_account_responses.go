// Code generated by go-swagger; DO NOT EDIT.

package organization_cloud_cost_management_accounts

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

// DeleteCloudCostManagementAccountReader is a Reader for the DeleteCloudCostManagementAccount structure.
type DeleteCloudCostManagementAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCloudCostManagementAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCloudCostManagementAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteCloudCostManagementAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCloudCostManagementAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteCloudCostManagementAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCloudCostManagementAccountNoContent creates a DeleteCloudCostManagementAccountNoContent with default headers values
func NewDeleteCloudCostManagementAccountNoContent() *DeleteCloudCostManagementAccountNoContent {
	return &DeleteCloudCostManagementAccountNoContent{}
}

/*
DeleteCloudCostManagementAccountNoContent describes a response with status code 204, with default header values.

CloudCostManagementAccount has been deleted.
*/
type DeleteCloudCostManagementAccountNoContent struct {
}

// IsSuccess returns true when this delete cloud cost management account no content response has a 2xx status code
func (o *DeleteCloudCostManagementAccountNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete cloud cost management account no content response has a 3xx status code
func (o *DeleteCloudCostManagementAccountNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cloud cost management account no content response has a 4xx status code
func (o *DeleteCloudCostManagementAccountNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete cloud cost management account no content response has a 5xx status code
func (o *DeleteCloudCostManagementAccountNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cloud cost management account no content response a status code equal to that given
func (o *DeleteCloudCostManagementAccountNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete cloud cost management account no content response
func (o *DeleteCloudCostManagementAccountNoContent) Code() int {
	return 204
}

func (o *DeleteCloudCostManagementAccountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/cloud_cost_management/accounts/{cloud_cost_management_account_canonical}][%d] deleteCloudCostManagementAccountNoContent", 204)
}

func (o *DeleteCloudCostManagementAccountNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/cloud_cost_management/accounts/{cloud_cost_management_account_canonical}][%d] deleteCloudCostManagementAccountNoContent", 204)
}

func (o *DeleteCloudCostManagementAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCloudCostManagementAccountForbidden creates a DeleteCloudCostManagementAccountForbidden with default headers values
func NewDeleteCloudCostManagementAccountForbidden() *DeleteCloudCostManagementAccountForbidden {
	return &DeleteCloudCostManagementAccountForbidden{}
}

/*
DeleteCloudCostManagementAccountForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type DeleteCloudCostManagementAccountForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete cloud cost management account forbidden response has a 2xx status code
func (o *DeleteCloudCostManagementAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cloud cost management account forbidden response has a 3xx status code
func (o *DeleteCloudCostManagementAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cloud cost management account forbidden response has a 4xx status code
func (o *DeleteCloudCostManagementAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cloud cost management account forbidden response has a 5xx status code
func (o *DeleteCloudCostManagementAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cloud cost management account forbidden response a status code equal to that given
func (o *DeleteCloudCostManagementAccountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete cloud cost management account forbidden response
func (o *DeleteCloudCostManagementAccountForbidden) Code() int {
	return 403
}

func (o *DeleteCloudCostManagementAccountForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/cloud_cost_management/accounts/{cloud_cost_management_account_canonical}][%d] deleteCloudCostManagementAccountForbidden %s", 403, payload)
}

func (o *DeleteCloudCostManagementAccountForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/cloud_cost_management/accounts/{cloud_cost_management_account_canonical}][%d] deleteCloudCostManagementAccountForbidden %s", 403, payload)
}

func (o *DeleteCloudCostManagementAccountForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteCloudCostManagementAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCloudCostManagementAccountNotFound creates a DeleteCloudCostManagementAccountNotFound with default headers values
func NewDeleteCloudCostManagementAccountNotFound() *DeleteCloudCostManagementAccountNotFound {
	return &DeleteCloudCostManagementAccountNotFound{}
}

/*
DeleteCloudCostManagementAccountNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type DeleteCloudCostManagementAccountNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete cloud cost management account not found response has a 2xx status code
func (o *DeleteCloudCostManagementAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cloud cost management account not found response has a 3xx status code
func (o *DeleteCloudCostManagementAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cloud cost management account not found response has a 4xx status code
func (o *DeleteCloudCostManagementAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cloud cost management account not found response has a 5xx status code
func (o *DeleteCloudCostManagementAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cloud cost management account not found response a status code equal to that given
func (o *DeleteCloudCostManagementAccountNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete cloud cost management account not found response
func (o *DeleteCloudCostManagementAccountNotFound) Code() int {
	return 404
}

func (o *DeleteCloudCostManagementAccountNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/cloud_cost_management/accounts/{cloud_cost_management_account_canonical}][%d] deleteCloudCostManagementAccountNotFound %s", 404, payload)
}

func (o *DeleteCloudCostManagementAccountNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/cloud_cost_management/accounts/{cloud_cost_management_account_canonical}][%d] deleteCloudCostManagementAccountNotFound %s", 404, payload)
}

func (o *DeleteCloudCostManagementAccountNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteCloudCostManagementAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCloudCostManagementAccountDefault creates a DeleteCloudCostManagementAccountDefault with default headers values
func NewDeleteCloudCostManagementAccountDefault(code int) *DeleteCloudCostManagementAccountDefault {
	return &DeleteCloudCostManagementAccountDefault{
		_statusCode: code,
	}
}

/*
DeleteCloudCostManagementAccountDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DeleteCloudCostManagementAccountDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete cloud cost management account default response has a 2xx status code
func (o *DeleteCloudCostManagementAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete cloud cost management account default response has a 3xx status code
func (o *DeleteCloudCostManagementAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete cloud cost management account default response has a 4xx status code
func (o *DeleteCloudCostManagementAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete cloud cost management account default response has a 5xx status code
func (o *DeleteCloudCostManagementAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete cloud cost management account default response a status code equal to that given
func (o *DeleteCloudCostManagementAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete cloud cost management account default response
func (o *DeleteCloudCostManagementAccountDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCloudCostManagementAccountDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/cloud_cost_management/accounts/{cloud_cost_management_account_canonical}][%d] deleteCloudCostManagementAccount default %s", o._statusCode, payload)
}

func (o *DeleteCloudCostManagementAccountDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/cloud_cost_management/accounts/{cloud_cost_management_account_canonical}][%d] deleteCloudCostManagementAccount default %s", o._statusCode, payload)
}

func (o *DeleteCloudCostManagementAccountDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteCloudCostManagementAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
