package middleware_test

import (
	"testing"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/stretchr/testify/assert"
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
