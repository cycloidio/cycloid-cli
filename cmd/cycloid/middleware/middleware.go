package middleware

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"time"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

type Middleware interface {
	UserLogin(org, email *string, password string) (*models.UserSession, *http.Response, error)
	UserLoginToOrg(org, email, password string) (*models.UserSession, *http.Response, error)
	UserSignup(username, email, password, fullName string) (*http.Response, error)
	RefreshToken(org, childOrg *string, token string) (*models.UserSession, *http.Response, error)

	ActivateLicence(org, licence string) (*http.Response, error)

	// cycloid
	GetAppVersion() (*models.AppVersion, *http.Response, error)
	GetStatus() (*models.GeneralStatus, *http.Response, error)

	// catalog_repositories
	CreateCatalogRepository(org, name, url, branch, cred, visibility, teamCanonical string) (*models.ServiceCatalogSource, *http.Response, error)
	ListCatalogRepositories(org string) ([]*models.ServiceCatalogSource, *http.Response, error)
	GetCatalogRepository(org, catalogRepo string) (*models.ServiceCatalogSource, *http.Response, error)
	DeleteCatalogRepository(org, catalogRepo string) (*http.Response, error)
	UpdateCatalogRepository(org, catalogRepo string, name, url, branch, cred string, visibility *string) (*models.ServiceCatalogSource, *http.Response, error)
	RefreshCatalogRepository(org, catalogRepo string) (*models.ServiceCatalogChanges, *http.Response, error)

	CreateConfigRepository(org, name, canonical, url, branch, cred string, setDefault bool) (*models.ConfigRepository, *http.Response, error)
	DeleteConfigRepository(org, configRepo string) (*http.Response, error)
	GetConfigRepository(org, configRepo string) (*models.ConfigRepository, *http.Response, error)
	ListConfigRepositories(org string) ([]*models.ConfigRepository, *http.Response, error)
	UpdateConfigRepository(org, configRepo, cred, name, url, branch string, setDefault bool) (*models.ConfigRepository, *http.Response, error)

	// stacks (service_catalogs)
	GetStack(org, ref string) (*models.ServiceCatalog, *http.Response, error)
	UpdateStack(org, ref, teamCanonical string, visibility *string) (*models.ServiceCatalog, *http.Response, error)
	ListStacks(org string) ([]*models.ServiceCatalog, *http.Response, error)
	ListStackUseCases(org, ref, versionTag, versionBranch, versionCommitHash string) ([]*StackUseCase, *http.Response, error)
	ListStackVersions(org, ref string) ([]*StackVersion, *http.Response, error)
	ResolveStackVersion(org, ref, stackVersion string) (uint32, string, error)
	ListBlueprints(org string) ([]*models.ServiceCatalog, *http.Response, error)
	CreateStackFromBlueprint(org, blueprintRef, name, stack, catalogRepository, useCase string) (*models.ServiceCatalog, *http.Response, error)

	// organization_credentials
	CreateCredential(org, name, credentialType string, rawCred *models.CredentialRaw, path, canonical, description string) (*models.Credential, *http.Response, error)
	UpdateCredential(org, name, credentialType string, rawCred *models.CredentialRaw, path, canonical, description string) (*models.Credential, *http.Response, error)
	DeleteCredential(org, credential string) (*http.Response, error)
	GetCredential(org, credential string) (*models.Credential, *http.Response, error)
	ListCredentials(org, credentialType string) ([]*models.CredentialSimple, *http.Response, error)

	// events
	SendEvent(org, eventType, title, message, severity string, tags map[string]string, color string) (*http.Response, error)
	ListEvents(org string, eventType, eventSeverity []string, begin, end uint64) ([]*models.Event, *http.Response, error)

	// external_backends
	DeleteExternalBackend(org string, externalBackend uint32) (*http.Response, error)
	CreateExternalBackends(org, project, env, purpose, credential string, isDefault bool, externalBackendConfig models.ExternalBackendConfiguration) (*models.ExternalBackend, *http.Response, error)
	ListExternalBackends(org string) ([]*models.ExternalBackend, *http.Response, error)
	GetExternalBackend(org string, externalBackend uint32) (*models.ExternalBackend, *http.Response, error)
	GetRemoteTFExternalBackend(org string) (*models.ExternalBackend, error)
	UpdateExternalBackend(org string, externalBackendID uint32, purpose, credential string, isDefault bool, externalBackendConfig models.ExternalBackendConfiguration) (*models.ExternalBackend, *http.Response, error)

	// organization_member
	DeleteMember(org string, id uint32) (*http.Response, error)
	GetMember(org string, id uint32) (*models.MemberOrg, *http.Response, error)
	InviteMember(org, email, role string) (*models.MemberOrg, *http.Response, error)
	ListMembers(org string) ([]*models.MemberOrg, *http.Response, error)
	ListInvites(org string) ([]*models.MemberOrg, *http.Response, error)
	UpdateMember(org string, id uint32, role string) (*models.MemberOrg, *http.Response, error)

	// organization_teams
	ListTeams(org string, teamNameFilter *string, createdAtFilter *uint64, memberIDFilter *uint32, orderBy *TeamOrderByParam) ([]*models.Team, *http.Response, error)
	GetTeam(org, team string) (*models.Team, *http.Response, error)
	CreateTeam(org string, name, team, owner *string, roles []string) (*models.Team, *http.Response, error)
	UpdateTeam(org string, name, team, owner *string, roles []string) (*models.Team, *http.Response, error)
	DeleteTeam(org, team string) (*http.Response, error)

	// organization_team_members
	ListTeamMembers(org string, team string) ([]*models.MemberTeam, *http.Response, error)
	GetTeamMember(org string, team string, memberID uint32) (*models.MemberTeam, *http.Response, error)
	AssignMemberToTeam(org, team string, username, email *string) (*models.MemberTeam, *http.Response, error)
	UnAssignMemberFromTeam(org, team string, memberID uint32) (*http.Response, error)

	// organizations
	CreateOrganization(name string) (*models.Organization, *http.Response, error)
	UpdateOrganization(org, name string) (*models.Organization, *http.Response, error)
	DeleteOrganization(org string) (*http.Response, error)
	GetOrganization(org string) (*models.Organization, *http.Response, error)
	ListOrganizations() ([]*models.Organization, *http.Response, error)
	ListOrganizationWorkers(org string) ([]*models.Worker, *http.Response, error)
	ListOrganizationChildrens(org string) ([]*models.Organization, *http.Response, error)
	CreateOrganizationChild(org, childOrg string, childOrgName *string) (*models.Organization, *http.Response, error)
	CreateOrUpdateSubscription(org string, plan SubscriptionPlan, expiresAt time.Time, membersCount uint64, overwrite bool) (*models.Subscription, *http.Response, error)

	// Organization Forms
	InterpolateFormsConfig(org, env, project, component, serviceCatalogRef, useCase string, inputs models.FormVariables) (*models.ServiceCatalogConfig, *http.Response, error)
	ValidateForm(org string, rawForms []byte) (*models.FormsValidationResult, *http.Response, error)

	// Organization pipelines
	GetOrgPipelines(org string, concoursePipeline, project, env *string, statuses []string) ([]*models.Pipeline, *http.Response, error)

	// Project Pipelines
	GetProjectPipelines(org, project string) ([]*models.Pipeline, *http.Response, error)

	// Environments pipelines
	GetEnvPipelines(org, project, env string) ([]*models.Pipeline, *http.Response, error)

	// Component pipelines
	PausePipeline(org, project, env, component, pipelineName string) (*http.Response, error)
	UnpausePipeline(org, project, env, component, pipelineName string) (*http.Response, error)
	DiffPipeline(org, project, env, component, pipelineName, yamlPipeline, yamlVariables string, checkCredentials bool) (*models.PipelineDiffs, *http.Response, error)
	CreatePipeline(org, project, env, pipeline, component, yamlPipeline, yamlVariables string, checkCredentials bool) (*models.Pipeline, *http.Response, error)
	RenamePipeline(org, project, env, component, pipeline, newName string) (*http.Response, error)
	SyncedPipeline(org, project, env, component, pipeline string) (*models.PipelineStatus, *http.Response, error)
	GetPipeline(org, project, env, component, pipeline string) (*models.Pipeline, *http.Response, error)
	UpdatePipeline(org, project, env, component, pipelineName, yamlPipeline, yamlVariables string, checkCredentials bool) (*models.Pipeline, *http.Response, error)
	DeletePipeline(org, project, env, component, pipeline string) (*http.Response, error)

	// Component pipelines jobs
	GetJobs(org, project, env, component, pipeline string) ([]*models.Job, *http.Response, error)
	GetJob(org, project, env, component, pipeline, job string) (*models.Job, *http.Response, error)
	PauseJob(org, project, env, component, pipeline, job string) (*http.Response, error)
	UnPauseJob(org, project, env, component, pipeline, job string) (*http.Response, error)
	ClearTaskCache(org, project, env, component, pipeline, job, step string) (*models.ClearTaskCache, *http.Response, error)

	// Component pipelines jobs build
	GetBuilds(org, project, env, component, pipeline, job string) ([]*models.Build, *http.Response, error)
	CreateBuild(org, project, env, component, pipeline, job string) (*models.Build, *http.Response, error)
	GetBuild(org, project, env, component, pipeline, job, buildID string) (*models.Build, *http.Response, error)
	RerunBuild(org, project, env, component, pipeline, job, buildID string) (*models.Build, *http.Response, error)
	AbortBuild(org, project, env, component, pipeline, job, buildID string) (*http.Response, error)
	GetBuildEvents(org, project, env, component, pipeline, buildID string) (*string, *http.Response, error)
	OpenBuildEventsStream(ctx context.Context, org, project, env, component, pipeline, buildID, lastEventID string) (io.ReadCloser, *http.Response, error)
	GetBuildPlan(org, project, env, component, pipeline, job, buildID string) (*models.PublicPlan, *http.Response, error)
	GetBuildPreparation(org, project, env, component, pipeline, job, buildID string) (*models.Preparation, *http.Response, error)
	GetBuildResources(org, project, env, component, pipeline, job, buildID string) (*models.BuildInputsOutputs, *http.Response, error)

	// Project
	CreateProject(org, projectName, project, description, configRepository, owner, team, color, icon string) (*models.Project, *http.Response, error)
	UpdateProject(org, projectName, project, description, configRepository, owner, team, color, icon, cloudProvider string) (*models.Project, *http.Response, error)
	DeleteProject(org, project string) (*http.Response, error)
	GetProject(org string, project string) (*models.Project, *http.Response, error)
	ListProjects(org string) ([]*models.Project, *http.Response, error)
	ListProjectsEnv(org, project string) ([]*models.Environment, *http.Response, error)

	// Env
	GetEnv(org, project, env string) (*models.Environment, *http.Response, error)
	CreateEnv(org, project, env, envName, color string) (*models.Environment, *http.Response, error)
	UpdateEnv(org, project, env, envName, color string) (*models.Environment, *http.Response, error)
	DeleteEnv(org, project, env string) (*http.Response, error)

	// Component
	CreateOrUpdateComponent(org, project, env, component, description, name, stackRef, versionTag, versionBranch, versionCommitHash, useCase, cloudProvider string, vars models.FormVariables) (*models.Component, *http.Response, error)
	ListComponents(org, project, env string) ([]*models.Component, *http.Response, error)
	GetComponent(org, project, env, component string) (*models.Component, *http.Response, error)
	MigrateComponent(org, project, env, component, targetProject, targetEnv, newCanonical, newName string) (*models.Component, *http.Response, error)
	DeleteComponent(org, project, env, component string) (*http.Response, error)
	GetComponentConfig(org, project, env, component string) (models.FormVariables, *http.Response, error)
	GetComponentStackConfig(org, project, env, component, useCase, versionTag, versionBranch, versionCommitHash string) (models.ServiceCatalogConfigs, *http.Response, error)

	DeleteRole(org, role string) (*http.Response, error)
	GetRole(org, role string) (*models.Role, *http.Response, error)
	ListRoles(org string) ([]*models.Role, *http.Response, error)
	CreateRole(org string, name, canonical, description *string, rules []*models.NewRule) (*models.NewRole, *http.Response, error)
	UpdateRole(org, roleCanonical string, name, canonical, description *string, rules []*models.NewRule) (*models.Role, *http.Response, error)

	// ApiKeys
	ListAPIKeys(org string) ([]*models.APIKey, *http.Response, error)
	GetAPIKey(org, canonical string) (*models.APIKey, *http.Response, error)
	CreateAPIKey(org, canonical, description, owner string, name *string, rules []*models.NewRule) (*models.APIKey, *http.Response, error)
	DeleteAPIKey(org, canonical string) (*http.Response, error)

	// ValidateInfraPolicies will validate the TF plan against OPA policies defined on the Cycloid server
	ValidateInfraPolicies(org, project, env string, plan []byte) (*models.InfraPoliciesValidationResult, *http.Response, error)
	CreateInfraPolicy(org, policyFile, policyCanonical, description, policyName, ownercanonical, severity string, enabled bool) (*models.InfraPolicy, *http.Response, error)
	DeleteInfraPolicy(org, policycanonical string) (*http.Response, error)
	ListInfraPolicies(org string) ([]*models.InfraPolicy, *http.Response, error)
	GetInfraPolicy(org, infraPolicy string) (*models.InfraPolicy, *http.Response, error)
	UpdateInfraPolicy(org, infraPolicy, policyFile, description, policyName, ownercanonical, severity string, enabled bool) (*models.InfraPolicy, *http.Response, error)

	// CostEstimation will consume the backend API endpoint for cost estimation
	CostEstimation(org string, plan []byte) (*models.CostEstimationResult, *http.Response, error)

	// Extra actions out of the api
	InitFirstOrg(org, userName, fullName, email, password, licence string, apiKeyCanonical *string) (*FirstOrgData, *http.Response, error)

	// Generic request for un-implemented routes
	GenericRequest(req Request, response any) (*http.Response, error)
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
	api           *common.APIClient
	GenericClient http.Client
}

func NewMiddleware(api *common.APIClient) Middleware {
	client := http.DefaultClient
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: api.Config.Insecure,
		},
	}
	return &middleware{
		api:           api,
		GenericClient: *client,
	}
}
