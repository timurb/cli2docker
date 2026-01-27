## Overview

`cli2docker shim` prints a shell shim script to stdout to run a Docker image as a CLI.

## Interfaces

- Command: `cli2docker shim`
- Flags: `--image`, `--mount-cwd`, `--mount-home`, `--mount-home-rw`

## Invariants

- For `--mount-cwd`, insert `-v "${PWD}:/work"` and `-w /work`.
- For `--mount-home`, resolve to `/home/node/<relative>` and mount read-only by default.

## Usage workflows

- Generate a shim and save it to a file in `PATH`.
- Generate a shim with `--mount-cwd` and run it against the current project.
- Generate a shim with `--mount-home` for config access.

## Key constraints

- Mounts are opt-in.
- `--mount-home` must be inside `$HOME`.
- Only one `$HOME` mount is supported.

## Related docs

- [spec](spec.md)
- [overview](../overview.md)
