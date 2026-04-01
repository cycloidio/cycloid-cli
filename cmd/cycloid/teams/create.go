package teams

import (
	stderrors "errors"
	"fmt"
	"net/http"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/cycloidio/cycloid-cli/internal/utils"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/spf13/cobra"
)

func NewCreateTeamCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Create a team.",
		Example: `cy team create --name "My cool team" --role organization-admin`,
		RunE:    createTeam,
		Args:    cobra.NoArgs,
	}

	cmd.Flags().Bool("update", false, "allow update if team exists")
	cmd.MarkFlagsOneRequired(
		cyargs.AddTeamNameFlag(cmd),
		cyargs.AddTeamFlag(cmd),
	)
	cmd.MarkFlagRequired(cyargs.AddTeamRolesFlag(cmd))
	cyargs.AddTeamOwnerFlag(cmd)
	return cmd
}

func createTeam(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	team, err := cyargs.GetTeam(cmd)
	if err != nil {
		return err
	}

	teamName, err := cyargs.GetTeamName(cmd)
	if err != nil {
		return err
	}

	teamOwner, err := cyargs.GetTeamOwner(cmd)
	if err != nil {
		return err
	}

	displayName, teamCanonical, err := middleware.NameOrCanonical(&teamName, &team)
	if err != nil {
		return err
	}

	roles, err := cyargs.GetTeamRoles(cmd)
	if err != nil {
		return err
	}

	allowUpdate, err := cmd.Flags().GetBool("update")
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return fmt.Errorf("failed to get printer for output type %q: %w", output, err)
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)
	var newTeam *models.Team

	if allowUpdate {
		currentTeam, _, err := m.GetTeam(org, teamCanonical)
		if err != nil {
			var apiErr *middleware.APIResponseError
			if !stderrors.As(err, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
				return printer.SmartPrint(p, nil, err, "failed to get team to check if it exists", printer.Options{}, cmd.OutOrStderr())
			}
		} else {
			newTeam, _, err = m.UpdateTeam(
				org, ptr.Ptr(utils.CoalesceNonZero(displayName, ptr.Value(currentTeam.Name))),
				currentTeam.Canonical, ptr.Ptr(utils.CoalesceNonZero(teamOwner, ptr.Value(ptr.Value(currentTeam.Owner).Username))), roles,
			)
			if err != nil {
				return printer.SmartPrint(p, nil, err, "failed to UpdateTeam", printer.Options{}, cmd.OutOrStderr())
			}

			return printer.SmartPrint(p, newTeam, nil, "", printer.Options{}, cmd.OutOrStdout())
		}
	}

	newTeam, _, err = m.CreateTeam(org, &displayName, &teamCanonical, &teamOwner, roles)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to CreateTeam", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, newTeam, nil, "", printer.Options{}, cmd.OutOrStdout())
}
