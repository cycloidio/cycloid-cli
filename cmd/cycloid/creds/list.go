package creds

import (
	"fmt"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  list,
	}

	WithFlagType(cmd)
	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	credT, err := cmd.Flags().GetString("type")
	if err != nil {
		return err
	}

	creds, err := m.ListCredentials(org, credT)
	if err != nil {
		return err
	}

	for _, c := range creds {
		fmt.Printf("id: %d    type: %s    path: %s  \n", *c.ID, *c.Type, *c.Path)

	}
	fmt.Println(creds)
	fmt.Printf("%+v\n", err)
	return nil
}
