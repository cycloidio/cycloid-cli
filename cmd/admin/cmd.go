package admin

import (
	"github.com/spf13/cobra"
)

// NewCommands returns the instance administration commands, reserved to the admins of the root
// organization
func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "admin",
		Short: "Administrate the instance, for the root organization admins only",
		Example: `
	# List the users without an organization
	cy --org my-root-org admin users list --without-organization

	# Delete a user, including one without an organization
	cy --org my-root-org admin users delete --email user@example.com
`,
	}

	cmd.AddCommand(NewUsersCommands())

	return cmd
}

// NewUsersCommands returns the commands managing the users of the whole instance
func NewUsersCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "users",
		Aliases: []string{"user"},
		Short:   "Manage the users of the whole instance",
	}

	cmd.AddCommand(NewListUsersCommand(), NewDeleteUserCommand())

	return cmd
}
