## Why

Bun is increasingly used to install and run JavaScript CLIs, but `cli2docker` currently assumes npm. Adding Bun support avoids forcing users to rewrite toolchains and enables faster install paths where Bun is preferred.

## What Changes

- Add a build option to select the package manager/runtime (`npm` default, `bun` optional) for `cli2docker build`.
- Generate a Bun-based Dockerfile path when Bun is selected (base image + install + entrypoint adjusted for Bun).
- Extend validation/help text and tests to cover the Bun build flow.

## Non-goals

- Auto-detecting the package manager from package metadata.
- Supporting additional managers (e.g., pnpm, yarn) in this change.
- Changing the `shim` subcommand behavior.

## Capabilities

### New Capabilities
- `<none>`: No new capability; this is an extension of existing build behavior.

### Modified Capabilities
- `build`: Add requirements to support a Bun-based build workflow and Dockerfile generation when selected.

## Impact

Build command flag surface, Dockerfile generation logic, and build tests; potential new base image requirements for Bun in generated Dockerfiles.
