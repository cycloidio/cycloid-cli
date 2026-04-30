package members

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "delete [id...]",
		Aliases: []string{"rm"},
		Args:    cyargs.RequireArgsOrFlag("id"),
		Short:   "Remove a user from the organization",
		Example: `
	# Remove a member from my-org organization using the --id flag
	cy --org my-org members delete --id 50

	# Remove multiple members using positional args
	cy --org my-org members delete 50 51 52
`,
		RunE: deleteMember,
	}

	cyargs.AddMemberIDFlag(cmd)

	return cmd
}

func deleteMember(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	if flagID, _ := cyargs.GetMemberID(cmd); flagID != 0 {
		idStr := strconv.FormatUint(uint64(flagID), 10)
		found := false
		for _, a := range args {
			if a == idStr {
				found = true
				break
			}
		}
		if !found {
			args = append(args, idStr)
		}
	}

	for _, arg := range args {
		id64, parseErr := strconv.ParseUint(arg, 10, 32)
		if parseErr != nil {
			return fmt.Errorf("invalid member ID %q: %w", arg, parseErr)
		}
		id := uint32(id64)
		_, err = m.DeleteMember(org, id)
		if err != nil {
			return cyout.Print(cmd, nil, err, fmt.Sprintf("unable to remove member %d", id))
		}
	}
	return nil
}
