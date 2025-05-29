package internal

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/version"
)

func Error(out io.Writer, msg string) {
	switch viper.GetString("verbosity") {
	case "info", "debug", "warning":
		// This is still dirty, we should detect if the current
		// terminal is able to support colors
		// But that would be for another PR.
		fmt.Fprintf(out, "\033[1;31merror:\033[0m %s", msg)
	default:
	}
}

func Warning(out io.Writer, msg string) {
	switch viper.GetString("verbosity") {
	case "info", "debug", "warning":
		// This is still dirty, we should detect if the current
		// terminal is able to support colors
		// But that would be for another PR.
		fmt.Fprintf(out, "\033[1;35mwarning:\033[0m %s", msg)
	default:
	}
}

func Debug(msg ...any) {
	switch viper.GetString("verbosity") {
	case "debug":
		// This is still dirty, we should detect if the current
		// terminal is able to support colors
		// But that would be for another PR.
		fmt.Fprintf(os.Stderr, "\033[1;34mdebug:\033[0m ")
		fmt.Fprintln(os.Stderr, msg...)
	default:
	}
}

func CheckAPIAndCLIVersion(cmd *cobra.Command, args []string) error {
	cliVersion := version.Version

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)
	d, err := m.GetAppVersion()

	if err != nil {
		Warning(cmd.ErrOrStderr(), "Unable to get the API version\n")
		return nil
	}

	apiVersion := fmt.Sprintf("%s", *d.Version)
	reg := regexp.MustCompile("^([^-]+)(-.*)$")
	apiVersion = reg.ReplaceAllString(apiVersion, "${1}")

	if cliVersion != apiVersion {
		Warning(cmd.ErrOrStderr(), fmt.Sprintf("CLI version %s does not match the API version. You should consider to download CLI version %s\n", cliVersion, apiVersion))
	}
	return nil
}
