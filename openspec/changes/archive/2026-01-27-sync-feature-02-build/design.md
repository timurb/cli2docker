## Context

`docs/feature-02-build.md` currently documents the `build` command, while OpenSpec contains only a partial spec for `build` (defaults derived from `--package`). The goal is to make OpenSpec the source of truth by translating the feature doc into requirements/scenarios and syncing them into `openspec/specs/build/spec.md`.

## Goals / Non-Goals

**Goals:**
- Convert `feature-02-build` content into OpenSpec requirements and scenarios.
- Update the existing `build` spec via a delta spec.
- Keep implementation unchanged.

**Non-Goals:**
- No changes to `cli2docker build` behavior or flags.
- No conversion of other features or ADRs.
- No documentation edits until after verification/archival.

## Decisions

### Decision 1: Use a delta spec for capability `build`
We will create a delta spec under `openspec/changes/sync-feature-02-build/specs/build/spec.md` and sync it into `openspec/specs/build/spec.md`. This keeps the change localized and preserves existing `build` requirements.

### Decision 2: Map feature doc sections to requirements
Sections in `feature-02-build` will be transformed as follows:
- **Interface/Behavior/Constraints/Errors/Security** → requirements with scenarios where testable.
- **In scope/Out of scope** → goals/non-goals (specs only for in-scope behavior).

## Risks / Trade-offs

- **Risk:** Feature doc may contain implicit behavior not captured as scenarios → **Mitigation:** review spec against code and run `verify` after syncing.
- **Risk:** Spec drift if docs remain after sync → **Mitigation:** remove `docs/feature-02-build.md` after verification and archive.

## Migration Plan

1. Write delta spec for `build` from `feature-02-build`.
2. Sync delta spec into `openspec/specs/build/spec.md`.
3. Verify change.
4. After archive, delete `docs/feature-02-build.md`.

## Open Questions

- None.
