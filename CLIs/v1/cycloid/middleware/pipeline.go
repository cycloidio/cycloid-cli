package middleware

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_pipelines"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
)

func (m *middleware) UnpausePipeline(org string, project string, env string) error {

	pipelineName := fmt.Sprintf("%s-%s", project, env)

	params := organization_pipelines.NewUnpausePipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)

	_, err := m.api.OrganizationPipelines.UnpausePipeline(params, root.ClientCredentials())
	// if err != nil {
	// 	return nil, err
	// }

	return err
}
