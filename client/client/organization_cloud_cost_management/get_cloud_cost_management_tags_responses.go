// Code generated by go-swagger; DO NOT EDIT.

package organization_cloud_cost_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// GetCloudCostManagementTagsReader is a Reader for the GetCloudCostManagementTags structure.
type GetCloudCostManagementTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudCostManagementTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCloudCostManagementTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetCloudCostManagementTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCloudCostManagementTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetCloudCostManagementTagsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCloudCostManagementTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCloudCostManagementTagsOK creates a GetCloudCostManagementTagsOK with default headers values
func NewGetCloudCostManagementTagsOK() *GetCloudCostManagementTagsOK {
	return &GetCloudCostManagementTagsOK{}
}

/*
GetCloudCostManagementTagsOK describes a response with status code 200, with default header values.

List of Cloud Cost Management records' tags.
*/
type GetCloudCostManagementTagsOK struct {
	Payload *GetCloudCostManagementTagsOKBody
}

// IsSuccess returns true when this get cloud cost management tags o k response has a 2xx status code
func (o *GetCloudCostManagementTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cloud cost management tags o k response has a 3xx status code
func (o *GetCloudCostManagementTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud cost management tags o k response has a 4xx status code
func (o *GetCloudCostManagementTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cloud cost management tags o k response has a 5xx status code
func (o *GetCloudCostManagementTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud cost management tags o k response a status code equal to that given
func (o *GetCloudCostManagementTagsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cloud cost management tags o k response
func (o *GetCloudCostManagementTagsOK) Code() int {
	return 200
}

func (o *GetCloudCostManagementTagsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tags][%d] getCloudCostManagementTagsOK %s", 200, payload)
}

func (o *GetCloudCostManagementTagsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tags][%d] getCloudCostManagementTagsOK %s", 200, payload)
}

func (o *GetCloudCostManagementTagsOK) GetPayload() *GetCloudCostManagementTagsOKBody {
	return o.Payload
}

func (o *GetCloudCostManagementTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCloudCostManagementTagsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudCostManagementTagsForbidden creates a GetCloudCostManagementTagsForbidden with default headers values
func NewGetCloudCostManagementTagsForbidden() *GetCloudCostManagementTagsForbidden {
	return &GetCloudCostManagementTagsForbidden{}
}

/*
GetCloudCostManagementTagsForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetCloudCostManagementTagsForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get cloud cost management tags forbidden response has a 2xx status code
func (o *GetCloudCostManagementTagsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud cost management tags forbidden response has a 3xx status code
func (o *GetCloudCostManagementTagsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud cost management tags forbidden response has a 4xx status code
func (o *GetCloudCostManagementTagsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cloud cost management tags forbidden response has a 5xx status code
func (o *GetCloudCostManagementTagsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud cost management tags forbidden response a status code equal to that given
func (o *GetCloudCostManagementTagsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get cloud cost management tags forbidden response
func (o *GetCloudCostManagementTagsForbidden) Code() int {
	return 403
}

func (o *GetCloudCostManagementTagsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tags][%d] getCloudCostManagementTagsForbidden %s", 403, payload)
}

func (o *GetCloudCostManagementTagsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tags][%d] getCloudCostManagementTagsForbidden %s", 403, payload)
}

func (o *GetCloudCostManagementTagsForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetCloudCostManagementTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCloudCostManagementTagsNotFound creates a GetCloudCostManagementTagsNotFound with default headers values
func NewGetCloudCostManagementTagsNotFound() *GetCloudCostManagementTagsNotFound {
	return &GetCloudCostManagementTagsNotFound{}
}

/*
GetCloudCostManagementTagsNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetCloudCostManagementTagsNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get cloud cost management tags not found response has a 2xx status code
func (o *GetCloudCostManagementTagsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud cost management tags not found response has a 3xx status code
func (o *GetCloudCostManagementTagsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud cost management tags not found response has a 4xx status code
func (o *GetCloudCostManagementTagsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cloud cost management tags not found response has a 5xx status code
func (o *GetCloudCostManagementTagsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud cost management tags not found response a status code equal to that given
func (o *GetCloudCostManagementTagsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get cloud cost management tags not found response
func (o *GetCloudCostManagementTagsNotFound) Code() int {
	return 404
}

func (o *GetCloudCostManagementTagsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tags][%d] getCloudCostManagementTagsNotFound %s", 404, payload)
}

func (o *GetCloudCostManagementTagsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tags][%d] getCloudCostManagementTagsNotFound %s", 404, payload)
}

func (o *GetCloudCostManagementTagsNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetCloudCostManagementTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCloudCostManagementTagsUnprocessableEntity creates a GetCloudCostManagementTagsUnprocessableEntity with default headers values
func NewGetCloudCostManagementTagsUnprocessableEntity() *GetCloudCostManagementTagsUnprocessableEntity {
	return &GetCloudCostManagementTagsUnprocessableEntity{}
}

/*
GetCloudCostManagementTagsUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetCloudCostManagementTagsUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get cloud cost management tags unprocessable entity response has a 2xx status code
func (o *GetCloudCostManagementTagsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud cost management tags unprocessable entity response has a 3xx status code
func (o *GetCloudCostManagementTagsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud cost management tags unprocessable entity response has a 4xx status code
func (o *GetCloudCostManagementTagsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cloud cost management tags unprocessable entity response has a 5xx status code
func (o *GetCloudCostManagementTagsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud cost management tags unprocessable entity response a status code equal to that given
func (o *GetCloudCostManagementTagsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get cloud cost management tags unprocessable entity response
func (o *GetCloudCostManagementTagsUnprocessableEntity) Code() int {
	return 422
}

func (o *GetCloudCostManagementTagsUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tags][%d] getCloudCostManagementTagsUnprocessableEntity %s", 422, payload)
}

func (o *GetCloudCostManagementTagsUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tags][%d] getCloudCostManagementTagsUnprocessableEntity %s", 422, payload)
}

func (o *GetCloudCostManagementTagsUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetCloudCostManagementTagsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCloudCostManagementTagsDefault creates a GetCloudCostManagementTagsDefault with default headers values
func NewGetCloudCostManagementTagsDefault(code int) *GetCloudCostManagementTagsDefault {
	return &GetCloudCostManagementTagsDefault{
		_statusCode: code,
	}
}

/*
GetCloudCostManagementTagsDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetCloudCostManagementTagsDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get cloud cost management tags default response has a 2xx status code
func (o *GetCloudCostManagementTagsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cloud cost management tags default response has a 3xx status code
func (o *GetCloudCostManagementTagsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cloud cost management tags default response has a 4xx status code
func (o *GetCloudCostManagementTagsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cloud cost management tags default response has a 5xx status code
func (o *GetCloudCostManagementTagsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cloud cost management tags default response a status code equal to that given
func (o *GetCloudCostManagementTagsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get cloud cost management tags default response
func (o *GetCloudCostManagementTagsDefault) Code() int {
	return o._statusCode
}

func (o *GetCloudCostManagementTagsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tags][%d] getCloudCostManagementTags default %s", o._statusCode, payload)
}

func (o *GetCloudCostManagementTagsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tags][%d] getCloudCostManagementTags default %s", o._statusCode, payload)
}

func (o *GetCloudCostManagementTagsDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetCloudCostManagementTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
GetCloudCostManagementTagsOKBody get cloud cost management tags o k body
swagger:model GetCloudCostManagementTagsOKBody
*/
type GetCloudCostManagementTagsOKBody struct {

	// data
	// Required: true
	Data []string `json:"data"`
}

// Validate validates this get cloud cost management tags o k body
func (o *GetCloudCostManagementTagsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCloudCostManagementTagsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getCloudCostManagementTagsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get cloud cost management tags o k body based on context it is used
func (o *GetCloudCostManagementTagsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCloudCostManagementTagsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCloudCostManagementTagsOKBody) UnmarshalBinary(b []byte) error {
	var res GetCloudCostManagementTagsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
