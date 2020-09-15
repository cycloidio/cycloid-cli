package middleware

import (
	"github.com/cycloidio/youdeploy-cli/client/client"
	"github.com/cycloidio/youdeploy-cli/client/models"
	strfmt "github.com/go-openapi/strfmt"
)

type Middleware interface {
	GetProject(org string, project string) (*models.Project, error)
	CreateProject(org, projectName, projectCanonical, env, pipelineTemplate, variables, description, cloudProvider, stackRef, usecase string, configRepo uint32) (*models.Project, error)
	UpdateProject(org, projectName, projectCanonical string, envs []string, description, cloudProvider, stackRef, owner string, configRepo uint32) (*models.Project, error)
	DeleteProjectEnv(org, project, env string) error
	DeleteProject(org, project string) error
	ListProjects(org string) ([]*models.ProjectsItem, error)

	UnpausePipeline(org string, project string, env string) error
	PausePipeline(org string, project string, env string) error
	UpdatePipeline(org string, project string, env string, pipeline string, variables string) (*models.Pipeline, error)
	CreatePipeline(org, project, env, pipeline, variables, usecase string) (*models.Pipeline, error)
	ClearTaskCachePipeline(org, project, env, job, task string) error
	DiffPipeline(org, project, env, pipeline, variables string) (*models.PipelineDiffs, error)
	GetPipelineJob(org, project, env, job string) (*models.Job, error)
	ListPipelineJobs(org, project, env string) ([]*models.Job, error)
	ListPipelineJobsBuilds(org, project, env, job string) ([]*models.Build, error)
	PausePipelineJob(org, project, env, job string) error
	UnpausePipelineJob(org, project, env, job string) error
	TriggerPipelineBuild(org, project, env, job string) error

	GetOrganization(org string) (*models.Organization, error)
	ListOrganizationWorkers(org string) ([]*models.Worker, error)
	ListOrganizations() ([]*models.OrganizationBasicInfo, error)

	ListCatalogRepositories(org string) ([]*models.ServiceCatalogSource, error)
	GetCatalogRepository(org string, catalogRepo uint32) (*models.ServiceCatalogSource, error)
	DeleteCatalogRepository(org string, catalogRepo uint32) error
	CreateCatalogRepository(org, name, url, branch string, cred uint32) (*models.ServiceCatalogSource, error)
	UpdateCatalogRepository(org string, catalogRepo uint32, name, url, branch string, cred uint32) (*models.ServiceCatalogSource, error)
	RefreshCatalogRepository(org string, catalogRepo uint32) (*models.ServiceCatalogSource, error)

	PushConfig(org string, project string, env string, configs map[string]strfmt.Base64) error
	ListConfigRepositories(org string) ([]*models.ConfigRepository, error)
	GetConfigRepository(org string, configRepo uint32) (*models.ConfigRepository, error)
	DeleteConfigRepository(org string, configRepo uint32) error
	CreateConfigRepository(org, name, url, branch string, setDefault bool, cred uint32) (*models.ConfigRepository, error)
	UpdateConfigRepository(org string, configRepo uint32, name, url, branch string, setDefault bool, cred uint32) (*models.ConfigRepository, error)

	ListExternalBackends(org string) ([]*models.ExternalBackend, error)
	DeleteExternalBackend(org string, externalBackend uint32) error
	CreateExternalBackends(org, project, env, purpose string, cred uint32, ebConfig models.ExternalBackendConfiguration) (*models.ExternalBackend, error)

	SendEvent(org, eventType, title, message, severity string, tags map[string]string, color string) error

	GetCredential(org string, cred uint32) (*models.Credential, error)
	DeleteCredential(org string, cred uint32) error
	ListCredentials(org, cType string) ([]*models.CredentialSimple, error)
	CreateCredential(org, name, cType string, rawCred *models.CredentialRaw, path, description string) error

	ListStacks(org string) ([]*models.ServiceCatalog, error)
	GetStack(org, ref string) (*models.ServiceCatalog, error)
}

type middleware struct {
	api *client.APIClient
}

func NewMiddleware(api *client.APIClient) Middleware {
	return &middleware{api: api}
}
