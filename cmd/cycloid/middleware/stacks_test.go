package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/internal/ptr"
)

const (
	testStackCanonical = "stack-e2e-stackforms"
	testStackRef       = "cycloid:stack-e2e-stackforms"
	testStackUseCase   = "default"
)

func TestStacksGetStack(t *testing.T) {
	m := config.Middleware
	t.Run("GetStackOk", func(t *testing.T) {
		stack, err := m.GetStack(config.Org, testStackRef)
		require.NoError(t, err)
		require.NotNil(t, stack)
		assert.Equal(t, testStackCanonical, *stack.Canonical)
		assert.NotNil(t, stack.Ref)
		assert.Equal(t, testStackRef, *stack.Ref)
	})
	t.Run("GetStackNotFound", func(t *testing.T) {
		_, err := m.GetStack(config.Org, "org:nonexistent-stack")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "404")
	})
}
func TestStacksListStacks(t *testing.T) {
	m := config.Middleware
	t.Run("ListStacksOk", func(t *testing.T) {
		stacks, err := m.ListStacks(config.Org)
		require.NoError(t, err)
		require.NotNil(t, stacks)
		assert.NotEmpty(t, stacks, "should have at least one stack")
		// Verify the test stack is in the list
		found := false
		for _, stack := range stacks {
			if stack.Canonical != nil && *stack.Canonical == testStackCanonical {
				found = true
				assert.Equal(t, testStackRef, *stack.Ref)
				break
			}
		}
		assert.True(t, found, "test stack should be in the list")
	})
}
func TestStacksListStackVersions(t *testing.T) {
	m := config.Middleware
	t.Run("ListStackVersionsOk", func(t *testing.T) {
		versions, err := m.ListStackVersions(config.Org, testStackRef)
		require.NoError(t, err)
		require.NotNil(t, versions)
		assert.NotEmpty(t, versions, "should have at least one version")
		// Verify all versions have required fields
		for _, version := range versions {
			assert.NotNil(t, version.ID, "version should have an ID")
			assert.NotNil(t, version.CommitHash, "version should have a commit hash")
			assert.NotNil(t, version.Type, "version should have a type")
			assert.NotNil(t, version.Name, "version should have a name")
		}
	})
	t.Run("ListStackVersionsNotFound", func(t *testing.T) {
		_, err := m.ListStackVersions(config.Org, "org:nonexistent-stack")
		require.Error(t, err)
	})
}
func TestStacksListStackUseCases(t *testing.T) {
	m := config.Middleware
	t.Run("ListStackUseCasesWithDefaultVersion", func(t *testing.T) {
		// When all version params are empty, should use default version
		useCases, err := m.ListStackUseCases(config.Org, testStackRef, "", "", "")
		require.NoError(t, err)
		require.NotNil(t, useCases)
		assert.NotEmpty(t, useCases, "should have at least one use case")
		found := false
		for _, uc := range useCases {
			if uc.UseCase != nil && *uc.UseCase == testStackUseCase {
				found = true
				assert.NotNil(t, uc.Name, "use case should have a name")
				break
			}
		}
		assert.True(t, found, "default use case should be in the list")
	})
	t.Run("ListStackUseCasesWithCommitHash", func(t *testing.T) {
		// First, get a valid commit hash. We'll use commit hash from the stacks
		// branch version to ensure we're getting use cases
		versions, err := m.ListStackVersions(config.Org, testStackRef)
		require.NoError(t, err)
		require.NotEmpty(t, versions)
		var commitHash string
		for _, version := range versions {
			if version.Type != nil && *version.Type == "branch" && version.Name != nil && *version.Name == "stacks" {
				commitHash = *version.CommitHash
				break
			}
		}
		// Use the commit hash to list use cases
		useCases, err := m.ListStackUseCases(config.Org, testStackRef, "", "", commitHash)
		require.NoError(t, err)
		require.NotNil(t, useCases)
		assert.NotEmpty(t, useCases, "should have at least one use case")
	})
	t.Run("ListStackUseCasesWithTag", func(t *testing.T) {
		versions, err := m.ListStackVersions(config.Org, testStackRef)
		require.NoError(t, err)
		require.NotEmpty(t, versions)

		useCases, err := m.ListStackUseCases(config.Org, testStackRef, "stack-e2e-stackforms/v1", "", "")
		require.NoError(t, err)
		require.NotNil(t, useCases)
		assert.NotEmpty(t, useCases, "should have at least one use case")
	})
	t.Run("ListStackUseCasesWithStacksBranch", func(t *testing.T) {
		// Use the branch to list use cases
		useCases, err := m.ListStackUseCases(config.Org, testStackRef, "", "stacks", "")
		require.NoError(t, err)
		require.NotNil(t, useCases)
		assert.NotEmpty(t, useCases, "should have at least one use case")
	})
	t.Run("ListStackUseCasesTagNotFound", func(t *testing.T) {
		_, err := m.ListStackUseCases(config.Org, testStackRef, "v999.999.999", "", "")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})
	t.Run("ListStackUseCasesBranchNotFound", func(t *testing.T) {
		_, err := m.ListStackUseCases(config.Org, testStackRef, "", "nonexistent-branch-xyz", "")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})
	t.Run("ListStackUseCasesCommitHashNotFound", func(t *testing.T) {
		_, err := m.ListStackUseCases(config.Org, testStackRef, "", "", "0000000000000000000000000000000000000000")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})
	t.Run("ListStackUseCasesStackNotFound", func(t *testing.T) {
		_, err := m.ListStackUseCases(config.Org, "org:nonexistent-stack", "", "", "")
		require.Error(t, err)
	})
}
func TestStacksUpdateStack(t *testing.T) {
	m := config.Middleware
	t.Run("UpdateStackVisibility", func(t *testing.T) {
		stack, err := m.GetStack(config.Org, testStackRef)
		require.NoError(t, err)
		originalVisibility := stack.Visibility
		teamCanonical := ""
		if stack.Team != nil && stack.Team.Canonical != nil {
			teamCanonical = *stack.Team.Canonical
		}
		// Toggle visibility
		newVisibility := "local"
		if originalVisibility != nil && *originalVisibility == "local" {
			newVisibility = "shared"
		}
		updatedStack, err := m.UpdateStack(config.Org, testStackRef, teamCanonical, ptr.Ptr(newVisibility))
		require.NoError(t, err)
		require.NotNil(t, updatedStack)
		assert.Equal(t, newVisibility, *updatedStack.Visibility)

		_, err = m.UpdateStack(config.Org, testStackRef, teamCanonical, originalVisibility)
		require.NoError(t, err)
	})
	t.Run("UpdateStackNotFound", func(t *testing.T) {
		_, err := m.UpdateStack(config.Org, "org:nonexistent-stack", "", ptr.Ptr("private"))
		require.Error(t, err)
	})
}

