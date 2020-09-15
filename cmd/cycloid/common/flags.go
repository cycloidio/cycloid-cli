package common

import (
	"github.com/spf13/cobra"
)

var projectFlag string
var envFlag string
var orgFlag string
var credFlag, idFlag uint32

func WithFlagOrg(cmd *cobra.Command) string {
	flagName := "org"
	cmd.PersistentFlags().StringVar(&orgFlag, flagName, "", "Org cannonical name")
	return flagName
}

func WithFlagProject(cmd *cobra.Command) string {
	flagName := "project"
	cmd.PersistentFlags().StringVar(&projectFlag, flagName, "", "Project cannonical name")
	return flagName
}

func WithFlagEnv(cmd *cobra.Command) string {
	flagName := "env"
	cmd.PersistentFlags().StringVar(&envFlag, flagName, "", "Environment")
	return flagName
}

func WithFlagID(cmd *cobra.Command) string {
	flagName := "id"
	// TODO  how make it nil or without any value in case we don't want any creds ?
	cmd.Flags().Uint32Var(&idFlag, flagName, 0, "id")
	return flagName
}

func WithFlagCred(cmd *cobra.Command) string {
	flagName := "cred"
	// TODO  how make it nil or without any value in case we don't want any creds ?
	cmd.Flags().Uint32Var(&credFlag, flagName, 0, "cred id")
	return flagName
}
