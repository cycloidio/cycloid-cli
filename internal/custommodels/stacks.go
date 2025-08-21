package custommodels

import "github.com/cycloidio/cycloid-cli/client/models"

type Blueprint struct {
	// Added use cases for blueprint make it more usable output
	// for MCP server
	UseCases *[]string `json:"use_cases"`
	models.ServiceCatalog
}
