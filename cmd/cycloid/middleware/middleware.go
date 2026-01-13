package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

type Middleware interface {
	UserLogin(org, email *string, password string) (*models.UserSession, error)
	UserLoginToOrg(org, email, password string) (*models.UserSession, error)
	UserSignup(username, email, password, fullName string) error
	RefreshToken(org, childOrg *string, token string) (*models.UserSession, error)

	ActivateLicence(org, licence string) error

	// cycloid
	GetAppVersion() (*models.AppVersion, error)
	GetStatus() (*models.GeneralStatus, error)

	// catalog_repositories
	CreateCatalogRepository(org, name, url, branch, cred, visibility, teamCanonical string) (*models.ServiceCatalogSource, error)
	ListCatalogRepositories(org string) ([]*models.ServiceCatalogSource, error)
	GetCatalogRepository(org, catalogRepo string) (*models.ServiceCatalogSource, error)
	DeleteCatalogRepository(org, catalogRepo string) error
	UpdateCatalogRepository(org, name, url, branch, catalogRepo, cred string) (*models.ServiceCatalogSource, error)
	RefreshCatalogRepository(org, catalogRepo string) (*models.ServiceCatalogChanges, error)

	CreateConfigRepository(org, name, canonical, url, branch, cred string, setDefault bool) (*models.ConfigRepository, error)
	DeleteConfigRepository(org, configRepo string) error
	GetConfigRepository(org, configRepo string) (*models.ConfigRepository, error)
	ListConfigRepositories(org string) ([]*models.ConfigRepository, error)
	UpdateConfigRepository(org, configRepo, cred, name, url, branch string, setDefault bool) (*models.ConfigRepository, error)

	// stacks (service_catalogs)
	GetStack(org, ref string) (*models.ServiceCatalog, error)
	UpdateStack(org, ref, teamCanonical string, visibility *string) (*models.ServiceCatalog, error)
	ListStacks(org string) ([]*models.ServiceCatalog, error)
	ListStackUseCases(org, ref, versionTag, versionBranch, versionCommitHash string) ([]*models.StackUseCase, error)
	ListStackVersions(org, ref string) ([]*models.ServiceCatalogSourceVersion, error)
	ListBlueprints(org string) ([]*models.ServiceCatalog, error)
	CreateStackFromBlueprint(org, blueprintRef, name, stack, catalogRepository, useCase string) (*models.ServiceCatalog, error)

	// organization_credentials
	CreateCredential(org, name, credentialType string, rawCred *models.CredentialRaw, path, canonical, description string) (*models.Credential, error)
	UpdateCredential(org, name, credentialType string, rawCred *models.CredentialRaw, path, canonical, description string) (*models.Credential, error)
	DeleteCredential(org, credential string) error
	GetCredential(org, credential string) (*models.Credential, error)
	ListCredentials(org, credentialType string) ([]*models.CredentialSimple, error)

	// events
	SendEvent(org, eventType, title, message, severity string, tags map[string]string, color string) error
	ListEvents(org string, eventType, eventSeverity []string, begin, end uint64) ([]*models.Event, error)

	// external_backends
	DeleteExternalBackend(org string, externalBackend uint32) error
	CreateExternalBackends(org, project, env, purpose, credential string, isDefault bool, externalBackendConfig models.ExternalBackendConfiguration) (*models.ExternalBackend, error)
	ListExternalBackends(org string) ([]*models.ExternalBackend, error)
	GetExternalBackend(org string, externalBackend uint32) (*models.ExternalBackend, error)
	GetRemoteTFExternalBackend(org string) (*models.ExternalBackend, error)
	UpdateExternalBackend(org string, externalBackendID uint32, purpose, credential string, isDefault bool, externalBackendConfig models.ExternalBackendConfiguration) (*models.ExternalBackend, error)

	// organization_member
	DeleteMember(org string, id uint32) error
	GetMember(org string, id uint32) (*models.MemberOrg, error)
	InviteMember(org, email, role string) (*models.MemberOrg, error)
	ListMembers(org string) ([]*models.MemberOrg, error)
	ListInvites(org string) ([]*models.MemberOrg, error)
	UpdateMember(org string, id uint32, role string) (*models.MemberOrg, error)

	// organizations
	CreateOrganization(name string) (*models.Organization, error)
	UpdateOrganization(org, name string) (*models.Organization, error)
	DeleteOrganization(org string) error
	GetOrganization(org string) (*models.Organization, error)
	ListOrganizations() ([]*models.Organization, error)
	ListOrganizationWorkers(org string) ([]*models.Worker, error)
	ListOrganizationChildrens(org string) ([]*models.Organization, error)
	CreateOrganizationChild(org, childOrg string, childOrgName *string) (*models.Organization, error)

	// Organization Forms
	InterpolateFormsConfig(org, env, project, component, serviceCatalogRef, useCase string, inputs models.FormVariables) (*models.ServiceCatalogConfig, error)
	ValidateForm(org string, rawForms []byte) (*models.FormsValidationResult, error)

	// Organization pipelines
	GetOrgPipelines(org string, concoursePipeline, project, env *string, statuses []string) ([]*models.Pipeline, error)

	// Project Pipelines
	GetProjectPipelines(org, project string) ([]*models.Pipeline, error)

	// Environments pipelines
	GetEnvPipelines(org, project, env string) ([]*models.Pipeline, error)

	// Component pipelines
	PausePipeline(org, project, env, component, pipelineName string) error
	UnpausePipeline(org, project, env, component, pipelineName string) error
	DiffPipeline(org, project, env, component, pipelineName, yamlPipeline, yamlVariables string, checkCredentials bool) (*models.PipelineDiffs, error)
	CreatePipeline(org, project, env, pipeline, component, yamlPipeline, yamlVariables string, checkCredentials bool) (*models.Pipeline, error)
	RenamePipeline(org, project, env, component, pipeline, newName string) error
	SyncedPipeline(org, project, env, component, pipeline string) (*models.PipelineStatus, error)
	GetPipeline(org, project, env, component, pipeline string) (*models.Pipeline, error)
	UpdatePipeline(org, project, env, component, pipelineName, yamlPipeline, yamlVariables string, checkCredentials bool) (*models.Pipeline, error)
	DeletePipeline(org, project, env, component, pipeline string) error

	// Component pipelines jobs
	GetJobs(org, project, env, component, pipeline string) ([]*models.Job, error)
	GetJob(org, project, env, component, pipeline, job string) (*models.Job, error)
	PauseJob(org, project, env, component, pipeline, job string) error
	UnPauseJob(org, project, env, component, pipeline, job string) error
	ClearTaskCache(org, project, env, component, pipeline, job, step string) (*models.ClearTaskCache, error)

	// Component pipelines jobs build
	GetBuilds(org, project, env, component, pipeline, job string) ([]*models.Build, error)
	CreateBuild(org, project, env, component, pipeline, job string) (*models.Build, error)
	GetBuild(org, project, env, component, pipeline, job, buildID string) (*models.Build, error)
	RerunBuild(org, project, env, component, pipeline, job, buildID string) (*models.Build, error)
	AbortBuild(org, project, env, component, pipeline, job, buildID string) error
	GetBuildEvents(org, project, env, component, pipeline, buildID string) (*string, error)
	GetBuildPlan(org, project, env, component, pipeline, job, buildID string) (*models.PublicPlan, error)
	GetBuildPreparation(org, project, env, component, pipeline, job, buildID string) (*models.Preparation, error)
	GetBuildResources(org, project, env, component, pipeline, job, buildID string) (*models.BuildInputsOutputs, error)

	// Project
	CreateProject(org, projectName, project, description, configRepository, owner, team, color, icon string) (*models.Project, error)
	UpdateProject(org, projectName, project, description, configRepository, owner, team, color, icon, cloudProvider string) (*models.Project, error)
	DeleteProject(org, project string) error
	GetProject(org string, project string) (*models.Project, error)
	ListProjects(org string) ([]*models.Project, error)
	ListProjectsEnv(org, project string) ([]*models.Environment, error)

	// Env
	GetEnv(org, project, env string) (*models.Environment, error)
	CreateEnv(org, project, env, envName, color string) (*models.Environment, error)
	UpdateEnv(org, project, env, envName, color string) (*models.Environment, error)
	DeleteEnv(org, project, env string) error

	// Component
	CreateComponent(org, project, env, component, description, componentName, serviceCatalogRef, versionTag, versionBranch, versionCommitHash, cloudProviderCanonical string) (*models.Component, error)
	UpdateComponent(org, project, env, component, description string, componentName *string) (*models.Component, error)
	CreateAndConfigureComponent(org, project, env, component, description, componentName, serviceCatalogRef, versionTag, versionBranch, versionCommitHash, useCase, cloudProviderCanonical string, vars models.FormVariables) (*models.Component, error)
	ConfigureComponent(org, project, env, component, useCase string, vars models.FormVariables) error
	ListComponents(org, project, env string) ([]*models.Component, error)
	GetComponent(org, project, env, component string) (*models.Component, error)
	MigrateComponent(org, project, env, component, targetProject, targetEnv, newCanonical, newName string) (*models.Component, error)
	DeleteComponent(org, project, env, component string) error
	GetComponentConfig(org, project, env, component string) (models.FormVariables, error)
	GetComponentStackConfig(org, project, env, component, useCase, versionTag, versionBranch, versionCommitHash string) (models.ServiceCatalogConfigs, error)

	DeleteRole(org, role string) error
	GetRole(org, role string) (*models.Role, error)
	ListRoles(org string) ([]*models.Role, error)

	// CreateKpi(name, kpiType, widget, org, project, job, env, config string) (*models.KPI, error)
	// DeleteKpi(org, kpi string) error
	// ListKpi(org, project, env string) ([]*models.KPI, error)

	// ApiKeys
	ListAPIKeys(org string) ([]*models.APIKey, error)
	GetAPIKey(org, canonical string) (*models.APIKey, error)
	CreateAPIKey(org, canonical, description, owner string, name *string, rules []*models.NewRule) (*models.APIKey, error)
	DeleteAPIKey(org, canonical string) error

	// ValidateInfraPolicies will validate the TF plan against OPA policies defined on the Cycloid server
	ValidateInfraPolicies(org, project, env string, plan []byte) (*models.InfraPoliciesValidationResult, error)
	CreateInfraPolicy(org, policyFile, policyCanonical, description, policyName, ownercanonical, severity string, enabled bool) (*models.InfraPolicy, error)
	DeleteInfraPolicy(org, policycanonical string) error
	ListInfraPolicies(org string) ([]*models.InfraPolicy, error)
	GetInfraPolicy(org, infraPolicy string) (*models.InfraPolicy, error)
	UpdateInfraPolicy(org, infraPolicy, policyFile, description, policyName, ownercanonical, severity string, enabled bool) (*models.InfraPolicy, error)

	// CostEstimation will consume the backend API endpoint for cost estimation
	CostEstimation(org string, plan []byte) (*models.CostEstimationResult, error)

	// Extra actions out of the api
	InitFirstOrg(org, userName, fullName, email, password, licence string, apiKeyCanonical *string) (*FirstOrgData, error)
}

type FirstOrgData struct {
	Org                 string
	Username            string
	FullName            string
	Email               string
	Password            string
	Token               string
	APIKey              *string
	CredentialCanonical *string
}

type middleware struct {
	api *common.APIClient
}

func NewMiddleware(api *common.APIClient) Middleware {
	return &middleware{api: api}
}
