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

// func TestCorrectUpdateMapField(t *testing.T) {
// 	inputs := []struct {
// 		Field    string
// 		Value    any
// 		Base     map[string]map[string]map[string]any
// 		Expected map[string]map[string]map[string]any
// 	}{
// 		{
// 			Field: "config.test.toto",
// 			Value: "tutu",
// 			Base:  make(map[string]map[string]map[string]any),
// 			Expected: map[string]map[string]map[string]any{
// 				"config": {
// 					"test": {"toto": "tutu"},
// 				},
// 			},
// 		},
// 		{
// 			Field: "config.test.toto",
// 			Value: true,
// 			Base:  make(map[string]map[string]map[string]any),
// 			Expected: map[string]map[string]map[string]any{
// 				"config": {
// 					"test": {"toto": true},
// 				},
// 			},
// 		},
// 		{
// 			Field: "config.test.toto",
// 			Value: 3,
// 			Base:  make(map[string]map[string]map[string]any),
// 			Expected: map[string]map[string]map[string]any{
// 				"config": {
// 					"test": {"toto": 3},
// 				},
// 			},
// 		},
// 	}
//
// 	for _, input := range inputs {
// 		err := common.UpdateMapField(input.Field, input.Value, input.Base)
// 		assert.Nil(t, err)
// 		assert.Equal(t, input.Expected, input.Base)
// 	}
// }
