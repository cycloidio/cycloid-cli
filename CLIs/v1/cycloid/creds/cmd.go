package creds

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "cred",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}

	cmd.AddCommand(NewCreateCommand(),
		NewUpdateCommand(),
		NewDeleteCommand(),
		NewListCommand(),
		NewGetCommand())

	return cmd
}

// /organizations/{organization_canonical}/credentials
// get: getCredentials
// Return all the Credentials
