## Why

`cli2docker build` currently requires explicit `--image` and `--bin` even though both can be inferred from `--package`. This is extra typing and a source of small but frequent mistakes.

## What Changes

- If `--image` is omitted, derive it from `--package` (`@scope/name` -> `scope/name`).
- If `--bin` is omitted, derive it from the package name without scope (`@scope/name` -> `name`).
- Emit a warning when auto-values are applied so the user can confirm or override them.

## Capabilities

### New Capabilities
- `build`: derive default `--image` and `--bin` values from `--package` when omitted.

### Modified Capabilities
None.

## Non-goals

- No registry or npm metadata lookups.
- No changes to `shim`.
- No changes to tag defaulting behavior.

## Impact

- `main.go`: add default-derivation and warning output in the build command path.
- `docs/feature-02-build.md`: update interface/behavior description (deferred until permitted).
