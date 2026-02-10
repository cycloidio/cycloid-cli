package teams

import (
	"fmt"
	"strconv"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
)

// anyMemberToID convert map any valid member Username or Email to its ID
func anyMembersToID(org, team string, members []string) ([]*uint32, error) {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	teamMembers, err := m.ListTeamMembers(org, team)
	if err != nil {
		return nil, fmt.Errorf("failed to convert members to ID, cannot list current team members: %w", err)
	}

	result := make([]*uint32, len(members))
	for i, member := range members {
		for _, teamMember := range teamMembers {
			if member == teamMember.FullName || member == teamMember.Email.String() {
				result[i] = teamMember.ID
				break
			}

			id, err := strconv.ParseUint(member, 10, 32)
			if err != nil {
				// not a uint
				continue
			}

			if uint32(id) == ptr.Value(teamMember.ID) {
				result[i] = teamMember.ID
				break
			}
		}

		if result[i] == nil {
			return result, fmt.Errorf("failed to match identifier %q to any member from team %q in org %q", member, team, org)
		}
	}

	return result, nil
}
