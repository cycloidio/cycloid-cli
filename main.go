package main

import (
	"fmt"
	"os"
	"plugin"
	"strconv"

	// Commented for now while figure out how do plugins
	// "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}
var version = 1

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// Commented for now while figure out how do plugins
	// if err := cmd.Execute(); err != nil {
	// 	log.Printf("%+v\n", err)
	// 	os.Exit(1)
	// }

	ver, err := strconv.Atoi(os.Getenv("V"))

	if err == nil {
		version = ver
	}

	fmt.Printf("Version %d\n", version)

	p, err := plugin.Open(fmt.Sprintf("plugins/v%d.so", version))
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("Load")
	if err != nil {
		panic(err)
	}
	// *v.(*int) = 7
	f.(func(*cobra.Command))(rootCmd) // prints "Hello, number 7"

	Execute()

}
