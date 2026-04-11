package version

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewPublishCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "publish <registry> <plugin>",
		Args:              cobra.ExactArgs(2),
		ValidArgsFunction: cyargs.CompleteRegistryPluginID,
		Short:             "[beta] Publish a new version of a plugin",
		Example: `
  # Publish an archive-based plugin version
  cy beta plugin registry plugin version publish my-registry my-plugin --url https://example.com/plugin-v1.2.tar.gz

  # Publish a Docker image-based plugin version
  cy beta plugin registry plugin version publish my-registry my-plugin --docker-image docker-registry:5000/org/plugin:42
`,
		RunE: publishVersion,
	}

	cyargs.AddURLFlag(cmd, "URL of the plugin version archive (mutually exclusive with --docker-image)")
	cyargs.AddDockerImageFlag(cmd)
	return cmd
}

func publishVersion(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	registryID, pluginID, err := resolveRegistryAndPlugin(org, args, m)
	if err != nil {
		return err
	}

	url, err := cyargs.GetURL(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get --url flag")
	}

	dockerImage, err := cyargs.GetDockerImage(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get --docker-image flag")
	}

	if url == "" && dockerImage == "" {
		return errors.New("one of --url or --docker-image is required")
	}
	if url != "" && dockerImage != "" {
		return errors.New("--url and --docker-image are mutually exclusive")
	}

	versionURL := url
	if dockerImage != "" {
		versionURL = dockerImage
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	result, _, err := m.CreatePluginVersion(org, registryID, pluginID, versionURL)
	return printer.SmartPrint(p, result, err, "unable to publish plugin version", printer.Options{}, cmd.OutOrStdout())
}
