## Context

`cli2docker build` generates a Dockerfile in a temp directory and immediately runs `docker build`. Users can’t see the generated Dockerfile, which makes debugging and provenance tracking harder. The change should expose the exact Dockerfile content while keeping the default build flow intact.

## Goals / Non-Goals

**Goals:**
- Provide a way to output the generated Dockerfile for inspection or archival.
- Allow a debug/provenance mode that generates the Dockerfile without running `docker build`.
- Ensure the emitted Dockerfile is exactly the same as the one used for real builds.

**Non-Goals:**
- Adding non-Docker build backends.
- Generating SBOMs or supply-chain attestations.
- Changing default build behavior when the new option is not used.

## Decisions

1. **Add a `--print-dockerfile` flag to `build`** that renders the Dockerfile to stdout and exits without invoking Docker.
   - *Alternatives considered:*
     - `--dockerfile <path>` plus `--no-build` (two flags, more surface area).
     - Always build and optionally print (stdout interleaves with build logs).
   - *Rationale:* One boolean flag is the smallest interface; users can redirect stdout to a file for provenance.

2. **Single source of truth for Dockerfile content** via a pure `renderDockerfile(opts buildFlags) string` helper.
   - `writeDockerfile` will use the helper, and `--print-dockerfile` will print the same string.
   - Ensures no divergence between printed and built Dockerfiles.

3. **Require Docker only when actually building.**
   - Move `ensureCommandFn("docker")` to run only in the build path (not in `--print-dockerfile` mode).
   - Allows debugging/provenance output on machines without Docker installed.

4. **Keep stdout strictly Dockerfile content in print-only mode.**
   - Derivation warnings remain on stderr; stdout is reserved for the Dockerfile so users can redirect cleanly.

## Risks / Trade-offs

- **[Output mixing]** Printing to stdout could mix with other output → `--print-dockerfile` exits before build logs, and existing warnings stay on stderr.
- **[Interface lock-in]** Choosing `--print-dockerfile` may limit future path-based output → users can redirect stdout; revisit if needed.
- **[Behavioral change]** Docker check is deferred until build path → update specs/tests to reflect no-Docker requirement in print-only mode.

## Migration Plan

- Add the new flag and render helper.
- Update build flow and tests.
- Update build specs/README to document the new mode.

## Open Questions

- Is `--print-dockerfile` the preferred flag name, or do we want `--dockerfile`/`--no-build` for explicitness?
