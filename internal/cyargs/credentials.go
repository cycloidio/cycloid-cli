package cyargs

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func AddCredentialDescriptionPersistentFlag(cmd *cobra.Command) string {
	flagName := "description"
	cmd.PersistentFlags().StringP(flagName, "d", "", "add an helpful description to your credential.")
	return flagName
}

func AddCredentialDescriptionFlag(cmd *cobra.Command) string {
	flagName := "description"
	cmd.Flags().StringP(flagName, "d", "", "add an helpful description to your credential.")
	return flagName
}

func GetCredentialDescription(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("description")
}

func AddCredentialCanonicalPersistentFlag(cmd *cobra.Command) string {
	flagName := "canonical"
	cmd.PersistentFlags().StringP(flagName, "c", "", "the canonical of your credential")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteCredentialCanonical)
	return flagName
}

func AddCredentialCanonicalFlag(cmd *cobra.Command) string {
	flagName := "canonical"
	cmd.Flags().StringP(flagName, "c", "", "the canonical of your credential")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteCredentialCanonical)
	return flagName
}

func CompleteCredentialCanonical(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	creds, err := m.ListCredentials(org, "")
	if err != nil {
		return cobra.AppendActiveHelp(nil, "completion failed: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	var completions = make([]cobra.Completion, len(creds))
	for index, cred := range creds {
		if cred.Canonical != nil && strings.HasPrefix(*cred.Canonical, toComplete) {
			completions[index] = cobra.CompletionWithDesc(*cred.Canonical, *cred.Name+": "+cred.Description)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func GetCredentialCanonical(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("canonical")
}

func AddCredentialPathPersistentFlag(cmd *cobra.Command) string {
	flagName := "path"
	cmd.PersistentFlags().StringP(flagName, "p", "", "the path of your credential")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteCredentialPath)
	return flagName
}

func AddCredentialPathFlag(cmd *cobra.Command) string {
	flagName := "path"
	cmd.Flags().StringP(flagName, "p", "", "the path of your credential")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteCredentialPath)
	return flagName
}

func CompleteCredentialPath(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	creds, err := m.ListCredentials(org, "")
	if err != nil {
		return cobra.AppendActiveHelp(nil, "completion failed: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	var completions = make([]cobra.Completion, len(creds))
	for index, cred := range creds {
		if cred.Path != nil && strings.HasPrefix(*cred.Path, toComplete) {
			completions[index] = cobra.CompletionWithDesc(*cred.Path, *cred.Name+": "+cred.Description)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func GetCredentialPath(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("path")
}

func AddCredentialNamePersistentFlag(cmd *cobra.Command) string {
	flagName := "name"
	cmd.PersistentFlags().StringP(flagName, "n", "", "the name of your credential")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteCredentialName)
	return flagName
}

func AddCredentialNameFlag(cmd *cobra.Command) string {
	flagName := "name"
	cmd.Flags().StringP(flagName, "n", "", "the name of your credential")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteCredentialName)
	return flagName
}

func CompleteCredentialName(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	creds, err := m.ListCredentials(org, "")
	if err != nil {
		return cobra.AppendActiveHelp(nil, "completion failed: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	var completions = make([]cobra.Completion, len(creds))
	for index, cred := range creds {
		if cred.Name != nil && strings.HasPrefix(*cred.Name, toComplete) {
			completions[index] = cobra.CompletionWithDesc(*cred.Name, *cred.Canonical+": "+cred.Description)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func GetCredentialName(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("name")
}

func AddCredentialCaCertFlag(cmd *cobra.Command) string {
	flagName := "ca-cert"
	cmd.Flags().String(flagName, "", "ca-cert")
	return flagName
}

func GetCredentialCaCert(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("ca-cert")
}

func AddCredentialAuthURLFlag(cmd *cobra.Command) string {
	flagName := "auth-url"
	cmd.Flags().String(flagName, "", "auth-url")
	return flagName
}

func GetCredentialAuthURL(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("auth-url")
}

func AddCredentialDomainIDFlag(cmd *cobra.Command) string {
	flagName := "domain-id"
	cmd.Flags().String(flagName, "", "domain-id")
	return flagName
}

func GetCredentialDomainID(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("domain-id")
}

func AddCredentialAccountNameFlag(cmd *cobra.Command) string {
	flagName := "account-name"
	cmd.Flags().String(flagName, "", "account-name")
	return flagName
}

func GetCredentialAccountName(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("account-name")
}

func AddCredentialClientIDFlag(cmd *cobra.Command) string {
	flagName := "client-id"
	cmd.Flags().String(flagName, "", "client-id")
	return flagName
}

func GetCredentialClientID(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("client-id")
}

func AddCredentialClientSecretFlag(cmd *cobra.Command) string {
	flagName := "client-secret"
	cmd.Flags().String(flagName, "", "client-secret")
	return flagName
}

func GetCredentialClientSecret(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("client-secret")
}

func AddCredentialSubscriptionIDFlag(cmd *cobra.Command) string {
	flagName := "subscription-id"
	cmd.Flags().String(flagName, "", "subscription-id")
	return flagName
}

func GetCredentialSubscriptionID(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("subscription-id")
}

func AddCredentialTenantIDFlag(cmd *cobra.Command) string {
	flagName := "tenant-id"
	cmd.Flags().String(flagName, "", "tenant-id")
	return flagName
}

func GetCredentialTenantID(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("tenant-id")
}

func AddCredentialAccessKeyFlag(cmd *cobra.Command) string {
	flagName := "access-key"
	cmd.Flags().String(flagName, "", "access-key")
	return flagName
}

func GetCredentialAccessKey(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("access-key")
}

func AddCredentialSecretKeyFlag(cmd *cobra.Command) string {
	flagName := "secret-key"
	cmd.Flags().String(flagName, "", "secret-key")
	return flagName
}

func GetCredentialSecretKey(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("secret-key")
}

func AddCredentialFieldFlag(cmd *cobra.Command) string {
	flagName := "field"
	cmd.Flags().StringToString(flagName, nil, "key=value")
	return flagName
}

func GetCredentialField(cmd *cobra.Command) (map[string]string, error) {
	return cmd.Flags().GetStringToString("field")
}

func AddCredentialFieldFileFlag(cmd *cobra.Command) string {
	flagName := "field-file"
	cmd.Flags().StringToString(flagName, nil, "key=/file/path")
	return flagName
}

func GetCredentialFieldFile(cmd *cobra.Command) (map[string]string, error) {
	return cmd.Flags().GetStringToString("field-file")
}

func AddCredentialTypeFlag(cmd *cobra.Command) string {
	flagName := "type"
	cmd.Flags().String(flagName, "", "type")
	cmd.RegisterFlagCompletionFunc("type", CompleteCredentialType)

	return flagName
}

func CompleteCredentialType(_ *cobra.Command, _ []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	credTypes := []cobra.Completion{
		"ssh",
		"basic_auth",
		"custom",
		"aws",
		"azure",
		"azure_storage",
		"gcp",
		"elasticsearch",
		"swift",
	}

	var comps = make([]cobra.Completion, len(credTypes))
	for index, credType := range credTypes {
		if strings.HasPrefix(credType, toComplete) {
			comps[index] = credType
		}
	}

	return comps, cobra.ShellCompDirectiveNoFileComp
}

func GetCredentialType(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("type")
}

func AddCredentialUsernameFlag(cmd *cobra.Command) string {
	flagName := "username"
	cmd.Flags().String(flagName, "", "username")
	return flagName
}

func GetCredentialUsername(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("username")
}

func AddCredentialPasswordFlag(cmd *cobra.Command) string {
	flagName := "password"
	cmd.Flags().String(flagName, "", "password")
	return flagName
}

func GetCredentialPassword(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("password")
}

func AddCredentialSSHKeyFlag(cmd *cobra.Command) string {
	flagName := "ssh-key"
	cmd.Flags().String(flagName, "", "ssh key path")
	return flagName
}

func GetCredentialSSHKey(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("ssh-key")
}

func AddCredentialJSONKeyFlag(cmd *cobra.Command) string {
	flagName := "json-key"
	cmd.Flags().String(flagName, "", "json key path")
	return flagName
}

func GetCredentialJSONKey(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("json-key")
}
