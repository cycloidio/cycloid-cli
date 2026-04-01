package roles

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRolesCommandTree(t *testing.T) {
	t.Parallel()
	root := NewCommands()
	assert.Equal(t, "roles", root.Name())

	names := make(map[string]int)
	for _, c := range root.Commands() {
		names[c.Name()]++
	}

	require.Equal(t, 1, names["list"])
	require.Equal(t, 1, names["get"])
	require.Equal(t, 1, names["create"])
	require.Equal(t, 1, names["update"])
	require.Equal(t, 1, names["delete"], "delete must be registered exactly once")
}

func TestRolesCreateHasUpdateFlag(t *testing.T) {
	t.Parallel()
	cmd := NewCreateCommand()
	f := cmd.Flags().Lookup("update")
	require.NotNil(t, f)
	assert.Equal(t, "bool", f.Value.Type())
}

func TestRolesUpdateCommandNoUpdateFlag(t *testing.T) {
	t.Parallel()
	cmd := NewUpdateCommand()
	// update subcommand is always upsert; no --update flag
	assert.Nil(t, cmd.Flags().Lookup("update"))
}
