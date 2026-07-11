package members

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [id...]",
		Args:  cyargs.RequireArgsOrFlag("id"),
		Short: "Get the organization member",
		Example: `
	# Get a member within my-org organization using the --id flag
	cy --org my-org members get --id 50

	# Get multiple members using positional args
	cy --org my-org members get 50 51 52
`,
		RunE: getMember,
	}

	cyargs.AddMemberIDFlag(cmd)
	cyout.RegisterModel(cmd, models.MemberOrg{})

	return cmd
}

func getMember(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

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

	if len(args) == 1 {
		id64, parseErr := strconv.ParseUint(args[0], 10, 32)
		if parseErr != nil {
			return fmt.Errorf("invalid member ID %q: %w", args[0], parseErr)
		}
		mb, _, err := m.GetMember(org, uint32(id64))
		return cyout.PrintWithOptions(cmd, mb, err, "unable to get member", memberTableOptions)
	}

	results := make([]*models.MemberOrg, 0, len(args))
	for _, arg := range args {
		id64, parseErr := strconv.ParseUint(arg, 10, 32)
		if parseErr != nil {
			return fmt.Errorf("invalid member ID %q: %w", arg, parseErr)
		}
		mb, _, err := m.GetMember(org, uint32(id64))
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, fmt.Sprintf("unable to get member %s", arg), memberTableOptions)
		}
		results = append(results, mb)
	}
	return cyout.PrintWithOptions(cmd, results, nil, "", memberTableOptions)
}
