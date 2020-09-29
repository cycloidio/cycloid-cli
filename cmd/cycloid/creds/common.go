package creds

import "github.com/spf13/cobra"

var (
	typeFlag        string
	nameFlag        string
	descriptionFlag string
	pathFlag        string
	fieldsFlags     map[string]string
	usernameFlag    string
	passwordFlag    string
	sshKeyFlag      string
)

func WithFlagField(cmd *cobra.Command) string {
	flagName := "field"
	cmd.Flags().StringToStringVar(&fieldsFlags, flagName, nil, "key=value")
	return flagName
}

func WithFlagType(cmd *cobra.Command) string {
	flagName := "type"
	cmd.Flags().StringVar(&typeFlag, flagName, "", "type")
	return flagName
}
func WithFlagUsername(cmd *cobra.Command) string {
	flagName := "username"
	cmd.Flags().StringVar(&usernameFlag, flagName, "", "username")
	return flagName
}
func WithFlagPassword(cmd *cobra.Command) string {
	flagName := "password"
	cmd.Flags().StringVar(&passwordFlag, flagName, "", "password")
	return flagName
}
func WithFlagSSHKey(cmd *cobra.Command) string {
	flagName := "ssh-key"
	cmd.Flags().StringVar(&sshKeyFlag, flagName, "", "ssh key path")
	return flagName
}
func WithPersistentFlagName(cmd *cobra.Command) string {
	flagName := "name"
	cmd.PersistentFlags().StringVar(&nameFlag, flagName, "", "name")
	return flagName
}
func WithPersistentFlagDescription(cmd *cobra.Command) string {
	flagName := "description"
	cmd.PersistentFlags().StringVar(&descriptionFlag, flagName, "", "Description")
	return flagName
}
func WithPersistentFlagPath(cmd *cobra.Command) string {
	flagName := "path"
	cmd.PersistentFlags().StringVar(&pathFlag, flagName, "", "Path")
	return flagName
}
