## Overview

`cli2docker` is a CLI wrapper with `build` and `shim` subcommands and standard help/usage.

## Interfaces

- Commands: `cli2docker`, `cli2docker build`, `cli2docker shim`
- Help: `--help`

## Invariants

- Invalid flags result in non-zero exit and stderr message.

## Usage workflows

- Show help for root or subcommands.
- Run a subcommand with its flags.

## Key constraints

- Command structure is fixed to root + `build` + `shim`.

## Related docs

- [spec](spec.md)
- [overview](../overview.md)
