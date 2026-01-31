## Context

`cli2docker build` already labels images with package and bin metadata, but builds from unversioned packages are not directly traceable. We need a minimal, stable timestamp recorded once per build and propagated into the image label and printed Dockerfile without changing the CLI interface.

## Goals / Non-Goals

**Goals:**
- Capture a single UTC build timestamp at command start and reuse it for all labeling.
- Add a deterministic label key/value in the generated Dockerfile and built image.
- Keep the build interface unchanged (no new flags or environment knobs).

**Non-Goals:**
- Determining the resolved package version for unversioned packages.
- Supporting reproducible builds or `SOURCE_DATE_EPOCH`.
- Modifying Docker’s `Created` field or registry timestamps.

## Decisions

- **Timestamp format**: Use RFC3339 in UTC (seconds precision) for readability and interoperability. Alternatives: Unix epoch (less readable), RFC3339Nano (too noisy).
- **Label key**: Use `io.cli2docker.build-timestamp` to match existing `io.cli2docker.*` labels. Alternative: `build-time` (less explicit), or nested keys (adds complexity).
- **Single capture point**: Record timestamp once at the start of `cli2docker build` and pass it through Dockerfile generation to avoid mismatched values across labels/print-only mode. Alternative: generate inside Dockerfile (would require shell logic and reduces determinism across outputs).
- **No CLI change**: Keep UX stable; this is additive metadata only.

## Risks / Trade-offs

- [Cache invalidation] New timestamp label changes on every build, reducing Docker layer reuse → acceptable because provenance is the priority and build is already per-package.
- [Printed Dockerfile reuse] Users who save `--print-dockerfile` output and build later will embed an old timestamp → document as expected behavior in specs.

## Migration Plan

- Add the new label in Dockerfile generation and image label set; no migration required for existing images.

## Open Questions

- Should we also include a Unix epoch label for machine parsing, or is RFC3339 sufficient?
