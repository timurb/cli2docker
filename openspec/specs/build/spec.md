## Purpose

Defines normative requirements for `cli2docker build`.

## Related docs

- [feature](feature.md)

## Requirements

### Requirement: Default image from package

If `--image` is omitted, the system SHALL derive the image name from `--package` after removing any version/tag suffix.

#### Scenario: Plain package name with version tag
- **WHEN** `--package` is `eslint@latest` and `--image` is omitted
- **THEN** image name is `cli/eslint`
- **AND** a warning is emitted to stderr

#### Scenario: Scoped package name with version
- **WHEN** `--package` is `@acme/eslint@1.2.3` and `--image` is omitted
- **THEN** image name is `cli/acme/eslint`
- **AND** a warning is emitted to stderr

### Requirement: Default bin from package

If `--bin` is omitted, the system SHALL derive the binary name from `--package` by removing any version/tag suffix and then using the package name without scope.

#### Scenario: Plain package name with version
- **WHEN** `--package` is `eslint@latest` and `--bin` is omitted
- **THEN** binary name is `eslint`
- **AND** a warning is emitted to stderr

#### Scenario: Scoped package name with version
- **WHEN** `--package` is `@acme/eslint@1.2.3` and `--bin` is omitted
- **THEN** binary name is `eslint`
- **AND** a warning is emitted to stderr

### Requirement: Explicit values override defaults

If `--image` or `--bin` is provided explicitly, the system SHALL use the explicit value.

#### Scenario: Explicit image and bin
- **WHEN** `--package` is `eslint` and `--image` is `acme/eslint` and `--bin` is `eslint-cli`
- **THEN** image name is `acme/eslint`
- **AND** binary name is `eslint-cli`

#### Scenario: Explicit image overrides prefix
- **WHEN** `--image` is `acme/eslint` and `--image-prefix` is set
- **THEN** image name is `acme/eslint`
- **AND** a warning is emitted that the prefix was ignored

### Requirement: Image prefix flag

The system SHALL accept `--image-prefix` to prefix the derived image name when `--image` is omitted.

#### Scenario: Prefix applied to derived image
- **WHEN** `--package` is `eslint`, `--image` is omitted, and `--image-prefix` is `cli/`
- **THEN** image name is `cli/eslint`

#### Scenario: Prefix applied to scoped derived image
- **WHEN** `--package` is `@acme/eslint`, `--image` is omitted, and `--image-prefix` is `cli/`
- **THEN** image name is `cli/acme/eslint`

#### Scenario: Default prefix
- **WHEN** `--package` is `eslint`, `--image` is omitted, and `--image-prefix` is not set
- **THEN** image name is `cli/eslint`

### Requirement: Required docker executable

The system SHALL require `docker` to be available in `PATH` before running the build workflow, except when `--print-dockerfile` is set.

#### Scenario: Docker missing
- **WHEN** the `docker` executable is not found in `PATH` and `--print-dockerfile` is not set
- **THEN** the command fails with an error indicating the missing required command

#### Scenario: Print-only without docker
- **WHEN** the `docker` executable is not found in `PATH` and `--print-dockerfile` is set
- **THEN** the command proceeds to print the Dockerfile

### Requirement: Build command interface

The system SHALL accept and process the build interface flags: `--package`, `--bin`, `--image`, `--image-prefix`, `--tag`, `--base`, `--user`, `--no-user`, `--no-cache`, `--package-manager`, `--print-dockerfile`.

#### Scenario: Flags are provided
- **WHEN** the user provides any supported build flags
- **THEN** the command uses them to configure the build

### Requirement: Package manager selection

The system SHALL accept `--package-manager` with values `npm` and `bun`. If omitted, the system SHALL default to `npm`.

#### Scenario: Default package manager
- **WHEN** `--package-manager` is omitted
- **THEN** the system uses the npm build flow

#### Scenario: Bun selected
- **WHEN** `--package-manager` is `bun`
- **THEN** the system uses the Bun build flow

#### Scenario: Invalid package manager
- **WHEN** `--package-manager` is `pnpm`
- **THEN** the command fails with an error indicating an invalid package manager

### Requirement: Package manager defaults for base image and user

The system SHALL default `--base` and `--user` based on the selected package manager when those flags are omitted.

#### Scenario: npm defaults
- **WHEN** `--package-manager` is omitted and `--base` and `--user` are omitted
- **THEN** the base image is `node:20-alpine`
- **AND** the runtime user is `node`

#### Scenario: Bun defaults
- **WHEN** `--package-manager` is `bun` and `--base` and `--user` are omitted
- **THEN** the base image is `oven/bun:1`
- **AND** the runtime user is `bun`

