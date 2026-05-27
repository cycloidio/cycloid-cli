package cloudaccounts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCloudAccountsCommandTree(t *testing.T) {
	t.Parallel()
	root := NewCommands()
	assert.Equal(t, "cloud-account", root.Name())

	names := make(map[string]int)
	for _, c := range root.Commands() {
		names[c.Name()]++
	}

	require.Equal(t, 1, names["list"])
	require.Equal(t, 1, names["get"])
	require.Equal(t, 1, names["create"])
	require.Equal(t, 1, names["update"])
	require.Equal(t, 1, names["delete"])
}

func TestCloudAccountCreateFlags(t *testing.T) {
	t.Parallel()
	cmd := NewCreateCommand()

	require.NotNil(t, cmd.Flags().Lookup("cloud-provider"), "cloud-provider flag must be registered")
	require.NotNil(t, cmd.Flags().Lookup("credential"), "credential flag must be registered")
	require.NotNil(t, cmd.Flags().Lookup("new-credential"), "new-credential flag must be registered")
}

func TestCloudAccountCreateHasUpdateFlag(t *testing.T) {
	t.Parallel()
	cmd := NewCreateCommand()
	f := cmd.Flags().Lookup("update")
	require.NotNil(t, f, "create must carry --update for upsert")
	assert.Equal(t, "bool", f.Value.Type())
}

func TestCloudAccountUpdateNoUpdateFlag(t *testing.T) {
	t.Parallel()
	cmd := NewUpdateCommand()
	// update subcommand is always an update; no upsert --update flag
	assert.Nil(t, cmd.Flags().Lookup("update"))
}

func TestCloudAccountUpdateRequiresCanonical(t *testing.T) {
	t.Parallel()
	cmd := NewUpdateCommand()
	f := cmd.Flags().Lookup("cloud-account")
	require.NotNil(t, f, "update must require --cloud-account (canonical)")
}

func TestCloudAccountNameFallsBackToCanonical(t *testing.T) {
	t.Parallel()
	// cloudAccountName() is exercised via the cobra command flags; verify
	// the function directly: when --name is empty it should use --cloud-account.
	cmd := NewCreateCommand()
	require.NoError(t, cmd.Flags().Set("cloud-account", "my-canonical"))
	require.NoError(t, cmd.Flags().Set("cloud-provider", "aws"))

	name, err := cloudAccountName(cmd)
	require.NoError(t, err)
	assert.Equal(t, "my-canonical", name, "name must fall back to canonical when --name is absent")
}
