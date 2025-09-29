package middleware_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cycloidio/cycloid-cli/internal/testcfg"
)

var config *testcfg.Config

// Put any preparation code here so that defer() can work
func runMain(ctx context.Context, main *testing.M) (int, error) {
	_ = ctx
	// Initialize global vars
	var err error
	config, err = testcfg.NewConfig("middleware")
	defer config.Cleanup()
	if err != nil {
		return 1, fmt.Errorf("Config setup failed for package middleware: %v", err)
	}

	os.Setenv("CY_API_URL", config.APIUrl)
	os.Setenv("CY_API_KEY", config.APIKey)
	os.Setenv("CY_ORG", config.Org)
	return main.Run(), nil
}

func TestMain(main *testing.M) {
	ctx := context.Background()
	code, err := runMain(ctx, main)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}
