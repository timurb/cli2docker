## Context

`docs/feature-01-cli.md` documents the root CLI structure and shared behavior, but OpenSpec has no `cli` capability. We need to add a spec that captures the CLI surface as requirements/scenarios without changing implementation.

## Goals / Non-Goals

**Goals:**
- Create an OpenSpec spec for the core CLI (`cli2docker`, `build`, `shim`).
- Translate interface and behavior from `feature-01-cli`.

**Non-Goals:**
- No changes to runtime behavior or flag parsing.
- No updates to other features (`build`, `shim`) beyond their own specs.
- No documentation edits during this change.

## Decisions

### Decision 1: New capability `cli`
Create `openspec/changes/sync-feature-01-cli/specs/cli/spec.md` and later sync to `openspec/specs/cli/spec.md`.

### Decision 2: Map doc sections to requirements
Map `Interface/Behavior/Constraints/Errors` into requirements with scenarios. Keep examples as scenarios where possible.

## Risks / Trade-offs

- **Risk:** CLI help/usage behavior is not fully specified → **Mitigation:** keep requirements minimal and aligned to existing doc.
- **Risk:** Spec drift if legacy docs remain → **Mitigation:** remove `docs/feature-01-cli.md` after archive.
