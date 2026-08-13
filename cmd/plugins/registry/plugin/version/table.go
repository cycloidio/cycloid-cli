package version

import (
	"fmt"
	"strings"

	"github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

var versionTableOptions = printer.Options{
	Columns:    []string{"ID", "Name", "URL", "Status", "Description", "Scope", "Error"},
	Identifier: "ID",
	Transform: func(obj interface{}) map[string]string {
		v, ok := obj.(*models.PluginVersion)
		if !ok {
			return map[string]string{}
		}
		scope := ""
		if len(v.Scope) > 0 {
			scope = strings.Join(v.Scope, ", ")
		}
		return map[string]string{
			"ID":          fmt.Sprintf("%d", ptr.Value(v.ID)),
			"Name":        ptrStr(v.Name),
			"URL":         ptrURI(v.URL),
			"Status":      ptrStr(v.Status),
			"Description": v.Description,
			"Scope":       scope,
			"Error":       v.Error,
		}
	},
}

func ptrStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func ptrURI(v *strfmt.URI) string {
	if v == nil {
		return ""
	}
	return string(*v)
}
