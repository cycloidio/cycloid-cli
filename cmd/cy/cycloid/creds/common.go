package creds

import "github.com/spf13/cobra"

var (
	typeFlag           string
	nameFlag           string
	descriptionFlag    string
	pathFlag           string
	fieldsFlags        map[string]string
	fieldsFileFlags    map[string]string
	usernameFlag       string
	accessKeyFlag      string
	secretKeyFlag      string
	clientIDFlag       string
	clientSecretFlag   string
	subscriptionIDFlag string
	accountNameFlag    string
	authUrlFlag        string
	domainIDFlag       string
	tenantIDFlag       string
	passwordFlag       string
	sshKeyFlag         string
	jsonKeyFlag        string
	caCertFlag         string
)

func WithFlagCaCert(cmd *cobra.Command) string {
	flagName := "ca-cert"
	cmd.Flags().StringVar(&caCertFlag, flagName, "", "ca-cert")
	return flagName
}
func WithFlagAuthUrl(cmd *cobra.Command) string {
	flagName := "auth-url"
	cmd.Flags().StringVar(&authUrlFlag, flagName, "", "auth-url")
	return flagName
}
func WithFlagDomainID(cmd *cobra.Command) string {
	flagName := "domain-id"
	cmd.Flags().StringVar(&domainIDFlag, flagName, "", "domain-id")
	return flagName
}
func WithFlagAccountName(cmd *cobra.Command) string {
	flagName := "account-name"
	cmd.Flags().StringVar(&accountNameFlag, flagName, "", "account-name")
	return flagName
}
func WithFlagClientID(cmd *cobra.Command) string {
	flagName := "client-id"
	cmd.Flags().StringVar(&clientIDFlag, flagName, "", "client-id")
	return flagName
}
func WithFlagClientSecret(cmd *cobra.Command) string {
	flagName := "client-secret"
	cmd.Flags().StringVar(&clientSecretFlag, flagName, "", "client-secret")
	return flagName
}
func WithFlagSubscriptionID(cmd *cobra.Command) string {
	flagName := "subscription-id"
	cmd.Flags().StringVar(&subscriptionIDFlag, flagName, "", "subscription-id")
	return flagName
}
func WithFlagTenantID(cmd *cobra.Command) string {
	flagName := "tenant-id"
	cmd.Flags().StringVar(&tenantIDFlag, flagName, "", "tenant-id")
	return flagName
}
func WithFlagAccessKey(cmd *cobra.Command) string {
	flagName := "access-key"
	cmd.Flags().StringVar(&accessKeyFlag, flagName, "", "access-key")
	return flagName
}
func WithFlagSecretKey(cmd *cobra.Command) string {
	flagName := "secret-key"
	cmd.Flags().StringVar(&secretKeyFlag, flagName, "", "secret-key")
	return flagName
}
func WithFlagField(cmd *cobra.Command) string {
	flagName := "field"
	cmd.Flags().StringToStringVar(&fieldsFlags, flagName, nil, "key=value")
	return flagName
}
func WithFlagFieldFile(cmd *cobra.Command) string {
	flagName := "field-file"
	cmd.Flags().StringToStringVar(&fieldsFileFlags, flagName, nil, "key=/file/path")
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
func WithFlagJsonKey(cmd *cobra.Command) string {
	flagName := "json-key"
	cmd.Flags().StringVar(&jsonKeyFlag, flagName, "", "json key path")
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
