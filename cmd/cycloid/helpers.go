package root

import (
	"os"

	"net/url"

	"github.com/cycloidio/youdeploy-cli/client/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/spf13/viper"

	strfmt "github.com/go-openapi/strfmt"
)

func NewAPI() *client.APIClient {

	cfg := client.DefaultTransportConfig()

	rawApiUrl := viper.GetString("api-url")

	apiUrl, err := url.Parse(rawApiUrl)
	if err == nil && apiUrl.Host != "" {
		cfg = cfg.WithHost(apiUrl.Host)
		cfg = cfg.WithSchemes([]string{apiUrl.Scheme})
		cfg = cfg.WithBasePath(apiUrl.Path)
	}

	// cfg = cfg.WithHost("http-api-staging.cycloid.io")
	// cfg = cfg.WithSchemes([]string{"https"})
	// cfg = cfg.WithHost("127.0.0.1:80")
	// cfg = cfg.WithSchemes([]string{"http"})
	// cfg = cfg.WithBasePath("/api")

	// cfg = cfg.WithHost(apiUrl.Host)

	// hostUrl = fmt.Sprintf("%s/%s/%s-%s.tfstate", project)

	api := client.NewHTTPClientWithConfig(strfmt.Default, cfg)

	// Hack because https://github.com/go-swagger/go-swagger/issues/1899
	// none of producers: map[application/json:0x7f7dff8da3d0 application/octet-stream:0x7f7dff8d8ff0 application/xml:0x7f7dff8db1d0 text/csv:0x7f7dff8d9da0 text/html:0x7f7dff8daa60 text/plain:0x7f7dff8daa60] registered. try application/vnd.cycloid.io.v1+json
	tr := api.Transport.(*httptransport.Runtime)
	tr.Producers["application/vnd.cycloid.io.v1+json"] = runtime.JSONProducer()
	// tr.DefaultAuthentication = httptransport.BearerToken("token")
	// api.SetTransport(tr)
	return api
}

func ClientCredentials() runtime.ClientAuthInfoWriter {
	token := os.Getenv("TOKEN")
	if token == "" {
		panic("env var TOKEN not found")
	}

	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		r.SetHeaderParam("Authorization", "Bearer "+token)
		return nil
	})
}
