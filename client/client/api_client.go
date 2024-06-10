// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/client/cost_estimation"
	"github.com/cycloidio/cycloid-cli/client/client/cycloid"
	"github.com/cycloidio/cycloid-cli/client/client/organization_api_keys"
	"github.com/cycloidio/cycloid-cli/client/client/organization_children"
	"github.com/cycloidio/cycloid-cli/client/client/organization_config_repositories"
	"github.com/cycloidio/cycloid-cli/client/client/organization_credentials"
	"github.com/cycloidio/cycloid-cli/client/client/organization_external_backends"
	"github.com/cycloidio/cycloid-cli/client/client/organization_forms"
	"github.com/cycloidio/cycloid-cli/client/client/organization_infrastructure_policies"
	"github.com/cycloidio/cycloid-cli/client/client/organization_invitations"
	"github.com/cycloidio/cycloid-cli/client/client/organization_kpis"
	"github.com/cycloidio/cycloid-cli/client/client/organization_members"
	"github.com/cycloidio/cycloid-cli/client/client/organization_pipelines"
	"github.com/cycloidio/cycloid-cli/client/client/organization_pipelines_jobs"
	"github.com/cycloidio/cycloid-cli/client/client/organization_pipelines_jobs_build"
	"github.com/cycloidio/cycloid-cli/client/client/organization_projects"
	"github.com/cycloidio/cycloid-cli/client/client/organization_roles"
	"github.com/cycloidio/cycloid-cli/client/client/organization_service_catalog_sources"
	"github.com/cycloidio/cycloid-cli/client/client/organization_workers"
	"github.com/cycloidio/cycloid-cli/client/client/organizations"
	"github.com/cycloidio/cycloid-cli/client/client/service_catalogs"
	"github.com/cycloidio/cycloid-cli/client/client/user"
)

// Default API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "http-api.cycloid.io"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *API {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *API {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *API {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(API)
	cli.Transport = transport
	cli.CostEstimation = cost_estimation.New(transport, formats)
	cli.Cycloid = cycloid.New(transport, formats)
	cli.OrganizationAPIKeys = organization_api_keys.New(transport, formats)
	cli.OrganizationChildren = organization_children.New(transport, formats)
	cli.OrganizationConfigRepositories = organization_config_repositories.New(transport, formats)
	cli.OrganizationCredentials = organization_credentials.New(transport, formats)
	cli.OrganizationExternalBackends = organization_external_backends.New(transport, formats)
	cli.OrganizationForms = organization_forms.New(transport, formats)
	cli.OrganizationInfrastructurePolicies = organization_infrastructure_policies.New(transport, formats)
	cli.OrganizationInvitations = organization_invitations.New(transport, formats)
	cli.OrganizationKpis = organization_kpis.New(transport, formats)
	cli.OrganizationMembers = organization_members.New(transport, formats)
	cli.OrganizationPipelines = organization_pipelines.New(transport, formats)
	cli.OrganizationPipelinesJobs = organization_pipelines_jobs.New(transport, formats)
	cli.OrganizationPipelinesJobsBuild = organization_pipelines_jobs_build.New(transport, formats)
	cli.OrganizationProjects = organization_projects.New(transport, formats)
	cli.OrganizationRoles = organization_roles.New(transport, formats)
	cli.OrganizationServiceCatalogSources = organization_service_catalog_sources.New(transport, formats)
	cli.OrganizationWorkers = organization_workers.New(transport, formats)
	cli.Organizations = organizations.New(transport, formats)
	cli.ServiceCatalogs = service_catalogs.New(transport, formats)
	cli.User = user.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// API is a client for API
type API struct {
	CostEstimation cost_estimation.ClientService

	Cycloid cycloid.ClientService

	OrganizationAPIKeys organization_api_keys.ClientService

	OrganizationChildren organization_children.ClientService

	OrganizationConfigRepositories organization_config_repositories.ClientService

	OrganizationCredentials organization_credentials.ClientService

	OrganizationExternalBackends organization_external_backends.ClientService

	OrganizationForms organization_forms.ClientService

	OrganizationInfrastructurePolicies organization_infrastructure_policies.ClientService

	OrganizationInvitations organization_invitations.ClientService

	OrganizationKpis organization_kpis.ClientService

	OrganizationMembers organization_members.ClientService

	OrganizationPipelines organization_pipelines.ClientService

	OrganizationPipelinesJobs organization_pipelines_jobs.ClientService

	OrganizationPipelinesJobsBuild organization_pipelines_jobs_build.ClientService

	OrganizationProjects organization_projects.ClientService

	OrganizationRoles organization_roles.ClientService

	OrganizationServiceCatalogSources organization_service_catalog_sources.ClientService

	OrganizationWorkers organization_workers.ClientService

	Organizations organizations.ClientService

	ServiceCatalogs service_catalogs.ClientService

	User user.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *API) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.CostEstimation.SetTransport(transport)
	c.Cycloid.SetTransport(transport)
	c.OrganizationAPIKeys.SetTransport(transport)
	c.OrganizationChildren.SetTransport(transport)
	c.OrganizationConfigRepositories.SetTransport(transport)
	c.OrganizationCredentials.SetTransport(transport)
	c.OrganizationExternalBackends.SetTransport(transport)
	c.OrganizationForms.SetTransport(transport)
	c.OrganizationInfrastructurePolicies.SetTransport(transport)
	c.OrganizationInvitations.SetTransport(transport)
	c.OrganizationKpis.SetTransport(transport)
	c.OrganizationMembers.SetTransport(transport)
	c.OrganizationPipelines.SetTransport(transport)
	c.OrganizationPipelinesJobs.SetTransport(transport)
	c.OrganizationPipelinesJobsBuild.SetTransport(transport)
	c.OrganizationProjects.SetTransport(transport)
	c.OrganizationRoles.SetTransport(transport)
	c.OrganizationServiceCatalogSources.SetTransport(transport)
	c.OrganizationWorkers.SetTransport(transport)
	c.Organizations.SetTransport(transport)
	c.ServiceCatalogs.SetTransport(transport)
	c.User.SetTransport(transport)
}
