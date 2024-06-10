package common

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"

	"net/url"

	"github.com/cycloidio/cycloid-cli/client/client"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/config"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	strfmt "github.com/go-openapi/strfmt"
)

var orgRe = regexp.MustCompile(`\(\$ organization_canonical \$\)`)
var envRe = regexp.MustCompile(`\(\$ environment \$\)`)
var projRe = regexp.MustCompile(`\(\$ project \$\)`)

func RequiredPersistentFlag(withFlag func(cmd *cobra.Command) string, cmd *cobra.Command) {
	flagName := withFlag(cmd)
	cmd.MarkPersistentFlagRequired(flagName)
}
func RequiredFlag(withFlag func(cmd *cobra.Command) string, cmd *cobra.Command) {
	flagName := withFlag(cmd)
	cmd.MarkFlagRequired(flagName)
}

type CycloidContext struct {
	Org     string
	Env     string
	Project string
}

func ReplaceCycloidVars(ctx CycloidContext, text []byte) []byte {
	if ctx.Org != "" {
		text = orgRe.ReplaceAll(text, []byte(ctx.Org))
	}
	if ctx.Env != "" {
		text = envRe.ReplaceAll(text, []byte(ctx.Env))
	}
	if ctx.Project != "" {
		text = projRe.ReplaceAll(text, []byte(ctx.Project))
	}
	return text
}

func ReplaceCycloidVarsString(ctx CycloidContext, text string) string {
	if ctx.Org != "" {
		text = orgRe.ReplaceAllString(text, ctx.Org)
	}
	if ctx.Env != "" {
		text = envRe.ReplaceAllString(text, ctx.Env)
	}
	if ctx.Project != "" {
		text = projRe.ReplaceAllString(text, ctx.Project)
	}
	return text
}

func IsInList(pattern string, list []string) bool {
	for _, x := range list {
		if x == pattern {
			return true
		}
	}
	return false
}

func GetPipelineName(project, env string) string {
	return fmt.Sprintf("%s-%s", project, env)
}

type APIConfig struct {
	URL      string
	Insecure bool
	Token    string
}

type APIOptions func(acfg *APIConfig)

func WithURL(u string) APIOptions {
	return func(acfg *APIConfig) {
		acfg.URL = u
	}
}

func WithInsecure(i bool) APIOptions {
	return func(acfg *APIConfig) {
		acfg.Insecure = i
	}
}

func WithToken(t string) APIOptions {
	return func(acfg *APIConfig) {
		acfg.Token = t
	}
}

type APIClient struct {
	*client.API

	Config APIConfig
}

func NewAPI(opts ...APIOptions) *APIClient {
	cfg := client.DefaultTransportConfig()

	acfg := APIConfig{
		URL:      viper.GetString("api-url"),
		Insecure: viper.GetBool("insecure"),
		Token:    "",
	}

	for _, o := range opts {
		o(&acfg)
	}

	apiUrl, err := url.Parse(acfg.URL)
	if err == nil && apiUrl.Host != "" {
		cfg = cfg.WithHost(apiUrl.Host)
		cfg = cfg.WithSchemes([]string{apiUrl.Scheme})
		cfg = cfg.WithBasePath(apiUrl.Path)
	}

	api := client.NewHTTPClientWithConfig(strfmt.Default, cfg)

	rt, err := httptransport.TLSTransport(httptransport.TLSClientOptions{InsecureSkipVerify: acfg.Insecure})
	if err != nil {
		// TODO: error handling ...
		fmt.Printf("unable to create round tripper: %v", err)
		return nil
	}

	// Hack because https://github.com/go-swagger/go-swagger/issues/1899
	// none of producers: map[application/json:0x7f7dff8da3d0 application/octet-stream:0x7f7dff8d8ff0 application/xml:0x7f7dff8db1d0 text/csv:0x7f7dff8d9da0 text/html:0x7f7dff8daa60 text/plain:0x7f7dff8daa60] registered. try application/vnd.cycloid.io.v1+json
	tr := api.Transport.(*httptransport.Runtime)
	tr.Producers["application/vnd.cycloid.io.v1+json"] = runtime.JSONProducer()
	tr.Transport = rt
	// tr.DefaultAuthentication = httptransport.BearerToken("token")
	// api.SetTransport(tr)
	return &APIClient{
		API: api,

		Config: acfg,
	}
}

func (a *APIClient) Credentials(org *string) runtime.ClientAuthInfoWriter {
	var token = a.Config.Token
	if token == "" {
		// we first try to get the token from the env variable
		token = os.Getenv("TOKEN")
	}
	// if the token is not set with env variable we try to fetch
	// him from the config (if the user is logged)
	if len(token) == 0 {
		// we fetch the running config
		config, _ := config.Read()

		if org == nil {
			return nil
		}

		// we try to find a token for this `org`
		if t, ok := config.Organizations[*org]; ok {
			token = t.Token
		} else {
			return nil
		}
	}
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		r.SetHeaderParam("Authorization", "Bearer "+token)
		return nil
	})
}

