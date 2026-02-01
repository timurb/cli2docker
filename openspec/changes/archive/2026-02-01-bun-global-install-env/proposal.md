## Why

When using Bun, global installs run as root inside the image and land under `/root`, so the runtime user cannot access the CLI at execution time. This breaks `cli2docker build` for Bun-based images today; we need a deterministic, non-home install path.

## What Changes

- Set Bun global install environment variables in generated Dockerfiles so global packages install into a system path and their binaries land on a standard PATH location.
- Keep npm behavior unchanged unless the same issue is observed and explicitly addressed in a follow-on change.

## Non-goals

- Multi-stage builds or image size optimizations.
- Changing npm global install behavior.
- Supporting additional package managers beyond npm and Bun.

## Capabilities

### New Capabilities

### Modified Capabilities
- `build`: The Bun Dockerfile generation SHALL set global install environment variables to avoid user home installs.

## Impact

- Affects Dockerfile generation for Bun builds.
- Updates build specs and tests to codify the new environment variables.
