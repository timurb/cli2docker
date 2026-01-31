## Context

`cli2docker build` derives a Docker image from an npm package, and `cli2docker shim` emits a runnable script, but neither artifact retains provenance. Today there is no simple way to answer "which package/version produced this image/shim?" without external bookkeeping. The change should add minimal, durable metadata without expanding the CLI surface.

## Goals / Non-Goals

**Goals:**
- Persist the originating package name, package version, and derived bin in the built image metadata.
- Include the same origin metadata in the generated shim output so the file is self-describing.
- Keep existing CLI flags and flows unchanged.

**Non-Goals:**
- No new subcommands or inspection tooling.
- No registry/version resolution beyond what is already provided.
- No changes to image tagging logic.

## Decisions

### Decision 1: Use image labels for provenance
**Choice:** Write package/bin provenance into Docker image labels during build.
**Why:** Labels are standard, queryable via `docker inspect`, and persist with the image. They don't require new CLI commands.
**Alternatives considered:**  
- Bake into image name or tag → pollutes naming, already reserved for user intent.  
- Write a sidecar file on disk → not durable once image is pushed/pulled.

### Decision 2: Store explicit package version only
**Choice:** Record the package spec as provided to `--package`, and extract a version only when explicitly present (e.g., `eslint@9.1.0`, `@scope/pkg@1.2.3`). Do not resolve versions automatically.
**Why:** Keeps behavior deterministic and avoids new network calls or registry lookups. We can revisit version resolution later.
**Alternatives considered:**  
- Resolve latest version via registry lookup → adds network dependency and makes builds sensitive to registry availability.

### Decision 3: Embed provenance as shim comments
**Choice:** Prepend a small comment block to the shim script with package/bin info.
**Why:** Keeps the shim self-describing without changing runtime behavior or requiring parsing at execution time.
**Alternatives considered:**  
- Add shim flags or env vars → expands CLI surface.  
- Store in an external metadata file → easy to lose.

### Decision 4: Normalize metadata keys
**Choice:** Use stable, namespaced label keys like `io.cli2docker.package` and `io.cli2docker.bin`.
**Why:** Avoids collisions and makes intent obvious.

## Risks / Trade-offs

- **Risk:** Users expect resolved versions even when none are specified. → **Mitigation:** document that only explicit versions are recorded unless we add registry resolution.
- **Risk:** Label format inconsistency across existing images. → **Mitigation:** labels are additive; no migration required.

## Migration Plan

- No migration required. New builds will include labels and shim comments. Existing images/shims remain unchanged.

## Open Questions

- (none for now)
