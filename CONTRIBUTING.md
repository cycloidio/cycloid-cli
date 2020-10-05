# If needed, update swagger or re-generate clients

```
sudo make generate-client
```

# Run fake api server

```
docker run -it -p 80:5000 gaell/simple-docker-app
```

# Build the cli

```
make build
```

# Give a try

```
./cy login --org my-org --email example@email.com --password my-password --api-url https://cycloid-api.local.tld
./cy external-backends list
```
