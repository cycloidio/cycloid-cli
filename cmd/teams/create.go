package teams

import (
	stderrors "errors"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
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

	displayName, teamCanonical, err := apiclient.NameOrCanonical(&teamName, &team)
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

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)
	var newTeam *models.Team

	if allowUpdate {
		currentTeam, _, err := m.GetTeam(org, teamCanonical)
		if err != nil {
			var apiErr *apiclient.APIResponseError
			if !stderrors.As(err, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
				return cyout.PrintWithOptions(cmd, nil, err, "failed to get team to check if it exists", printer.Options{})
			}
		} else {
			newTeam, _, err = m.UpdateTeam(
				org, ptr.Ptr(utils.CoalesceNonZero(displayName, ptr.Value(currentTeam.Name))),
				currentTeam.Canonical, ptr.Ptr(utils.CoalesceNonZero(teamOwner, ptr.Value(ptr.Value(currentTeam.Owner).Username))), roles,
			)
			return cyout.PrintWithOptions(cmd, newTeam, err, "failed to UpdateTeam", printer.Options{})
		}
	}

	newTeam, _, err = m.CreateTeam(org, &displayName, &teamCanonical, &teamOwner, roles)
	return cyout.PrintWithOptions(cmd, newTeam, err, "failed to CreateTeam", printer.Options{})
}
