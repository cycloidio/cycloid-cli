package cyout

import (
	"reflect"
	"strings"

	"github.com/spf13/cobra"
)

const annotationFields = "cyout.fields"

// RegisterModel stores the exported field names of model in cmd.Annotations.
// The global --output completion function reads these to offer model-aware suggestions.
// Call this in NewXCommand() for every command that outputs structured data.
func RegisterModel(cmd *cobra.Command, model interface{}) {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	fields := make([]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			fields = append(fields, strings.ToLower(t.Field(i).Name))
		}
	}
	if cmd.Annotations == nil {
		cmd.Annotations = map[string]string{}
	}
	cmd.Annotations[annotationFields] = strings.Join(fields, ",")
}

// GetModelFields reads the field names stored by RegisterModel from a command's annotations.
// Returns nil if RegisterModel was not called for this command.
func GetModelFields(cmd *cobra.Command) []string {
	raw, ok := cmd.Annotations[annotationFields]
	if !ok || raw == "" {
		return nil
	}
	return strings.Split(raw, ",")
}
