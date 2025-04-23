package middleware_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cycloidio/cycloid-cli/internal/test_utils"
)

/*
1. Start the backend / mysql / vault
2. Init the backend
3. Create first org / admin / apiKey
4. Do the testing
5. Drop the tables
6. Repeat from 3
*/

func runMain(ctx context.Context, m *testing.M) (int, error) {
	config := test_utils.NewDefaultBackendComposeConfig()
	compose, err := test_utils.NewBackendCompose(ctx, config)
	if err != nil {
		return 1, err
	}
	defer func() {
		err := compose.Cleanup(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	return m.Run(), nil
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	exitCode, err := runMain(ctx, m)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(exitCode)
}

func TestBackend(t *testing.T) {
	fmt.Println("cool")
}
