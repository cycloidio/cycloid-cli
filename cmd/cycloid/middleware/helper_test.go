package middleware_test

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func RandomCanonical(baseName string) string {
	var size = 4
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

	b := make([]rune, size)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return strings.ToLower(baseName) + "-" + string(b)
}

// SetupTestProject will create a project with a random canonical derived from identifier
// and return the project, the function to defer for its deletion and error.
// The func will always be returned so even if err != nil, defer the func.
func (config *TestConfig) SetupTestProject(identifier string) (*models.Project, func(), error) {
	var (
		project          = RandomCanonical(identifier)
		description      = "Testing project " + identifier
		configRepository = configRepository
		owner            = ""
		team             = ""
		color            = "default"
		icon             = "inventory"
	)

	m := config.Middleware

	deleteFunc := func() {
		err := m.DeleteProject(config.Org, project)
		if err != nil {
			log.Fatalf("cannot cleanup projet '%s' for test '%s': %s", project, identifier, err)
			return
		}
	}

	out, err := m.CreateProject(config.Org, project, project, description, configRepository, owner, team, color, icon)
	if err != nil {
		return nil, deleteFunc, fmt.Errorf("failed to setup test project: %s", err)
	}

	return out, deleteFunc, nil
}

// setupTestProject will create an env with a random canonical derived from identifier
// and return the env, the function to defer for its deletion and error.
// The func will always be returned so even if err != nil, defer the func.
func (config *TestConfig) SetupTestEnv(identifier, project string) (*models.Environment, func(), error) {
	var (
		env   = RandomCanonical(identifier)
		color = "default"
	)

	m := config.Middleware

	deleteFunc := func() {
		err := m.DeleteEnv(config.Org, project, env)
		if err != nil {
			log.Fatalf("cannot cleanup env '%s' for test '%s': %s", env, identifier, err)
			return
		}
	}

	out, err := m.CreateEnv(config.Org, project, env, env, color)
	if err != nil {
		return nil, deleteFunc, fmt.Errorf("failed to setup test environment: %s", err)
	}

	return out, deleteFunc, nil
}

// setupTestProject will create an component with a random canonical derived from identifier
// and return the component, the function to defer for its deletion and error.
// The func will always be returned so even if err != nil, defer the func.
func (config *TestConfig) SetupTestComponent(project, env, identifier, stackRef, useCase string, inputs *models.FormVariables) (*models.Component, func(), error) {
	component := RandomCanonical(identifier)
	deleteFunc := func() {
		if err := m.DeleteComponent(config.Org, project, env, component); err != nil {
			log.Fatalf("failed to cleanup component for test '%s': %s", identifier, err)
		}
	}

	createdComponent, err := m.CreateComponent(
		config.Org, project, env, component, "", &component, &stackRef, &useCase, nil, inputs,
	)
	if err != nil {
		return nil, deleteFunc, fmt.Errorf("failed to setup component '%s' for test '%s':\n%v", component, identifier, err)
	}

	return createdComponent, deleteFunc, nil
}
