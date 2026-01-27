## Overview (Black Box)

`cli2docker build` builds a Docker image for an npm CLI package using a generated Dockerfile.

## Goals / Non-Goals

**Goals:**
- Build an image from an npm package and CLI entrypoint.
- Allow basic configuration of base image, user, tag, and cache.

**Non-Goals:**
- Multi-arch builds.
- Registry publishing/signing.
- Arbitrary Dockerfile generation.

## Entities

- **Image**: Docker image name + tag.
- **Package**: npm package name.
- **Binary**: CLI entrypoint name.

## Interfaces

- Command: `cli2docker build`
- Flags: `--package`, `--bin`, `--image`, `--tag`, `--base`, `--user`, `--no-user`, `--no-cache`

## Events / Triggers

- User runs `cli2docker build ...`

## Behavior (White Box)

- Generates a Dockerfile and runs `docker build`.
- Default tag is `latest` if omitted.
- `--no-user` omits `USER` from Dockerfile; `--no-cache` passes `--no-cache` to Docker.

## Requirements (Test Cases)

### Requirement: Default image from package

If `--image` is omitted, the system SHALL derive the image name from `--package`.

#### Scenario: Plain package name
- **WHEN** `--package` is `eslint` and `--image` is omitted
- **THEN** image name is `eslint`
- **AND** a warning is emitted to stderr

#### Scenario: Scoped package name
- **WHEN** `--package` is `@acme/eslint` and `--image` is omitted
- **THEN** image name is `acme/eslint`
- **AND** a warning is emitted to stderr

### Requirement: Default bin from package

If `--bin` is omitted, the system SHALL derive the binary name from `--package` by using the package name without scope.

#### Scenario: Plain package name
- **WHEN** `--package` is `eslint` and `--bin` is omitted
- **THEN** binary name is `eslint`
- **AND** a warning is emitted to stderr

#### Scenario: Scoped package name
- **WHEN** `--package` is `@acme/eslint` and `--bin` is omitted
- **THEN** binary name is `eslint`
- **AND** a warning is emitted to stderr

### Requirement: Explicit values override defaults

If `--image` or `--bin` is provided explicitly, the system SHALL use the explicit value.

#### Scenario: Explicit image and bin
- **WHEN** `--package` is `eslint` and `--image` is `acme/eslint` and `--bin` is `eslint-cli`
- **THEN** image name is `acme/eslint`
- **AND** binary name is `eslint-cli`

### Requirement: Required docker executable

The system SHALL require `docker` to be available in `PATH` before running the build workflow.

#### Scenario: Docker missing
- **WHEN** the `docker` executable is not found in `PATH`
- **THEN** the command fails with an error indicating the missing required command

### Requirement: Build command interface

The system SHALL accept and process the build interface flags: `--package`, `--bin`, `--image`, `--tag`, `--base`, `--user`, `--no-user`, `--no-cache`.

#### Scenario: Flags are provided
- **WHEN** the user provides any supported build flags
- **THEN** the command uses them to configure the build

### Requirement: Required inputs

The system SHALL require `--package` to be provided; `--bin` and `--image` are optional with defaults.

#### Scenario: Missing package
- **WHEN** `--package` is missing
- **THEN** the command fails with a required flag error

### Requirement: Build workflow

The system SHALL build a Docker image for an npm CLI tool using a generated Dockerfile.

#### Scenario: Build from npm package
- **WHEN** `build` is executed with required inputs
- **THEN** the system generates a Dockerfile that installs the npm package
- **AND** the system runs `docker build` to create the image

### Requirement: Dockerfile content

The generated Dockerfile SHALL include: `FROM <base>`, `npm install -g <package>`, `ENTRYPOINT ["<bin>"]`, and `USER <user>` unless `--no-user` is set.

#### Scenario: Default user
- **WHEN** `--no-user` is not set
- **THEN** the Dockerfile includes `USER <user>`

#### Scenario: No user
- **WHEN** `--no-user` is set
- **THEN** the Dockerfile does not include a `USER` instruction

### Requirement: Build cache control

If `--no-cache` is set, the system SHALL disable the Docker build cache.

#### Scenario: No cache
- **WHEN** `--no-cache` is set
- **THEN** `docker build` is executed with `--no-cache`

### Requirement: Tag defaulting

If `--tag` is omitted, the system SHALL use `latest`.

#### Scenario: Missing tag
- **WHEN** `--tag` is omitted
- **THEN** the image tag is `latest`

### Requirement: Build constraints

The system SHALL require a running Docker daemon and a published npm package providing the specified binary.

#### Scenario: Docker daemon unavailable
- **WHEN** the Docker daemon is not available
- **THEN** the build fails

### Requirement: Build errors surface

The system SHALL surface errors from `docker build` and invalid flags.

#### Scenario: Docker build error
- **WHEN** `docker build` fails
- **THEN** the error is returned to the user
