package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewLoginCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "login",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// It should also take the api url
// Or gettocken
// /user/login
// post: login
// Authenticate a user and return a new JWT token.
