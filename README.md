# youdeploy-cli
Source code of Cycloid command line to use yd API


# Generate client from swagger

```
rm -rf ./clients/api-v1/ ; mkdir ./clients/api-v1
docker-compose run swagger generate client --spec=swagger-files/api-v1.yml --default-produces="application/vnd.cycloid.io.v1+json" --target=./clients/api-v1 --name=api-v1 --tags=Cycloid --tags="Organization External Backends"
```

# Generate plugins

```
cd plugins
for i in $(seq 1 1); do go build -buildmode=plugin -o "gen/v${i}.so" "v${i}.go"; done
```

# Run fake api server

```
docker run -it -p 80:5000 gaell/simple-docker-app
```

# Build and run cli using plugin version 1

```
make build &&  V=1 ./cy external-backends list
```

