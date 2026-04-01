package middleware_test

import (
	"errors"
	"log"
	"strconv"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func TestComponentCRUD(t *testing.T) {
	m := config.Middleware
	for index := range 2 { // Made only two to speed up tests
		var (
			componentName        = "Test Component " + strconv.Itoa(index)
			component            = "test-component-" + strconv.Itoa(index)
			componentDescription = "My cool component"
			stackRef             = config.Org + ":stack-e2e-stackforms"
			useCase              = "default"
			newVar               = models.FormVariables{
				"can two sections have same name with different caps ?": {
					"can two groups have same name with different caps ?": {
						"group1": "EDITED",
					},
				},
			}
			formVars = &models.FormVariables{
				"can two sections have same name with different caps ?": {
					"can two groups have same name with different caps ?": {
						"group1": "osef",
					},
				},
				"CAN TWO SECTIONS HAVE SAME NAME WITH DIFFERENT CAPS ?": {
					"CAN TWO GROUPS HAVE SAME NAME WITH DIFFERENT CAPS ?": {
						"group2": "osef",
					},
				},
				"section spaces AND CAPS": {
					"group spaces AND CAPS": {
						"no_spaces_no_caps": "osef",
					},
				},
				"section with spaces": {
					"group with spaces": {
						"no_spaces": "osef",
					},
				},
				"types": {
					"tests": {
						"array": []any{
							"hello",
							"world",
							false,
							index,
							1.1,
						},
						"bool":    true,
						"float":   0.1,
						"integer": 1,
						"map": map[string]any{
							"array": []any{
								"hello",
								"world",
							},
							"bool":    false,
							"float":   0.1,
							"integer": 1,
							"nested": map[string]string{
								"map": "hello",
							},
							"string": "string",
						},
						"string": "stringValue1",
					},
				},
			}
		)

		var createdComponent *models.Component
		var err, errList error
		for range 3 { // retries due to concurenccy bug in backend
			createdComponent, _, err = m.GetComponent(config.Org, *config.Project.Canonical, *config.Environment.Canonical, component)
			if err == nil {
				errList = nil
				break
			}

			createdComponent, _, err = m.CreateOrUpdateComponent(config.Org, *config.Project.Canonical, *config.Environment.Canonical, component, componentDescription, componentName, stackRef, "", "", "", useCase, "", *formVars)
			if err != nil {
				errList = errors.Join(errList, err)
				continue
			}

			errList = nil
			break
		}

		if errList != nil {
			t.Errorf("Failed to create component '%s':\n%v", component, err)
		}

		defer func() {
			_, err := m.DeleteComponent(config.Org, *config.Project.Canonical, *config.Environment.Canonical, *createdComponent.Canonical)
			if err != nil {
				log.Fatalf("Failed to delete component '%s': %v", *createdComponent.Canonical, err)
				return
			}
		}()

		var (
			newDescription   = "New desc"
			newComponentName = "New name" + strconv.Itoa(index)
		)
		errList, err = nil, nil
		for range 3 {
			_, _, err = m.CreateOrUpdateComponent(config.Org, *config.Project.Canonical, *config.Environment.Canonical, *createdComponent.Canonical, newDescription, newComponentName, stackRef, "", "", *config.CatalogRepoVersionStacks.CommitHash, useCase, "", newVar)
			if err != nil {
				errList = errors.Join(errList, err)
				continue
			}

			errList = nil
			break
		}
		if errList != nil {
			t.Errorf("Failed to update component '%s':\n%v", *createdComponent.Canonical, err)
		}

		// TODO: Fix after issue: https://linear.app/cycloid/issue/BE-801/invalid-response-for-updatecomponent
		// assert.NotNil(t, updatedComponent, "response should not be nil.")
		// assert.NotNil(t, updatedComponent.Name, "Name should not be nil.")
		// assert.Equal(t, newComponentName, *updatedComponent.Name)
		// assert.Equal(t, newDescription, *updatedComponent.Canonical)
	}

	_, _, err := m.ListComponents(config.Org, *config.Project.Canonical, *config.Environment.Canonical)
	if err != nil {
		t.Errorf("Failed to list components in project '%s':\n%v", *config.Project.Canonical, err)
	}
}
