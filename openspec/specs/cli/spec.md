## Overview (Black Box)

`cli2docker` is a CLI wrapper with `build` and `shim` subcommands and standard help/usage.

## Goals / Non-Goals

**Goals:**
- Provide a root command and subcommands `build` and `shim`.
- Parse flags and render help/usage consistently.

**Non-Goals:**
- Implement build/shim behavior (covered by their specs).

## Entities

- **CLI**: root command and subcommands.
- **Flags**: command-line options for subcommands.

## Interfaces

- Commands: `cli2docker`, `cli2docker build`, `cli2docker shim`
- Help: `--help`

## Events / Triggers

- User invokes any CLI command.

## Behavior (White Box)

- Uses `spf13/cobra` for flag parsing and help/usage rendering.
- Invalid flags result in non-zero exit and stderr message.

## Requirements (Test Cases)

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
