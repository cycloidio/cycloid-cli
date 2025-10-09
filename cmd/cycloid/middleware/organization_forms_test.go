package middleware_test

import (
	"testing"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func TestValidateFormYamlAnchors(t *testing.T) {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var org = "cycloid-sandbox"
	var rawForms = []byte(
		`---
		shared:
		- &some_anchor
			key: hello
			name: hello
			type: string
			widget: simple_text

		sections: 
		- name: "section"
		  groups:
			- name: "group"
			  technologies: [pipeline]
				vars:
				- *some_anchor
		`,
	)

	_, err := m.ValidateForm(org, rawForms)
	if err != nil {
		t.Fatalf("Using yaml anchor should not error:\nerr: '%s'\nyaml files: '%v'", err, rawForms)
	}
}
