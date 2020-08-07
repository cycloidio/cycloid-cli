package root

import (
	"github.com/cycloidio/youdeploy-cli/client/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

func NewAPI() *client.APIClient {
	cfg := client.DefaultTransportConfig()
	cfg = cfg.WithHost("127.0.0.1")
	cfg = cfg.WithSchemes([]string{"http"})

	api := client.NewHTTPClientWithConfig(strfmt.Default, cfg)

	// Hack because https://github.com/go-swagger/go-swagger/issues/1899
	// none of producers: map[application/json:0x7f7dff8da3d0 application/octet-stream:0x7f7dff8d8ff0 application/xml:0x7f7dff8db1d0 text/csv:0x7f7dff8d9da0 text/html:0x7f7dff8daa60 text/plain:0x7f7dff8daa60] registered. try application/vnd.cycloid.io.v1+json
	tr := api.Transport.(*httptransport.Runtime)
	tr.Producers["application/vnd.cycloid.io.v1+json"] = runtime.JSONProducer()
	api.SetTransport(tr)

	return api
}
