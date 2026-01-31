## Why

`cli2docker build` already derives `--image` from `--package`, but there is no way to apply a consistent namespace/prefix across images. Adding `--image-prefix` makes it easy to keep images grouped without requiring explicit `--image` every time.

## What Changes

- Add `--image-prefix` to `cli2docker build` with default `cli/`.
- When `--image` is omitted, apply the prefix to the derived image name.
- When `--image` is provided explicitly, ignore `--image-prefix` and emit a warning.

## Capabilities

### New Capabilities
None.

### Modified Capabilities
- `build`: add prefix behavior to image name derivation.

## Non-goals

- No changes to `shim`.
- No changes to image tagging behavior.
- No automatic registry detection or namespace inference.

## Impact

- `openspec/specs/build/spec.md` (delta spec).
- `main.go` CLI flag handling and image derivation.
- `README.md` usage examples.
