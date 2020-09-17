package main

import (
	"fmt"
	"os"
	"plugin"
	"strings"

	models "github.com/cycloidio/youdeploy-cli/client/models"
	"github.com/cycloidio/youdeploy-cli/lookup"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/runtime"
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
		spew.Dump(err)
		apiErr, ok := err.(*runtime.APIError)
		if ok {
			fmt.Println("runtime.APIError")
			spew.Dump(apiErr.Error())
			r := apiErr.Response.(runtime.ClientResponse)
			spew.Dump(r.Message())
			fmt.Println("---debug---")

			spew.Dump(apiErr.Response)

			// fmt.Println(err.GetPayload())
			//
			// unexpectedSuccess := result.(*GetCredentialsDefault)
		}
		// unexpectedSuccess := result.(*GetCredentialsDefault)

		// _ = organization_credentials.GetCredentialsDefault{}

		// Create a generic interface to get errors payload from struct like organization_credentials.GetCredentialsDefault
		type errorP interface {
			GetPayload() *models.ErrorPayload
		}

		var errorPayload errorP
		errorPayload, ok = err.(errorP)

		// _ = models.ErrorPayload{}
		// errorPayload, ok := err.(*organization_credentials.GetCredentialsDefault)
		// errorPayload, ok := err.(*models.ErrorPayload)
		if ok {
			fmt.Println("models.ErrorPayload")
			// spew.Dump(errorPayload.GetPayload().Errors)

			for _, e := range errorPayload.GetPayload().Errors {
				fmt.Println(*e.Message)
				fmt.Println(*e.Code)
				for _, d := range e.Details {
					fmt.Println(d)
				}
			}
		}
		spew.Dump(err.Error())
		fmt.Printf("%+v\n", err.Error())
		// fmt.Println(err)

		// fmt.Println(err.Type())

		// s := reflect.ValueOf(&err.Elem()
		// typeOfT := s.Type()
		//
		// for i := 0; i < s.NumField(); i++ {
		// 	f := s.Field(i)
		// 	fmt.Printf("%d: %s %s = %v\n", i,
		// 		typeOfT.Field(i).Name, f.Type(), f.Interface())
		// }

		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(initConfig)
	viper.SetEnvPrefix("CY")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	rootCmd.PersistentFlags().StringVarP(&userOutput, "output", "o", "table", "The formatting style for command output [json|yaml|table].")
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))

	rootCmd.PersistentFlags().StringP("verbosity", "v", "warning", "Override the default verbosity for this command. VERBOSITY must be one of: debug, info, warning, error, critical, none.")
	viper.BindPFlag("verbosity", rootCmd.PersistentFlags().Lookup("verbosity"))
	viper.SetDefault("verbosity", "warning")

	rootCmd.PersistentFlags().String("api-url", "", ".....")
	viper.BindPFlag("api-url", rootCmd.PersistentFlags().Lookup("api-url"))

	rootCmd.PersistentFlags().String("cy-plugin-dir", "/tmp/cy-plugins", "directory where the CLI plugins are stored")
	viper.BindPFlag("cy-plugin-dir", rootCmd.PersistentFlags().Lookup("cy-plugin-dir"))

}

func main() {

	var err error

	v, err := lookup.GetAPIVersion()
	if err != nil {
		panic(err)
	}

	version := v.Version
	// Plugin not found locally, lookup for the plugin
	pluginPath, err := lookup.LookupPlugin(version)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Running plugin version %s\n", version)
	p, err := plugin.Open(pluginPath)
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("AttachCommands")
	if err != nil {
		panic(err)
	}

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
