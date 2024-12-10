package pipelines

import (
	"encoding/json"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

// Swagger definition of GetStackConfig does not return a model https://github.com/cycloidio/youdeploy-http-api/issue/1676.
// only an interface{}. As we are looking for a specific field only, we use this hack to go through different map layers
// to get the pipeline variables path
type ppVarPath struct {
	Destination string `json:"destination"`
}
type ppVar struct {
	Variables ppVarPath `json:"variables"`
}
type pp struct {
	Pipeline ppVar `json:"pipeline"`
}

// GetPipelineVarsPath get pipeline variables path to use in a config repository
// from a specified project usecase
func GetPipelineVarsPath(m middleware.Middleware, org, project, usecase string) (string, error) {

	// Get stack ref
	proj, err := m.GetProject(org, project)
	if err != nil {
		return "", err
	}

	// Get stack config
	sc, err := m.GetStackConfig(org, *proj.ServiceCatalog.Ref)
	if err != nil {
		return "", err
	}

	scJson, err := json.Marshal(sc)
	if err != nil {
		return "", err
	}

	var scMap map[string]pp
	err = json.Unmarshal(scJson, &scMap)
	if err != nil {
		return "", err
	}

	return scMap[usecase].Pipeline.Variables.Destination, nil
}
