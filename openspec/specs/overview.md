## Overview / Context

`cli2docker` packages npm CLI tools into Docker images and provides a shim to run them as local commands.

## Problems addressed

- Run npm CLIs without installing Node.js on the host.
- Provide consistent CLI behavior via Docker images.

## Goals / Non-Goals

**Goals:**
- Build images for npm CLI tools.
- Provide a shim that runs those images as local commands.
- Keep interaction via a single CLI entrypoint.

**Non-Goals:**
- Multi-arch builds or registry publishing.
- Managing multiple host environments or runtime sandboxing.

## Usage workflows

- Build an image from an npm package.
- Generate a shim and install it in `PATH`.

## Success criteria

- `cli2docker build` produces a runnable image for a published npm CLI.
- `cli2docker shim` produces a working shim script.
- README contains usage documentation and examples.

## Architecture

- **CLI entrypoint**: `cli2docker`
- **Capabilities**: `cli`, `build`, `shim`

## Constraints

- Requires a working Docker daemon.
- Shim mounts are opt-in and limited to `$PWD` and a single `$HOME` path.

## Invariants

- `build` always uses a generated Dockerfile.
- `shim` always emits a shell script to stdout.

## Design decisions (from ADRs)

- `shim` mounts current directory only when `--mount-cwd` is set; mount point is `/work`.
- `shim` home mounts are limited to a single path inside `$HOME`, read-only by default, with `--mount-home-rw` for write access.
- `shim` prints its script to stdout; users choose where to save it.
- CLI uses `spf13/cobra` for subcommands and flag parsing.
- Multi-language prototypes were evaluated; Go was chosen (historical, superseded).

## Capabilities

- [CLI](cli/feature.md)
- [Build](build/feature.md)
- [Shim](shim/feature.md)

## Related specs

- [CLI requirements](cli/spec.md)
- [Build requirements](build/spec.md)
- [Shim requirements](shim/spec.md)
