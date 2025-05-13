package common

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func FormUseCaseToFormVars(stackConfig models.ServiceCatalogConfigs, useCaseName string) (*models.FormVariables, error) {
	useCaseData, ok := stackConfig[useCaseName]
	if !ok {
		return nil, fmt.Errorf("cannot find usecase named '%s' in stack config", useCaseName)
	}

	output := make(models.FormVariables)

	for _, section := range useCaseData.Forms.Sections {
		for _, group := range section.Groups {
			for _, widget := range group.Vars {
				if widget.Default != nil {
					output[*section.Name][*group.Name][*widget.Key] = widget.Default
				}
			}
		}
	}

	return nil, nil
}
