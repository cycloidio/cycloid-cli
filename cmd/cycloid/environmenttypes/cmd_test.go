package environmenttypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnvironmentTypesCommandTree(t *testing.T) {
	t.Parallel()
	root := NewCommands()
	assert.Equal(t, "environment-type", root.Name())

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

func TestEnvironmentTypeCreateFlags(t *testing.T) {
	t.Parallel()
	cmd := NewCreateCommand()

	require.NotNil(t, cmd.Flags().Lookup("color"), "color flag must be registered")
	require.NotNil(t, cmd.Flags().Lookup("environment-type-name"), "environment-type-name flag must be registered")
	require.NotNil(t, cmd.Flags().Lookup("environment-type"), "environment-type (canonical) flag must be registered")
}

func TestEnvironmentTypeCreateHasUpdateFlag(t *testing.T) {
	t.Parallel()
	cmd := NewCreateCommand()
	f := cmd.Flags().Lookup("update")
	require.NotNil(t, f, "create must carry --update for upsert")
	assert.Equal(t, "bool", f.Value.Type())
}

func TestEnvironmentTypeUpdateNoUpdateFlag(t *testing.T) {
	t.Parallel()
	cmd := NewUpdateCommand()
	// update subcommand is always an update; no upsert --update flag
	assert.Nil(t, cmd.Flags().Lookup("update"))
}

func TestAppendUniqueArg(t *testing.T) {
	t.Parallel()

	t.Run("appends new value", func(t *testing.T) {
		result := appendUniqueArg([]string{"a", "b"}, "c")
		assert.Equal(t, []string{"a", "b", "c"}, result)
	})

	t.Run("deduplicates existing value", func(t *testing.T) {
		result := appendUniqueArg([]string{"a", "b"}, "b")
		assert.Equal(t, []string{"a", "b"}, result)
	})

	t.Run("works on empty slice", func(t *testing.T) {
		result := appendUniqueArg(nil, "x")
		assert.Equal(t, []string{"x"}, result)
	})
}
