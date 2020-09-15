package main

import (
	"fmt"
	"os"
	"plugin"
	"strconv"
	"strings"

	// Commented for now while figure out how do plugins
	// "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	models "github.com/cycloidio/youdeploy-cli/client/models"

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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	rootCmd.PersistentFlags().StringVarP(&userOutput, "output", "o", "table", "The formatting style for command output [json|yaml|table].")
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))

	rootCmd.PersistentFlags().StringP("verbosity", "v", "warning", "Override the default verbosity for this command. VERBOSITY must be one of: debug, info, warning, error, critical, none.")
	viper.BindPFlag("verbosity", rootCmd.PersistentFlags().Lookup("verbosity"))
	viper.SetDefault("verbosity", "warning")

	rootCmd.PersistentFlags().String("api-url", "", ".....")
	viper.BindPFlag("api-url", rootCmd.PersistentFlags().Lookup("api-url"))
	// rootCmd.PersistentFlags().Bool("version", false, "Display the version of this tool.")

	// rootCmd.PersistentFlags().BoolP("quiet", "q", false, "....") // also -q ?
	// viper.BindPFlag("quiet", rootCmd.PersistentFlags().Lookup("quiet"))

	// rootCmd.PersistentFlags().Bool("debug", false, "Turn on debug logging.")
	// viper.BindPFlag("useDebug", rootCmd.PersistentFlags().Lookup("debug"))

	// rootCmd.PersistentFlags().Bool("no-verify-ssl", false, ".....")
	// viper.BindPFlag("noVerifySSL", rootCmd.PersistentFlags().Lookup("noVerifySSL"))

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
