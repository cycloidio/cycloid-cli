package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var blurp string
var bli string

var rootCmd = &cobra.Command{
	Use:   "foop",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("curl http://127.0.0.1/foop blurp=%s bli=%s\n", blurp, bli)
	},
}

func Load(cmd *cobra.Command) {
	rootCmd.Flags().StringVar(&blurp, "blurp", "", "")
	rootCmd.Flags().StringVar(&bli, "bli", "", "")
	cmd.AddCommand(rootCmd)
}