#### Scenario: Explicit overrides
- **WHEN** `--package-manager` is `bun`, `--base` is `acme/bun`, and `--user` is `root`
- **THEN** the base image is `acme/bun`
- **AND** the runtime user is `root`

### Requirement: Required inputs

The system SHALL require `--package` to be provided; `--bin` and `--image` are optional with defaults.

#### Scenario: Missing package
- **WHEN** `--package` is missing
- **THEN** the command fails with a required flag error

### Requirement: Build workflow

The system SHALL build a Docker image for an npm or Bun CLI tool using a generated Dockerfile, unless `--print-dockerfile` is set.

#### Scenario: Build from npm package
- **WHEN** `build` is executed with required inputs and `--package-manager` is omitted
- **THEN** the system generates a Dockerfile that installs the npm package
- **AND** the system runs `docker build` to create the image

#### Scenario: Build from Bun package
- **WHEN** `build` is executed with required inputs and `--package-manager` is `bun`
- **THEN** the system generates a Dockerfile that installs the package using Bun
- **AND** the system runs `docker build` to create the image

#### Scenario: Print-only build
- **WHEN** `build` is executed with required inputs and `--print-dockerfile` is set
- **THEN** the system generates a Dockerfile and writes it to stdout
- **AND** the system does not run `docker build`

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

### Requirement: Dockerfile content

The generated Dockerfile SHALL include: `FROM <base>`, an install command based on the selected package manager, `ENTRYPOINT ["<bin>"]`, and `USER <user>` unless `--no-user` is set.

#### Scenario: Npm install
- **WHEN** `--package-manager` is omitted
- **THEN** the Dockerfile includes `npm install -g <package>`

#### Scenario: Bun install
- **WHEN** `--package-manager` is `bun`
- **THEN** the Dockerfile includes `bun add -g <package>`

#### Scenario: Default user
- **WHEN** `--no-user` is not set
- **THEN** the Dockerfile includes `USER <user>`

#### Scenario: No user
- **WHEN** `--no-user` is set
- **THEN** the Dockerfile does not include a `USER` instruction

### Requirement: Image origin labels

The system SHALL set image labels on the built image for the originating package and bin. The system SHALL set `io.cli2docker.package` and `io.cli2docker.bin` for every build, and SHALL set `io.cli2docker.package-version` only when an explicit version is present in `--package`. The system SHALL set `io.cli2docker.build-timestamp` for every build to a single RFC3339 UTC timestamp captured at the start of `cli2docker build`.

#### Scenario: Package without explicit version
- **WHEN** `--package` is `eslint` and `cli2docker build` completes
- **THEN** the image labels include `io.cli2docker.package=eslint` and `io.cli2docker.bin=eslint`
- **AND** the image label `io.cli2docker.package-version` is not set
- **AND** the image label `io.cli2docker.build-timestamp` is set to an RFC3339 UTC timestamp

#### Scenario: Package with explicit version
- **WHEN** `--package` is `@acme/eslint@1.2.3`, `--bin` is `eslint-cli`, and `cli2docker build` completes
- **THEN** the image labels include `io.cli2docker.package=@acme/eslint`, `io.cli2docker.package-version=1.2.3`, and `io.cli2docker.bin=eslint-cli`
- **AND** the image label `io.cli2docker.build-timestamp` is set to an RFC3339 UTC timestamp

#### Scenario: Dockerfile contains origin labels
- **WHEN** `cli2docker build` generates a Dockerfile for `--package` `eslint` with derived `--bin`
- **THEN** the Dockerfile includes a `LABEL` instruction containing `io.cli2docker.package=eslint` and `io.cli2docker.bin=eslint`
- **AND** the Dockerfile does not include `io.cli2docker.package-version`
- **AND** the Dockerfile includes `io.cli2docker.build-timestamp` set to an RFC3339 UTC timestamp

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

The system SHALL require a running Docker daemon and a published npm package providing the specified binary when `--print-dockerfile` is not set.

#### Scenario: Docker daemon unavailable
- **WHEN** the Docker daemon is not available and `--print-dockerfile` is not set
- **THEN** the build fails

#### Scenario: Print-only without daemon
- **WHEN** `--print-dockerfile` is set and the Docker daemon is not available
- **THEN** the Dockerfile is printed and the build does not fail due to the daemon

### Requirement: Build errors surface

The system SHALL surface errors from `docker build` and invalid flags.

#### Scenario: Docker build error
- **WHEN** `docker build` fails
- **THEN** the error is returned to the user
