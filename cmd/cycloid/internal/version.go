package internal

import (
	"fmt"
	"os"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/youdeploy-cli/internal/version"
	"github.com/spf13/cobra"
)

func warning(msg string) {
	fmt.Fprintf(os.Stderr, "\033[1;35m%s\033[0m\n", msg)
}

func CheckAPIAndCLIVersion(cmd *cobra.Command, args []string) error {
	cliVersion := version.Version

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)
	d, err := m.GetAppVersion()
	if err != nil {
		warning("Warning: Unable to get the API version\n")
		return nil
	}
	apiVersion := *d.Version

	if cliVersion != apiVersion {
		warning(fmt.Sprintf("Warning: CLI version does not match the API version. You should consider to download CLI version %s\n", apiVersion))
	}

	return nil
}
