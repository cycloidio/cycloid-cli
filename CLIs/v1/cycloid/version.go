package root

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  version,
	}
	return cmd

}

func version(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	d, err := m.GetAppVersion()

	fmt.Printf("ver: %s    rev: %s    branch: %s  \n", *d.Version, *d.Revision, *d.Branch)

	fmt.Printf("%+v\n", err)
	return nil
}
