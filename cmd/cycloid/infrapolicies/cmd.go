package infrapolicies

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

var (
	example = `
	# Manage infrapolicies and validate Terraform plan against
	cy --org my-org api-key validate
`
	short = "Manage infrapolicies"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use: "infrapolicy",
		Aliases: []string{
			"infra-policies",
			"infra-policy",
			"ip",
			"infrapolicies",
		},
		Example: example,
		Short:   short,
	}

	common.RequiredFlag(common.WithFlagProject, cmd)
	common.RequiredFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewValidateCommand())
	return cmd
}
