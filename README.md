# Cycloid CLI (cy)

This repository hosts the source code of Cycloid command line to use Cycloid APIs.

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

## Getting started

Before playing with the CLI, you first need to authenticate a user into the Cycloid API:

```
cy login --org my-org --email example@email.com --password my-password --api-url https://cycloid-api.local.tld
```

From there, you can now explore the various commands using the `--help` flag for each command / subcommand.

## Common actions

:construction:
<!-- This is where we could add some useful examples: create a user, etc. -->
