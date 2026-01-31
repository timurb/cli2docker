## Context

`cli2docker build` derives default image/bin names from `--package`. Today it uses the raw package string, so a package spec with a version/tag (e.g., `@fission-ai/openspec@latest`) leaks `@latest` into the derived image name and yields an invalid Docker reference. The package string must still be preserved for `npm install -g` so users can pin versions.

## Goals / Non-Goals

**Goals:**
- Allow npm package specs with version/tag suffixes while keeping default image/bin derivation valid.
- Keep explicit `--image`, `--bin`, and `--tag` overrides unchanged.
- Minimize changes and avoid new dependencies.

**Non-Goals:**
- Full npm package-spec parsing (git URLs, aliases, file paths).
- Changing tag defaulting behavior.
- Accepting invalid Docker image references.

## Decisions

1) Strip the version/tag suffix using a minimal string rule: remove the substring from the last `@` when it is not the leading scope marker.
- **Rationale:** npm package specs only use `@` for scope and version; the last `@` denotes version/tag for both scoped and unscoped packages.
- **Alternatives:** 
  - Add a package-spec parser dependency: more correctness for exotic specs, but adds dependency and complexity not justified by scope.
  - Special-case scoped packages with a more complex split: higher complexity without clear benefit for the targeted bug.

2) Apply version stripping only when deriving defaults (`imageFromPackage`, `packageBaseName`) and keep the original `opts.Package` for Dockerfile installation.
- **Rationale:** Users expect version pinning in `npm install -g`. Only the derived image/bin should be normalized.
- **Alternative:** Normalize the package string globally, which would break version pinning.

## Risks / Trade-offs

- **Risk:** Non-standard package specs with `@` in unexpected positions may be normalized incorrectly.  
  **Mitigation:** Limit stripping to derived defaults only; document that `--package` expects npm registry specs.
- **Risk:** Behavior changes for odd inputs like `name@tag` might surprise users who relied on the invalid image reference.  
  **Mitigation:** Explicit overrides (`--image`, `--bin`) remain available and unchanged.

## Migration Plan

- No migration required. Backwards compatible for packages without versions; existing failures with versioned packages become valid.
- Update/add tests around `imageFromPackage` and `packageBaseName` derivation with versioned specs.

## Open Questions

- None.
