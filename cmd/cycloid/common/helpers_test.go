package common_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

func TestGenerateCanonical(t *testing.T) {
	inputs := []struct {
		Name      string
		Canonical string
	}{
		{
			Name:      "Test Project",
			Canonical: "test-project",
		},
		{
			Name:      "test project",
			Canonical: "test-project",
		},
		{
			Name:      "Test-Project",
			Canonical: "test-project",
		},
		{
			Name:      "test-project",
			Canonical: "test-project",
		},
		{
			Name:      "test          project",
			Canonical: "test-project",
		},
	}

	for _, input := range inputs {
		assert.Equal(t, input.Canonical, common.GenerateCanonical(input.Name))
	}
}
