package common

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"net/url"

	"github.com/cycloidio/cycloid-cli/client/client"
	"github.com/cycloidio/cycloid-cli/config"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
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

func NewAPI() *client.APIClient {
	cfg := client.DefaultTransportConfig()

	rawApiUrl := viper.GetString("api-url")
	insecure := viper.GetBool("insecure")

	apiUrl, err := url.Parse(rawApiUrl)
	if err == nil && apiUrl.Host != "" {
		cfg = cfg.WithHost(apiUrl.Host)
		cfg = cfg.WithSchemes([]string{apiUrl.Scheme})
		cfg = cfg.WithBasePath(apiUrl.Path)
	}

	api := client.NewHTTPClientWithConfig(strfmt.Default, cfg)

	rt, err := httptransport.TLSTransport(httptransport.TLSClientOptions{InsecureSkipVerify: insecure})
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
	return api
}

func ClientCredentials(org *string) runtime.ClientAuthInfoWriter {
	var token string
	// we first try to get the token from the env variable
	token = os.Getenv("TOKEN")
	// if the token is not set with env variable we try to fetch
	// him from the config (if the user is logged)
	if len(token) == 0 {
		// we fetch the running config
		config, _ := config.ReadConfig()
		if org == nil {
			// if there is no org, it means we need to use
			// the user token (the first token in the config)
			token = config.Token
		} else {
			// we try to find a token for this `org`
			if t, ok := config.Organizations[*org]; ok {
				token = t.Token
			} else {
				return nil
			}
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
