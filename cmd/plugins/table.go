package plugins

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

var pluginTableOptions = printer.Options{
	Columns:    []string{"ID", "Name", "Owned", "Orphaned", "Status", "UUID", "Registry"},
	Identifier: "ID",
	Transform: func(obj interface{}) map[string]string {
		p, ok := obj.(*models.Plugin)
		if !ok {
			return map[string]string{}
		}
		status := ""
		uuid := ""
		if p.Install != nil {
			status = ptrStr(p.Install.Status)
			if p.Install.UUID != nil {
				uuid = p.Install.UUID.String()
			}
		}
		regName := ""
		if p.Registry != nil {
			regName = ptrStr(p.Registry.Name)
		}
		return map[string]string{
			"ID":       fmt.Sprintf("%d", ptr.Value(p.ID)),
			"Name":     ptrStr(p.Name),
			"Owned":    fmt.Sprintf("%t", ptr.Value(p.Owned)),
			"Orphaned": fmt.Sprintf("%t", p.Orphaned),
			"Status":   status,
			"UUID":     uuid,
			"Registry": regName,
		}
	},
}

func ptrStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
