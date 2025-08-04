package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestBlueprintRefFormatValidation tests blueprint reference format validation
func TestBlueprintRefFormatValidation(t *testing.T) {
	tests := []struct {
		name         string
		blueprintRef string
		isValid      bool
	}{
		{
			name:         "ValidBlueprintRef",
			blueprintRef: "repo:blueprint-canonical",
			isValid:      true,
		},
		{
			name:         "ValidBlueprintRefWithHyphens",
			blueprintRef: "my-repo:my-blueprint-canonical",
			isValid:      true,
		},
		{
			name:         "ValidBlueprintRefWithUnderscores",
			blueprintRef: "my_repo:my_blueprint_canonical",
			isValid:      true,
		},
		{
			name:         "EmptyBlueprintRef",
			blueprintRef: "",
			isValid:      false,
		},
		{
			name:         "BlueprintRefWithoutColon",
			blueprintRef: "invalid-ref",
			isValid:      false,
		},
		{
			name:         "BlueprintRefWithMultipleColons",
			blueprintRef: "repo:blueprint:extra",
			isValid:      false,
		},
		{
			name:         "BlueprintRefStartingWithColon",
			blueprintRef: ":blueprint",
			isValid:      false,
		},
		{
			name:         "BlueprintRefEndingWithColon",
			blueprintRef: "repo:",
			isValid:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := validateBlueprintRefFormat(tt.blueprintRef)
			assert.Equal(t, tt.isValid, isValid)
		})
	}
}

// validateBlueprintRefFormat validates the format of a blueprint reference
func validateBlueprintRefFormat(blueprintRef string) bool {
	if blueprintRef == "" {
		return false
	}

	// Check if it contains exactly one colon
	colonCount := 0
	for _, char := range blueprintRef {
		if char == ':' {
			colonCount++
		}
	}

	if colonCount != 1 {
		return false
	}

	// Split by colon and check both parts are non-empty
	parts := []string{}
	for i, char := range blueprintRef {
		if char == ':' {
			parts = append(parts, blueprintRef[:i], blueprintRef[i+1:])
			break
		}
	}

	if len(parts) != 2 {
		return false
	}

	return parts[0] != "" && parts[1] != ""
}

// TestBlueprintUseCasesExtraction tests extraction of use cases from blueprint config
func TestBlueprintUseCasesExtraction(t *testing.T) {
	tests := []struct {
		name             string
		config           map[string]interface{}
		expectedUseCases []string
	}{
		{
			name: "SingleUseCase",
			config: map[string]interface{}{
				"production": map[string]interface{}{
					"sections": []interface{}{},
				},
			},
			expectedUseCases: []string{"production"},
		},
		{
			name: "MultipleUseCases",
			config: map[string]interface{}{
				"production": map[string]interface{}{
					"sections": []interface{}{},
				},
				"development": map[string]interface{}{
					"sections": []interface{}{},
				},
				"staging": map[string]interface{}{
					"sections": []interface{}{},
				},
			},
			expectedUseCases: []string{"production", "development", "staging"},
		},
		{
			name:             "EmptyConfig",
			config:           map[string]interface{}{},
			expectedUseCases: []string{},
		},
		{
			name: "ConfigWithNonUseCaseKeys",
			config: map[string]interface{}{
				"shared": map[string]interface{}{
					"vars": []interface{}{},
				},
				"production": map[string]interface{}{
					"sections": []interface{}{},
				},
			},
			expectedUseCases: []string{"shared", "production"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCases := extractUseCasesFromConfig(tt.config)
			require.ElementsMatch(t, tt.expectedUseCases, useCases,
				"Extracted use cases should match expected use cases")
		})
	}
}

// extractUseCasesFromConfig extracts use case names from blueprint config
func extractUseCasesFromConfig(config map[string]interface{}) []string {
	var useCases []string
	for useCase := range config {
		useCases = append(useCases, useCase)
	}
	return useCases
}

