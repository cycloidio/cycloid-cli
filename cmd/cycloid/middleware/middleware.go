package middleware

import (
	"github.com/cycloidio/youdeploy-cli/client/client"
	"github.com/cycloidio/youdeploy-cli/client/models"
)

type Middleware interface {
	GetProject(org string, project string) (*models.Project, error)
	UnpausePipeline(org string, project string, env string) error
}

type middleware struct {
	api *client.APIClient
}

func NewMiddleware(api *client.APIClient) Middleware {
	return &middleware{api: api}
}
