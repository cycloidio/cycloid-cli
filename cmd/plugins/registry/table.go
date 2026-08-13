package registry

import (
	"fmt"

	"github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

var registryTableOptions = printer.Options{
	Columns:    []string{"ID", "Name", "URL", "Status", "Access", "Created At", "Updated At"},
	Identifier: "ID",
	Transform: func(obj interface{}) map[string]string {
		r, ok := obj.(*models.PluginRegistry)
		if !ok {
			return map[string]string{}
		}
		return map[string]string{
			"ID":         fmt.Sprintf("%d", ptr.Value(r.ID)),
			"Name":       ptrStr(r.Name),
			"URL":        ptrURI(r.URL),
			"Status":     ptrStr(r.Status),
			"Access":     fmt.Sprintf("%t", ptr.Value(r.Access)),
			"Created At": fmt.Sprintf("%d", ptr.Value(r.CreatedAt)),
			"Updated At": fmt.Sprintf("%d", ptr.Value(r.UpdatedAt)),
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
