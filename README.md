# Cycloid CLI (cy)

This repository hosts the source code of Cycloid command line to use Cycloid APIs.

## Getting started

Before playing with the CLI, you first need to authenticate a user into the Cycloid API:

```
cy login --org my-org --email example@email.com --password my-password --api-url https://cycloid-api.local.tld
```

From there, you can now explore the various commands using the `--help` flag for each command / subcommand.

```shell
$ cy help
Cy is a CLI for Cycloid framework. Learn more at https://www.cycloid.io/.

Usage:
  cy [command]

Available Commands:
  catalog-repository Manage the catalog repositories
  config-repository  Manage the catalog repositories
  credential         Manage the credentials
  event              Manage the events
  external-backend   manage external backends
  help               Help about any command
  login              Login against the Cycloid console
  members            Manage members from the organization
  organization       Manage the organizations
  pipeline           Manage the pipelines
  project            Manage the projects
  stack              Manage the stacks
  version            Get the version of the consumed API
```

## Installation

### From the sources

You can clone the repository and build from sources (we use `Go Modules` (1.13+)):

```shell
git clone git@github.com:cycloidio/cycloid-cli.git
cd cycloid-cli
```

Build and install the binary

```
make build
sudo mv cy /usr/local/bin
cy --version
```

### From the releases page

You can download the latest Linux binary from the [release](https://github.com/cycloidio/cycloid-cli/releases) page.

## Common actions

:construction:
<!-- This is where we could add some useful examples: create a user, etc. -->

## Contributing

Have a look to our [CONTRIBUTING.md](CONTRIBUTING.md)

## License

This project is under MIT License, see [LICENSE](LICENSE)
