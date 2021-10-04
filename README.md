# Cycloid CLI (cy)

This repository hosts the source code of Cycloid command line to use Cycloid APIs.

## Installation

Precompiled binaries for released versions are available in the [release](https://github.com/cycloidio/cycloid-cli/releases) page on Github. Using the latest production release binary is the recommended way of installing the Cycloid CLI.

## Getting started

Before playing with the CLI, you will first need to authenticate yourself into the Cycloid API using an API key as described in our [official documentation](http://docs.cycloid.io/cli.md).

```
cy login --org my-org --api-key "<MY_API_KEY>"
```

From there, you can now explore the various commands using the `--help` flag for each command / subcommand.

```shell
$ cy help
Cy is a CLI for Cycloid framework. Learn more at https://www.cycloid.io/.

Usage:
  cy [command]

Available Commands:
  api-key            Manage organization API keys
  catalog-repository Manage the catalog repositories
  completion         Output shell completion for the given shell (bash or zsh)
  config-repository  Manage the catalog repositories
  credential         Manage the credentials
  event              Manage the events
  external-backend   manage external backends
  help               Help about any command
  infrapolicy        Manage infrapolicies
  login              Login against the Cycloid console
  members            Manage members from the organization
  organization       Manage the organizations
  pipeline           Manage the pipelines
  project            Manage the projects
  roles              Manage roles from the organization
  stack              Manage the stacks
  status             Get the status of the Cycloid services
  terracost          Use terracost feature
  version            Get the version of the consumed API
```

## Common actions

### Get the Cycloid services in an unhealthy state

```
cy status -o json | jq '.[] | select( .status != "Success")'
```

### Create config repository using credential named "Git Config"

```
GIT_CRED=$(cy --org $ORG credential list -o json | jq '.[] | select( .name == ""Git Config") | .canonical')
cy --org myorg  config-repository create --branch master --cred $GIT_CRED --name "lab-config" --default --url "git@github.com:org/repo.git"
```

### Invite members with "Admin" role

```
ADMIN_ROLE=$(cy --org $ORG  roles list -o json | jq '.[] | select( .name | contains("Admin")) | .canonical')
cy  --org myorg  members invite --role $ADMIN_ROLE --email foo@email.com
```

:construction:
<!-- This is where we could add some useful examples: create a user, etc. -->

## Building from source

To build the CLI from source code, first ensure that have a working Go environment with version 1.13 or greater installed and the `make` command available.
After that, you can clone the repository yourself, build using `make build` and move the built binary where you want it to be (eg. in `/usr/local/bin` for a global install):

```shell
git clone git@github.com:cycloidio/cycloid-cli.git
cd cycloid-cli

make build
sudo mv cy /usr/local/bin
sudo chmod +x /usr/local/bin/cy
cy --version
```

## Contributing

Have a look to our [CONTRIBUTING.md](CONTRIBUTING.md)

## Testing for cycloid developers

Take a look at our [DEVELOPING_TIPS.md](DEVELOPING_TIPS.md). It Gives some insights on how to perform ci testing using a local be.

Note! It is meant for used of cycloid developers only, since it requires access to private cycloid repositories.

## License

This project is under MIT License, see [LICENSE](LICENSE)
