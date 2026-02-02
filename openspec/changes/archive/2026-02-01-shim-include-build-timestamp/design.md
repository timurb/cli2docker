## Context

`cli2docker shim` prints a self-contained shell script and includes origin metadata as comment lines derived from image labels. The image already carries a build timestamp label (`io.cli2docker.build-timestamp`) from `cli2docker build`, but the shim comment set omits it, breaking provenance tracing for unversioned packages.

## Goals / Non-Goals

**Goals:**
- Emit `io.cli2docker.build-timestamp` as a shim comment when the label is present.
- Keep shim output deterministic and backward-compatible when the label is absent.

**Non-Goals:**
- Changing how build timestamps are generated or stored on images.
- Introducing new CLI flags or altering docker run semantics.
- Modifying Docker image `Created` metadata.

## Decisions

- **Extend shim origin comments to include build timestamp.** This keeps provenance metadata in one place, aligns shim output with image labels, and avoids new surface area.
- **No new label reads or CLI flags.** Reuse existing `readImageLabels` output and only extend comment rendering. Alternative (flag-gated output) adds unnecessary complexity for a small, always-useful metadata line.

## Risks / Trade-offs

- **[Risk]** Some images may not have the build timestamp label → **Mitigation:** Emit the comment only when present; no change in behavior otherwise.
- **[Trade-off]** Slightly larger shim header → **Mitigation:** Single extra line; negligible impact.

## Migration Plan

No migration required. The change only affects newly generated shim output and is backward compatible.

## Open Questions

None.
