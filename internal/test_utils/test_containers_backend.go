package test_utils

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	backendBaseConfig = `---
start-timeout: 10s
max-header-size: "1MB"
port: 3001
host: 0.0.0.0
schema:
- http
cleanup-timeout: 1ms

# This defines if the logger have to be run in dev-mode or not.
# The log-level will define the overall log being displayed.
#
# Possible levels are the custom "NONE" to disable logging, and:
# https://godoc.org/go.uber.org/zap/zapcore#Level
#
# If a wrong loglevel is given, "INFO" will be taken.
#
# For more information about the different mode, please visit:
# https://godoc.org/go.uber.org/zap#NewDevelopmentConfig
# https://godoc.org/go.uber.org/zap#NewProductionConfig
log-dev-mode: true
log-level: "INFO"

db-host: 172.42.0.2
db-port: 3306
db-user: root
db-pwd: youdeploy
db-name: youdeploy_public_test
db-max-conns: 15
db-max-idle-conns: 10
db-max-lifetime-conn: 5m

concourse-url: http://172.42.0.5
concourse-port: 8080
concourse-username: yd-concourse
concourse-team: main
concourse-password: Oogiemaep6iebeibue5viucheePaeX7Y

vault-role-id: custom-role-id
vault-secret-id: custom-secret-id
vault-url: http://172.42.0.7:8200

frontend-base-url: "http://localhost:3000"
backend-base-url: "http://localhost:3001"

redis-uri: redis://172.42.0.10:6379

email-smtp-svr-addr: 172.42.0.8:1025
email-smtp-username: admin
email-smtp-password: admin
email-addr-from: "Cycloid Platform <noreply@cycloid.io>"
email-addr-return-path: "admin+ydbounce@cycloid.io"
email-dev-mode: true
crypto-signing-key: totally-random-secret-key
jwt-keys:
  - 2f2122de-63f2-4eec-9c6f-c6abb3e1f007:7cdyHps2tYDp6e7VKPEstE5sDMQbK6WLyN3GmTsF7x7QpE6ZP5ra6yfVSkvXakbB
local-auth-enabled: true
azure-tenant-id: b677a6b8-f2e7-4551-8849-f45dc7b730de
azure-client-id: b4471929-8cd4-4bc9-b534-b8315bc43d99
saml-sp-certificate-path: services/authentication/saml/testdata/certificate.pem
saml-sp-private-key-path: services/authentication/saml/testdata/private-key.pem
# This file (docker/saml-idp.xml) is served from https://samltest.id/saml/idp
# as alternative we can set the saml-idp-metadata-url parmeter to that URL,
# but the server is down at the time or writing this comment.
saml-idp-metadata-path: docker/saml-idp.xml
google-client-id: 741192805913-s10ibou8065iofnb9rcir9269skiqts9.apps.googleusercontent.com
github-client-id: 6a94210b44f4a612952e
github-client-secret: 8dfd349e8f1260f3f1a3f6ebc7862b59fca36690
contact-us-form-url: https://www.cycloid.io/contact-us
tell-us-why-licence-form-url: https://www.cycloid.io/contact-us

cost-explorer-es-url: http://172.42.0.11:9200
cost-explorer-es-max-bulk-bytes: 5000000
cost-explorer-es-bulk-increase-factor: 1.5
cost-explorer-es-bulk-decrease-factor: 0.7
cost-explorer-es-retry-period-seconds: 3

worker-queues: [emails, hubspot, cost_explorer, checks, terracost]
worker-run-internal: true
worker-run-scheduler: true

# CRM
hubspot-private-app-token: "pat-eu1-d697999c-cd30-4bc8-a9fb-81699eeeaeba"

# Sentry
# sentry-dsn set to an empty string disables sentry
sentry-dsn: ""
sentry-env: "dev"
sentry-enable-tracing: true
sentry-sample-rate: 1.0 # Range [0.0, 1.0]. 1.0 captures 100% of transactions for tracing. Lower this value in production.

telemetry-flush-timeout: 2s

# Jaeger
# jaeger-endpoint set to an empty string disables jaeger
jaeger-endpoint: "http://172.42.0.19:4318/v1/traces"
`
)

func GetBackend(ctx context.Context, registry, version, dbHost, dbPort, dbName, dbUser, dbPassword, ccHost, ccPort, ccUser, ccPassword, ccTeam, redisUrl, vaultUrl, vaultRoleId, vaultSecretId string) (*testcontainers.Container, error) {
	backendEnv := map[string]string{
		"DB_HOST":            dbHost,
		"DB_NAME":            dbName,
		"DB_USER":            dbUser,
		"DB_PWD":             dbPassword,
		"CONCOURSE_URL":      ccHost,
		"CONCOURSE_PORT":     ccPort,
		"CONCOURSE_USERNAME": ccUser,
		"CONCOURSE_PASSWORD": ccPassword,
		"CONCOURSE_TEAM":     ccTeam,
	}
	backendContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        fmt.Sprintf("%s:%s", registry, version),
			ExposedPorts: []string{"3001"},
			// Networks:     []string{"cli-tests"},
			Env: backendEnv,
			Files: []testcontainers.ContainerFile{
				{
					Reader:            strings.NewReader(backendBaseConfig),
					ContainerFilePath: "/ci/config.yml",
					FileMode:          0440,
				},
			},
			WaitingFor: wait.
				ForHTTP("/status/database").
				WithPollInterval(2 * time.Second),
			AlwaysPullImage: true,
			WorkingDir:      "/go/src/github.com/cycloidio/youdeploy-http-api",
			Entrypoint: []string{
				"bash", "-ec", `
echo -e "# \e[33mDB migrate ...\e[0m"
timeout 120 bash -c '
  until /go/youdeploy-http-api migrate up --config-file /ci/config.yml --migrations-dir /opt/migrations --db-name=cycloid && echo "ok"; do
    >&2 echo -e "Waiting for DB migrations"
    sleep 1
  done
'
echo -e "Running Cycloid API"
exec /go/youdeploy-http-api server --config-file /ci/config.yml
	`, // TODO: fix entrypoint
			},
		},
		Started: true,
	})
	return &backendContainer, err
}
