package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/json"
	"github.com/cycloidio/cycloid-cli/printer/yaml"
)

func TestGetPrinter(t *testing.T) {
	t.Run("SuccessYAML", func(t *testing.T) {
		p, err := GetPrinter("yaml")
		require.NoError(t, err)
		assert.Implements(t, (*printer.Printer)(nil), p)
		assert.IsType(t, yaml.YAML{}, p)
	})
	t.Run("SuccessJSON", func(t *testing.T) {
		p, err := GetPrinter("json")
		require.NoError(t, err)
		assert.Implements(t, (*printer.Printer)(nil), p)
		assert.IsType(t, json.JSON{}, p)
	})
	t.Run("FieldExtractor", func(t *testing.T) {
		// Unknown output values fall through to the field extractor, not an error
		p, err := GetPrinter("not a printer")
		require.NoError(t, err)
		assert.NotNil(t, p)
		assert.Implements(t, (*printer.Printer)(nil), p)
	})
}
