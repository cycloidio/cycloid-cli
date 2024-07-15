package common

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	projectFlag, envFlag, credFlag, canFlag string
	idFlag                                  uint32
)

func GetOrg(cmd *cobra.Command) (org string, err error) {
	org = viper.GetString("org")
	if org == "" {
		return "", errors.New("org is not set, use --org flag or CY_ORG env var")
	}

	if viper.GetString("verbosity") == "debug" {
		fmt.Fprintln(os.Stderr, "\033[1;34mdebug:\033[0m using org:", org)
	}

	return org, nil
}

func WithFlagProject(cmd *cobra.Command) string {
	flagName := "project"
	cmd.PersistentFlags().StringVar(&projectFlag, flagName, "", "Project canonical name")
	return flagName
}

func WithFlagEnv(cmd *cobra.Command) string {
	flagName := "env"
	cmd.PersistentFlags().StringVar(&envFlag, flagName, "", "Environment")
	return flagName
}

func WithFlagCan(cmd *cobra.Command) string {
	flagName := "canonical"
	// TODO  how make it nil or without any value in case we don't want any creds ?
	cmd.Flags().StringVar(&canFlag, flagName, "", "canonical")
	return flagName
}

func WithPersistentFlagCan(cmd *cobra.Command) string {
	flagName := "canonical"
	// TODO  how make it nil or without any value in case we don't want any creds ?
	cmd.PersistentFlags().StringVar(&canFlag, flagName, "", "canonical")
	return flagName
}

func WithFlagCred(cmd *cobra.Command) string {
	flagName := "cred"
	// TODO  how make it nil or without any value in case we don't want any creds ?
	cmd.Flags().StringVar(&credFlag, flagName, "", "cred canonical")
	return flagName
}

func WithFlagID(cmd *cobra.Command) string {
	flagName := "id"
	cmd.Flags().Uint32Var(&idFlag, flagName, 0, "id")
	return flagName
}
