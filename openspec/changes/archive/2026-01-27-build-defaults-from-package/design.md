## Context

The `build` command currently requires explicit `--image` and `--bin`. We want to derive defaults from `--package` when those flags are omitted, without introducing network lookups.

## Goals / Non-goals

**Goals:**
- Derive `--image` and `--bin` defaults from `--package` when omitted.
- Emit warnings on stderr when auto-values are applied.
- Keep explicit flag values as the source of truth.

**Non-goals:**
- Query npm registry or package metadata.
- Change tag defaulting or build workflow beyond the new defaults.
- Update docs in this change (deferred).

## Decisions

### Decision 1: Derive defaults in the CLI layer after flag parsing

We will remove `--image` and `--bin` from required flags and compute defaults inside `RunE` before calling `buildWithOptions`. This keeps derivation at the boundary and keeps core build logic unchanged.

### Decision 2: Simple, deterministic package parsing

Implement small helpers:
- `packageBaseName(pkg string) string` -> returns name without scope (`@scope/name` -> `name`).
- `imageFromPackage(pkg string) string` -> returns `pkg` without leading `@` for scoped packages (`@scope/name` -> `scope/name`).

These helpers are pure and do not validate beyond string parsing (prototype profile).

### Decision 3: Warning output

When a default is applied, emit a warning to `stderr`, e.g.:
- `warning: --image not set, using derived value "<value>"`
- `warning: --bin not set, using derived value "<value>"`
