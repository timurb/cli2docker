## Why

`cli2docker shim` currently emits `docker run` arguments using Unix-oriented host assumptions such as `${PWD}` and direct bind-mount path concatenation. On Windows hosts, Docker CLI path and workdir arguments are shell- and platform-sensitive, so `--mount-cwd`, `--mount-home`, and XDG mounts can resolve to values that Docker interprets incorrectly or rejects.

## What Changes

- Define Windows-compatible requirements for shim-generated `docker run` arguments that reference host paths.
- Ensure `--mount-cwd`, `--mount-home`, and `--mount-xdg` produce Docker-compatible host path arguments when the shim is executed from supported Windows POSIX environments, specifically `git-bash` and `WSL`.
- Clarify expected behavior for working-directory related arguments so the container still starts in `/work` when `--mount-cwd` is enabled on Windows.
- Add Windows-focused acceptance scenarios for cwd, home, and XDG mount handling.

## Non-goals

- Add native `PowerShell` or `cmd.exe` shim generation or shell-specific argument handling.
- Solve unrelated Docker Desktop integration or host file-sharing configuration problems outside the generated command line.
- Expand `build` behavior or introduce new package-manager features.

## Capabilities

### New Capabilities
- none

### Modified Capabilities
- `shim`: host path and workdir related `docker run` arguments must remain valid when the generated shim is executed from supported Windows POSIX environments (`git-bash`, `WSL`).

## Impact

Shim command generation in `main.go` (notably `buildShimExecLine` and mount path resolution), Windows-specific shim tests, README platform guidance, and delta requirements for `openspec/specs/shim/spec.md`.
