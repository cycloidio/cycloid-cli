package main

import (
	"fmt"
	"os"
	"plugin"
	"strconv"

	// Commented for now while figure out how do plugins
	// "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile    string
	userOutput string
	verbosity  string

	rootCmd = &cobra.Command{
		Use:   "cy",
		Short: "Cycloid CLI",
		Long:  `Cy is a CLI for Cycloid framework. Learn more at https://www.cycloid.io/.`,
	}
)

var version = 1

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(initConfig)
	viper.SetEnvPrefix("CY")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	rootCmd.PersistentFlags().StringVarP(&userOutput, "output", "o", "", "The formatting style for command output [json|yaml|text|table].")
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	viper.SetDefault("output", "text")

	// rootCmd.PersistentFlags().StringVarP(&verbosity, "verbosity", "v", "", "Override the default verbosity for this command. VERBOSITY must be one of: debug, info, warning, error, critical, none.")
	rootCmd.PersistentFlags().StringP("verbosity", "v", "warning", "Override the default verbosity for this command. VERBOSITY must be one of: debug, info, warning, error, critical, none.")
	viper.BindPFlag("verbosity", rootCmd.PersistentFlags().Lookup("verbosity"))
	viper.SetDefault("verbosity", "warning")

	rootCmd.PersistentFlags().Bool("version", false, "Display the version of this tool.")
	// viper.BindPFlag("version", rootCmd.PersistentFlags().Lookup("version"))

	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "....") // also -q ?
	viper.BindPFlag("quiet", rootCmd.PersistentFlags().Lookup("quiet"))

	rootCmd.PersistentFlags().Bool("debug", false, "Turn on debug logging.")
	viper.BindPFlag("useDebug", rootCmd.PersistentFlags().Lookup("debug"))

	rootCmd.PersistentFlags().Bool("no-verify-ssl", false, ".....")
	viper.BindPFlag("noVerifySSL", rootCmd.PersistentFlags().Lookup("noVerifySSL"))

	// --log-http
	//    Log all HTTP server requests and responses to stderr. Overrides the
	//    default core/log_http property value for this command invocation.

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

	f, err := p.Lookup("AttachCommands")
	if err != nil {
		panic(err)
	}
	// *v.(*int) = 7
	// f.(func(*cobra.Command))(rootCmd) // prints "Hello, number 7"
	// rootCmd = f.(func() *cobra.Command)()
	f.(func(*cobra.Command))(rootCmd)

	Execute()

}

// func er(msg interface{}) {
// 	fmt.Println("Error:", msg)
// 	os.Exit(1)
// }

// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := homedir.Dir()
// 		if err != nil {
// 			er(err)
// 		}
//
// 		// Search config in home directory with name ".cobra" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".cobra")
// 	}
//
// 	viper.AutomaticEnv()
//
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
