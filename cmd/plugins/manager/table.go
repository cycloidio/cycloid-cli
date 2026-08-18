package manager

import (
	"fmt"

	"github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

var managerTableOptions = printer.Options{
	Columns:    []string{"ID", "Name", "URL", "Status", "Invite Status", "Created At", "Updated At"},
	Identifier: "ID",
	Transform: func(obj interface{}) map[string]string {
		m, ok := obj.(*models.PluginManager)
		if !ok {
			return map[string]string{}
		}
		return map[string]string{
			"ID":            fmt.Sprintf("%d", ptr.Value(m.ID)),
			"Name":          ptrStr(m.Name),
			"URL":           ptrURI(m.URL),
			"Status":        ptrStr(m.Status),
			"Invite Status": ptrStr(m.InviteStatus),
			"Created At":    fmt.Sprintf("%d", ptr.Value(m.CreatedAt)),
			"Updated At":    fmt.Sprintf("%d", ptr.Value(m.UpdatedAt)),
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
