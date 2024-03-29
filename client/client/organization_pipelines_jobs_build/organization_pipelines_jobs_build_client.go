// Code generated by go-swagger; DO NOT EDIT.

package organization_pipelines_jobs_build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new organization pipelines jobs build API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organization pipelines jobs build API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AbortBuild Abort a specific build.
*/
func (a *Client) AbortBuild(params *AbortBuildParams, authInfo runtime.ClientAuthInfoWriter) (*AbortBuildNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAbortBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "abortBuild",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/abort",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AbortBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AbortBuildNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AbortBuildDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateBuild Create a new build for the job
*/
func (a *Client) CreateBuild(params *CreateBuildParams, authInfo runtime.ClientAuthInfoWriter) (*CreateBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createBuild",
		Method:             "POST",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateBuildOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateBuildDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBuild Get the information of the build.
*/
func (a *Client) GetBuild(params *GetBuildParams, authInfo runtime.ClientAuthInfoWriter) (*GetBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBuild",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBuildOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBuildDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBuildPlan Get the plan of the build.
*/
func (a *Client) GetBuildPlan(params *GetBuildPlanParams, authInfo runtime.ClientAuthInfoWriter) (*GetBuildPlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBuildPlanParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBuildPlan",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/plan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBuildPlanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBuildPlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBuildPlanDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBuildPreparation Get the preparation of the Build.
*/
func (a *Client) GetBuildPreparation(params *GetBuildPreparationParams, authInfo runtime.ClientAuthInfoWriter) (*GetBuildPreparationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBuildPreparationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBuildPreparation",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/preparation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBuildPreparationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBuildPreparationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBuildPreparationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBuildResources Get the resources of the build.
*/
func (a *Client) GetBuildResources(params *GetBuildResourcesParams, authInfo runtime.ClientAuthInfoWriter) (*GetBuildResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBuildResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBuildResources",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/resources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBuildResourcesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBuildResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBuildResourcesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBuilds Get the pipeline job's builds that the authenticated user has access to.
*/
func (a *Client) GetBuilds(params *GetBuildsParams, authInfo runtime.ClientAuthInfoWriter) (*GetBuildsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBuildsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBuilds",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBuildsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBuildsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBuildsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RerunBuild Reruns a specific build.
*/
func (a *Client) RerunBuild(params *RerunBuildParams, authInfo runtime.ClientAuthInfoWriter) (*RerunBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRerunBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "rerunBuild",
		Method:             "POST",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RerunBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RerunBuildOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RerunBuildDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
