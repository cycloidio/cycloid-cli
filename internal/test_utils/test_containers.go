package test_utils

import (
	"context"
	"fmt"
	"log"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/modules/redis"
	"github.com/testcontainers/testcontainers-go/modules/vault"
	"github.com/testcontainers/testcontainers-go/network"
)

type BackendComposeConfig struct {
	RedisVersion         string
	MysqlVersion         string
	MysqlConfig          string
	BackendImageRegistry string
	BackendImageVersion  string
	PostgresVersion      string
	PostgresUser         string
	PostgresPassword     string
	PostgresDatabase     string
	VaultRootToken       string
	VaultVersion         string
	ConcourseVersion     string
}

func NewDefaultBackendComposeConfig() BackendComposeConfig {
	return BackendComposeConfig{
		RedisVersion: "7.4.1-alpine",
		MysqlVersion: "8.4.3",
		MysqlConfig: `
[client]
default-character-set = utf8mb4

[mysql]
default-character-set = utf8mb4

[mysqld]
collation_server = utf8mb4_unicode_ci
character_set_server = utf8mb4
		`,
		BackendImageRegistry: "661913936052.dkr.ecr.eu-west-1.amazonaws.com/youdeploy-http-api",
		BackendImageVersion:  "staging",
		PostgresVersion:      "15.4",
		PostgresUser:         "concourse",
		PostgresPassword:     "concourse",
		PostgresDatabase:     "concourse",
		VaultRootToken:       "cycloid",
		VaultVersion:         "1.18.1",
		ConcourseVersion:     "7.9.1",
	}
}

type BackendCompose struct {
	Network   *testcontainers.DockerNetwork
	Redis     *redis.RedisContainer
	Postgres  *postgres.PostgresContainer
	Mysql     *mysql.MySQLContainer
	Vault     *vault.VaultContainer
	Concourse testcontainers.Container
	Backend   testcontainers.Container
}

func (c *BackendCompose) Cleanup(ctx context.Context) error {
	if c.Network != nil {
		if err := c.Network.Remove(ctx); err != nil {
			return err
		}
	}

	if c.Redis != nil {
		if err := c.Redis.Terminate(ctx); err != nil {
			return err
		}
	}

	if c.Postgres != nil {
		if err := c.Postgres.Terminate(ctx); err != nil {
			return err
		}
	}

	if c.Mysql != nil {
		if err := c.Mysql.Terminate(ctx); err != nil {
			return err
		}
	}

	if c.Vault != nil {
		if err := c.Vault.Terminate(ctx); err != nil {
			return err
		}
	}

	if c.Concourse != nil {
		if err := c.Concourse.Terminate(ctx); err != nil {
			return err
		}
	}

	if c.Backend != nil {
		if err := c.Backend.Terminate(ctx); err != nil {
			return err
		}
	}

	return nil
}

func NewBackendCompose(ctx context.Context, config BackendComposeConfig) (*BackendCompose, error) {
	var compose = BackendCompose{}

	net, err := network.New(ctx)
	if err != nil {
		return &compose, err
	}
	compose.Network = net

	redisContainer, err := GetRedis(ctx, config.RedisVersion)
	if err != nil {
		return &compose, err
	}
	compose.Redis = redisContainer

	mysqlContainer, err := GetMysql(ctx, config.MysqlVersion, config.MysqlConfig)
	if err != nil {
		return &compose, err
	}
	compose.Mysql = mysqlContainer

	vaultContainer, err := GetVault(ctx, config.VaultVersion, config.VaultRootToken)
	if err != nil {
		return &compose, err
	}
	compose.Vault = vaultContainer

	postgresContainer, err := GetPostgres(ctx, config.PostgresVersion, "", config.PostgresDatabase, config.PostgresUser, config.PostgresPassword)
	if err != nil {
		return &compose, err
	}
	compose.Postgres = postgresContainer

	vaultHost, err := vaultContainer.Host(ctx)
	if err != nil {
		return &compose, fmt.Errorf("failed to get vault host address: %v", err)
	}

	postgresHost, err := postgresContainer.Host(ctx)
	if err != nil {
		return &compose, fmt.Errorf("failed to get postgres host address: %v", err)
	}

	mysqlIP, err := mysqlContainer.ContainerIP(ctx)
	if err != nil {
		return &compose, fmt.Errorf("failed to get mysql ip address: %v", err)
	}

	// Only fill the IP, the backend will use the port
	mysqlHost := mysqlIP

	concourseContainer, err := GetConcourse(ctx, config.ConcourseVersion, "concourse", "concourse", config.PostgresDatabase, config.PostgresUser, config.PostgresPassword, postgresHost, vaultHost, config.VaultRootToken, "/cycloid")
	if err != nil {
		return &compose, err
	}
	compose.Concourse = concourseContainer

	concourseIP, err := concourseContainer.ContainerIP(ctx)
	if err != nil {
		log.Println(err)
		return &compose, err
	}

	concourseHost := "http://" + concourseIP + ":8000"

	redisIP, err := redisContainer.ContainerIP(ctx)
	if err != nil {
		log.Println(err)
		return &compose, err
	}
	redisUrl := fmt.Sprintf("redis://%s:6379", redisIP)

	vaultHost, err = vaultContainer.ContainerIP(ctx)
	if err != nil {
		log.Println(err)
		return &compose, err
	}
	vaultUrl := fmt.Sprintf("http://%s:%s", vaultHost, "8200")

	backendContainer, err := GetBackend(ctx, config.BackendImageRegistry, config.BackendImageVersion, mysqlHost, "3306", "cycloid", "cycloid", "cycloid", concourseHost, "8080", "concourse", "concourse", "main", redisUrl, vaultUrl, "cycloid", "cycloid")
	if err != nil {
		return &compose, err
	}
	compose.Backend = *backendContainer

	return &compose, nil
}
