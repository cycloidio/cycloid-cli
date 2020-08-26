package middleware

import (
	"github.com/cycloidio/youdeploy-cli/client/client"
	"github.com/cycloidio/youdeploy-cli/client/models"
	strfmt "github.com/go-openapi/strfmt"
)

type Middleware interface {
	GetProject(org string, project string) (*models.Project, error)
	CreateProject(org, projectName, projectCanonical, env, pipelineTemplate, variables, description, cloudProvider, stackRef, usecase string, configRepo uint32) (*models.Project, error)
	UnpausePipeline(org string, project string, env string) error
	PausePipeline(org string, project string, env string) error
	UpdatePipeline(org string, project string, env string, pipeline string, variables string) (*models.Pipeline, error)
	PushConfig(org string, project string, env string, configs map[string]strfmt.Base64) error
}

type middleware struct {
	api *client.APIClient
}

func NewMiddleware(api *client.APIClient) Middleware {
	return &middleware{api: api}
}
