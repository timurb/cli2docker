## Why

After building an image and generating a shim, there is no durable record of which npm package the image was built from. This makes audits, updates, and debugging harder. We need minimal provenance that survives beyond the build output.

## What Changes

- Record the originating npm package, package version, and derived bin as image metadata during `cli2docker build`.
- Include the same origin metadata in the generated shim output so the shim file is self-describing.
- Document how to read the stored origin metadata (e.g., via image inspection or shim header).

## Non-goals

- No new subcommands (e.g., no `inspect` command).
- No automatic registry resolution.
- No changes to image tagging behavior.

## Capabilities

### New Capabilities
- (none)

### Modified Capabilities
- `build`: store origin metadata (package, version, bin) for the built image derived from `--package`.
- `shim`: emit origin metadata (package, version, bin) in the generated shim output.

## Impact

- Code: `build` and `shim` command paths in `main.go`.
- Tests: update/add tests around image metadata and shim output.
- Users: new metadata (including version) visible in images and shim files; no CLI surface changes.
