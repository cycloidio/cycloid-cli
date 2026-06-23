package teams

import (
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewTeamsCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "teams",
		Aliases: []string{"team"},
		Short:   "Commands to manage teams",
	}
	cmd.AddCommand(
		NewTeamMembersCommand(),
		NewCreateTeamCommand(),
		NewUpdateTeamCommand(),
		NewGetTeamCommand(),
		NewDeleteTeamCommand(),
		NewListTeamCommand(),
	)
	return cmd
}

var teamTableOptions = printer.Options{
	Columns:    []string{"Canonical", "Name", "MemberCount", "Roles"},
	Identifier: "Canonical",
	Transform: func(obj interface{}) map[string]string {
		t, ok := obj.(*models.Team)
		if !ok {
			return map[string]string{}
		}
		canonical := ""
		if t.Canonical != nil {
			canonical = *t.Canonical
		}
		name := ""
		if t.Name != nil {
			name = *t.Name
		}
		memberCount := ""
		if t.MemberCount != nil {
			memberCount = strconv.FormatUint(uint64(*t.MemberCount), 10)
		}
		roleNames := make([]string, 0, len(t.Roles))
		for _, r := range t.Roles {
			if r.Name != nil {
				roleNames = append(roleNames, *r.Name)
			}
		}
		return map[string]string{
			"Canonical":   canonical,
			"Name":        name,
			"MemberCount": memberCount,
			"Roles":       strings.Join(roleNames, ", "),
		}
	},
}
