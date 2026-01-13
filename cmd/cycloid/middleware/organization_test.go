package middleware_test

import (
	"log"
	"testing"
)

// We need to ensure the Child org CRUD works before kicking off tests.
func TestChildOrgCrud(t *testing.T) {
	// config, err := testcfg.NewConfig()
	// defer config.Cleanup()
	// if err != nil {
	// 	t.Errorf("Config setup failed: %v", err)
	// }
	m := config.Middleware

	var childOrg = "test-create-child-org"
	defer func() {
		err := m.DeleteOrganization(childOrg)
		if err != nil {
			log.Fatalf("Failed to delete org '%s': %v", childOrg, err)
			return
		}
	}()

	_, err := m.CreateOrganizationChild(config.Org, childOrg, nil)
	if err != nil {
		t.Errorf("Failed to create org '%s': %v", childOrg, err)
		return
	}

	var newName = "test-update-child-org"
	newOrg, err := m.UpdateOrganization(childOrg, newName)
	if err != nil {
		t.Errorf("Failed to update org '%s': %v", childOrg, err)
	}

	if *newOrg.Name != newName {
		t.Errorf("Update org failed.\nExpected: %s | Got: %s", newName, *newOrg.Name)
	}

	_, err = m.GetOrganization(childOrg)
	if err != nil {
		t.Errorf("Org '%s' is not created: %v", childOrg, err)
		return
	}
}
