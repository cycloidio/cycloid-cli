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

# Add changelog entry

Currently we use `changie` to manage the changelog. Each PR merge will result in a file at changelog/unreleased/ with a certain format.

To simplify the usage of `changie` there's a yaml file, which describes the different file's format, possibilities, etc. To create a new entry you just need to use the following makefile target:

```
make new-changelog-entry
```

For more details on `changie` at https://changie.dev/guide/quick-start/