package common

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cycloidio/cycloid-cli/client/client"
	"github.com/cycloidio/cycloid-cli/config"
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

	if !strings.HasPrefix(acfg.URL, "https://") {
		if strings.Contains(acfg.URL, "localhost") {
			// This handles the weird case of localhost:3001
			// being interpreted as scheme=localhost by url.Parse
			acfg.URL = "http://" + acfg.URL
		}
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
		API:    api,
		Config: acfg,
	}
}

func (a *APIClient) GetToken(org string) string {
	token := a.Config.Token

	// we first try to get the token from the env variable
	if token == "" {
		for _, env_var := range []string{"CY_API_KEY", "CY_API_TOKEN", "TOKEN"} {
			token, ok := os.LookupEnv(env_var)

			// Still display warning for future deprecation
			if ok && env_var == "TOKEN" {
				fmt.Fprintln(os.Stderr, "TOKEN env var is deprecated, please use CY_API_KEY instead")
			}

			if ok && len(token) != 0 {
				break
			}
		}
	}

	// if the token is not set with env variable we try to fetch
	// him from the config (if the user is logged)
	if len(token) == 0 {
		// we fetch the running config
		config, _ := config.Read()

		// we try to find a token for this `org`
		if t, ok := config.Organizations[org]; ok {
			token = t.Token
		}
	}

	return token
}

func (a *APIClient) Credentials(org *string) runtime.ClientAuthInfoWriter {
	if org == nil {
		return nil
	}

	token := a.GetToken(*org)

	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		if len(token) == 0 {
			return errors.New("No API_KEY was provided, please provide one by CY_API_KEY env var or using cy login.")
		}

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

// Update map 'm' with field 'field' to 'value'
// the field must be in dot notation
// e.g. field='one.nested.key' value='myValue'
// If the map is nil, it will be created
func UpdateMapField(field string, value string, m map[string]map[string]map[string]interface{}) error {
	keys := strings.Split(field, ".")

	if len(keys) != 3 {
		return errors.New("key=val update failed, you can only update a value using `section.group.var=value` syntax")
	}

	if m == nil {
		m = make(map[string]map[string]map[string]any)
	}

	// Try to detect JSON first
	// we strip value for space and newline in begin/end of the string
	trimmedValue := strings.TrimSpace(value)
	if strings.HasPrefix(trimmedValue, "[") && strings.HasSuffix(trimmedValue, "]") || strings.HasPrefix(trimmedValue, "{") && strings.HasSuffix(trimmedValue, "}") {
		var data interface{}
		err := json.Unmarshal([]byte(trimmedValue), &data)
		if err != nil {
			return errors.Wrapf(err, "invalid JSON value in key=val update with value '%s'", trimmedValue)
		}

		m[keys[0]][keys[1]][keys[2]] = data
		return nil
	}

	// We will prioritize the use of quotes to explicitly define strings values
	// This allow users to circumvent issues in case of strings that could be parsed
	// as other types
	if strings.HasPrefix(trimmedValue, `"`) && strings.HasSuffix(trimmedValue, `"`) ||
		strings.HasPrefix(trimmedValue, "'") && strings.HasSuffix(trimmedValue, "'") {
		m[keys[0]][keys[1]][keys[2]] = trimmedValue[1 : len(trimmedValue)-1]
		return nil
	}

	// Detect standard types
	// numbers, we do all as float since JSON doesn't care
	// Important! We parse number firsts, since 1 and 0 are considered bools by strconv.ParseBool
	float, err := strconv.ParseFloat(value, 64)
	if err == nil {
		m[keys[0]][keys[1]][keys[2]] = float
		return nil
	}

	// bools
	boolean, err := strconv.ParseBool(value)
	if err == nil {
		m[keys[0]][keys[1]][keys[2]] = boolean
		return nil
	}

	// null
	if strings.ToLower(value) == "null" {
		m[keys[0]][keys[1]][keys[2]] = nil
		return nil
	}

	// if all type conversion failed, consider the value as string
	m[keys[0]][keys[1]][keys[2]] = value
	return nil
}
