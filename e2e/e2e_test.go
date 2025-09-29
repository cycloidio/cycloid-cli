package e2e_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cycloidio/cycloid-cli/internal/testcfg"
)

var config *testcfg.Config

// Put any preparation code here so that defer() can work
func runMain(main *testing.M) (int, error) {
	// We must wait a bit that the middleware test are done initializing the config
	// Otherwise there will be conflcts -_-
	var err error
	config, err = testcfg.NewConfig("e2e")
	defer config.Cleanup()
	if err != nil {
		return 1, fmt.Errorf("test config setup failed for e2e tests: %w", err)
	}

	os.Setenv("CY_API_URL", config.APIUrl)
	os.Setenv("CY_API_KEY", config.APIKey)
	os.Setenv("CY_ORG", config.Org)

	return main.Run(), nil
}

func TestMain(main *testing.M) {
	code, err := runMain(main)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}
