package stacks_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/stacks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateDoc(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var b bytes.Buffer

		exp, err := ioutil.ReadFile("./testdata/stacks/success.md")
		require.Nil(t, err)

		err = stacks.GenerateDoc("./testdata/stacks/success.yml", &b)
		require.Nil(t, err)

		assert.Equal(t, exp, b.Bytes())
	})
	t.Run("Error", func(t *testing.T) {
		var b bytes.Buffer

		err := stacks.GenerateDoc("./testdata/stacks/error.yml", &b)
		require.NotNil(t, err)
	})
}
