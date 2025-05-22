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
	// Initialize global vars
	var err error
	config, err = testcfg.NewConfig()
	defer config.Cleanup()
	if err != nil {
		return 1, fmt.Errorf("test config setup failed: %v", err)
	}

	return main.Run(), nil
}

func TestMain(main *testing.M) {
	code, err := runMain(main)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}
