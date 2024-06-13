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

func TestUpdateMapField(t *testing.T) {
	inputs := []struct {
		Field    string
		Value    any
		Base     map[string]any
		Expected map[string]any
	}{
		{
			Field: "test",
			Value: "value",
			Base:  make(map[string]any),
			Expected: map[string]any{
				"test": "value",
			},
		},
		{
			Field: "test.nested.value",
			Value: "hello",
			Base:  make(map[string]any),
			Expected: map[string]any{
				"test": map[string]any{
					"nested": map[string]any{
						"value": "hello",
					},
				},
			},
		},
		{
			Field: "changed",
			Value: "yes",
			Base: map[string]any{
				"changed": "notChanged",
			},
			Expected: map[string]any{
				"changed": "yes",
			},
		},
	}

	for _, input := range inputs {
		err := common.UpdateMapField(input.Field, input.Value, input.Base)
		assert.Nil(t, err)
		assert.Equal(t, input.Expected, input.Base)
	}
}
