package middleware_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cycloidio/cycloid-cli/internal/test_utils"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
)

var (
	backendImageRegistry = "661913936052.dkr.ecr.eu-west-1.amazonaws.com/youdeploy-http-api"
	backendImageVersion  = "staging"
	mysqlVersion         = "8.4.3"
	mysqlConfig          = `
[client]
default-character-set = utf8mb4

[mysql]
default-character-set = utf8mb4

[mysqld]
collation_server = utf8mb4_unicode_ci
character_set_server = utf8mb4
`
	redisVersion     = "7.4.1-alpine"
	postgresVersion  = "15.4"
	postgresUser     = "concourse"
	postgresPassword = "concourse"
	postgresDatabase = "concourse"
	vaultRootToken   = "rootToken"
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
	net, err := network.New(ctx)
	if err != nil {
		return 1, err
	}

	defer func() {
		if err := net.Remove(ctx); err != nil {
			log.Printf("failed to remove network: %s", err)
		}
	}()

	if err != nil {
		log.Println(err)
		return 1, err
	}

	redisContainer, err := test_utils.GetRedis(ctx, redisVersion)
	defer func() {
		if err := testcontainers.TerminateContainer(redisContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		return 1, err
	}

	mysqlContainer, err := test_utils.GetMysql(ctx, mysqlVersion, mysqlConfig)
	defer func() {
		if err := testcontainers.TerminateContainer(mysqlContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		return 1, err
	}

	vaultContainer, err := test_utils.GetVault(ctx, "1.18.1", vaultRootToken)
	defer func() {
		if err := testcontainers.TerminateContainer(vaultContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		return 1, err
	}

	postgresContainer, err := test_utils.GetPostgres(ctx, postgresVersion, "", postgresDatabase, postgresUser, postgresPassword)
	defer func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		return 1, err
	}

	vaultHost, err := vaultContainer.Host(ctx)
	if err != nil {
		fmt.Println("failed to get vault host address", err)
	}

	postgresHost, err := postgresContainer.Host(ctx)
	if err != nil {
		fmt.Println("failed to get postgres host address", err)
	}

	mysqlHost, err := mysqlContainer.Host(ctx)
	if err != nil {
		fmt.Println("failed to get mysql host address", err)
	}

	concourseContainer, err := test_utils.GetConcourse(ctx, "", "concourse", "concourse", postgresDatabase, postgresUser, postgresPassword, postgresHost, vaultHost, vaultRootToken, "/cycloid")
	defer func() {
		if err := testcontainers.TerminateContainer(*concourseContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Println(err)
		return 1, err
	}

	concourseHost, err := concourseContainer.Host(ctx)
	if err != nil {
		log.Println(err)
		return 1, err
	}

	redisUrl, err := redisContainer.ConnectionString(ctx)
	if err != nil {
		log.Println(err)
		return 1, err
	}

	vaultUrl, err := vaultContainer.HttpHostAddress(ctx)
	if err != nil {
		log.Println(err)
		return 1, err
	}

	backendContainer, err := test_utils.GetBackend(ctx, backendImageRegistry, backendImageVersion, mysqlHost, "3306", "cycloid", "cycloid", "cycloid", concourseHost, "8080", "concourse", "concourse", "main", redisUrl, vaultUrl, "cycloid", "cycloid")
	defer func() {
		if err := testcontainers.TerminateContainer(*backendContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Println(err)
		return 1, err
	}

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
