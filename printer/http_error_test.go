package printer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/printer"
)

func TestFirstNLinesFromBytes(t *testing.T) {
	t.Run("FewerThanN", func(t *testing.T) {
		got := printer.FirstNLinesFromBytes([]byte("a\nb"), 10)
		assert.Equal(t, "a\nb", got)
	})
	t.Run("ExactlyN", func(t *testing.T) {
		got := printer.FirstNLinesFromBytes([]byte("1\n2\n3"), 3)
		assert.Equal(t, "1\n2\n3", got)
	})
	t.Run("TruncatesAfterN", func(t *testing.T) {
		got := printer.FirstNLinesFromBytes([]byte("1\n2\n3\n4\n5"), 3)
		assert.Equal(t, "1\n2\n3", got)
	})
	t.Run("ZeroN", func(t *testing.T) {
		assert.Equal(t, "", printer.FirstNLinesFromBytes([]byte("a\nb"), 0))
	})
	t.Run("TrailingSpaceTrimmed", func(t *testing.T) {
		got := printer.FirstNLinesFromBytes([]byte("x\ny\n  \n"), 10)
		assert.Equal(t, "x\ny", got)
	})
}
