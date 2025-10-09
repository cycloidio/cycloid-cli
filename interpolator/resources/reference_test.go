package resources_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/viper"

	"github.com/cycloidio/cycloid-cli/interpolator/resources"
)

func TestParseURI(t *testing.T) {
	org := "myOrg"
	project := "myProject"
	env := "myEnv"
	component := "myComponent"
	tests := []struct {
		name    string
		uri     string
		want    *resources.Reference
		wantErr bool
	}{
		{
			"simple_credentials",
			"cy://org/some_org/cred/some_cred?format=json",
			&resources.Reference{
				Path: "/organizations/some_org/credentials/some_cred?format=json",
				Params: map[string][]string{
					"format": {"json"},
				},
			},
			false,
		},
		{
			"expand",
			"cy://org/org/project/project/env/env/component/component",
			&resources.Reference{
				Path:   "/organizations/org/projects/project/environments/env/components/component",
				Params: map[string][]string{},
			},
			false,
		},
		{
			"expandAndVars",
			"cy://org/{org}/project/{project}/env/{env}/component/{component}",
			&resources.Reference{
				Path:   fmt.Sprintf("/organizations/%s/projects/%s/environments/%s/components/%s", org, project, env, component),
				Params: map[string][]string{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Set("org", org)
			viper.Set("project", project)
			viper.Set("env", env)
			viper.Set("component", component)
			got, err := resources.NewResourceReference(tt.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("\nhave: %v\nwant: %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nhave: %v\nwant: %v", got, tt.want)
			}
		})
	}
}
