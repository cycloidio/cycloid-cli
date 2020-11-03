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

# Running E2E tests

E2E tests tend to be idempotent as possible but we recommend to use it again a dedicated Cycloid server. Before running the test, you can specify a few environment variables:

  * CY_API_URL: the URL of the Cycloid API server
  * CY_TEST_EMAIL: the test email
  * CY_TEST_PASSWORD: the test password
  * CY_TEST_ORG: the test organization
