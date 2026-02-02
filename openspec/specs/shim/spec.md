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

The system SHALL include the image origin metadata as comments in the generated shim output when the image labels are present. The metadata SHALL use the same label keys as the image (`io.cli2docker.package`, `io.cli2docker.package-version`, `io.cli2docker.bin`, `io.cli2docker.build-timestamp`).

#### Scenario: Shim includes package and bin
- **WHEN** `cli2docker shim --image <image>` is run and the image labels include `io.cli2docker.package=eslint` and `io.cli2docker.bin=eslint`
- **THEN** the shim output includes comment lines with `io.cli2docker.package=eslint` and `io.cli2docker.bin=eslint`

#### Scenario: Shim includes explicit version
- **WHEN** `cli2docker shim --image <image>` is run and the image labels include `io.cli2docker.package=@acme/eslint`, `io.cli2docker.package-version=1.2.3`, and `io.cli2docker.bin=eslint-cli`
- **THEN** the shim output includes comment lines with `io.cli2docker.package=@acme/eslint`, `io.cli2docker.package-version=1.2.3`, and `io.cli2docker.bin=eslint-cli`

#### Scenario: Shim includes build timestamp
- **WHEN** `cli2docker shim --image <image>` is run and the image labels include `io.cli2docker.build-timestamp=2026-02-01T00:00:00Z`
- **THEN** the shim output includes a comment line with `io.cli2docker.build-timestamp=2026-02-01T00:00:00Z`

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

If `--mount-home` is set, the system SHALL mount a path from `$HOME` to the container user's home directory read-only by default. The container user's home directory SHALL be `/root` when the user label is `root`, otherwise `/home/<user>`. The user label SHALL be read from `io.cli2docker.user` on the image.

#### Scenario: Mount home read-only
- **WHEN** `--mount-home` is set and `--mount-home-rw` is not set and the image user is `bun`
- **THEN** the shim includes a `:ro` mount targeting `/home/bun/<relative>` for the resolved path

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

### Requirement: XDG mount flag
The system SHALL support `--mount-xdg` to mount XDG-style app directories under the container user's home directory. The container user's home directory SHALL be `/root` when the user label is `root`, otherwise `/home/<user>`. The user label SHALL be read from `io.cli2docker.user` on the image.

#### Scenario: Default mounts for all XDG dirs
- **WHEN** `cli2docker shim --image <image> --mount-xdg --mount-xdg-app eslint` is run and the image user is `bun`
- **THEN** the shim includes mounts for:
  - `$HOME/.config/eslint` to `/home/bun/.config/eslint`
  - `$HOME/.cache/eslint` to `/home/bun/.cache/eslint`
  - `$HOME/.local/share/eslint` to `/home/bun/.local/share/eslint`
  - `$HOME/.local/state/eslint` to `/home/bun/.local/state/eslint`

### Requirement: XDG app subpath defaulting
If `--mount-xdg` is set and `--mount-xdg-app` is omitted, the system SHALL derive the app subpath from image labels: prefer `io.cli2docker.bin`; otherwise use the last path segment of `io.cli2docker.package`.

#### Scenario: Default app from bin label
- **WHEN** `cli2docker shim --image <image> --mount-xdg` is run and the image labels include `io.cli2docker.bin=eslint`
- **THEN** the shim uses `eslint` as the app subpath for all XDG mounts

#### Scenario: Default app from package label
- **WHEN** `cli2docker shim --image <image> --mount-xdg` is run, `io.cli2docker.bin` is absent, and the image labels include `io.cli2docker.package=@acme/eslint`
- **THEN** the shim uses `eslint` as the app subpath for all XDG mounts

#### Scenario: Missing labels without explicit app
- **WHEN** `cli2docker shim --image <image> --mount-xdg` is run, `--mount-xdg-app` is omitted, and neither `io.cli2docker.bin` nor `io.cli2docker.package` is present
- **THEN** the command fails with an error indicating the XDG app name could not be derived

### Requirement: XDG app subpath override
If `--mount-xdg-app` is provided, the system SHALL use it as the app subpath for all XDG mounts, regardless of image labels.

#### Scenario: Explicit app subpath
- **WHEN** `cli2docker shim --image <image> --mount-xdg --mount-xdg-app mytool` is run and the image labels include `io.cli2docker.bin=eslint`
- **THEN** the shim uses `mytool` as the app subpath for all XDG mounts

### Requirement: XDG dirs selection
The system SHALL support `--mount-xdg-dirs` as a comma-separated list of `config`, `cache`, `data`, and `state`. If omitted, all four directories SHALL be mounted.

#### Scenario: Select subset of dirs
- **WHEN** `cli2docker shim --image <image> --mount-xdg --mount-xdg-app eslint --mount-xdg-dirs config,cache` is run
- **THEN** the shim includes mounts only for the config and cache XDG dirs

#### Scenario: Default to all dirs
- **WHEN** `cli2docker shim --image <image> --mount-xdg --mount-xdg-app eslint` is run and `--mount-xdg-dirs` is omitted
- **THEN** the shim includes mounts for config, cache, data, and state

#### Scenario: Invalid XDG dir name
- **WHEN** `cli2docker shim --image <image> --mount-xdg --mount-xdg-app eslint --mount-xdg-dirs logs` is run
- **THEN** the command fails with an error indicating an invalid XDG dir name

### Requirement: XDG base dir resolution
The system SHALL resolve XDG base dirs from environment variables when set; otherwise use defaults under `$HOME`.

#### Scenario: XDG override for config
- **WHEN** `XDG_CONFIG_HOME=$HOME/.config-alt` and `cli2docker shim --image <image> --mount-xdg --mount-xdg-app eslint --mount-xdg-dirs config` is run and the image user is `bun`
- **THEN** the shim mounts `$HOME/.config-alt/eslint` to `/home/bun/.config/eslint`

### Requirement: XDG base dir validation
The system SHALL reject any XDG base dir that resolves outside `$HOME`.

#### Scenario: XDG override outside home
- **WHEN** `XDG_CACHE_HOME=/tmp/cache` and `cli2docker shim --image <image> --mount-xdg --mount-xdg-app eslint --mount-xdg-dirs cache` is run
- **THEN** the command fails with an error indicating the XDG base dir must be within `$HOME`

### Requirement: User label required
The system SHALL require the image to include `io.cli2docker.user` and fail with an error if the label is missing.

#### Scenario: Missing user label
- **WHEN** `cli2docker shim --image <image> --mount-home` is run and the image lacks `io.cli2docker.user`
- **THEN** the command fails with an error indicating the missing user label

### Requirement: XDG read-write option
The system SHALL mount XDG directories read-only by default, and read-write when `--mount-xdg-rw` is set.

#### Scenario: Default read-only
- **WHEN** `cli2docker shim --image <image> --mount-xdg --mount-xdg-app eslint --mount-xdg-dirs config` is run without `--mount-xdg-rw`
- **THEN** the shim includes a `:ro` mount for the config dir

#### Scenario: Read-write opt-in
- **WHEN** `cli2docker shim --image <image> --mount-xdg --mount-xdg-app eslint --mount-xdg-dirs config --mount-xdg-rw` is run
- **THEN** the shim includes a `:rw` mount for the config dir

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
