## ADDED Requirements

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

## MODIFIED Requirements

### Requirement: Mount home path
If `--mount-home` is set, the system SHALL mount a path from `$HOME` to the container user's home directory read-only by default. The container user's home directory SHALL be `/root` when the user label is `root`, otherwise `/home/<user>`. The user label SHALL be read from `io.cli2docker.user` on the image.

#### Scenario: Mount home read-only
- **WHEN** `--mount-home` is set and `--mount-home-rw` is not set and the image user is `bun`
- **THEN** the shim includes a `:ro` mount targeting `/home/bun/<relative>` for the resolved path
