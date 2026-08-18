package cyargs

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

func AddBootstrapUserFlags(cmd *cobra.Command) {
	cmd.Flags().String("username", "", "admin username for the first user")
	cmd.Flags().String("full-name", "", "full name for the first user")
	cmd.Flags().String("email", "", "email for the first user")
	cmd.Flags().String("password", "", "password for the first user")
	cmd.Flags().Bool("password-stdin", false, "read the password from stdin")

	_ = cmd.MarkFlagRequired("username")
	_ = cmd.MarkFlagRequired("full-name")
	_ = cmd.MarkFlagRequired("email")

	cmd.MarkFlagsMutuallyExclusive("password", "password-stdin")
}

func AddBootstrapAPIKeyCanonicalFlag(cmd *cobra.Command) string {
	flagName := "api-key-canonical"
	cmd.Flags().String(flagName, "", "optional canonical for an admin API key and matching credential")
	return flagName
}

func GetBootstrapPassword(cmd *cobra.Command) (string, error) {
	fromStdin, err := cmd.Flags().GetBool("password-stdin")
	if err != nil {
		return "", err
	}
	password, err := cmd.Flags().GetString("password")
	if err != nil {
		return "", err
	}

	if fromStdin && password != "" {
		return "", fmt.Errorf("only one of --password and --password-stdin may be set")
	}
	if fromStdin {
		content, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return "", fmt.Errorf("failed to read password from stdin: %w", err)
		}
		password = strings.TrimSpace(string(content))
		if password == "" {
			return "", fmt.Errorf("stdin looks empty, please provide a password")
		}
		return password, nil
	}
	if password == "" {
		return "", fmt.Errorf("password is required: pass --password or --password-stdin")
	}
	return password, nil
}

func GetBootstrapAPIKeyCanonical(cmd *cobra.Command) (*string, error) {
	canonical, err := cmd.Flags().GetString("api-key-canonical")
	if err != nil {
		return nil, err
	}
	if canonical == "" {
		return nil, nil
	}
	return &canonical, nil
}
