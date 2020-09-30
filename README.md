# Cycloid CLI (cy)

This repository hosts the source code of Cycloid command line to use Cycloid APIs.

## Installation

### From the sources

You can clone the repository and build from sources (we use `Go Modules` (1.13+)):

```shell
git clone git@github.com:cycloidio/youdeploy-cli.git
cd youdeploy-cli
```

To ease the build workflow, we use `docker` to build the CLI. Make sure that docker daemon is up and running.

```
make build
sudo mv cy /usr/local/bin
```

### From the releases page

:construction:


## Getting started

Before playing with the CLI, you first need to authenticate a user into the Cycloid API:

```
cy login --org my-org --email example@email.com --password my-password --api-url https://cycloid-api.local.tld
```

From there, you can now explore the various commands using the `--help` flag for each command / subcommand.

## Common actions

:construction:
<!-- This is where we could add some useful examples: create a user, etc. -->
