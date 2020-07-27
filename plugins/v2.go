package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var bar string
var rootCmd = &cobra.Command{
	Use:   "foo",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("curl http://127.0.0.1/foo bar=%s\n", bar)
	},
}

func Load(cmd *cobra.Command) {
	rootCmd.Flags().StringVar(&bar, "bar", "", "")
	cmd.AddCommand(rootCmd)
}
