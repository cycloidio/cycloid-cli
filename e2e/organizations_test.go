package e2e_test

import (
	"encoding/json"
	"slices"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/testcfg"
	"github.com/matryer/is"
)

func TestOrganizations(t *testing.T) {
	t.Run("SuccessOrganizationsGet", func(t *testing.T) {
		is := is.New(t)
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"organization",
			"get",
		})
		is.NoErr(cmdErr) // cmd should not fail

		var outOrg *models.Organization
		err := json.Unmarshal([]byte(cmdOut), &outOrg)
		is.NoErr(err)                           // JSON output should be a valid model
		is.Equal(config.Org, *outOrg.Canonical) // org canonicals should match
	})

	childOrg := testcfg.RandomCanonical("e2e-child")
	t.Run("SuccessOrganizationsCreateChild", func(t *testing.T) {
		is := is.New(t)
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"organization",
			"create",
			"--name", childOrg,
			"--child-of", config.Org,
		})
		is.NoErr(cmdErr) // command should not fail
		defer t.Run("SuccessDeleteOrg", func(t *testing.T) {
			is := is.New(t)
			_, deleteErr := executeCommand([]string{
				"--output", "json",
				"organization",
				"delete",
				"--org", childOrg,
			})
			is.NoErr(deleteErr) // command should not fail
		})

		var outOrg *models.Organization
		err := json.Unmarshal([]byte(cmdOut), &outOrg)
		is.NoErr(err)                         // JSON output should be a valid model
		is.Equal(childOrg, *outOrg.Canonical) // org canonicals should match

		t.Run("SuccessOrganizationsListChildrens", func(t *testing.T) {
			is := is.New(t)
			cmdOut, cmdErr := executeCommand([]string{
				"--output", "json",
				"organization",
				"list-childrens",
				"--org", config.Org,
			})
			is.NoErr(cmdErr) // command should not fail

			var outOrgs []*models.Organization
			err := json.Unmarshal([]byte(cmdOut), &outOrgs)
			is.NoErr(err) // output should be deserializable orgs

			// We try to find the new child org
			index := slices.IndexFunc(outOrgs, func(o *models.Organization) bool {
				return *o.Canonical == childOrg
			})
			is.True(index != -1) // if -1, means that we can't find our new child org in list. Should be there.
		})

		t.Run("SuccessOrganizationsListWorkers", func(t *testing.T) {
			is := is.New(t)
			cmdOut, cmdErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"organization",
				"list-workers",
			})
			is.NoErr(cmdErr) // cmd should not fail

			var workers []*models.Worker
			err := json.Unmarshal([]byte(cmdOut), &workers)
			is.NoErr(err) // output should be deserializable
		})
	})
}
