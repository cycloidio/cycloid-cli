package e2e_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestExecuteCommandStdin(t *testing.T) {
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			in, err := io.ReadAll(cmd.InOrStdin())
			if err != nil {
				return err
			}

			cmd.OutOrStdout().Write(in)
			return nil
		},
	}

	expected := `My stdin`
	stdoutBuf := new(bytes.Buffer)
	reader := strings.NewReader(expected)
	_, err := reader.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}

	cmd.SetIn(reader)
	cmd.SetOut(stdoutBuf)
	err = cmd.Execute()
	if err != nil {
		t.Fatalf("command failed: %s", err)
	}

	stdout, err := io.ReadAll(stdoutBuf)
	if err != nil {
		t.Fatalf("failed to read cmd output: %s", err)
	}

	assert.Equal(t, expected, string(stdout))
}
