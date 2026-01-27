## Purpose

Defines the CLI surface for `cli2docker`, including the root command and primary subcommands.

## Requirements

### Requirement: Root command

The system SHALL provide a root CLI command named `cli2docker`.

#### Scenario: Help for root command
- **WHEN** the user runs `cli2docker --help`
- **THEN** the CLI prints usage/help for the root command

### Requirement: Subcommands

The system SHALL provide `build` and `shim` subcommands.

#### Scenario: Build help
- **WHEN** the user runs `cli2docker build --help`
- **THEN** the CLI prints usage/help for the build command

#### Scenario: Shim help
- **WHEN** the user runs `cli2docker shim --help`
- **THEN** the CLI prints usage/help for the shim command

### Requirement: Flag parsing and usage

The system SHALL parse flags and render usage using `spf13/cobra`.

#### Scenario: Invalid flags
- **WHEN** a user passes an unknown flag
- **THEN** the CLI exits non-zero and prints an error to stderr
