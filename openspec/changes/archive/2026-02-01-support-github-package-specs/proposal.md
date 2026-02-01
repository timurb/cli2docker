## Why

`cli2docker build` derives default `--image` and `--bin` from `--package`. For npm
git specs like `github:tobi/qmd`, the derived defaults contain `:` and `/`, which
produces invalid Docker image references and unusable defaults. We need a
defined behavior so users can build from common `github:` specs without manual
workarounds.

## What Changes

- Accept npm git spec form `github:<owner>/<repo>` as a valid `--package` input.
- Derive default `--image` and `--bin` from the repository name portion of the
  `github:` spec.
- Provide a clear error if a supported git spec cannot produce valid defaults.
- Document the supported git spec form and any limitations.

## Non-goals

- Full npm package-spec parsing (e.g., arbitrary git URLs, tarballs, file paths).
- Registry lookups to resolve package metadata or binaries.
- Auto-detecting the CLI binary from package metadata.

## Capabilities

### New Capabilities

<!-- none -->

### Modified Capabilities

- `build`: Accept `github:` package specs and define default derivation behavior.

## Impact

- Build option parsing and default derivation helpers.
- Build-related tests and spec coverage.
- README usage examples/notes.
