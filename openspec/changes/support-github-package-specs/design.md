## Context

`cli2docker build` derives default `--image` and `--bin` from `--package` via
helpers that only understand npm registry specs (scoped/unscoped with optional
version tags). npm also allows git shorthand specs like `github:owner/repo`.
Today those specs flow into default derivation unchanged, producing invalid
Docker image references because of the `github:` prefix.

## Goals / Non-Goals

**Goals:**
- Support `github:<owner>/<repo>` in `--package` for default derivation.
- Derive a valid default image and bin from the repo portion.
- Fail clearly when a `github:` spec cannot be parsed into valid defaults.

**Non-Goals:**
- Full npm package spec parsing (git URLs, tarballs, file paths, aliases).
- Registry lookups to infer binaries or metadata.
- Changing how the package string is passed to npm/bun installs.

## Decisions

1) Minimal parsing of `github:` shorthand only.
   - Detect `strings.HasPrefix(pkg, "github:")`.
   - Parse the remainder as `<owner>/<repo>` with an optional `#<ref>` suffix
     that is stripped for default derivation (same as `@version`) but preserved
     for metadata.
   - If the remainder does not contain exactly one `/` or the owner/repo segment
     is empty, return a clear error.
   - Rationale: keeps scope tight and avoids pulling in a full package‑spec
     parser dependency.
   - Alternatives:
     - Full npm spec parsing library → more correctness, higher complexity.
     - Accept arbitrary git URLs → ambiguous rules and more validation.

2) Default derivation mapping for `github:` specs mirrors scoped package behavior.
   - Treat `github:foo/bar` as equivalent to `@foo/bar` for default derivation.
   - `imageFromPackage("github:foo/bar")` → `foo/bar` (then apply
     `--image-prefix` if present).
   - `packageBaseName("github:foo/bar")` → `bar`.
   - Treat `github:foo/bar#ref` as equivalent to `@foo/bar@ref` for defaults:
     strip `#ref` before deriving `--image`/`--bin`.
   - Rationale: preserves namespace uniqueness for images while keeping bin
     consistent with scoped package defaults.
   - Alternative: image = `bar` only → simpler but increases collisions.

3) Errors only when defaults are needed.
   - Parsing errors should surface only when `--image`/`--bin` are omitted and
     defaults must be derived.
   - If users provide explicit `--image` and `--bin`, allow any `--package`
     string to pass through to npm/bun unchanged.
   - Rationale: aligns with current behavior (defaults are the only consumer of
     the derived name logic).

4) Implementation touchpoints stay localized.
   - Extend `imageFromPackage` and `packageBaseName` to handle `github:` specs.
   - Keep `renderDockerfile` unchanged: `npm install -g <package>` should still
     receive the original spec.
   - Add unit tests for both helpers covering `github:owner/repo` and
     `github:owner/repo#ref`.

5) Treat git ref as explicit version metadata for labels.
   - If `--package` is `github:owner/repo#ref`, set
     `io.cli2docker.package` to `github:owner/repo` and
     `io.cli2docker.package-version` to `ref`, with a note that it is a git ref.
   - Rationale: keep provenance explicit without pretending the ref is a semver.

6) Align `github:` shorthand with documented npm/bun behavior.
   - Accept `github:owner/repo` with optional `#ref` only.
   - Do not accept `github:owner/repo.git`; advise using a full git URL when
     `.git` is required.
   - Rationale: match standard shorthand forms without expanding scope.

## Risks / Trade-offs

- [Rejects valid but uncommon git specs] → Document the supported form and
  advise explicit `--image`/`--bin` for anything else.
- [Ambiguous parsing for edge cases like missing owner/repo] → Explicit error
  with expected format.

## Migration Plan

- No migration required. This is additive and backward‑compatible.

<!-- No open questions. -->
