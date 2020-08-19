// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cycloidio/youdeploy-cli/client/client/cycloid"
	"github.com/cycloidio/youdeploy-cli/client/client/organization_config_repositories"
	"github.com/cycloidio/youdeploy-cli/client/client/organization_credentials"
	"github.com/cycloidio/youdeploy-cli/client/client/organization_external_backends"
	"github.com/cycloidio/youdeploy-cli/client/client/organization_projects"
	"github.com/cycloidio/youdeploy-cli/client/client/organizations"
	"github.com/cycloidio/youdeploy-cli/client/client/service_catalogs"
)

// Default API client HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "http-api-stoplight.cycloid.io"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new API client HTTP client.
func NewHTTPClient(formats strfmt.Registry) *APIClient {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new API client HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *APIClient {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new API client client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *APIClient {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(APIClient)
	cli.Transport = transport

	cli.Cycloid = cycloid.New(transport, formats)

	cli.OrganizationConfigRepositories = organization_config_repositories.New(transport, formats)

	cli.OrganizationCredentials = organization_credentials.New(transport, formats)

	cli.OrganizationExternalBackends = organization_external_backends.New(transport, formats)

	cli.OrganizationProjects = organization_projects.New(transport, formats)

	cli.Organizations = organizations.New(transport, formats)

	cli.ServiceCatalogs = service_catalogs.New(transport, formats)

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

// APIClient is a client for API client
type APIClient struct {
	Cycloid *cycloid.Client

	OrganizationConfigRepositories *organization_config_repositories.Client

	OrganizationCredentials *organization_credentials.Client

	OrganizationExternalBackends *organization_external_backends.Client

	OrganizationProjects *organization_projects.Client

	Organizations *organizations.Client

	ServiceCatalogs *service_catalogs.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *APIClient) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Cycloid.SetTransport(transport)

	c.OrganizationConfigRepositories.SetTransport(transport)

	c.OrganizationCredentials.SetTransport(transport)

	c.OrganizationExternalBackends.SetTransport(transport)

	c.OrganizationProjects.SetTransport(transport)

	c.Organizations.SetTransport(transport)

	c.ServiceCatalogs.SetTransport(transport)

}
