package cyargs

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

func AddLicenceKeyFlag(cmd *cobra.Command) string {
	flagName := "key"
	cmd.Flags().String(flagName, "", "licence key value")
	return flagName
}

func AddLicenceKeyFileFlag(cmd *cobra.Command) string {
	flagName := "key-file"
	cmd.Flags().String(flagName, "", "path to a file containing the licence key")
	cmd.RegisterFlagCompletionFunc(flagName, completeFilePath)
	return flagName
}

func AddLicenceFlag(cmd *cobra.Command) string {
	flagName := "licence"
	cmd.Flags().String(flagName, "", "licence key value")
	return flagName
}

func AddLicenceFileFlag(cmd *cobra.Command) string {
	flagName := "licence-file"
	cmd.Flags().String(flagName, "", "path to a file containing the licence key")
	cmd.RegisterFlagCompletionFunc(flagName, completeFilePath)
	return flagName
}

func GetLicenceKey(cmd *cobra.Command) (string, error) {
	return readSecretValue(cmd, "key", "key-file")
}

func GetLicence(cmd *cobra.Command) (string, error) {
	return readSecretValue(cmd, "licence", "licence-file")
}

func readSecretValue(cmd *cobra.Command, valueFlag, fileFlag string) (string, error) {
	value, err := cmd.Flags().GetString(valueFlag)
	if err != nil {
		return "", err
	}
	file, err := cmd.Flags().GetString(fileFlag)
	if err != nil {
		return "", err
	}

	if value != "" && file != "" {
		return "", fmt.Errorf("only one of --%s and --%s may be set", valueFlag, fileFlag)
	}
	if value != "" {
		return strings.TrimSpace(value), nil
	}
	if file != "" {
		content, err := os.ReadFile(file)
		if err != nil {
			return "", fmt.Errorf("failed to read licence from %q: %w", file, err)
		}
		key := strings.TrimSpace(string(content))
		if key == "" {
			return "", fmt.Errorf("licence file %q is empty", file)
		}
		return key, nil
	}
	if common.DetectStdinInput() {
		content, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return "", fmt.Errorf("failed to read licence from stdin: %w", err)
		}
		key := strings.TrimSpace(string(content))
		if key == "" {
			return "", fmt.Errorf("stdin looks empty, please provide a licence key")
		}
		return key, nil
	}

	return "", fmt.Errorf(
		"licence key required: pass --%s, --%s, or pipe the key via stdin",
		valueFlag, fileFlag,
	)
}

func completeFilePath(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	return nil, cobra.ShellCompDirectiveFilterFileExt
}
