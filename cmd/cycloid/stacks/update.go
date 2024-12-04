package stacks

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update",
		Short: "update a stack",
		Long:  "The CLI will fetch the current value of the --stack-ref and update with the field you will provide",
		Example: `# Update one stack visibility
cy stacks update --stack-ref org:myStack --visibility shared

# Full args example
cy stacks update \
	--stack-ref my:stack-ref \
	--name "Stack display name" \
	--author "authorCanonical" \
	--description "" \
	--keyword "keyword1,keyword2" \
	--keyword "keyword3" \
	--image "https://url-to-img.example.com/img" \
	--visibility "local" \
	--team "teamCanonical"
`,
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	cmd.Flags().String("stack-ref", "", "stack reference, format 'org:stack-canonical'")
	cmd.MarkFlagRequired("stack-ref")

	cmd.Flags().String("name", "", "update the stack display name")
	cmd.Flags().String("author", "", "update the stack author")
	cmd.Flags().String("description", "", "update the stack description")
	cmd.Flags().StringSlice("keyword", []string{}, "update the stack keywords (will replace keywords, not append them.)")
	cmd.Flags().String("image", "", "update the stack image, must be a valid URL")
	cmd.Flags().String("visibility", "", "update stack visibility")
	cmd.Flags().String("team", "", "update the maintainer team canonical")

	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := common.GetOrg(cmd)
	if err != nil {
		return err
	}

	stackRef, err := cmd.Flags().GetString("stack-ref")
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// Fetch the current stack state
	stack, err := m.GetStack(org, stackRef)
	if err != nil {
		return printer.SmartPrint(p, nil, err, fmt.Sprintf("failed to retrieve the stack with stack ref: %s", stackRef), printer.Options{}, cmd.OutOrStderr())
	}

	// Manage optional parameters
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	if name == "" {
		name = *stack.Name
	}

	author, err := cmd.Flags().GetString("author")
	if err != nil {
		return err
	}

	if author == "" {
		author = *stack.Author
	}

	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return err
	}

	if description == "" {
		description = *stack.Description
	}

	var imageUrl strfmt.URI
	image, err := cmd.Flags().GetString("image")
	if err != nil {
		return err
	}

	if image == "" {
		imageUrl = stack.Image
	} else {
		imageUrl = strfmt.URI(image)
	}

	visibility, err := cmd.Flags().GetString("visibility")
	if err != nil {
		return err
	}

	if visibility == "" {
		visibility = *stack.Visibility
	}

	team, err := cmd.Flags().GetString("team")
	if err != nil {
		return err
	}

	if team == "" {
		team = *stack.Team.Canonical
	}

	keywords, err := cmd.Flags().GetStringSlice("keyword")
	if err != nil {
		return err
	}

	if keywords == nil {
		keywords = stack.Keywords
	}

	s, err := m.UpdateStack(org, stackRef, name, *stack.Canonical, author, description, visibility, stack.ServiceCatalogSourceCanonical, team, imageUrl, keywords, stack.Technologies, stack.Dependencies)
	return printer.SmartPrint(p, s, err, fmt.Sprintf("fail to update stack with ref: %s", stackRef), printer.Options{}, cmd.OutOrStdout())
}
