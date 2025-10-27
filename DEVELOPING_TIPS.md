# Cycloid Developer Tips

> [!NOTE] These are internal notes meant for cycloid develovers only since they
> require access to private git repositories.

This file gives some tips to how to test, change or upgrade cli, for cycloid developers.

The cli pipeline is available in the [cycloid-stacks](https://github.com/cycloidio/cycloid-stacks/tree/stacks/cycloid-cli) and the prod and staging cli can be checked [here](https://console.cycloid.io/organizations/cycloid/projects/cycloid-cli).

## Prepare a CLI update

Right now, the CLI release process and version is tied to the backend release. 
This will be changed once the CLI codebase will be merged in the backend, until then,
we will have to make do.

So the release process is as follows:
- The backend is merged and release in staging/production
- An automatic PR is made with the client update on [this component](https://console.cycloid.io/organizations/cycloid/projects/cycloid-cli/environments/automatic_bump/components/apibump/overview)
- The automatic PR will contain the client update and write the version in the 
  [`./client/version`](./client/version) file.

> [!WARNING]
> The version tagging will occur when the PR with the updated client version with
> the updated [`./client/version`](./client/version) file will be merged on develop
>
> This means that every feature PR merged after this one will not be integrated
> in the release

- Once the Client update PR merged, the version is tagged the pipeline on 
  [this component must go to completion on the release job](https://console.cycloid.io/organizations/cycloid/projects/cycloid-cli/environments/staging/components/develop/overview)
- To release in prod, use the [`merge-develop-to-master`](https://console.cycloid.io/organizations/cycloid/projects/cycloid-cli/environments/prod/components/master/pipelines/cycloid-cli-prod-master/jobs/merge-develop-to-master/builds/465851136#all) button on the prod pipeline.
- Then the release process will test, build and release on github.

So we advise the following workflow:

1. Until a client version update occurs, you can merge any feature PR on develop
2. Once the client PR update is created, and if you need to develop fixes and logic
   linked to that client update, create a branch from the client update PR.
   (basically treating the client update PR as the new develop)
3. Once you are done with all the changes linked to this version and the test passes,
   you can merge the client update PR and proceed to release in production.

If you need to create the PR update yourself manually to anticipate a version not release yet,
you can generate the client update using the `make client-generate` target.

### Make a manual client update

- Download the correct `swagger.yml` and put it at the repository's root at `./swagger.yml`
- Use `make client-generate`
- Wait for the backend release and add the correct version in `./client/version`

Ping devops if you have any questions or doubt on the release process.

## Add a new commands to the cli

To add a new command you can start by verifying the api endpoint to implement in the [API docs](https://docs.cycloid.io/api/index.html).

While you can follow this guide, I advise you to look up other implemented routes.

### Implement the middleware method

1. Ensure your client is up to date, if the backend version is not release yet,
   see [the above section](#make-a-manual-client-update)

2. Define the method that implements the endpoint in the middleware

    1. Create the method in the middleware interface [here](./cmd/cycloid/middleware/middleware.go)
       The naming convention of the function should reflect the naming on the API.
       If you have a getComponent route on the API, the middleware should be a GetComponent function.

> [!NOTE]
> Only exception, if a GET route list resources, we rename it to ListResource
> so the `getComponents` should be named `ListComponents`.
  
    2. The function should accept all required parameters as argument in function.
        a. Optional argument should be a pointer, a nil value means absent value.
        b. Required values must be plain values
        c. If a request has a body, use `body.Validate(strfmt.Default)`

    3. Implement the API call either:
        a. By using the generated client parameters, look any other implemented
           like [getProject](./cmd/cycloid/middleware/organization_projects.go)
        b. If the route can't be implemented using the client, use `net/http`.
    
    4. All route should return either some data and error, or just an error:
        a. Delete methods -> `return error`
        b. Create/Get methods -> `return (*models.object, error)`
        c. List methods -> `return ([]*models.object, error)`

    5. Do not validate payload, it has generated too much issues in the past.

#### Middleware testing

Add in the middleware package a basic happy path test to ensure
that given correct parameters, the API responds.

This is meant mostly to detect any API dift or issues. This way it will enable
developper to differenciate more easily server and client issues.

Testing setup is done on the [TestMain function](./cmd/cycloid/middleware/middleware_test.go)
that will use the [testcfg package](./internal/testcfg/config.go) to initialize
the backend, provision the required tools and give you all the required context
for testing (api key, org, api url, basic stuff like project/env/component, config
and catalog repo, etc...).

Lookup how existing test are made and follow the same logic.

### Implement the actual CLI command

Now you need to define the cobra command at `cmd/cycloid/<feature>/` so that
the method previously defined can be used in the cli.

   Each feature is structured as follows:

```tree
./cmd/cycloid/<featurename>/
./cmd/cycloid/<featurename>/cmd.go # list and combine all the subcommands
./cmd/cycloid/<featurename>/get.go # each command has its own file
./cmd/cycloid/<featurename>/create.go
./cmd/cycloid/<featurename>/update.go
./cmd/cycloid/<featurename>/list.go
./cmd/cycloid/<featurename>/common.go # add common logic specific to this series of commands
```
   A feature can have subcommands in that case structure it like this:
  
```tree
./cmd/cycloid/<featurename>/
./cmd/cycloid/<featurename>/cmd.go # list and combine all the subcommands
./cmd/cycloid/<featurename>/get.go # each command has its own file
./cmd/cycloid/<featurename>/create.go
./cmd/cycloid/<featurename>/update.go
./cmd/cycloid/<featurename>/list.go
./cmd/cycloid/<featurename>/sub_command.go # add your subcommand here, it must behave like the cmd.go file
./cmd/cycloid/<featurename>/sub_command_get.go # prefix all sub sub_command with the sub_command name
./cmd/cycloid/<featurename>/sub_command_create.go 
./cmd/cycloid/<featurename>/common.go # add common logic specific to this series of commands
```

Command convention is as follow:

`cy <feature> <verb> <args_and_flags....>`

Example for components:
```bash
cy component create
cy component update
cy component get
cy component list
```

You should start by adding the type of command to the set of list of
available commands in `cmd.go`

Then create a new file on this folder, where you will specify the method
that will return the cobra command that defines the cli command to implement
with the flags possible to use.

Finally you define a method that will be run by this cobra command that will
take as argument the multiple flags and pass them to the middleware
interface method that you previsouly created.

#### Implementation directives

Use and implemant **all flags, arguments and completion definitions and functions**
in the [cyargs package](./internal/cyargs), each flag must be declared and its value
must be retrieved using a function in the cyargs package.

Only if a flag is declared on only one specific command you can declare them locally.

Look up how flags and completion are implemented for example [the component flag](./internal/cyargs/component.go).

The command should:
1. Retrieve all flag
2. Process any pre-requisite (file read, default values management, etc...)
3. Do the action against the API.
4. Use the [printer](./printer/printer.go) to print the output of the API to the console
   a. If no error, with valid payload => output on stdout using `cmd.OutOrStdout()`
   b. If any error, print error and details on stderr using `cmd.OutOrStderr()`

This about the UX for the command line:
- Add the completion
- Try to infer values as much as you can (use env vars, context, fetching existing values)
- Make action idempotent as much as possible
- Implement upsert (like `cy component create --update`)

### Testing commands

To test commands, do it in the [e2e test packages](./e2e/e2e_test.go). It has
all the required function an context for testing the command directly.

It also uses the [testcfg package](./internal/testcfg/config.go) to initialize
the backend, provision the required tools and give you all the required context
for testing (api key, org, api url, basic stuff like project/env/component, config
and catalog repo, etc...).

Add tests for your command, what you should be testing:
- All basic happy path
- Test that all arguments, flags and env var works to configure your action
- Test that the output match expectations
- Try to test and catch edge cases

As usual, look up other test commands for examples.

Tests in CI are not executed in parallel due to some backend issues with concurrency
around git.

### Add a changelog

Once your changes are done, add a changelog to it:

```
make new-changelog-entry
```

## CLI local testing

### Requirements

To perform local test you need:
- Access to the Cycloid console on `cycloid` via the API
- A valid API key on `CY_SAAS_API_KEY`

To start the test, login to docker to pull the backend image:

```bash
make docker-login docker-pull
```

Start the local backend:

```bash
make be-start
```

You can start the tests using `go test ./...` or

```bash
make test
```

You can cleanup the backend with:

```bash
make be-stop
```

And reset it (stop + start):
```bash
make be-reset
```