// GenerateCanonical will generate a canonical from the
// name passed in input
func GenerateCanonical(name string) string {
	var (
		extraspaces = regexp.MustCompile(`\s+`)
		spaces      = regexp.MustCompile(`[^a-zA-z0-9_\-]`)
		canonical   string
	)

	canonical = extraspaces.ReplaceAllString(name, " ")
	canonical = spaces.ReplaceAllString(canonical, "-")

	return strings.ToLower(canonical)
}

// From a *models.FormEntity, retrieve a value associated with its type
// if getCurrent is true, we return the current value in priority
// otherwise, we get the default value
// if no default is set, we get a zeroed value of the correct type
// Return nil if the type is invalid.
// TODO: Could this be better with generics ?
func EntityGetValue(entity *models.FormEntity, getCurrent bool) any {
	switch *entity.Type {
	case "string":
		if getCurrent { // Try to get current value if asked.
			value, ok := entity.Current.(string)
			if ok {
				return value
			}
		}

		// Try to get the default value
		value, ok := entity.Default.(string)
		if ok {
			return value
		}

		// Else return a valid typed zeroed value
		return ""
	case "integer":
		if getCurrent { // Try to get current value if asked.
			value, ok := entity.Current.(int64)
			if ok {
				return value
			}
		}

		// Try to get the default value
		value, ok := entity.Default.(int64)
		if ok {
			return value
		}

		// Else return a valid typed zeroed value
		return 0
	case "float":
		if getCurrent { // Try to get current value if asked.
			value, ok := entity.Current.(float64)
			if ok {
				return value
			}
		}

		// Try to get the default value
		value, ok := entity.Default.(float64)
		if ok {
			return value
		}

		// Else return a valid typed zeroed value
		return 0
	case "boolean":
		if getCurrent { // Try to get current value if asked.
			value, ok := entity.Current.(bool)
			if ok {
				return value
			}
		}

		// Try to get the default value
		value, ok := entity.Default.(bool)
		if ok {
			return value
		}

		// Else return a valid typed zeroed value
		return false
	case "array":
		if getCurrent { // Try to get current value if asked.
			value, ok := entity.Current.([]any)
			if ok {
				return value
			}
		}

		// Try to get the default value
		value, ok := entity.Default.([]any)
		if ok {
			return value
		}

		// Else return a valid typed zeroed value
		return []any{}
	case "map":
		if getCurrent { // Try to get current value if asked.
			value, ok := entity.Current.(map[string]any)
			if ok {
				return value
			}
		}

		// Try to get the default value
		value, ok := entity.Default.(map[string]any)
		if ok {
			return value
		}

		// Else return a valid typed zeroed value
		return make(map[string]any)
	default:
		return nil
	}
}

func ParseFormsConfig(conf *models.ProjectEnvironmentConfig, useCase string, getCurrent bool) (vars map[string]map[string]map[string]any, err error) {
	form, err := GetFormsUseCase(conf.Forms.UseCases, useCase)
	if err != nil {
		return nil, errors.Wrap(err, "failed to extract forms data from project config.")
	}

	vars = make(map[string]map[string]map[string]any)
	for _, section := range form.Sections {
		if section == nil {
			continue
		}

		var groups = make(map[string]map[string]any)
		for _, group := range section.Groups {
			if group == nil {
				continue
			}

			vars := make(map[string]any)

			for _, varEntity := range group.Vars {
				if varEntity == nil {
					continue
				}

				value := EntityGetValue(varEntity, getCurrent)
				// We have to strings.ToLower() the keys otherwise, it will not be
				// recognized as input for a create-env
				vars[strings.ToLower(*varEntity.Name)] = value
			}

			groups[strings.ToLower(*group.Name)] = vars
		}

		vars[strings.ToLower(*section.Name)] = groups
	}

	return vars, nil
}

func GetFormsUseCase(formsUseCases []*models.FormUseCase, useCase string) (*models.FormUseCase, error) {
	if formsUseCases == nil {
		return nil, fmt.Errorf("got empty forms use case")
	}

	for _, form := range formsUseCases {
		if *form.Name == useCase {
			return form, nil
		}
	}

	return nil, errors.New(fmt.Sprint("failed to find usecase:", useCase, "in form input:", formsUseCases))
}

// Update map 'm' with field 'field' to 'value'
// the field must be in dot notation
// e.g. field='one.nested.key' value='myValue'
// If the map is nil, it will be created
func UpdateMapField(field string, value any, m map[string]any) error {
	keys := strings.Split(field, ".")

	if m == nil {
		m = make(map[string]any)
	}

	if len(keys) == 1 {
		m[keys[0]] = value
		return nil
	}

	child, exists := m[keys[0]]
	if exists && reflect.ValueOf(child).Kind() == reflect.Map {
		childMap, ok := child.(map[string]any)
		if !ok {
			return fmt.Errorf("failed to parse nested map: %v\n%v", child, childMap)
		}
		return UpdateMapField(strings.Join(keys[1:], "."), value, childMap)
	}

	child = make(map[string]any)
	err := UpdateMapField(strings.Join(keys[1:], "."), value, child.(map[string]any))
	if err != nil {
		return err
	}

	m[keys[0]] = child

	return nil
}
