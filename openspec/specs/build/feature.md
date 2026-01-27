## Overview

`cli2docker build` builds a Docker image for an npm CLI package using a generated Dockerfile.

## Interfaces

- Command: `cli2docker build`
- Flags: `--package`, `--bin`, `--image`, `--tag`, `--base`, `--user`, `--no-user`, `--no-cache`

## Invariants

- Default tag is `latest` if omitted.
- `--no-user` omits `USER` from Dockerfile.
- `--no-cache` passes `--no-cache` to Docker.
- Runtime user defaults to `node` unless `--no-user` is set.

## Usage workflows

- Build an image from an npm package and tag it.
- Build an image without dropping privileges (`--no-user`).

## Key constraints

- Requires a running Docker daemon.
- Requires the npm package to publish the specified binary.

## Related docs

- [spec](spec.md)
- [overview](../overview.md)
