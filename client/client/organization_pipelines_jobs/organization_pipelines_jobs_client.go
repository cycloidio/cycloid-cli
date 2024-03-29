// Code generated by go-swagger; DO NOT EDIT.

package organization_pipelines_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new organization pipelines jobs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organization pipelines jobs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ClearTaskCache Clear task cache
*/
func (a *Client) ClearTaskCache(params *ClearTaskCacheParams, authInfo runtime.ClientAuthInfoWriter) (*ClearTaskCacheOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClearTaskCacheParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "clearTaskCache",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/tasks/{step_name}/cache",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ClearTaskCacheReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ClearTaskCacheOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ClearTaskCacheDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetJob Get the information of the job.
*/
func (a *Client) GetJob(params *GetJobParams, authInfo runtime.ClientAuthInfoWriter) (*GetJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJob",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJobReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetJobDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetJobs Get the jobs of the pipeline that the authenticated user has access to.
*/
func (a *Client) GetJobs(params *GetJobsParams, authInfo runtime.ClientAuthInfoWriter) (*GetJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJobsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJobs",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJobsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetJobsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetJobsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PauseJob Pause a job
*/
func (a *Client) PauseJob(params *PauseJobParams, authInfo runtime.ClientAuthInfoWriter) (*PauseJobNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPauseJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pauseJob",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/pause",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PauseJobReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PauseJobNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PauseJobDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UnpauseJob Unpause a job
*/
func (a *Client) UnpauseJob(params *UnpauseJobParams, authInfo runtime.ClientAuthInfoWriter) (*UnpauseJobNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnpauseJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "unpauseJob",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/unpause",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnpauseJobReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UnpauseJobNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UnpauseJobDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
