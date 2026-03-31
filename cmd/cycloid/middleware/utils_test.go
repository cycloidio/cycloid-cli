package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func TestToCanonical(t *testing.T) {
	testCanonicals := []struct {
		testCase string
		name     string
		expect   string
	}{
		{"AllowedChar", "azAZ02-azAZ02", "azaz02-azaz02"},
		{"SpaceConversionChar", "azAZ02 azAZ02", "azaz02_azaz02"},
		{"RemoveUnallowedChar", "az@@$$02 azŷùù&éZ02", "az02_azz02"},
		{"EndWith-", "canonical-", "canonical"},
		{"StartWith-", "-canonical", "canonical"},
		{"EndWith_", "canonical_", "canonical"},
		{"StartWith_", "_canonical", "canonical"},
	}

	for _, tc := range testCanonicals {
		t.Run(tc.testCase, func(t *testing.T) {
			assert.Equal(t, tc.expect, middleware.ToCanonical(tc.name))
		})
	}
}

func TestNameOrCanonical(t *testing.T) {
	testCases := []struct {
		name           string
		inputName      string
		inputCanonical string
		expectedName   string
		expectedCan    string
		expectErr      bool
	}{
		{
			name:         "name-only infers canonical",
			inputName:    "My Cool Project",
			expectedName: "My Cool Project",
			expectedCan:  "my_cool_project",
			expectErr:    false,
		},
		{
			name:           "canonical-only infers name",
			inputCanonical: "my-project",
			expectedName:   "My-project",
			expectedCan:    "my-project",
			expectErr:      false,
		},
		{
			name:           "both provided keep values",
			inputName:      "Existing Name",
			inputCanonical: "existing-canonical",
			expectedName:   "Existing Name",
			expectedCan:    "existing-canonical",
			expectErr:      false,
		},
		{
			name:      "both empty return error",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotName, gotCanonical, err := middleware.NameOrCanonical(&tc.inputName, &tc.inputCanonical)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedName, gotName)
			assert.Equal(t, tc.expectedCan, gotCanonical)
		})
	}
}
