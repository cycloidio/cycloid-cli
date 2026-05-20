package testcfg

// CLI bootstrap provisioning for be-reset / TestMain.
//
// Validated manually and via e2e (PR #441): `cy beta bootstrap-first-org` works
// against staging. testcfg still calls middleware.InitFirstOrg directly because
// in-process CLI bootstrap left viper/CY_API_KEY state that broke some middleware
// tests (403 Need to refresh token). Re-enable by swapping the provision block
// in config.go when we want testcfg to exercise the cobra path again.
//
// Uncomment bootstrapFirstOrgCLI and use it from NewConfig provisionAPI block:
//
//	init, err := bootstrapFirstOrgCLI(bootstrapFirstOrgParams{
//		APIURL:          config.APIUrl,
//		Org:             config.Org,
//		Username:        username,
//		FullName:        fullName,
//		Email:           email,
//		Password:        password,
//		Licence:         licence,
//		APIKeyCanonical: apiKeyCanonical,
//	})
//	if err != nil {
//		return nil, fmt.Errorf("failed to init console: %w", err)
//	}
//	config.APIKey = *init.APIKey
//	api.Config.Token = *init.APIKey
//
// /*
//
//	import (
//		"bytes"
//		"encoding/json"
//		"fmt"
//
//		"github.com/cycloidio/cycloid-cli/cmd"
//		"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
//	)
//
//	type bootstrapFirstOrgParams struct {
//		APIURL          string
//		Org             string
//		Username        string
//		FullName        string
//		Email           string
//		Password        string
//		Licence         string
//		APIKeyCanonical string
//	}
//
//	func bootstrapFirstOrgCLI(params bootstrapFirstOrgParams) (*middleware.FirstOrgData, error) {
//		rootCmd := cmd.NewRootCommand()
//
//		stdout := new(bytes.Buffer)
//		stderr := new(bytes.Buffer)
//		rootCmd.SetOut(stdout)
//		rootCmd.SetErr(stderr)
//		rootCmd.SetArgs([]string{
//			"--api-url", params.APIURL,
//			"--output", "json",
//			"--org", params.Org,
//			"beta", "bootstrap-first-org",
//			"--username", params.Username,
//			"--full-name", params.FullName,
//			"--email", params.Email,
//			"--password", params.Password,
//			"--licence", params.Licence,
//			"--api-key-canonical", params.APIKeyCanonical,
//		})
//
//		if err := rootCmd.Execute(); err != nil {
//			return nil, fmt.Errorf("bootstrap-first-org CLI failed: %w (stderr: %s)", err, stderr.String())
//		}
//
//		var result middleware.FirstOrgData
//		if err := json.Unmarshal(stdout.Bytes(), &result); err != nil {
//			return nil, fmt.Errorf("failed to decode bootstrap-first-org JSON output: %w (stdout: %s)", err, stdout.String())
//		}
//		if result.APIKey == nil || *result.APIKey == "" {
//			return nil, fmt.Errorf("bootstrap-first-org succeeded but APIKey is missing in output")
//		}
//
//		return &result, nil
//	}
//
// */
