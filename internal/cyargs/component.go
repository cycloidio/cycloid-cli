package cyargs

import (
	"strings"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func AddComponentDescriptionFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("description", "d", "", "set the description of the component")
}

func GetComponentDescription(cmd *cobra.Command) (*string, error) {
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return nil, err
	}

	return &description, nil
}

func AddCloudProviderFlag(cmd *cobra.Command) {
	cmd.Flags().String("cloud-provider", "", "set the cloud provider of the component")
}

func GetCloudProvider(cmd *cobra.Command) (*string, error) {
	cloudProvider, err := cmd.Flags().GetString("cloud-provider")
	if err != nil {
		return nil, err
	}

	return &cloudProvider, nil
}

func AddComponentStackRefFlag(cmd *cobra.Command) string {
	flagName := "stack-ref"
	cmd.Flags().StringP(flagName, "s", "", "set the stack ref of the component in format org:stack-canonical")
	cmd.MarkFlagRequired(flagName)
	cmd.RegisterFlagCompletionFunc(flagName, CompleteStackRef)
	return flagName
}

func GetComponentStackRef(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("stack-ref")
}

// AddUseCaseFlag will add the use-case flag with completion and return the flag name
func AddUseCaseFlag(cmd *cobra.Command) string {
	flagName := "use-case"
	cmd.Flags().StringP(flagName, "u", "", "set the use-case of the component")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteUseCase)
	return flagName
}

func CompleteUseCase(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}

	stackRef, err := GetStackRef(cmd)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	stackConfig, err := m.GetStackConfig(org, stackRef)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}

	var useCases []string
	for useCase, stack := range stackConfig {
		if strings.HasPrefix(useCase, toComplete) {
			desc := *stack.Name
			if stack.Description != nil {
				desc = desc + " - " + *stack.Description
			}
			useCases = append(useCases, cobra.CompletionWithDesc(useCase, desc))
		}
	}

	return useCases, cobra.ShellCompDirectiveNoFileComp
}

func GetUseCase(cmd *cobra.Command) (*string, error) {
	useCase, err := cmd.Flags().GetString("use-case")
	if err != nil {
		return nil, err
	}

	return &useCase, nil
}
