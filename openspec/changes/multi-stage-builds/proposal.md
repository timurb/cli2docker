## Why

Generated images currently install packages directly in a single stage, which makes it hard to separate build-time artifacts and can lead to larger images. Multi-stage builds would let us install and copy only the required runtime artifacts, improving reproducibility and size without changing user workflows.

## What Changes

- Generate multi-stage Dockerfiles for `cli2docker build`, with a build stage that installs the package and a runtime stage that copies only the needed artifacts.
- Keep CLI surface area minimal (no new required flags), but allow opting out if needed.

## Non-goals

- Full image optimization (e.g., stripping binaries, tree-shaking JS, or dependency pruning).
- Changing package manager semantics beyond what is required for multi-stage layout.
- Adding support for additional package managers.

## Capabilities

### New Capabilities

### Modified Capabilities
- `build`: Dockerfile generation SHALL support multi-stage builds to isolate installation from runtime.

## Impact

- Affects Dockerfile generation logic and related tests.
- Potentially affects image size and build time.
