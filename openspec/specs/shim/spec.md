## Purpose

Defines normative requirements for `cli2docker shim`.

## Related docs

- [feature](feature.md)

## Requirements

### Requirement: Generate shim to stdout

The system SHALL print a shim script to stdout.

#### Scenario: Shim output
- **WHEN** the user runs `cli2docker shim --image <image>`
- **THEN** the command outputs a shell script to stdout

### Requirement: Shim includes origin metadata

The system SHALL include the image origin metadata as comments in the generated shim output when the image labels are present. The metadata SHALL use the same label keys as the image (`io.cli2docker.package`, `io.cli2docker.package-version`, `io.cli2docker.bin`).

#### Scenario: Shim includes package and bin
- **WHEN** `cli2docker shim --image <image>` is run and the image labels include `io.cli2docker.package=eslint` and `io.cli2docker.bin=eslint`
- **THEN** the shim output includes comment lines with `io.cli2docker.package=eslint` and `io.cli2docker.bin=eslint`

#### Scenario: Shim includes explicit version
- **WHEN** `cli2docker shim --image <image>` is run and the image labels include `io.cli2docker.package=@acme/eslint`, `io.cli2docker.package-version=1.2.3`, and `io.cli2docker.bin=eslint-cli`
- **THEN** the shim output includes comment lines with `io.cli2docker.package=@acme/eslint`, `io.cli2docker.package-version=1.2.3`, and `io.cli2docker.bin=eslint-cli`

### Requirement: Required image flag

The system SHALL require `--image` for `shim`.

#### Scenario: Missing image
- **WHEN** `--image` is omitted
- **THEN** the command fails with a required flag error

### Requirement: Docker required

The system SHALL require `docker` to be available in `PATH` before generating the shim.

#### Scenario: Docker missing
- **WHEN** the `docker` executable is not found in `PATH`
- **THEN** the command fails with an error indicating the missing required command

### Requirement: Mount current directory

If `--mount-cwd` is set, the system SHALL mount `$PWD` to `/work` and set working directory to `/work`.

#### Scenario: Mount cwd
- **WHEN** `--mount-cwd` is set
- **THEN** the shim includes `-v "${PWD}:/work"` and `-w /work`

### Requirement: Mount home path

If `--mount-home` is set, the system SHALL mount a path from `$HOME` to `/home/node/<relative>` read-only by default.

#### Scenario: Mount home read-only
- **WHEN** `--mount-home` is set and `--mount-home-rw` is not set
- **THEN** the shim includes a `:ro` mount for the resolved path

### Requirement: Mount home read-write

If `--mount-home-rw` is set, the system SHALL mount the `--mount-home` path read-write.

#### Scenario: Mount home read-write
- **WHEN** `--mount-home` is set and `--mount-home-rw` is set
- **THEN** the shim includes a `:rw` mount for the resolved path

### Requirement: Home path validation

The system SHALL reject `--mount-home` paths that are empty or outside `$HOME`.

#### Scenario: Empty mount-home
- **WHEN** `--mount-home` is empty or whitespace
- **THEN** the command fails with an error indicating an empty mount-home path

#### Scenario: Mount-home outside $HOME
- **WHEN** `--mount-home` resolves outside `$HOME`
- **THEN** the command fails with an error indicating the path must be within `$HOME`

### Requirement: Shim hardening defaults

The system SHALL include `--cap-drop=ALL`, `--security-opt=no-new-privileges`, and `--read-only` in the shim `docker run` command by default.

#### Scenario: Default hardening flags
- **WHEN** `cli2docker shim --image <image>` is run without opt-out flags
- **THEN** the shim includes `--cap-drop=ALL`, `--security-opt=no-new-privileges`, and `--read-only`

### Requirement: Opt-out flags for hardening

The system SHALL provide explicit opt-out flags that remove the default hardening flags from the shim.

#### Scenario: Disable cap drop
- **WHEN** `cli2docker shim --image <image> --no-drop-caps` is run
- **THEN** the shim does not include `--cap-drop=ALL`

#### Scenario: Disable no-new-privileges
- **WHEN** `cli2docker shim --image <image> --allow-new-privileges` is run
- **THEN** the shim does not include `--security-opt=no-new-privileges`

#### Scenario: Disable read-only rootfs
- **WHEN** `cli2docker shim --image <image> --no-read-only` is run
- **THEN** the shim does not include `--read-only`

### Requirement: Read-only warning

The system SHALL emit a warning to stderr stating that read-only mode is experimental and indicating the opt-out flag.

#### Scenario: Warning on default read-only
- **WHEN** `cli2docker shim --image <image>` is run
- **THEN** stderr contains a warning that read-only mode is experimental and can be disabled with `--no-read-only`
