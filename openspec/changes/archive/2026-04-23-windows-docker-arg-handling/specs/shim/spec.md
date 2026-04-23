## ADDED Requirements

### Requirement: Windows POSIX shell mount compatibility

The system SHALL generate mount-related `docker run` arguments that remain valid when the POSIX shim is executed from supported Windows POSIX environments. Supported Windows POSIX environments SHALL be `git-bash` and `WSL`.

#### Scenario: Git Bash preserves container-side paths
- **WHEN** a shim generated with any of `--mount-cwd`, `--mount-home`, or `--mount-xdg` is executed from `git-bash`
- **THEN** the Docker invocation preserves container-side paths such as `/work` and `/home/<user>/...` unchanged
- **AND** host-side bind mount source paths are passed in Docker-compatible Windows form

#### Scenario: WSL keeps Linux-style host paths
- **WHEN** a shim generated with any of `--mount-cwd`, `--mount-home`, or `--mount-xdg` is executed from `WSL`
- **THEN** the Docker invocation uses the Linux-style host paths visible inside `WSL`
- **AND** container-side paths such as `/work` and `/home/<user>/...` remain unchanged

## MODIFIED Requirements

### Requirement: Mount current directory

If `--mount-cwd` is set, the system SHALL mount the current working directory to `/work` as a bind mount and set working directory to `/work`. When the shim is executed from a supported Windows POSIX environment, the host current working directory SHALL be serialized in a Docker-compatible form for that environment.

#### Scenario: Mount cwd
- **WHEN** `--mount-cwd` is set
- **THEN** the shim includes a bind mount from the current working directory to `/work`
- **AND** the shim includes `-w /work`

#### Scenario: Mount cwd from git-bash
- **WHEN** `--mount-cwd` is set and the shim is executed from `git-bash` on Windows
- **THEN** the shim converts the runtime current working directory to a Docker-compatible Windows path and mounts it to `/work`
- **AND** the shim includes `-w /work`

### Requirement: Mount home path

If `--mount-home` is set, the system SHALL mount a path from `$HOME` to the container user's home directory read-only by default as a bind mount. The container user's home directory SHALL be `/root` when the user label is `root`, otherwise `/home/<user>`. The user label SHALL be read from `io.cli2docker.user` on the image. When the shim is generated on Windows, the resolved host path SHALL be rendered in Docker-compatible Windows form.

#### Scenario: Mount home read-only
- **WHEN** `--mount-home` is set and `--mount-home-rw` is not set and the image user is `bun`
- **THEN** the shim includes a read-only bind mount targeting `/home/bun/<relative>` for the resolved path

#### Scenario: Mount home read-only on Windows
- **WHEN** the generator host home directory is `C:\\Users\\alice`, `--mount-home .config` is set, and the image user is `bun`
- **THEN** the shim includes a read-only bind mount from `C:/Users/alice/.config` to `/home/bun/.config`

### Requirement: Mount home read-write

If `--mount-home-rw` is set, the system SHALL mount the `--mount-home` path read-write as a bind mount.

#### Scenario: Mount home read-write
- **WHEN** `--mount-home` is set and `--mount-home-rw` is set
- **THEN** the shim includes a read-write bind mount for the resolved path

### Requirement: XDG mount flag

The system SHALL support `--mount-xdg` to mount XDG-style app directories under the container user's home directory. The container user's home directory SHALL be `/root` when the user label is `root`, otherwise `/home/<user>`. The user label SHALL be read from `io.cli2docker.user` on the image. Each XDG directory SHALL be rendered as a bind mount, and when the shim is generated on Windows each resolved host path SHALL be rendered in Docker-compatible Windows form.

#### Scenario: Default mounts for all XDG dirs
- **WHEN** `cli2docker shim --image <image> --mount-xdg --mount-xdg-app eslint` is run and the image user is `bun`
- **THEN** the shim includes mounts for:
  - `$HOME/.config/eslint` to `/home/bun/.config/eslint`
  - `$HOME/.cache/eslint` to `/home/bun/.cache/eslint`
  - `$HOME/.local/share/eslint` to `/home/bun/.local/share/eslint`
  - `$HOME/.local/state/eslint` to `/home/bun/.local/state/eslint`

#### Scenario: Windows XDG host paths are Docker-compatible
- **WHEN** the generator host home directory is `C:\\Users\\alice` and `cli2docker shim --image <image> --mount-xdg --mount-xdg-app eslint --mount-xdg-dirs config` is run
- **THEN** the shim includes a bind mount from `C:/Users/alice/.config/eslint` to `/home/<user>/.config/eslint`

### Requirement: XDG base dir resolution

The system SHALL resolve XDG base dirs from environment variables when set; otherwise use defaults under `$HOME`. When resolved on Windows, the resulting host path SHALL be normalized to Docker-compatible Windows form before being rendered into the shim.

#### Scenario: XDG override for config
- **WHEN** `XDG_CONFIG_HOME=$HOME/.config-alt` and `cli2docker shim --image <image> --mount-xdg --mount-xdg-app eslint --mount-xdg-dirs config` is run and the image user is `bun`
- **THEN** the shim mounts `$HOME/.config-alt/eslint` to `/home/bun/.config/eslint`

#### Scenario: XDG override for config on Windows
- **WHEN** `XDG_CONFIG_HOME` is set to a path under `$HOME` that resolves to `C:\\Users\\alice\\.config-alt` on Windows and `cli2docker shim --image <image> --mount-xdg --mount-xdg-app eslint --mount-xdg-dirs config` is run and the image user is `bun`
- **THEN** the shim mounts `C:/Users/alice/.config-alt/eslint` to `/home/bun/.config/eslint`

### Requirement: XDG read-write option

The system SHALL mount XDG directories read-only by default as bind mounts, and read-write when `--mount-xdg-rw` is set.

#### Scenario: Default read-only
- **WHEN** `cli2docker shim --image <image> --mount-xdg --mount-xdg-app eslint --mount-xdg-dirs config` is run without `--mount-xdg-rw`
- **THEN** the shim includes a read-only bind mount for the config dir

#### Scenario: Read-write opt-in
- **WHEN** `cli2docker shim --image <image> --mount-xdg --mount-xdg-app eslint --mount-xdg-dirs config --mount-xdg-rw` is run
- **THEN** the shim includes a read-write bind mount for the config dir