// TestCreateStackFromBlueprintInputValidation tests input validation for blueprint creation
func TestCreateStackFromBlueprintInputValidation(t *testing.T) {
	tests := []struct {
		name                          string
		org                           string
		blueprintRef                  string
		stackName                     string
		stackCanonical                string
		serviceCatalogSourceCanonical string
		useCase                       string
		shouldBeValid                 bool
	}{
		{
			name:                          "ValidInputs",
			org:                           "test-org",
			blueprintRef:                  "repo:blueprint-canonical",
			stackName:                     "Test Stack",
			stackCanonical:                "test-stack",
			serviceCatalogSourceCanonical: "test-catalog",
			useCase:                       "production",
			shouldBeValid:                 true,
		},
		{
			name:                          "EmptyBlueprintRef",
			org:                           "test-org",
			blueprintRef:                  "",
			stackName:                     "Test Stack",
			stackCanonical:                "test-stack",
			serviceCatalogSourceCanonical: "test-catalog",
			useCase:                       "production",
			shouldBeValid:                 false,
		},
		{
			name:                          "EmptyStackName",
			org:                           "test-org",
			blueprintRef:                  "repo:blueprint-canonical",
			stackName:                     "",
			stackCanonical:                "test-stack",
			serviceCatalogSourceCanonical: "test-catalog",
			useCase:                       "production",
			shouldBeValid:                 false,
		},
		{
			name:                          "EmptyStackCanonical",
			org:                           "test-org",
			blueprintRef:                  "repo:blueprint-canonical",
			stackName:                     "Test Stack",
			stackCanonical:                "",
			serviceCatalogSourceCanonical: "test-catalog",
			useCase:                       "production",
			shouldBeValid:                 false,
		},
		{
			name:                          "EmptyServiceCatalogSourceCanonical",
			org:                           "test-org",
			blueprintRef:                  "repo:blueprint-canonical",
			stackName:                     "Test Stack",
			stackCanonical:                "test-stack",
			serviceCatalogSourceCanonical: "",
			useCase:                       "production",
			shouldBeValid:                 false,
		},
		{
			name:                          "EmptyUseCase",
			org:                           "test-org",
			blueprintRef:                  "repo:blueprint-canonical",
			stackName:                     "Test Stack",
			stackCanonical:                "test-stack",
			serviceCatalogSourceCanonical: "test-catalog",
			useCase:                       "",
			shouldBeValid:                 false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := validateCreateStackFromBlueprintInputs(
				tt.org, tt.blueprintRef, tt.stackName, tt.stackCanonical,
				tt.serviceCatalogSourceCanonical, tt.useCase)
			assert.Equal(t, tt.shouldBeValid, isValid)
		})
	}
}

// validateCreateStackFromBlueprintInputs validates all required inputs for creating a stack from blueprint
func validateCreateStackFromBlueprintInputs(org, blueprintRef, stackName, stackCanonical, serviceCatalogSourceCanonical, useCase string) bool {
	return org != "" && blueprintRef != "" && stackName != "" &&
		stackCanonical != "" && serviceCatalogSourceCanonical != "" && useCase != ""
}

// TestBlueprintOutputFormatting tests the formatting of blueprint output data
func TestBlueprintOutputFormatting(t *testing.T) {
	tests := []struct {
		name             string
		useCases         []string
		expectedUseCases string
	}{
		{
			name:             "BlueprintWithUseCases",
			useCases:         []string{"production", "development"},
			expectedUseCases: "production, development",
		},
		{
			name:             "BlueprintWithoutUseCases",
			useCases:         []string{},
			expectedUseCases: "",
		},
		{
			name:             "BlueprintWithSingleUseCase",
			useCases:         []string{"production"},
			expectedUseCases: "production",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test use cases formatting
			useCasesStr := formatUseCases(tt.useCases)
			assert.Equal(t, tt.expectedUseCases, useCasesStr)
		})
	}
}

// formatUseCases formats a slice of use cases into a comma-separated string
func formatUseCases(useCases []string) string {
	if len(useCases) == 0 {
		return ""
	}

	result := ""
	for i, useCase := range useCases {
		if i > 0 {
			result += ", "
		}
		result += useCase
	}
	return result
}
