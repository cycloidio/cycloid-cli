package middleware_test

import (
	"log"
	"strconv"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func TestComponentCRUD(t *testing.T) {
	// setup
	t.Parallel()
	config, err := getTestConfig()
	if err != nil {
		t.Fatalf("Config setup failed: %v", err)
	}
	m := config.Middleware

	var (
		projectName      = "Test CRUD component"
		project          = "test-crud-component"
		description      = "Testing components"
		configRepository = configRepository
		owner            = ""
		team             = ""
		color            = "blue"
		icon             = "planet"
	)

	defer func() {
		err := m.DeleteProject(config.Org, project)
		if err != nil {
			log.Fatalf("Failed to decomission project '%s' from CRUD tests: %v", project, err)
			return
		}
	}()

	createdProject, err := m.CreateProject(config.Org, projectName, project, description, configRepository, owner, team, color, icon)
	if err != nil {
		t.Fatalf("Failed to create pre-requisite project, create project CRUD tests: %v", err)
	}

	var (
		env      = "test"
		envName  = "Test"
		envColor = "red"
	)

	defer func() {
		err := m.DeleteEnv(config.Org, project, env)
		if err != nil {
			log.Fatalf("Failed to delete env '%s': %v", env, err)
			return
		}
	}()

	_, err = m.CreateEnv(config.Org, *createdProject.Canonical, env, envName, envColor)
	if err != nil {
		t.Fatalf("Failed to create env '%s': %v", env, err)
	}
	// end setup

	// update
	for index := range 2 { // Made only two to speed up tests
		var (
			componentName        = "Test Component " + strconv.Itoa(index)
			component            = "test-component-" + strconv.Itoa(index)
			componentDescription = "My cool component"
			stackRef             = "cycloid:stack-e2e-stackforms"
			useCase              = "default"
			newVar               = models.FormVariables{
				"can two sections have same name with different caps ?": {
					"can two groups have same name with different caps ?": {
						"group1": "EDITED",
					},
				},
			}
			formVars = models.FormVariables{
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

		createdComponent, err := m.CreateComponent(
			config.Org, project, env, component, componentDescription, &componentName, &stackRef, &useCase, nil, &formVars,
		)
		if err != nil {
			t.Fatalf("Failed to create component '%s':\n%v", component, err)
		}
		defer func() {
			err := m.DeleteComponent(config.Org, project, env, *createdComponent.Canonical)
			if err != nil {
				log.Fatalf("Failed to delete component '%s': %v", *createdComponent.Canonical, err)
				return
			}
		}()

		var (
			newDescription   = "New desc"
			newComponentName = "New name" + strconv.Itoa(index)
		)
		_, err = m.UpdateComponent(config.Org, project, env, *createdComponent.Canonical, newDescription, &newComponentName, &useCase, &newVar)
		if err != nil {
			t.Fatalf("Failed to update component '%s':\n%v", *createdComponent.Canonical, err)
		}

		// TODO: Fix after issue: https://linear.app/cycloid/issue/BE-801/invalid-response-for-updatecomponent
		// assert.NotNil(t, updatedComponent, "response should not be nil.")
		// assert.NotNil(t, updatedComponent.Name, "Name should not be nil.")
		// assert.Equal(t, newComponentName, *updatedComponent.Name)
		// assert.Equal(t, newDescription, *updatedComponent.Canonical)
	}

	_, err = m.GetComponents(config.Org, project, env)
	if err != nil {
		t.Fatalf("Failed to list components in project '%s':\n%v", project, err)
	}
}
