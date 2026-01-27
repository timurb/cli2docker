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
