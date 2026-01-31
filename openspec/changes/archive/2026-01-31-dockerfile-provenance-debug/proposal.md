## Why

Users can’t currently see the generated Dockerfile, which makes it hard to debug build issues and to capture build provenance. Exposing the Dockerfile lets users inspect exactly what will be built and archive it alongside artifacts.

## What Changes

- Add a build option to emit the generated Dockerfile to stdout.
- Add or update a build mode to generate the Dockerfile without running `docker build` (debug/provenance use).
- Ensure the emitted Dockerfile matches the one used for actual builds (no divergence).

## Non-goals

- No changes to the default build flow when the new option is not used.
- No full supply-chain attestation or SBOM generation.
- No support for alternate build backends (still Docker).

## Capabilities

### New Capabilities
- (none)

### Modified Capabilities
- `build`: add requirements for emitting the generated Dockerfile and for a no-build/debug mode tied to that output.

## Impact

Build command flag handling, Dockerfile generation/output plumbing, CLI help text, and tests for the new behavior.
