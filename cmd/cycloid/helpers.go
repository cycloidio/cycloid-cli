package root

import (
	"github.com/cycloidio/youdeploy-cli/client/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

func NewAPI() *client.APIClient {
	cfg := client.DefaultTransportConfig()
	cfg = cfg.WithHost("http-api-staging.cycloid.io")
	cfg = cfg.WithSchemes([]string{"https"})
	// cfg = cfg.WithHost("127.0.0.1")
	// cfg = cfg.WithSchemes([]string{"http"})

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
	token := "eyJhbGciOiJIUzI1NiIsImtpZCI6IjJmMjEyMmRlLTYzZjItNGVlYy05YzZmLWM2YWJiM2UxZjAwNyIsInR5cCI6IkpXVCJ9.eyJjeWNsb2lkIjp7InVzZXIiOnsiaWQiOjIsInVzZXJuYW1lIjoiY3ljbG9pZF9zZXJhZiIsImdpdmVuX25hbWUiOiJKdWxpZW4iLCJmYW1pbHlfbmFtZSI6IlN5eCIsInBpY3R1cmVfdXJsIjoiaHR0cHM6Ly9hdmF0YXJzMi5naXRodWJ1c2VyY29udGVudC5jb20vdS8zOTMzMjQ_cz00MDBcdTAwMjZ1PTIxNTE5ZmQwYzUyMDI3NTgxMWYyZTNmYmIxZmIxZmE4ZTQxZTM2MDBcdTAwMjZ2PTQiLCJsb2NhbGUiOiJlbiJ9LCJvcmdhbml6YXRpb24iOnsiaWQiOjE1MiwiY2Fub25pY2FsIjoiY3ljbG9pZC1kZW1vIiwibmFtZSI6IkN5Y2xvaWQtZGVtbyIsImJsb2NrZWQiOltdLCJoYXNfY2hpbGRyZW4iOmZhbHNlfSwicGVybWlzc2lvbnMiOnsib3JnYW5pemF0aW9uOmFkbWluIjpbXX19LCJhdWQiOiJjdXN0b21lciIsImV4cCI6MTU5ODQzNjQ5NCwianRpIjoiZWVkMzJhZTItNDhlMi00N2EzLTg4ZTEtYzUyMzY3YTk3ZDIxIiwiaWF0IjoxNTk4MzUwMDk0LCJpc3MiOiJodHRwczovL2N5Y2xvaWQuaW8iLCJuYmYiOjE1OTgzNTAwOTQsInN1YiI6Imh0dHBzOi8vY3ljbG9pZC5pbyJ9.GieyC-QtOubmQO5XAwnfhiKZUcDi9adOM4hlfIruGx8"

	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		r.SetHeaderParam("Authorization", "Bearer "+token)
		return nil
	})
}
