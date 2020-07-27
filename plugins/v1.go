package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "foo",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("curl http://127.0.0.1/foo\n")
	},
}

func Load(cmd *cobra.Command) {
	cmd.AddCommand(rootCmd)
}
