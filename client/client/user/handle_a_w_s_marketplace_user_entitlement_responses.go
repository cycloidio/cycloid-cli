// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// HandleAWSMarketplaceUserEntitlementReader is a Reader for the HandleAWSMarketplaceUserEntitlement structure.
type HandleAWSMarketplaceUserEntitlementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HandleAWSMarketplaceUserEntitlementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 301:
		result := NewHandleAWSMarketplaceUserEntitlementMovedPermanently()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHandleAWSMarketplaceUserEntitlementDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHandleAWSMarketplaceUserEntitlementMovedPermanently creates a HandleAWSMarketplaceUserEntitlementMovedPermanently with default headers values
func NewHandleAWSMarketplaceUserEntitlementMovedPermanently() *HandleAWSMarketplaceUserEntitlementMovedPermanently {
	return &HandleAWSMarketplaceUserEntitlementMovedPermanently{}
}

/*HandleAWSMarketplaceUserEntitlementMovedPermanently handles this case with default header values.

The user is redirected based on his account state.
*/
type HandleAWSMarketplaceUserEntitlementMovedPermanently struct {
	Location string
}

func (o *HandleAWSMarketplaceUserEntitlementMovedPermanently) Error() string {
	return fmt.Sprintf("[POST /user/aws_marketplace/entitlement][%d] handleAWSMarketplaceUserEntitlementMovedPermanently ", 301)
}

func (o *HandleAWSMarketplaceUserEntitlementMovedPermanently) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}

// NewHandleAWSMarketplaceUserEntitlementDefault creates a HandleAWSMarketplaceUserEntitlementDefault with default headers values
func NewHandleAWSMarketplaceUserEntitlementDefault(code int) *HandleAWSMarketplaceUserEntitlementDefault {
	return &HandleAWSMarketplaceUserEntitlementDefault{
		_statusCode: code,
	}
}

/*HandleAWSMarketplaceUserEntitlementDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type HandleAWSMarketplaceUserEntitlementDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the handle a w s marketplace user entitlement default response
func (o *HandleAWSMarketplaceUserEntitlementDefault) Code() int {
	return o._statusCode
}

func (o *HandleAWSMarketplaceUserEntitlementDefault) Error() string {
	return fmt.Sprintf("[POST /user/aws_marketplace/entitlement][%d] handleAWSMarketplaceUserEntitlement default  %+v", o._statusCode, o.Payload)
}

func (o *HandleAWSMarketplaceUserEntitlementDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *HandleAWSMarketplaceUserEntitlementDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertUint64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "uint64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
