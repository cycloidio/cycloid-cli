package output

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/config"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
)

// NewOutputCmd returns the cobra command for managing the default output format.
func NewOutputCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "output",
		Short: "Manage the default output format",
		Long: `Manage the default --output format persisted in the config file.

The effective output format is resolved in this priority order:
  1. --jq flag
  2. --output / -o flag
  3. CY_OUTPUT environment variable
  4. Value stored by 'cy output set' (config file)
  5. "table" (built-in default)`,
		Args: cobra.NoArgs,
	}
	cmd.AddCommand(
		newSetCmd(),
		newGetCmd(),
		newResetCmd(),
	)
	return cmd
}

func newSetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set <value>",
		Short: "Persist a default output format",
		Args:  cobra.ExactArgs(1),
		Example: `cy output set table:border
cy output set json
cy output set table=canonical,name`,
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, _ := config.Read()
			if conf == nil {
				conf = &config.Config{}
			}
			conf.Output = args[0]
			if err := config.Write(conf); err != nil {
				return fmt.Errorf("unable to save output config: %w", err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Default output set to %q\n", args[0])
			return nil
		},
	}
}

func newGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Show the current effective default output format",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			effective, err := cyargs.GetOutput(cmd)
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), effective)
			return nil
		},
	}
}

func newResetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Short: "Remove the persisted default output format",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, _ := config.Read()
			if conf == nil {
				conf = &config.Config{}
			}
			conf.Output = ""
			if err := config.Write(conf); err != nil {
				return fmt.Errorf("unable to save output config: %w", err)
			}
			fmt.Fprintln(cmd.OutOrStdout(), "Default output reset to \"table\"")
			return nil
		},
	}
}
