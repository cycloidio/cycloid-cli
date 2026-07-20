package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Regression test for CLI-144: backend commit 7089ff5d20 (PROD-745) changed
// ServiceCatalog.dependencies from a list of ServiceCatalogDependency objects
// to a list of stack canonicals ([]string). Decoding must not error on the
// new shape.
func TestServiceCatalog_UnmarshalJSON_Dependencies(t *testing.T) {
	payload := []byte(`{
		"author": "cycloid",
		"canonical": "my-stack",
		"directory": "my-stack",
		"form_enabled": true,
		"id": 1,
		"keywords": [],
		"latest": true,
		"name": "My Stack",
		"organization_canonical": "org-root",
		"ref": "main",
		"trusted": true,
		"version": "1.0.0",
		"visibility": "shared",
		"dependencies": ["stack-a", "stack-b"]
	}`)

	var sc ServiceCatalog
	err := json.Unmarshal(payload, &sc)
	assert.NoError(t, err)
	assert.Equal(t, []string{"stack-a", "stack-b"}, sc.Dependencies)
}
