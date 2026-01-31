## Why

After running `cli2docker build --package "@fission-ai/openspec@latest"` and generation of shim running the shim fails because the derived image name includes the package version suffix (`@latest`), producing an invalid Docker reference. Fixing this allows standard npm version syntax while preserving the default image derivation behavior.

## What Changes

- Strip npm version/tag suffixes (e.g., `@latest`, `@1.2.3`) when deriving default image and binary names from `--package`.
- Keep `--image`, `--bin`, and `--tag` overrides unchanged.

## Non-goals

- Changing how `--tag` defaulting works.
- Allowing invalid Docker image references.

## Capabilities

### New Capabilities
None.

### Modified Capabilities
- `build`: Default image/bin derivation ignores npm version/tag suffixes in `--package`.

## Impact

Build command parsing and default derivation logic; related tests and documentation for `build`.
