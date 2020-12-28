package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client"
	"github.com/cycloidio/cycloid-cli/client/models"
	strfmt "github.com/go-openapi/strfmt"
)

type Middleware interface {
	GetAppVersion() (*models.AppVersion, error)
	GetStatus() (*models.GeneralStatus, error)

	CreateCatalogRepository(org, name, url, branch string, cred uint32) (*models.ServiceCatalogSource, error)
	DeleteCatalogRepository(org string, catalogRepo uint32) error
	GetCatalogRepository(org string, catalogRepo uint32) (*models.ServiceCatalogSource, error)
	ListCatalogRepositories(org string) ([]*models.ServiceCatalogSource, error)
	RefreshCatalogRepository(org string, catalogRepo uint32) (*models.ServiceCatalogSource, error)
	UpdateCatalogRepository(org string, catalogRepo uint32, name, url, branch string, cred uint32) (*models.ServiceCatalogSource, error)

	CreateConfigRepository(org, name, url, branch string, setDefault bool, cred uint32) (*models.ConfigRepository, error)
	DeleteConfigRepository(org string, configRepo uint32) error
	GetConfigRepository(org string, configRepo uint32) (*models.ConfigRepository, error)
	ListConfigRepositories(org string) ([]*models.ConfigRepository, error)
	PushConfig(org string, project string, env string, configs map[string]strfmt.Base64) error
	UpdateConfigRepository(org string, configRepo uint32, name, url, branch string, setDefault bool, cred uint32) (*models.ConfigRepository, error)

	CreateCredential(org, name, cType string, rawCred *models.CredentialRaw, path, description string) error
	DeleteCredential(org string, cred uint32) error
	GetCredential(org string, cred uint32) (*models.Credential, error)
	ListCredentials(org, cType string) ([]*models.CredentialSimple, error)

	SendEvent(org, eventType, title, message, severity string, tags map[string]string, color string) error

	DeleteExternalBackend(org string, externalBackend uint32) error
	CreateExternalBackends(org, project, env, purpose string, cred uint32, ebConfig models.ExternalBackendConfiguration) (*models.ExternalBackend, error)
	ListExternalBackends(org string) ([]*models.ExternalBackend, error)

	ValidateForm(org string, rawForms []byte) (*models.FormsValidationResult, error)

	// Login methods
	// Login is the method used to log the user into the Cycloid console
	Login(email, password string) (*models.UserSession, error)

	// LoginOrg is the used to log the user into a Cycloid organization
	LoginOrg(org, child string) (*models.UserSession, error)

	DeleteMember(org string, name string) error
	GetMember(org string, name string) (*models.MemberOrg, error)
	InviteMember(org string, email string, roleID uint32) error
	ListMembers(org string) ([]*models.MemberOrg, error)
	ListInvites(org string) ([]*models.Invitation, error)
	UpdateMembers(org string, name string, roleID uint32) (*models.MemberOrg, error)

	CreateOrganization(name string, canonical string) (*models.Organization, error)
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

	CreateProject(org, projectName, projectCanonical, env, pipelineTemplate, variables, description, stackRef, usecase string, configRepo uint32) (*models.Project, error)
	DeleteProjectEnv(org, project, env string) error
	DeleteProject(org, project string) error
	GetProject(org string, project string) (*models.Project, error)
	ListProjects(org string) ([]*models.ProjectsItem, error)
	UpdateProject(org, projectName, projectCanonical string, envs []*models.NewEnvironment, description, stackRef, owner string, configRepo uint32) (*models.Project, error)

	DeleteRole(org string, id uint32) error
	GetRole(org string, id uint32) (*models.Role, error)
	ListRoles(org string) ([]*models.Role, error)

	GetStack(org, ref string) (*models.ServiceCatalog, error)
	ListStacks(org string) ([]*models.ServiceCatalog, error)

	// API keys method
	// CreateAPIKey will request API to generate and return an API key
	CreateAPIKey(org, name, canonical, description string, roleID uint32) (*models.APIKey, error)

	// ListAPIKey will request API to list generated API keys
	ListAPIKey(org string) ([]*models.APIKey, error)

	// GetAPIKey will request API to get a specified generated API key by its canonical
	GetAPIKey(org, canonical string) (*models.APIKey, error)

	// DeleteAPIKey will request API to delete a specified generated API key
	DeleteAPIKey(org, canonical string) error
}

type middleware struct {
	api *client.APIClient
}

func NewMiddleware(api *client.APIClient) Middleware {
	return &middleware{api: api}
}
