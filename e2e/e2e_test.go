package e2e_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/cycloidio/cycloid-cli/internal/testcfg"
)

var config *testcfg.Config

// Put any preparation code here so that defer() can work
func runMain(main *testing.M) (int, error) {
	// We must wait a bit that the middleware test are done initializing the config
	// Otherwise there will be conflcts -_-
	var err error
	time.Sleep(time.Duration(3) * time.Second)
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
