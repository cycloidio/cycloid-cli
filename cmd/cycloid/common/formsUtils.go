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
		useCaseName := useCase.Name
		if result[useCaseName] == nil {
			result[useCaseName] = make(map[string]map[string]interface{})
		}

		for _, group := range useCase.Groups {
			groupName := group.Name
			if result[useCaseName][groupName] == nil {
				result[useCaseName][groupName] = make(map[string]interface{})
			}

			for _, widget := range group.Vars {
				widgetName := widget.Key
				value := widget.GetValue(useDefaults)
				if value != nil {
					result[useCaseName][groupName][widgetName] = value
				}
			}
		}
	}

	return result
}
