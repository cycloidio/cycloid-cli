package roles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCommandsHasRoleAlias(t *testing.T) {
	cmd := NewCommands()
	assert.Contains(t, cmd.Aliases, "role")
}
