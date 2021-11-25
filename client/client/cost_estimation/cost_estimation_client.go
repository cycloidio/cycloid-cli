// Code generated by go-swagger; DO NOT EDIT.

package cost_estimation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new cost estimation API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cost estimation API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CostEstimateForms Generate a set of configs based on the forms inputs
*/
func (a *Client) CostEstimateForms(params *CostEstimateFormsParams, authInfo runtime.ClientAuthInfoWriter) (*CostEstimateFormsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCostEstimateFormsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "costEstimateForms",
		Method:             "POST",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/forms/estimate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CostEstimateFormsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CostEstimateFormsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CostEstimateFormsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CostEstimateTfPlan Estimate costs of a Terraform plan in JSON format.
*/
func (a *Client) CostEstimateTfPlan(params *CostEstimateTfPlanParams, authInfo runtime.ClientAuthInfoWriter) (*CostEstimateTfPlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCostEstimateTfPlanParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "costEstimateTfPlan",
		Method:             "POST",
		PathPattern:        "/organizations/{organization_canonical}/cost_estimation/tfplan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CostEstimateTfPlanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CostEstimateTfPlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CostEstimateTfPlanDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
