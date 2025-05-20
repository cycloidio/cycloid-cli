package cy_args

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

func AddStackRefFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("stack-ref", "s", "", "set the stack ref of the component in format org:stack-canonical")
	cmd.MarkFlagRequired("stack-ref")
	cmd.RegisterFlagCompletionFunc("stack-ref", func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		api := common.NewAPI()
		m := middleware.NewMiddleware(api)

		org, err := GetOrg(cmd)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "completion failed: "+err.Error()),
				cobra.ShellCompDirectiveNoFileComp
		}

		stacks, err := m.ListStacks(org)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "failed to list stacks in org '"+org+"': "+err.Error()),
				cobra.ShellCompDirectiveNoFileComp
		}

		var stackRefs = make([]string, len(stacks))
		for index, stack := range stacks {
			if stack.Ref != nil && strings.HasPrefix(*stack.Ref, toComplete) {
				desc := *stack.Name
				if stack.Description != "" {
					desc = desc + " - " + stack.Description
				}
				stackRefs[index] = cobra.CompletionWithDesc(*stack.Ref, desc)
			}
		}

		return stackRefs, cobra.ShellCompDirectiveNoFileComp
	})
}

func GetStackRef(cmd *cobra.Command) (*string, error) {
	stackRef, err := cmd.Flags().GetString("stack-ref")
	if err != nil {
		return nil, err
	}

	return &stackRef, nil
}

// AddUseCaseFlag will add the use-case flag with completion and return the flag name
func AddUseCaseFlag(cmd *cobra.Command) string {
	flagName := "use-case"
	cmd.Flags().StringP("use-case", "u", "", "set the use-case of the component")
	cmd.RegisterFlagCompletionFunc("use-case", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		api := common.NewAPI()
		m := middleware.NewMiddleware(api)

		org, err := GetOrg(cmd)
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveNoFileComp
		}

		stackRef, err := GetStackRef(cmd)
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveNoFileComp
		}

		stackConfig, err := m.GetStackConfig(org, *stackRef)
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
	})

	return flagName
}

func GetUseCase(cmd *cobra.Command) (*string, error) {
	useCase, err := cmd.Flags().GetString("use-case")
	if err != nil {
		return nil, err
	}

	return &useCase, nil
}
