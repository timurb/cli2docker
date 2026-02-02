## Why

We need shims to be self-describing enough to trace back to the exact build provenance. Today the shim omits the build timestamp label, which breaks the “shim → image → source” chain for unversioned installs.

## What Changes

- Include `io.cli2docker.build-timestamp` as a comment line in generated shim output when the label exists on the image.
- Update shim origin metadata requirements to include the build timestamp label alongside package/bin metadata.

## Capabilities

### New Capabilities

- None.

### Modified Capabilities

- `shim`: require the shim output to include `io.cli2docker.build-timestamp` when present on the image.

## Non-goals

- Changing how build timestamps are generated or stored on images.
- Modifying Docker image `Created` metadata or registry timestamps.
- Adding new CLI flags for shim output.

## Impact

- `main.go`: shim comment emission (`originCommentLines`) and related tests.
- `openspec/specs/shim/spec.md`: requirement update to include build timestamp in shim output.
