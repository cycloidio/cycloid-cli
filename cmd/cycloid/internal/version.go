package internal

import (
	"fmt"
	"io"
	"regexp"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/version"
	"github.com/spf13/cobra"
)

func warning(out io.Writer, msg string) {
	fmt.Fprintf(out, "\033[1;35m%s\033[0m\n", msg)
}

func CheckAPIAndCLIVersion(cmd *cobra.Command, args []string) error {
	cliVersion := version.Version

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)
	d, err := m.GetAppVersion()

	if err != nil {
		warning(cmd.ErrOrStderr(), "Warning: Unable to get the API version\n")
		return nil
	}

	apiVersion := fmt.Sprintf("%s", *d.Version)
	reg := regexp.MustCompile("^([^-]+)(-.*)$")
	apiVersion = reg.ReplaceAllString(apiVersion, "${1}")

	if cliVersion != apiVersion {
		warning(cmd.ErrOrStderr(), fmt.Sprintf("Warning: CLI version %s does not match the API version. You should consider to download CLI version %s\n", cliVersion, apiVersion))
	}
	return nil
}
