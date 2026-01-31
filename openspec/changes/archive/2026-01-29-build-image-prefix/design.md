## Context

`cli2docker build` already derives `--image` from `--package`, but there is no built‑in namespace/prefix. Users must repeat image names manually. We want a simple prefix flag with minimal behavioral impact.

## Goals / Non-Goals

**Goals:**
- Add `--image-prefix` with default `cli/`.
- Apply prefix only when `--image` is not explicitly set.
- Emit a warning when `--image-prefix` is ignored due to explicit `--image`.

**Non-Goals:**
- No changes to tag derivation.
- No registry detection or automatic namespace inference.
- No changes to `shim`.

## Decisions

### Decision 1: Apply prefix during image derivation only
Prefix is applied only in the derived path (`--image` absent). This keeps explicit `--image` as the source of truth.

### Decision 2: Warning on ignored prefix
If both `--image` and `--image-prefix` are set, emit a warning that the prefix was ignored.

## Risks / Trade-offs

- **Risk:** Users may expect prefix to apply even with `--image` → **Mitigation:** explicit warning.
- **Risk:** Default prefix may be undesired in some environments → **Mitigation:** allow override via `--image`.
