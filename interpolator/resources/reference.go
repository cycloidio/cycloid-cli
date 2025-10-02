package resources

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

func ExpandShortcuts(uri string) string {
	replacer := strings.NewReplacer(
		"/org/", "/organizations/",
		"/organization/", "/organizations/",
		"/project/", "/projects/",
		"/env/", "/environments/",
		"/environment/", "/environments/",
		"/cred/", "/credentials/",
		"/credential/", "/credentials/",
		"/component/", "/components/",
	)

	return replacer.Replace(uri)
}

type Reference struct {
	Path   string
	Params url.Values
}

// NewResourceReference will parse cy:// uri and turn it into a resource reference.
// a basic var interpolation will be made on {org}/{organization}, {project}, {env},
// {environment}, {component} to use the current cli context.
func NewResourceReference(uri string) (*Reference, error) {
	// expand shortcuts like org -> organizations
	expandedURI := ExpandShortcuts(uri)

	// expand context like {org} -> CY_ORG value
	var (
		v           = viper.GetViper()
		org         = v.GetString("org")
		project     = v.GetString("project")
		env         = v.GetString("env")
		component   = v.GetString("component")
		varReplacer = strings.NewReplacer(
			"{org}", org,
			"{organization}", org,
			"{project}", project,
			"{env}", env,
			"{environment}", env,
			"{component}", component,
		)
	)
	finalURI := varReplacer.Replace(expandedURI)

	resourceURL, err := url.Parse(finalURI)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url %q: %w", finalURI, err)
	}

	query, err := url.ParseQuery(resourceURL.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to parse query parameters on %q: %w", finalURI, err)
	}

	return &Reference{
		Path:   strings.Replace(finalURI, "cy://", "/", 1),
		Params: query,
	}, nil
}
