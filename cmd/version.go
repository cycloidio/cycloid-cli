package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	clientversion "github.com/cycloidio/cycloid-cli/internal/version"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

func NewVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Args:  cobra.NoArgs,
		Short: "Get the client and server version",
		Example: `
	# get the version in JSON format
	cy version -o json
`,
		RunE: getVersion,
	}
	return cmd
}

// versionOutput is a single flat object. The server fields keep the pre-CLI-149
// top-level keys (version, revision, branch) so existing JSON consumers keep working;
// api_url, client_* and error are additive.
type versionOutput struct {
	Version        string `json:"version,omitempty"`
	Revision       string `json:"revision,omitempty"`
	Branch         string `json:"branch,omitempty"`
	APIURL         string `json:"api_url"`
	ClientVersion  string `json:"client_version"`
	ClientRevision string `json:"client_revision"`
	ClientBranch   string `json:"client_branch"`
	Error          string `json:"error,omitempty"`
}

func versionOrDefault(v, fallback string) string {
	if v == "" {
		return fallback
	}
	return v
}

func (o *versionOutput) setClientInfo() {
	o.ClientVersion = versionOrDefault(clientversion.Version, "dev")
	o.ClientRevision = versionOrDefault(clientversion.Revision, "unknown")
	o.ClientBranch = versionOrDefault(clientversion.Branch, "unknown")
}

// buildVersionOutput assembles the flat client + server version object. fetch is injected
// so the rendering logic can be unit-tested without a live API call. The fetch error is
// returned alongside the (still populated) output.
func buildVersionOutput(apiURL string, fetch func() (*models.AppVersion, *http.Response, error)) (*versionOutput, error) {
	out := &versionOutput{APIURL: apiURL}
	out.setClientInfo()

	d, _, err := fetch()
	if err != nil {
		out.Error = err.Error()
		return out, err
	}

	out.Version = ptr.Value(d.Version)
	out.Revision = ptr.Value(d.Revision)
	out.Branch = ptr.Value(d.Branch)
	return out, nil
}

// versionTableColumns: the server columns come first in the order the pre-CLI-149
// reflection-based table rendered models.AppVersion (struct field order).
var versionTableColumns = []string{
	"Branch", "Revision", "Version",
	"APIURL", "ClientVersion", "ClientRevision", "ClientBranch",
}

func versionTableOptions() printer.Options {
	return printer.Options{
		Columns: versionTableColumns,
		Transform: func(obj interface{}) map[string]string {
			v, ok := obj.(*versionOutput)
			if !ok {
				return nil
			}
			return map[string]string{
				"Branch":         v.Branch,
				"Revision":       v.Revision,
				"Version":        v.Version,
				"APIURL":         v.APIURL,
				"ClientVersion":  v.ClientVersion,
				"ClientRevision": v.ClientRevision,
				"ClientBranch":   v.ClientBranch,
			}
		},
	}
}

func getVersion(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	out, fetchErr := buildVersionOutput(api.Config.URL, m.GetAppVersion)
	if err := cyout.PrintWithOptions(cmd, out, nil, "", versionTableOptions()); err != nil {
		return err
	}
	if fetchErr != nil {
		// main.go prints this once as "Error: ..." on stderr.
		return fmt.Errorf("unable to get API version from %s: %w", out.APIURL, fetchErr)
	}
	return nil
}
