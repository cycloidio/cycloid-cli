package catalogRepositories

import (
	"time"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewTemplateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "template",
		Short: "template a stack for a catalog repository",
		Example: `
  Create a stack from a template using a template repository
  cy --org my-org cr template \
    -r my-catalog-repo \
    -s source-catalog:ref \
    --stack-name "My stack stack-name" \
    --canonical my-stack-canonical \
    --use-case stack-use-case
`,
		RunE: NewStackFromTemplate,
	}

	timestamp := uint64(time.Now().Unix())

	cmd.Flags().StringP("service-catalog-reference", "r", "", "[required] catalog reference")
	cmd.MarkFlagRequired("service-catalog-reference")
	cmd.Flags().StringP("service-catalog-source-canonical", "s", "", "[required] catalog source")
	cmd.MarkFlagRequired("service-catalog-source-canonical")
	cmd.Flags().StringP("stack-name", "n", "", "[required] stack stack-name")
	cmd.MarkFlagRequired("stack-name")
	cmd.Flags().StringP("stack-canonical", "c", "", "[required] stack canonical")
	cmd.MarkFlagRequired("stack-canonical")
	cmd.Flags().StringP("use-case", "u", "", "[required] stack use-case")
	cmd.MarkFlagRequired("use-case")
	cmd.Flags().StringP("author", "a", "", "stack author")
	cmd.Flags().Uint64("created-at", timestamp, "stack created-at")
	cmd.Flags().Uint64("updated-at", timestamp, "stack updated-at")

	return cmd
}

func NewStackFromTemplate(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	service_catalog_reference, err := cmd.Flags().GetString("service-catalog-reference")
	if err != nil {
		return err
	}

	stack_name, err := cmd.Flags().GetString("stack-name")
	if err != nil {
		return err
	}

	stack_canonical, err := cmd.Flags().GetString("stack-canonical")
	if err != nil {
		return err
	}

	author, err := cmd.Flags().GetString("author")
	if err != nil {
		return err
	}

	service_catalog_source_canonical, err := cmd.Flags().GetString("service-catalog-source-canonical")
	if err != nil {
		return err
	}

	use_case, err := cmd.Flags().GetString("use-case")
	if err != nil {
		return err
	}

	created_at, err := cmd.Flags().GetUint64("created-at")
	if err != nil {
		return err
	}

	updated_at, err := cmd.Flags().GetUint64("updated-at")
	if err != nil {
		return err
	}

	// Get output settings
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// Send result to printer
	stack, err := m.NewServiceCatalogFromTemplate(org, service_catalog_reference, stack_name, stack_canonical, author, service_catalog_source_canonical, use_case, created_at, updated_at)
	return printer.SmartPrint(p, stack, err, "failed to template stack", printer.Options{}, cmd.OutOrStdout())
}
