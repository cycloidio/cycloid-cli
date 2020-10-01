package main

import (
	"fmt"
	"os"
	"strings"

	models "github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd"
	"github.com/cycloidio/cycloid-cli/internal/version"

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

	versionString = fmt.Sprintf("%s, revision %s, branch %s, date %s; go %s", version.Version, version.Revision, version.Branch, version.BuildDate, version.GoVersion)

	rootCmd = &cobra.Command{
		Version: versionString,
		Use:     "cy",
		Short:   "Cycloid CLI",
		Long:    `Cy is a CLI for Cycloid framework. Learn more at https://www.cycloid.io/.`,
	}
)

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

	rootCmd.PersistentFlags().String("api-url", "https://http-api.cycloid.io", "Specify the HTTP url of Cycloid API to use eg https://http-api.cycloid.io. This can also be given by CY_API_URL environment variable.")
	viper.BindPFlag("api-url", rootCmd.PersistentFlags().Lookup("api-url"))
}

func main() {
	cmd.AttachCommands(rootCmd)

	Execute()
}
