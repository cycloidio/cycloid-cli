package stacks

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

// BlueprintSummary represents a simplified view of a blueprint for listing
type BlueprintSummary struct {
	Name        string `json:"name"`
	Ref         string `json:"ref"`
	UseCases    string `json:"usecases"`
	Description string `json:"description"`
}

// BlueprintWithUseCases extends ServiceCatalog with usecases for JSON output
type BlueprintWithUseCases struct {
	*models.ServiceCatalog
	UseCases []string `json:"usecases,omitempty"`
}

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "list",
		Aliases: []string{
			"ls",
		},
		Args:  cobra.NoArgs,
		Short: "list the stacks",
		Example: `cy --org my-org stack list
cy --org my-org stack list --blueprint`,
		RunE: list,
	}

	cmd.Flags().Bool("blueprint", false, "list only blueprint stacks (templates)")

	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	blueprint, err := cmd.Flags().GetBool("blueprint")
	if err != nil {
		return err
	}

	// Initialize middleware after all arguments are collected
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	if blueprint {
		stacks, err := m.ListBlueprints(org)
		if err != nil {
			return printer.SmartPrint(p, nil, err, "failed to list blueprints from API", printer.Options{}, cmd.OutOrStdout())
		}

		// For JSON output, return all fields with usecases added
		if output == "json" {
			blueprintsWithUseCases := make([]*BlueprintWithUseCases, len(stacks))
			for i, stack := range stacks {
				// Get usecases from stack config
				usecases := []string{}
				if stack.Ref != nil {
					config, err := m.GetStackConfig(org, *stack.Ref)
					if err == nil {
						// Extract usecase names from config
						for usecase := range config {
							usecases = append(usecases, usecase)
						}
					}
				}

				blueprintsWithUseCases[i] = &BlueprintWithUseCases{
					ServiceCatalog: stack,
					UseCases:       usecases,
				}
			}
			return printer.SmartPrint(p, blueprintsWithUseCases, nil, "", printer.Options{}, cmd.OutOrStdout())
		}

		// For table output, use simplified format
		blueprints := make([]*BlueprintSummary, len(stacks))
		for i, stack := range stacks {
			// Get usecases from stack config
			usecases := []string{}
			if stack.Ref != nil {
				config, err := m.GetStackConfig(org, *stack.Ref)
				if err == nil {
					// Extract usecase names from config
					for usecase := range config {
						usecases = append(usecases, usecase)
					}
				}
			}

			// Join usecases into a comma-separated string
			usecasesStr := strings.Join(usecases, ", ")

			blueprints[i] = &BlueprintSummary{
				Name:        getStringValue(stack.Name),
				Ref:         getStringValue(stack.Ref),
				UseCases:    usecasesStr,
				Description: stack.Description,
			}
		}

		return printer.SmartPrint(p, blueprints, nil, "", printer.Options{}, cmd.OutOrStdout())
	} else {
		stacks, err := m.ListStacks(org)
		if err != nil {
			return printer.SmartPrint(p, nil, err, "failed to list stacks from API", printer.Options{}, cmd.OutOrStdout())
		}

		return printer.SmartPrint(p, stacks, nil, "", printer.Options{}, cmd.OutOrStdout())
	}
}

// getStringValue safely extracts string value from a string pointer
func getStringValue(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}
