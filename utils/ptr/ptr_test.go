package ptr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

func TestPtrOrNil(t *testing.T) {
	t.Run("Bool", func(t *testing.T) {
		v := false
		res := ptr.PtrOrNil(v)
		assert.Nil(t, res)

		v = true
		res = ptr.PtrOrNil(v)
		assert.True(t, *res)
	})

	t.Run("Int", func(t *testing.T) {
		v := 0
		res := ptr.PtrOrNil(v)
		assert.Nil(t, res)

		v = 1
		res = ptr.PtrOrNil(v)
		assert.Equal(t, 1, *res)
	})

	t.Run("String", func(t *testing.T) {
		v := ""
		res := ptr.PtrOrNil(v)
		assert.Nil(t, res)

		v = "a"
		res = ptr.PtrOrNil(v)
		assert.Equal(t, "a", *res)
	})

	t.Run("Struct", func(t *testing.T) {
		type S struct {
			A string
			B int
		}
		var v S
		res := ptr.PtrOrNil(v)
		assert.Nil(t, res)

		v.A = "a"
		res = ptr.PtrOrNil(v)
		assert.Equal(t, S{A: "a"}, *res)
	})
}
