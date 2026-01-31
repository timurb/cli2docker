## ADDED Requirements

### Requirement: Print Dockerfile output
The system SHALL output the generated Dockerfile to stdout when `--print-dockerfile` is set and MUST NOT invoke `docker build`.

#### Scenario: Print-only mode
- **WHEN** `build` is executed with required inputs and `--print-dockerfile` is set
- **THEN** stdout contains the generated Dockerfile
- **AND** `docker build` is not executed

#### Scenario: Warnings go to stderr
- **WHEN** `--print-dockerfile` is set and derived defaults are used for `--image` or `--bin`
- **THEN** derivation warnings are written to stderr
- **AND** stdout contains only the Dockerfile content

## MODIFIED Requirements

### Requirement: Build command interface

The system SHALL accept and process the build interface flags: `--package`, `--bin`, `--image`, `--image-prefix`, `--tag`, `--base`, `--user`, `--no-user`, `--no-cache`, `--package-manager`, `--print-dockerfile`.

#### Scenario: Flags are provided
- **WHEN** the user provides any supported build flags
- **THEN** the command uses them to configure the build

### Requirement: Required docker executable

The system SHALL require `docker` to be available in `PATH` before running the build workflow, except when `--print-dockerfile` is set.

#### Scenario: Docker missing
- **WHEN** the `docker` executable is not found in `PATH` and `--print-dockerfile` is not set
- **THEN** the command fails with an error indicating the missing required command

#### Scenario: Print-only without docker
- **WHEN** the `docker` executable is not found in `PATH` and `--print-dockerfile` is set
- **THEN** the command proceeds to print the Dockerfile

### Requirement: Build workflow

The system SHALL build a Docker image for an npm or Bun CLI tool using a generated Dockerfile, unless `--print-dockerfile` is set.

#### Scenario: Build from npm package
- **WHEN** `build` is executed with required inputs, `--package-manager` is omitted, and `--print-dockerfile` is not set
- **THEN** the system generates a Dockerfile that installs the npm package
- **AND** the system runs `docker build` to create the image

#### Scenario: Build from Bun package
- **WHEN** `build` is executed with required inputs, `--package-manager` is `bun`, and `--print-dockerfile` is not set
- **THEN** the system generates a Dockerfile that installs the package using Bun
- **AND** the system runs `docker build` to create the image

#### Scenario: Print-only build
- **WHEN** `build` is executed with required inputs and `--print-dockerfile` is set
- **THEN** the system generates a Dockerfile and writes it to stdout
- **AND** the system does not run `docker build`

### Requirement: Build constraints

The system SHALL require a running Docker daemon and a published npm package providing the specified binary when `--print-dockerfile` is not set.

#### Scenario: Docker daemon unavailable
- **WHEN** the Docker daemon is not available and `--print-dockerfile` is not set
- **THEN** the build fails

#### Scenario: Print-only without daemon
- **WHEN** `--print-dockerfile` is set and the Docker daemon is not available
- **THEN** the Dockerfile is printed and the build does not fail due to the daemon
