package common

// Define what we need to extract only
type Widget struct {
	Widget      string      `json:"widget,omitempty"`
	Default     interface{} `json:"default"`
	Type        string      `json:"type,omitempty"`
	Description string      `json:"description,omitempty"`
	Key         string      `json:"key,omitempty"`
	Name        string      `json:"name,omitempty"`
	Current     interface{} `json:"current,omitempty"`
}

// Retrieve the value from a Stackforms widget definition
func (w *Widget) GetValue(useDefaults bool) interface{} {
	if w.Current == nil && useDefaults {
		return w.Default
	} else {
		return w.Current
	}
}

type UseCase struct {
	Name     string    `json:"name"`
	Sections []Section `json:"sections"`
}

type Section struct {
	Name   string  `json:"name"`
	Groups []Group `json:"groups"`
}

type Group struct {
	Name         string   `json:"name"`
	Technologies []string `json:"technologies,omitempty"`
	Vars         []Widget `json:"vars"`
}

// Convert a UseCase to a suitable Stackforms Input for create/update env API call
func UseCaseToFormInput(useCase UseCase, useDefaults bool) map[string]map[string]map[string]interface{} {
	result := make(map[string]map[string]map[string]interface{})

	for _, useCase := range useCase.Sections {
		if result[useCase.Name] == nil {
			result[useCase.Name] = make(map[string]map[string]interface{})
		}

		for _, group := range useCase.Groups {
			if result[useCase.Name][group.Name] == nil {
				result[useCase.Name][group.Name] = make(map[string]interface{})
			}

			for _, widget := range group.Vars {
				result[useCase.Name][group.Name][widget.Name] = widget.GetValue(useDefaults)
			}
		}
	}

	return result
}
