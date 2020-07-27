package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var blar string
var rootCmd = &cobra.Command{
	Use:   "foo",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("curl http://127.0.0.1/foo blar=%s\n", blar)
	},
}

func Load(cmd *cobra.Command) {
	rootCmd.Flags().StringVar(&blar, "blar", "", "")
	cmd.AddCommand(rootCmd)
}
