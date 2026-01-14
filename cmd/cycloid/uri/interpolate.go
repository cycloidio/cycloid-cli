package uri

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/interpolator"
	"github.com/cycloidio/cycloid-cli/interpolator/parsers"
	"github.com/cycloidio/cycloid-cli/interpolator/resolvers/httpresolver"
)

func NewInterpolateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "interpolate [files...]",
		Short: "Interpolate Cycloid URI references in a file.",
		Long: `This command will parse files for Cycloid URI references and interpolate
them with the requested resource

This command is meant to inject values like credentials, inventory values, terraform outputs
or any Cycloid API resource at runtime (for example in pipelines).
` + interpolator.Docs,
		Example: `Interpolate from stdin and output content on stdout
  echo "ssh: |- cy://org/some_org/credentials/some_ssh?key=.raw.ssh_key" | cy uri interpolate

Inject credential in multiple files in place
  cy uri interpolate -i file1.yaml file2.yaml

Crawl through all the files in the current directory, ignore .git directory
  cy uri interpolate --in-place --recurse . --ignore .git
`,
		Args: cobra.MatchAll(
			cyargs.ValidateFSArguments,
		),
		RunE: interpolate,
	}

	cyargs.AddFSRecurseFlag(cmd)
	cyargs.AddFSIgnoreFlag(cmd)
	cyargs.AddInPlaceFlag(cmd)
	cyargs.AddOutputDirectoryFlag(cmd)
	return cmd
}

func interpolate(cmd *cobra.Command, args []string) error {
	recurse, err := cyargs.GetFSRecurseFlag(cmd)
	if err != nil {
		return err
	}

	ignores, err := cyargs.GetFSIgnoreFlag(cmd)
	if err != nil {
		return err
	}

	paths := args
	if recurse {
		paths, err = RecurseFS(args, ignores)
		if err != nil {
			return err
		}
	}

	inPlace, err := cyargs.GetInPlaceFlag(cmd)
	if err != nil {
		return err
	}

	targetDir, err := cyargs.GetOutputDirectoryFlag(cmd)
	if err != nil {
		return err
	}

	// Ensure the target dir is created
	if targetDir != "" {
		err = os.MkdirAll(targetDir, 0750)
		if err != nil {
			return fmt.Errorf("failed to create target directory %q: %w", targetDir, err)
		}
	}

	resolver, err := httpresolver.NewHTTPResolver()
	if err != nil {
		return err
	}

	// Manage stdin input
	if len(args) == 0 {
		if !common.DetectStdinInput() {
			return fmt.Errorf("stdin looks empty, please fill stdin or a filename as argument")
		}

		stdin, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return err
		}

		if len(stdin) == 0 {
			return fmt.Errorf("failed to read from stdin, looks empty")
		}

		out, err := parsers.ReplaceFile(resolver, string(stdin))
		if err != nil {
			return fmt.Errorf("failed to interpolate from stdin: %w", err)
		}
		cleanedOut := strings.TrimRight(out, " \t\n")

		_, err = fmt.Fprintln(cmd.OutOrStdout(), cleanedOut)
		if err != nil {
			return fmt.Errorf("failed to print result to stdout: %w", err)
		}

		return nil
	}

	output := []string{}
	for _, filename := range paths {
		// wrap logic in a func to trigger defers on end of loop
		stats, err := os.Stat(filename)
		if err != nil {
			return fmt.Errorf("file not found for interpolation %q: %w", stats, err)
		}

		content, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("failed to open file named %q: %w", filename, err)
		}

		out, err := parsers.ReplaceFile(resolver, string(content))
		if err != nil {
			return fmt.Errorf("failed to interpolate file %q: %w", filename, err)
		}

		if inPlace {
			err := os.WriteFile(filename, []byte(out+"\n"), stats.Mode())
			if err != nil {
				return fmt.Errorf("failed to write to file %q during interpolation, file content may be lost: %w", filename, err)
			}
		} else if targetDir != "" {
			// Send to target dir
			var target string
			if filepath.IsLocal(filename) {
				target = filepath.Join(targetDir, filename)
			} else {
				target = filepath.Join(targetDir, filepath.Base(filename))
			}

			err = os.WriteFile(target, []byte(out+"\n"), 0640)
			if err != nil {
				return fmt.Errorf("failed to write file %q in dir %q: %w", filename, targetDir, err)
			}
		} else {
			output = append(output, out)
		}
	}

	if !inPlace {
		_, err := fmt.Fprintln(cmd.OutOrStdout(), strings.Join(output, "\n"))
		if err != nil {
			return fmt.Errorf("failed to print result to stdout: %w", err)
		}
	}

	return nil
}
