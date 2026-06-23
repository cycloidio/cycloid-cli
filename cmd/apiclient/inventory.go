package apiclient

import (
	"net/http"
	"strconv"

	"github.com/cycloidio/cycloid-cli/gen/models"
)

// ListInventoryResources lists inventory resources for an organization.
//
// Supported LHS filter attributes:
//   - resources_provider, resources_type, resources_name (display_name), resources_module,
//     resources_mode, resources_label — resource-level fields
//   - project_canonical, environment_canonical, component_canonical — via linked TF state
//   - service_catalog_canonical, service_catalog_ref — via linked component
//   - organization_canonical — organization scope
//
// Any unknown attribute key is treated as a JSON path filter against sr.attributes
// and sr.custom_attributes (e.g. "tags.env[eq]=prod").
func (m *middleware) ListInventoryResources(org string, filters ...LHSFilter) ([]*models.InventoryResource, *http.Response, error) {
	var result []*models.InventoryResource
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "inventory"},
		LHSFilters:   filters,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// ListInventoryOutputs lists terraform state outputs for an organization.
//
// Supported LHS filter attributes:
//   - output_key — output key name
//   - output_type — output type filter
//   - output_is_pinned — boolean, "true" or "false"
//   - project_canonical, environment_canonical, component_canonical — via linked TF state
//   - service_catalog_canonical — via linked component
//   - organization_canonical — organization scope
//
// Any unknown attribute key is treated as a JSON path filter against ou.value
// (e.g. "instance_ip[eq]=1.2.3.4").
func (m *middleware) ListInventoryOutputs(org string, filters ...LHSFilter) ([]*InventoryOutput, *http.Response, error) {
	var result []*InventoryOutput
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "inventory", "outputs"},
		LHSFilters:   filters,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// CreateInventoryResource creates an inventory resource for an organization.
func (m *middleware) CreateInventoryResource(org string, body *models.NewInventoryResource) (*models.InventoryResource, *http.Response, error) {
	var result *models.InventoryResource
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "inventory", "resources"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// DeleteInventoryResource deletes an inventory resource by ID.
func (m *middleware) DeleteInventoryResource(org string, id uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "inventory", "resources", strconv.FormatUint(uint64(id), 10)},
	}, nil)
	return resp, err
}