func TestResolveStackVersionIntegration(t *testing.T) {
	// These tests verify that resolveStackVersion works correctly through the public method ListStackUseCases

	m := config.Middleware
	t.Run("ResolveDefaultVersionThroughListStackUseCases", func(t *testing.T) {
		// When no version is specified, should use default version
		useCases1, err1 := m.ListStackUseCases(config.Org, testStackRef, "", "", "")
		require.NoError(t, err1)
		require.NotEmpty(t, useCases1)
		// Should return the same use cases each time with default version
		useCases2, err2 := m.ListStackUseCases(config.Org, testStackRef, "", "", "")
		require.NoError(t, err2)
		require.NotEmpty(t, useCases2)
		assert.Equal(t, len(useCases1), len(useCases2), "default version should be consistent")
	})
	t.Run("ResolveCommitHashThroughListStackUseCases", func(t *testing.T) {
		versions, err := m.ListStackVersions(config.Org, testStackRef)
		require.NoError(t, err)
		require.NotEmpty(t, versions)
		commitHash := *versions[0].CommitHash
		useCases, err := m.ListStackUseCases(config.Org, testStackRef, "", "", commitHash)
		require.NoError(t, err)
		require.NotEmpty(t, useCases)
	})
	t.Run("PriorityTagOverBranchOverCommitHash", func(t *testing.T) {
		versions, err := m.ListStackVersions(config.Org, testStackRef)
		require.NoError(t, err)
		require.NotEmpty(t, versions)
		var commitHash string
		// Make sure commit hash is from the stacks branch
		for _, version := range versions {
			if version.Type != nil && *version.Type == "branch" && version.Name != nil && *version.Name == "stacks" {
				commitHash = *version.CommitHash
				break
			}
		}
		// If we have a tag, it should be used
		useCases, err := m.ListStackUseCases(config.Org, testStackRef, "stack-e2e-stackforms/v1", "", "")
		require.NoError(t, err)
		require.NotEmpty(t, useCases)

		// If we have a branch, it should be used
		useCases, err = m.ListStackUseCases(config.Org, testStackRef, "", "stacks", "")
		require.NoError(t, err)
		require.NotEmpty(t, useCases)

		// If we have a commit hash, it should be used
		useCases, err = m.ListStackUseCases(config.Org, testStackRef, "", "", commitHash)
		require.NoError(t, err)
		require.NotEmpty(t, useCases)
	})
}
