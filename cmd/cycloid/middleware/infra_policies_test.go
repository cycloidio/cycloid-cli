package middleware_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/internal/testcfg"
)

const testOPAPolicy = `package cycloid

default allow := false

allow if {
	input.resource == "project"
}
`

func TestInfraPoliciesCRUD(t *testing.T) {
	t.Skip()
	m := config.Middleware

	canonical := testcfg.RandomCanonical("test-policy")
	name := "Test Infra Policy"

	policyFile, err := os.CreateTemp("", "test-policy-*.rego")
	require.NoError(t, err)
	defer os.Remove(policyFile.Name())
	_, err = policyFile.WriteString(testOPAPolicy)
	require.NoError(t, err)
	policyFile.Close()

	created, _, err := m.CreateInfraPolicy(config.Org, policyFile.Name(), canonical, "test policy", name, "administrator", "critical", true)
	require.NoError(t, err, "CreateInfraPolicy should succeed")
	require.NotNil(t, created)

	defer func() {
		_, err := m.DeleteInfraPolicy(config.Org, canonical)
		require.NoError(t, err, "DeleteInfraPolicy should succeed")
	}()

	got, _, err := m.GetInfraPolicy(config.Org, canonical)
	require.NoError(t, err, "GetInfraPolicy should succeed")
	assert.Equal(t, canonical, *got.Canonical)

	list, _, err := m.ListInfraPolicies(config.Org)
	require.NoError(t, err, "ListInfraPolicies should succeed")
	assert.NotEmpty(t, list)

	updated, _, err := m.UpdateInfraPolicy(config.Org, canonical, policyFile.Name(), "updated description", name, "administrator", "critical", true)
	require.NoError(t, err, "UpdateInfraPolicy should succeed")
	require.NotNil(t, updated.Description)
	assert.Equal(t, "updated description", *updated.Description)
}
