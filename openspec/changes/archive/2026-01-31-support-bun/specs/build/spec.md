## ADDED Requirements

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

## MODIFIED Requirements

### Requirement: Build command interface
The system SHALL accept and process the build interface flags: `--package`, `--bin`, `--image`, `--image-prefix`, `--tag`, `--base`, `--user`, `--no-user`, `--no-cache`, `--package-manager`.

#### Scenario: Flags are provided
- **WHEN** the user provides any supported build flags
- **THEN** the command uses them to configure the build

### Requirement: Build workflow
The system SHALL build a Docker image for an npm or Bun CLI tool using a generated Dockerfile.

#### Scenario: Build from npm package
- **WHEN** `build` is executed with required inputs and `--package-manager` is omitted
- **THEN** the system generates a Dockerfile that installs the npm package
- **AND** the system runs `docker build` to create the image

#### Scenario: Build from Bun package
- **WHEN** `build` is executed with required inputs and `--package-manager` is `bun`
- **THEN** the system generates a Dockerfile that installs the package using Bun
- **AND** the system runs `docker build` to create the image

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
