// Code generated by go-swagger; DO NOT EDIT.

package organization_infrastructure

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

// GetRunningInfraAWSElasticacheClustersReader is a Reader for the GetRunningInfraAWSElasticacheClusters structure.
type GetRunningInfraAWSElasticacheClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunningInfraAWSElasticacheClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunningInfraAWSElasticacheClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRunningInfraAWSElasticacheClustersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunningInfraAWSElasticacheClustersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetRunningInfraAWSElasticacheClustersUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRunningInfraAWSElasticacheClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRunningInfraAWSElasticacheClustersOK creates a GetRunningInfraAWSElasticacheClustersOK with default headers values
func NewGetRunningInfraAWSElasticacheClustersOK() *GetRunningInfraAWSElasticacheClustersOK {
	return &GetRunningInfraAWSElasticacheClustersOK{}
}

/*
GetRunningInfraAWSElasticacheClustersOK describes a response with status code 200, with default header values.

The list of AWS elasticache clusters which matches the scope specified by the filter.
*/
type GetRunningInfraAWSElasticacheClustersOK struct {
	Payload *GetRunningInfraAWSElasticacheClustersOKBody
}

// IsSuccess returns true when this get running infra a w s elasticache clusters o k response has a 2xx status code
func (o *GetRunningInfraAWSElasticacheClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get running infra a w s elasticache clusters o k response has a 3xx status code
func (o *GetRunningInfraAWSElasticacheClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get running infra a w s elasticache clusters o k response has a 4xx status code
func (o *GetRunningInfraAWSElasticacheClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get running infra a w s elasticache clusters o k response has a 5xx status code
func (o *GetRunningInfraAWSElasticacheClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get running infra a w s elasticache clusters o k response a status code equal to that given
func (o *GetRunningInfraAWSElasticacheClustersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get running infra a w s elasticache clusters o k response
func (o *GetRunningInfraAWSElasticacheClustersOK) Code() int {
	return 200
}

func (o *GetRunningInfraAWSElasticacheClustersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/infrastructures/aws/elasticache_clusters][%d] getRunningInfraAWSElasticacheClustersOK %s", 200, payload)
}

func (o *GetRunningInfraAWSElasticacheClustersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/infrastructures/aws/elasticache_clusters][%d] getRunningInfraAWSElasticacheClustersOK %s", 200, payload)
}

func (o *GetRunningInfraAWSElasticacheClustersOK) GetPayload() *GetRunningInfraAWSElasticacheClustersOKBody {
	return o.Payload
}

func (o *GetRunningInfraAWSElasticacheClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetRunningInfraAWSElasticacheClustersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunningInfraAWSElasticacheClustersForbidden creates a GetRunningInfraAWSElasticacheClustersForbidden with default headers values
func NewGetRunningInfraAWSElasticacheClustersForbidden() *GetRunningInfraAWSElasticacheClustersForbidden {
	return &GetRunningInfraAWSElasticacheClustersForbidden{}
}

/*
GetRunningInfraAWSElasticacheClustersForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetRunningInfraAWSElasticacheClustersForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get running infra a w s elasticache clusters forbidden response has a 2xx status code
func (o *GetRunningInfraAWSElasticacheClustersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get running infra a w s elasticache clusters forbidden response has a 3xx status code
func (o *GetRunningInfraAWSElasticacheClustersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get running infra a w s elasticache clusters forbidden response has a 4xx status code
func (o *GetRunningInfraAWSElasticacheClustersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get running infra a w s elasticache clusters forbidden response has a 5xx status code
func (o *GetRunningInfraAWSElasticacheClustersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get running infra a w s elasticache clusters forbidden response a status code equal to that given
func (o *GetRunningInfraAWSElasticacheClustersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get running infra a w s elasticache clusters forbidden response
func (o *GetRunningInfraAWSElasticacheClustersForbidden) Code() int {
	return 403
}

func (o *GetRunningInfraAWSElasticacheClustersForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/infrastructures/aws/elasticache_clusters][%d] getRunningInfraAWSElasticacheClustersForbidden %s", 403, payload)
}

func (o *GetRunningInfraAWSElasticacheClustersForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/infrastructures/aws/elasticache_clusters][%d] getRunningInfraAWSElasticacheClustersForbidden %s", 403, payload)
}

func (o *GetRunningInfraAWSElasticacheClustersForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetRunningInfraAWSElasticacheClustersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRunningInfraAWSElasticacheClustersNotFound creates a GetRunningInfraAWSElasticacheClustersNotFound with default headers values
func NewGetRunningInfraAWSElasticacheClustersNotFound() *GetRunningInfraAWSElasticacheClustersNotFound {
	return &GetRunningInfraAWSElasticacheClustersNotFound{}
}

/*
GetRunningInfraAWSElasticacheClustersNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetRunningInfraAWSElasticacheClustersNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get running infra a w s elasticache clusters not found response has a 2xx status code
func (o *GetRunningInfraAWSElasticacheClustersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get running infra a w s elasticache clusters not found response has a 3xx status code
func (o *GetRunningInfraAWSElasticacheClustersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get running infra a w s elasticache clusters not found response has a 4xx status code
func (o *GetRunningInfraAWSElasticacheClustersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get running infra a w s elasticache clusters not found response has a 5xx status code
func (o *GetRunningInfraAWSElasticacheClustersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get running infra a w s elasticache clusters not found response a status code equal to that given
func (o *GetRunningInfraAWSElasticacheClustersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get running infra a w s elasticache clusters not found response
func (o *GetRunningInfraAWSElasticacheClustersNotFound) Code() int {
	return 404
}

func (o *GetRunningInfraAWSElasticacheClustersNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/infrastructures/aws/elasticache_clusters][%d] getRunningInfraAWSElasticacheClustersNotFound %s", 404, payload)
}

func (o *GetRunningInfraAWSElasticacheClustersNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/infrastructures/aws/elasticache_clusters][%d] getRunningInfraAWSElasticacheClustersNotFound %s", 404, payload)
}

func (o *GetRunningInfraAWSElasticacheClustersNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetRunningInfraAWSElasticacheClustersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRunningInfraAWSElasticacheClustersUnprocessableEntity creates a GetRunningInfraAWSElasticacheClustersUnprocessableEntity with default headers values
func NewGetRunningInfraAWSElasticacheClustersUnprocessableEntity() *GetRunningInfraAWSElasticacheClustersUnprocessableEntity {
	return &GetRunningInfraAWSElasticacheClustersUnprocessableEntity{}
}

/*
GetRunningInfraAWSElasticacheClustersUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetRunningInfraAWSElasticacheClustersUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get running infra a w s elasticache clusters unprocessable entity response has a 2xx status code
func (o *GetRunningInfraAWSElasticacheClustersUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get running infra a w s elasticache clusters unprocessable entity response has a 3xx status code
func (o *GetRunningInfraAWSElasticacheClustersUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get running infra a w s elasticache clusters unprocessable entity response has a 4xx status code
func (o *GetRunningInfraAWSElasticacheClustersUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get running infra a w s elasticache clusters unprocessable entity response has a 5xx status code
func (o *GetRunningInfraAWSElasticacheClustersUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get running infra a w s elasticache clusters unprocessable entity response a status code equal to that given
func (o *GetRunningInfraAWSElasticacheClustersUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get running infra a w s elasticache clusters unprocessable entity response
func (o *GetRunningInfraAWSElasticacheClustersUnprocessableEntity) Code() int {
	return 422
}

func (o *GetRunningInfraAWSElasticacheClustersUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/infrastructures/aws/elasticache_clusters][%d] getRunningInfraAWSElasticacheClustersUnprocessableEntity %s", 422, payload)
}

func (o *GetRunningInfraAWSElasticacheClustersUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/infrastructures/aws/elasticache_clusters][%d] getRunningInfraAWSElasticacheClustersUnprocessableEntity %s", 422, payload)
}

func (o *GetRunningInfraAWSElasticacheClustersUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetRunningInfraAWSElasticacheClustersUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRunningInfraAWSElasticacheClustersDefault creates a GetRunningInfraAWSElasticacheClustersDefault with default headers values
func NewGetRunningInfraAWSElasticacheClustersDefault(code int) *GetRunningInfraAWSElasticacheClustersDefault {
	return &GetRunningInfraAWSElasticacheClustersDefault{
		_statusCode: code,
	}
}

/*
GetRunningInfraAWSElasticacheClustersDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetRunningInfraAWSElasticacheClustersDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get running infra a w s elasticache clusters default response has a 2xx status code
func (o *GetRunningInfraAWSElasticacheClustersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get running infra a w s elasticache clusters default response has a 3xx status code
func (o *GetRunningInfraAWSElasticacheClustersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get running infra a w s elasticache clusters default response has a 4xx status code
func (o *GetRunningInfraAWSElasticacheClustersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get running infra a w s elasticache clusters default response has a 5xx status code
func (o *GetRunningInfraAWSElasticacheClustersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get running infra a w s elasticache clusters default response a status code equal to that given
func (o *GetRunningInfraAWSElasticacheClustersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get running infra a w s elasticache clusters default response
func (o *GetRunningInfraAWSElasticacheClustersDefault) Code() int {
	return o._statusCode
}

func (o *GetRunningInfraAWSElasticacheClustersDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/infrastructures/aws/elasticache_clusters][%d] getRunningInfraAWSElasticacheClusters default %s", o._statusCode, payload)
}

func (o *GetRunningInfraAWSElasticacheClustersDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/infrastructures/aws/elasticache_clusters][%d] getRunningInfraAWSElasticacheClusters default %s", o._statusCode, payload)
}

func (o *GetRunningInfraAWSElasticacheClustersDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetRunningInfraAWSElasticacheClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetRunningInfraAWSElasticacheClustersOKBody get running infra a w s elasticache clusters o k body
swagger:model GetRunningInfraAWSElasticacheClustersOKBody
*/
type GetRunningInfraAWSElasticacheClustersOKBody struct {

	// data
	// Required: true
	Data []models.AWSInfrastructureResourceElasticacheCluster `json:"data"`

	// pagination
	Pagination *models.PaginationAWS `json:"pagination,omitempty"`
}

// Validate validates this get running infra a w s elasticache clusters o k body
func (o *GetRunningInfraAWSElasticacheClustersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRunningInfraAWSElasticacheClustersOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getRunningInfraAWSElasticacheClustersOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

func (o *GetRunningInfraAWSElasticacheClustersOKBody) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(o.Pagination) { // not required
		return nil
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getRunningInfraAWSElasticacheClustersOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getRunningInfraAWSElasticacheClustersOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get running infra a w s elasticache clusters o k body based on the context it is used
func (o *GetRunningInfraAWSElasticacheClustersOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRunningInfraAWSElasticacheClustersOKBody) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if o.Pagination != nil {

		if swag.IsZero(o.Pagination) { // not required
			return nil
		}

		if err := o.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getRunningInfraAWSElasticacheClustersOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getRunningInfraAWSElasticacheClustersOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRunningInfraAWSElasticacheClustersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRunningInfraAWSElasticacheClustersOKBody) UnmarshalBinary(b []byte) error {
	var res GetRunningInfraAWSElasticacheClustersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
