package yamlformatter_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/interpolator/formatters"
	"github.com/cycloidio/cycloid-cli/interpolator/resolvers/mockresolver"
	"github.com/cycloidio/cycloid-cli/interpolator/resources"
)

func TestYAMLFormatting(t *testing.T) {
	mo := mockresolver.NewMockResolverWithDefault()

	tcs := []struct {
		name   string
		ref    *resources.Reference
		expect string
	}{
		{
			"TestSimpleStringOk",
			&resources.Reference{
				Path:   "/simple/string",
				Params: url.Values{"output": {"yaml"}},
			},
			"simple",
		},
		{
			"TestSimpleNumOk",
			&resources.Reference{
				Path:   "/simple/num",
				Params: url.Values{"output": {"yaml"}},
			},
			"1\n",
		},
		{
			"TestSimpleBoolOk",
			&resources.Reference{
				Path:   "/simple/bool",
				Params: url.Values{"output": {"yaml"}},
			},
			"true\n",
		},
		{
			"TestSimpleObjectOk",
			&resources.Reference{
				Path:   "/simple/object",
				Params: url.Values{"output": {"yaml"}},
			},
			"str: simple\nnum: 1\nboolean: true\n",
		},
		{
			"TestListStringOk",
			&resources.Reference{
				Path:   "/list/string",
				Params: url.Values{"output": {"yaml"}},
			},
			"- one\n- two\n- three\n",
		},
		{
			"TestListNumOk",
			&resources.Reference{
				Path:   "/list/num",
				Params: url.Values{"output": {"yaml"}},
			},
			"- 1\n- 2\n- 3\n",
		},
		{
			"TestListBoolOk",
			&resources.Reference{
				Path:   "/list/bool",
				Params: url.Values{"output": {"yaml"}},
			},
			"- true\n- false\n- true\n",
		},
		{
			"TestListObjectOk",
			&resources.Reference{
				Path:   "/list/object",
				Params: url.Values{"output": {"yaml"}},
			},
			`- str: one
  num: 1
  boolean: true
- str: two
  num: 2
  boolean: false
- str: three
  num: 3
  boolean: true
`,
		},
		{
			"TestBigObjectIndentedOk",
			&resources.Reference{
				Path: "/simple/bigObject",
				Params: url.Values{
					"output":      {"yaml"},
					"indent_size": {"4"},
				},
			},
			`obj:
    str: simple
    num: 3
    boolean: false
list:
  - hello
  - world
listObj:
  - str: one
    num: 1
    boolean: true
  - str: two
    num: 2
    boolean: false
  - str: three
    num: 3
    boolean: true
`,
		},
	}
	for _, tc := range tcs {
		data, err := mo.Resolve(tc.ref)
		assert.NoError(t, err, "the resolution should not fail, check mockresolver")
		yamlFormatter := formatters.New(tc.ref.Params)
		result, err := yamlFormatter.Format(data)
		assert.NoError(t, err, "Formatting should not fail")
		assert.Equal(t, tc.expect, result, "result should match expectation")
	}
}
