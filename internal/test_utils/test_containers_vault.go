package test_utils

import (
	"context"
	"strings"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/vault"
)

func GetVault(ctx context.Context, version, rootToken string) (*vault.VaultContainer, error) {
	vaultConfig := `
path "cycloid/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}

path "sys/policy/cycloid/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}

path "auth/approle/role/cycloid-*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}

path "auth/approle/role/" {
  capabilities = ["read", "list"]
}

path "policies" {
  capabilities = ["read", "list"]
}

path "auth/token/create" {
  capabilities = ["create"]
}

path "auth/token/renew-self" {
  capabilities = ["create"]
}
`
	vaultContainer, err := vault.Run(ctx,
		"hashicorp/vault:"+version,
		vault.WithToken(rootToken),
		vault.WithInitCommand(
			"auth enable approle",
			"write sys/mounts/cycloid type=kv",
			"write sys/policy/cycloid policy=@policy.hcl",
			"write auth/approle/role/cycloid token_ttl=20m token_max_ttl=1h policies=cycloid",
			"write auth/approle/role/cycloid/role-id role_id=cycloid",
			"write auth/approle/role/cycloid/custom-secret-id secret_id=cycloid",
		),
		testcontainers.CustomizeRequest(testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Networks: []string{"cli-tests"},
				Files: []testcontainers.ContainerFile{
					{
						Reader:            strings.NewReader(vaultConfig),
						ContainerFilePath: "/policy.hcl",
						FileMode:          0644,
					},
				},
			},
		}),
	)
	return vaultContainer, err
}
