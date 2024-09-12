package infrapolicies

import "github.com/spf13/cobra"

var (
	planPathFlag    string
	policyPathFlag  string
	nameFlag        string
	ownerFlag       string
	severityFlag    string
	canonicalFlag   string
	descriptionFlag string
	enabledFlag     bool
)

func WithFlagPlanPath(cmd *cobra.Command) string {
	flagName := "plan-path"
	cmd.PersistentFlags().StringVar(&planPathFlag, flagName, "", "Path to the terraform plan result")
	return flagName
}

func WithFlagPolicyPath(cmd *cobra.Command) string {
	flagName := "policy-path"
	cmd.PersistentFlags().StringVar(&policyPathFlag, flagName, "", "Path to the infraPolicy rego file")
	return flagName
}

func WithFlagName(cmd *cobra.Command) string {
	flagName := "name"
	cmd.PersistentFlags().StringVar(&nameFlag, flagName, "", "Name of the infraPolicy")
	return flagName
}

func WithFlagOwner(cmd *cobra.Command) string {
	flagName := "owner"
	cmd.PersistentFlags().StringVar(&ownerFlag, flagName, "", "InfraPolicy's owner canonical")
	return flagName
}

func WithFlagSeverity(cmd *cobra.Command) string {
	flagName := "severity"
	cmd.PersistentFlags().StringVar(&severityFlag, flagName, "", "InfraPolicy's severity. Should be of type: critical|warning|advisory ")
	return flagName
}

func WithFlagcanonical(cmd *cobra.Command) string {
	flagName := "canonical"
	cmd.PersistentFlags().StringVar(&canonicalFlag, flagName, "", "InfraPolicy's canonical")
	return flagName
}

func WithFlagDescription(cmd *cobra.Command) string {
	flagName := "description"
	cmd.PersistentFlags().StringVar(&descriptionFlag, flagName, "", "InfraPolicy's description")
	return flagName
}

func WithFlagEnabled(cmd *cobra.Command) string {
	flagName := "enabled"
	cmd.PersistentFlags().BoolVar(&enabledFlag, flagName, false, "Whether to enable or not the infraPolicy. Note! You have to specify enabled=true|false enabled false|true doesn't work")
	return flagName
}
