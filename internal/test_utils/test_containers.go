package test_utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/modules/redis"
	"github.com/testcontainers/testcontainers-go/wait"
)

func GetRedis(ctx context.Context, version string) (*redis.RedisContainer, error) {
	return redis.Run(ctx, "redis:"+version)
}

func GetMysql(ctx context.Context, version, config string) (*mysql.MySQLContainer, error) {
	configFile, err := os.CreateTemp("", "my-*.cnf")
	if err != nil {
		return nil, fmt.Errorf("failed to create mysql conf file: %s", err)
	}

	_, err = configFile.WriteString(config)
	if err != nil {
		return nil, fmt.Errorf("failed to write config for mysql: %s", err)
	}

	mysqlC, err := mysql.Run(ctx, fmt.Sprintf("mysql:%s", version),
		mysql.WithConfigFile(configFile.Name()),
		mysql.WithDatabase("cycloid"),
		mysql.WithUsername("cycloid"),
		mysql.WithPassword("cycloid"),
	)
	if err != nil {
		return nil, err
	}

	defer os.Remove(configFile.Name())

	return mysqlC, err
}

func GetPostgres(ctx context.Context, version, config, database, user, password string) (*postgres.PostgresContainer, error) {
	return postgres.Run(ctx,
		"postgres:"+version,
		postgres.WithUsername(user),
		postgres.WithPassword(password),
		postgres.WithDatabase(database),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
}
