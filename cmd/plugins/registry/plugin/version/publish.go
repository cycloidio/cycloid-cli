package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewPublishCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "publish",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cyargs.CompletePluginVersionID,
		Short:             "Publish a new version of a plugin",
		Example: `
  # Publish an archive-based plugin version
  cy plugin registry plugin version publish --registry my-registry --plugin my-plugin --url https://example.com/plugin-v1.2.tar.gz

  # Publish a Docker image-based plugin version
  cy plugin registry plugin version publish --registry my-registry --plugin my-plugin --docker-image docker-registry:5000/org/plugin:42
`,
		RunE: publishVersion,
	}

	_ = cmd.MarkFlagRequired(cyargs.AddRegistryFlag(cmd))
	_ = cmd.MarkFlagRequired(cyargs.AddPluginFlag(cmd))
	cyargs.AddURLFlag(cmd, "URL of the plugin version archive (mutually exclusive with --docker-image)")
	cyargs.AddDockerImageFlag(cmd)
	return cmd
}

func publishVersion(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	url, err := cyargs.GetURL(cmd)
	if err != nil {
		return fmt.Errorf("unable to get --url flag: %w", err)
	}

	dockerImage, err := cyargs.GetDockerImage(cmd)
	if err != nil {
		return fmt.Errorf("unable to get --docker-image flag: %w", err)
	}

	if url == "" && dockerImage == "" {
		return fmt.Errorf("one of --url or --docker-image is required")
	}
	if url != "" && dockerImage != "" {
		return fmt.Errorf("--url and --docker-image are mutually exclusive")
	}

	versionURL := url
	if dockerImage != "" {
		versionURL = dockerImage
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	registryID, pluginID, err := resolveRegistryAndPlugin(org, cmd, m)
	if err != nil {
		return err
	}

	result, _, err := m.CreatePluginVersion(org, registryID, pluginID, versionURL)
	return cyout.PrintWithOptions(cmd, result, err, "unable to publish plugin version", printer.Options{})
}
