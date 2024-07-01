package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
)

type Middleware interface {
	GetAppVersion() (*models.AppVersion, error)
	GetStatus() (*models.GeneralStatus, error)

	CreateCatalogRepository(org, name, url, branch, cred string) (*models.ServiceCatalogSource, error)
	DeleteCatalogRepository(org, catalogRepo string) error
	GetCatalogRepository(org, catalogRepo string) (*models.ServiceCatalogSource, error)
	ListCatalogRepositories(org string) ([]*models.ServiceCatalogSource, error)
	RefreshCatalogRepository(org, catalogRepo string) (*models.ServiceCatalogChanges, error)
	UpdateCatalogRepository(org, name, url, branch, catalogRepo, cred string) (*models.ServiceCatalogSource, error)
	GetStackConfig(org, ref string) (interface{}, error)

	CreateConfigRepository(org, name, url, branch, cred string, setDefault bool) (*models.ConfigRepository, error)
	DeleteConfigRepository(org, configRepo string) error
	GetConfigRepository(org, configRepo string) (*models.ConfigRepository, error)
	ListConfigRepositories(org string) ([]*models.ConfigRepository, error)
	PushConfig(org string, project string, env string, configs map[string]strfmt.Base64) error
	UpdateConfigRepository(org, configRepo, cred, name, url, branch string, setDefault bool) (*models.ConfigRepository, error)

	CreateCredential(org, name, cType string, rawCred *models.CredentialRaw, path, can, description string) (*models.Credential, error)
	UpdateCredential(org, name, cType string, rawCred *models.CredentialRaw, path, can, description string) (*models.Credential, error)
	DeleteCredential(org, cred string) error
	GetCredential(org, cred string) (*models.Credential, error)
	ListCredentials(org, cType string) ([]*models.CredentialSimple, error)

	SendEvent(org, eventType, title, message, severity string, tags map[string]string, color string) error

	DeleteExternalBackend(org string, externalBackend uint32) error
	CreateExternalBackends(org, project, env, purpose, cred string, def bool, ebConfig models.ExternalBackendConfiguration) (*models.ExternalBackend, error)
	ListExternalBackends(org string) ([]*models.ExternalBackend, error)
	GetExternalBackend(org string, externalBackend uint32) (*models.ExternalBackend, error)
	GetRemoteTFExternalBackend(org string) (*models.ExternalBackend, error)
	UpdateExternalBackend(org string, externalBackendID uint32, purpose, cred string, def bool, ebConfig models.ExternalBackendConfiguration) (*models.ExternalBackend, error)

	// Organization Forms
	CreateFormsConfig(org string, project string, serviceCatalogRef string, inputs []*models.FormInput) (map[string]any, error)
	ValidateForm(org string, rawForms []byte) (*models.FormsValidationResult, error)

	DeleteMember(org string, name string) error
	GetMember(org string, name string) (*models.MemberOrg, error)
	InviteMember(org, email, role string) error
	DeleteInvite(org string, invite string) error
	ListMembers(org string) ([]*models.MemberOrg, error)
	ListInvites(org string) ([]*models.Invitation, error)
	UpdateMembers(org, name, role string) (*models.MemberOrg, error)

	CreateOrganization(name string) (*models.Organization, error)
	UpdateOrganization(org, name string) (*models.Organization, error)
	DeleteOrganization(org string) error
	GetOrganization(org string) (*models.Organization, error)
	ListOrganizations() ([]*models.Organization, error)
	ListOrganizationWorkers(org string) ([]*models.Worker, error)
	ListOrganizationChildrens(org string) ([]*models.Organization, error)
	CreateOrganizationChild(org, porg string) (*models.Organization, error)

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
	GetPipeline(org, project, env string) (*models.Pipeline, error)
	SyncedPipeline(org, project, env string) (*models.PipelineStatus, error)

	CreateProject(org, projectName, projectCanonical, env, pipelineTemplate, variables, description, stackRef, usecase, configRepo string) (*models.Project, error)
	CreateEmptyProject(org, projectName, projectCanonical, description, stackRef, configRepo string) (*models.Project, error)
	DeleteProjectEnv(org, project, env string) error
	DeleteProject(org, project string) error
	GetProject(org string, project string) (*models.Project, error)
	GetProjectConfig(org string, project string, environment string) (*models.ProjectEnvironmentConfig, error)
	ListProjects(org string) ([]*models.Project, error)
	UpdateProject(org, projectName, projectCanonical string, envs []*models.NewEnvironment, description, stackRef, owner, configRepo string, inputs []*models.FormInput, updatedAt uint64) (*models.Project, error)

	DeleteRole(org, role string) error
	GetRole(org, role string) (*models.Role, error)
	ListRoles(org string) ([]*models.Role, error)

	GetStack(org, ref string) (*models.ServiceCatalog, error)
	ListStacks(org string) ([]*models.ServiceCatalog, error)

	CreateKpi(name, kpiType, widget, org, project, job, env, config string) (*models.KPI, error)
	DeleteKpi(org, kpi string) error
	ListKpi(org, project, env string) ([]*models.KPI, error)

	// API keys method
	// ListAPIKey will request API to list generated API keys
	ListAPIKey(org string) ([]*models.APIKey, error)

	// GetAPIKey will request API to get a specified generated API key by its canonical
	GetAPIKey(org, canonical string) (*models.APIKey, error)

	// DeleteAPIKey will request API to delete a specified generated API key
	DeleteAPIKey(org, canonical string) error

	// ValidateInfraPolicies will validate the TF plan against OPA policies defined on the Cycloid server
	ValidateInfraPolicies(org, project, env string, plan []byte) (*models.InfraPoliciesValidationResult, error)
	// CreateInfraPolicy will create a new infraPolicy with the repo file supplied
	CreateInfraPolicy(org, policyFile, policyCanonical, description, policyName, ownerCannonical, severity string, enabled bool) (*models.InfraPolicy, error)
	DeleteInfraPolicy(org, policyCannonical string) error
	ListInfraPolicies(org string) ([]*models.InfraPolicy, error)
	GetInfraPolicy(org, infraPolicy string) (*models.InfraPolicy, error)
	UpdateInfraPolicy(org, infraPolicy, policyFile, description, policyName, ownerCannonical, severity string, enabled bool) (*models.InfraPolicy, error)

	// CostEstimation will consume the backend API endpoint for cost estimation
	CostEstimation(org string, plan []byte) (*models.CostEstimationResult, error)
}

type middleware struct {
	api *common.APIClient
}

func NewMiddleware(api *common.APIClient) Middleware {
	return &middleware{api: api}
}
