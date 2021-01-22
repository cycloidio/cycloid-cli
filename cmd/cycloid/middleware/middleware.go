package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client"
	"github.com/cycloidio/cycloid-cli/client/models"
	strfmt "github.com/go-openapi/strfmt"
)

type Middleware interface {
	GetAppVersion() (*models.AppVersion, error)
	GetStatus() (*models.GeneralStatus, error)

	CreateCatalogRepository(org, name, url, branch, cred string) (*models.ServiceCatalogSource, error)
	DeleteCatalogRepository(org, catalogRepo string) error
	GetCatalogRepository(org, catalogRepo string) (*models.ServiceCatalogSource, error)
	ListCatalogRepositories(org string) ([]*models.ServiceCatalogSource, error)
	RefreshCatalogRepository(org, catalogRepo string) (*models.ServiceCatalogSource, error)
	UpdateCatalogRepository(org, name, url, branch, catalogRepo, cred string) (*models.ServiceCatalogSource, error)

	CreateConfigRepository(org, name, url, branch, cred string, setDefault bool) (*models.ConfigRepository, error)
	DeleteConfigRepository(org, configRepo string) error
	GetConfigRepository(org, configRepo string) (*models.ConfigRepository, error)
	ListConfigRepositories(org string) ([]*models.ConfigRepository, error)
	PushConfig(org string, project string, env string, configs map[string]strfmt.Base64) error
	UpdateConfigRepository(org, configRepo, name, url, branch, cred string, setDefault bool) (*models.ConfigRepository, error)

	CreateCredential(org, name, cType string, rawCred *models.CredentialRaw, path, description string) error
	DeleteCredential(org, cred string) error
	GetCredential(org, cred string) (*models.Credential, error)
	ListCredentials(org, cType string) ([]*models.CredentialSimple, error)

	SendEvent(org, eventType, title, message, severity string, tags map[string]string, color string) error

	DeleteExternalBackend(org string, externalBackend uint32) error
	CreateExternalBackends(org, project, env, purpose, cred string, ebConfig models.ExternalBackendConfiguration) (*models.ExternalBackend, error)
	ListExternalBackends(org string) ([]*models.ExternalBackend, error)

	ValidateForm(org string, rawForms []byte) (*models.FormsValidationResult, error)

	// Login methods
	// Login is the method used to log the user into the Cycloid console
	Login(email, password string) (*models.UserSession, error)

	// LoginOrg is the used to log the user into a Cycloid organization
	LoginOrg(org, child string) (*models.UserSession, error)

	DeleteMember(org string, name string) error
	GetMember(org string, name string) (*models.MemberOrg, error)
	InviteMember(org, email, role string) error
	ListMembers(org string) ([]*models.MemberOrg, error)
	ListInvites(org string) ([]*models.Invitation, error)
	UpdateMembers(org, name, role string) (*models.MemberOrg, error)

	CreateOrganization(name string) (*models.Organization, error)
	DeleteOrganization(org string) error
	GetOrganization(org string) (*models.Organization, error)
	ListOrganizations() ([]*models.Organization, error)
	ListOrganizationWorkers(org string) ([]*models.Worker, error)

	ClearTaskCachePipeline(org, project, env, job, task string) error
	CreatePipeline(org, project, env, pipeline, variables, usecase string) (*models.Pipeline, error)
	DiffPipeline(org, project, env, pipeline, variables string) (*models.PipelineDiffs, error)
	GetPipelineJob(org, project, env, job string) (*models.Job, error)
	ListPipelineJobsBuilds(org, project, env, job string) ([]*models.Build, error)
	ListPipelineJobs(org, project, env string) ([]*models.Job, error)
	PausePipelineJob(org, project, env, job string) error
	PausePipeline(org string, project string, env string) error
	TriggerPipelineBuild(org, project, env, job string) error
	UnpausePipelineJob(org, project, env, job string) error
	UnpausePipeline(org string, project string, env string) error
	UpdatePipeline(org string, project string, env string, pipeline string, variables string) (*models.Pipeline, error)
	ListPipelines(org string) ([]*models.Pipeline, error)

	CreateProject(org, projectName, projectCanonical, env, pipelineTemplate, variables, description, stackRef, usecase, configRepo string) (*models.Project, error)
	DeleteProjectEnv(org, project, env string) error
	DeleteProject(org, project string) error
	GetProject(org string, project string) (*models.Project, error)
	ListProjects(org string) ([]*models.ProjectsItem, error)
	UpdateProject(org, projectName, projectCanonical string, envs []*models.NewEnvironment, description, stackRef, owner, configRepo string) (*models.Project, error)

	DeleteRole(org, role string) error
	GetRole(org, role string) (*models.Role, error)
	ListRoles(org string) ([]*models.Role, error)

	GetStack(org, ref string) (*models.ServiceCatalog, error)
	ListStacks(org string) ([]*models.ServiceCatalog, error)

	// API keys method
	// CreateAPIKey will request API to generate and return an API key
	CreateAPIKey(org, name, description, role string) (*models.APIKey, error)

	// ListAPIKey will request API to list generated API keys
	ListAPIKey(org string) ([]*models.APIKey, error)

	// GetAPIKey will request API to get a specified generated API key by its canonical
	GetAPIKey(org, canonical string) (*models.APIKey, error)

	// DeleteAPIKey will request API to delete a specified generated API key
	DeleteAPIKey(org, canonical string) error

	// ValidateInfraPolicies will validate the TF plan against
	// OPA policies defined on the Cycloid server
	ValidateInfraPolicies(org, project, env string, plan []byte) (*models.InfraPoliciesValidationResult, error)

	// CostEstimation will consume the backend API endpoint for cost estimation
	CostEstimation(org string, plan []byte) (*models.CostEstimationResult, error)
}

type middleware struct {
	api *client.APIClient
}

func NewMiddleware(api *client.APIClient) Middleware {
	return &middleware{api: api}
}
