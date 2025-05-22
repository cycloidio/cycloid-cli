package e2e

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
	config, err = testcfg.NewConfig()
	defer config.Cleanup()
	if err != nil {
		return 1, fmt.Errorf("Config setup failed: %v", err)
	}

	log.Printf("Starting tests with config:\nurl: %s\norg: %s", config.APIUrl, config.Org)
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
