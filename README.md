# youdeploy-cli
Source code of Cycloid command line to use yd API


# Generate client from swagger

```
for version in $(seq 1 3); do
rm -rf ./clients/api-v${version}/ ; mkdir ./clients/api-v${version}
docker-compose run swagger generate client --spec=swagger-files/api-v${version}.yml --default-produces="application/vnd.cycloid.io.v1+json" --target=./clients/api-v${version} --name=api-v${version} --tags=Cycloid --tags="Organization External Backends"
done
```

# Generate plugins

```
cd plugins
for i in $(seq 1 3); do go build -buildmode=plugin -o "gen/v${i}.so" "v${i}.go"; done
```

# Run fake api server

```
docker run -it -p 80:5000 gaell/simple-docker-app
```

# Build and run cli using plugin version 1

```
make build &&  V=1 ./cy external-backends list
```

# Usecases

V=1 mean we are connecting to Cycloid api v1. So we load v1 plugin
This should be done later by a curl on /version I think

## Usecase 1 url typo change

```
v1 '/organizations/{organization_canonical}/external_backend'
    flag pproject
>
v2 '/organizations/{organization_canonical}/external_backends'
    flag pproject
>
v3 '/organizations/{organization_canonical}/external_backends'
    flag project
```

Expected
```
make build &&  V=1 ./cy external-backends list --pproject foo
# > "GET /organizations/cycloid/external_backend?environment=website&pproject=prod

V=2 ./cy external-backends list --pproject foo
# > "GET /organizations/cycloid/external_backends?environment=website&pproject=prod

V=3 ./cy external-backends list --pproject foo
V=3 ./cy external-backends list --project foo
# > "GET /organizations/cycloid/external_backends?environment=website&project=prod
```
