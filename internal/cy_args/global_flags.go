package cy_args

import "github.com/spf13/cobra"

func GetOutput(cmd *cobra.Command) (string, error) {
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return "", err
	}

	return output, nil
}
