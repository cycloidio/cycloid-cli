package cyargs

import (
	"fmt"
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
		return cobra.AppendActiveHelp(nil, "missing org for completion: "+err.Error()), cobra.ShellCompDirectiveError
	}

	var stackRef string
	stackRef, _ = GetStackRef(cmd)
	if stackRef == "" { // Check for blueprint ref in case of cy stack create
		stackRef, err = GetBlueprintRef(cmd)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "missing stack-ref or blueprint ref for completion: "+err.Error()), cobra.ShellCompDirectiveNoFileComp
		}
	}

	// Try to get the stack version flags
	tag, branch, hash, err := GetStackVersionFlags(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to read stack version flags: "+err.Error()), cobra.ShellCompDirectiveError
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// Try to get the stack use cases using the provided flags
	// If none are provided yet, this will use default version
	stackUseCases, err := m.ListStackUseCases(org, stackRef, tag, branch, hash)
	if err != nil {
		// During completion, the version flags might not be set yet
		// Return a helpful message to guide the user
		return cobra.AppendActiveHelp(
				nil,
				fmt.Sprintf("Cannot detect version for stack %q: %s. Check and refresh your catalog repository.", stackRef, err.Error()),
			),
			cobra.ShellCompDirectiveNoFileComp
	}

	var useCases []string
	for _, useCase := range stackUseCases {
		if strings.HasPrefix(*useCase.UseCase, toComplete) {
			desc := *useCase.Name
			if useCase.Description != "" {
				desc = desc + " - " + useCase.Description
			}
			useCases = append(useCases, cobra.CompletionWithDesc(*useCase.UseCase, desc))
		}
	}

	return useCases, cobra.ShellCompDirectiveNoFileComp
}

func GetUseCase(cmd *cobra.Command) (string, error) {
	useCase, err := cmd.Flags().GetString("use-case")
	if err != nil {
		return "", err
	}

	return useCase, nil
}
