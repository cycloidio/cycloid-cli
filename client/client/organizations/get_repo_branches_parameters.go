// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetRepoBranchesParams creates a new GetRepoBranchesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRepoBranchesParams() *GetRepoBranchesParams {
	return &GetRepoBranchesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepoBranchesParamsWithTimeout creates a new GetRepoBranchesParams object
// with the ability to set a timeout on a request.
func NewGetRepoBranchesParamsWithTimeout(timeout time.Duration) *GetRepoBranchesParams {
	return &GetRepoBranchesParams{
		timeout: timeout,
	}
}

// NewGetRepoBranchesParamsWithContext creates a new GetRepoBranchesParams object
// with the ability to set a context for a request.
func NewGetRepoBranchesParamsWithContext(ctx context.Context) *GetRepoBranchesParams {
	return &GetRepoBranchesParams{
		Context: ctx,
	}
}

// NewGetRepoBranchesParamsWithHTTPClient creates a new GetRepoBranchesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRepoBranchesParamsWithHTTPClient(client *http.Client) *GetRepoBranchesParams {
	return &GetRepoBranchesParams{
		HTTPClient: client,
	}
}

/*
GetRepoBranchesParams contains all the parameters to send to the API endpoint

	for the get repo branches operation.

	Typically these are written to a http.Request.
*/
type GetRepoBranchesParams struct {

	/* CredentialCanonical.

	   A Credential canonical
	*/
	CredentialCanonical *string

	/* GitURL.

	   Git URL to repository
	*/
	GitURL string

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repo branches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepoBranchesParams) WithDefaults() *GetRepoBranchesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repo branches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepoBranchesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repo branches params
func (o *GetRepoBranchesParams) WithTimeout(timeout time.Duration) *GetRepoBranchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repo branches params
func (o *GetRepoBranchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repo branches params
func (o *GetRepoBranchesParams) WithContext(ctx context.Context) *GetRepoBranchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repo branches params
func (o *GetRepoBranchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repo branches params
func (o *GetRepoBranchesParams) WithHTTPClient(client *http.Client) *GetRepoBranchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repo branches params
func (o *GetRepoBranchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentialCanonical adds the credentialCanonical to the get repo branches params
func (o *GetRepoBranchesParams) WithCredentialCanonical(credentialCanonical *string) *GetRepoBranchesParams {
	o.SetCredentialCanonical(credentialCanonical)
	return o
}

// SetCredentialCanonical adds the credentialCanonical to the get repo branches params
func (o *GetRepoBranchesParams) SetCredentialCanonical(credentialCanonical *string) {
	o.CredentialCanonical = credentialCanonical
}

// WithGitURL adds the gitURL to the get repo branches params
func (o *GetRepoBranchesParams) WithGitURL(gitURL string) *GetRepoBranchesParams {
	o.SetGitURL(gitURL)
	return o
}

// SetGitURL adds the gitUrl to the get repo branches params
func (o *GetRepoBranchesParams) SetGitURL(gitURL string) {
	o.GitURL = gitURL
}

// WithOrganizationCanonical adds the organizationCanonical to the get repo branches params
func (o *GetRepoBranchesParams) WithOrganizationCanonical(organizationCanonical string) *GetRepoBranchesParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get repo branches params
func (o *GetRepoBranchesParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepoBranchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CredentialCanonical != nil {

		// query param credential_canonical
		var qrCredentialCanonical string

		if o.CredentialCanonical != nil {
			qrCredentialCanonical = *o.CredentialCanonical
		}
		qCredentialCanonical := qrCredentialCanonical
		if qCredentialCanonical != "" {

			if err := r.SetQueryParam("credential_canonical", qCredentialCanonical); err != nil {
				return err
			}
		}
	}

	// query param git_url
	qrGitURL := o.GitURL
	qGitURL := qrGitURL
	if qGitURL != "" {

		if err := r.SetQueryParam("git_url", qGitURL); err != nil {
			return err
		}
	}

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
