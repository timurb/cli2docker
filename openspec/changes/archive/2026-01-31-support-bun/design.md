## Context

`cli2docker build` currently generates a Dockerfile that uses npm (`npm install -g`) on a Node base image (`node:20-alpine`) and defaults to the `node` user. The change adds an optional Bun-based build path while preserving current defaults and behavior for npm.

## Goals / Non-Goals

**Goals:**
- Allow `cli2docker build` to generate a Bun-based Dockerfile via an explicit flag.
- Keep npm as the default with no behavior changes when the flag is omitted.
- Make defaults for base image and user sensible for the selected package manager while still allowing overrides.

**Non-Goals:**
- Auto-detecting the package manager from package metadata.
- Adding support for other package managers in this change.
- Changing the `shim` command or its output.

## Decisions

- **Flag name and shape:** Introduce `--package-manager` with allowed values `npm` (default) and `bun`.
  - **Why:** The change is specifically about how packages are installed; “package manager” is explicit and avoids overloading “runtime.”
  - **Alternative:** `--runtime` (ambiguous for Node vs Bun runtime) or `--bun` (less extensible if future managers are added).

- **Default base image and user are runtime-aware:** If `--package-manager=bun` is selected and the user did not explicitly set `--base`/`--user`, default to a Bun base image (e.g., `oven/bun:1`) and user `bun`.
  - **Why:** The current defaults assume a Node image and `node` user; these are likely invalid for Bun images.
  - **Alternative:** Keep current defaults and require explicit `--base`/`--user` for Bun, but that makes Bun support brittle and less ergonomic.

- **Dockerfile generation split by package manager:** Keep a shared Dockerfile generator but branch the install and environment lines:
  - npm: current `ENV NPM_CONFIG_*` + `npm install -g <package>`.
  - bun: `RUN bun add -g <package>` and omit npm-specific env.
  - Labels and entrypoint remain identical.
  - **Why:** Minimal change to existing flow, clear separation, and easy to test.
  - **Alternative:** Template the entire Dockerfile with string substitutions; more code for little benefit.

- **Validation:** Validate `--package-manager` early in `RunE` with a clear error on unknown values.
  - **Why:** Fail fast at the boundary and keep internal logic simple.
  - **Alternative:** Defer validation inside Dockerfile generator; that mixes shell concerns with core logic.

## Risks / Trade-offs

- **Base image expectations may drift** → Mitigation: keep defaults but document that `--base`/`--user` are overrideable; test with the chosen base image.
- **User mismatch when custom base used** → Mitigation: respect explicit `--user` and `--no-user`; do not guess.
- **Bun install behavior differs from npm** → Mitigation: restrict change to install command and entrypoint; leave functional testing to integration tests.

## Migration Plan

- Backward compatible: npm remains the default with identical output when `--package-manager` is omitted.
- Release note: document the new flag and defaults for Bun builds.
- No data migrations.

## Open Questions

- Should the flag be `--package-manager` or `--runtime`? (Preference: `--package-manager`.)
- What exact default Bun base image tag should we use (e.g., `oven/bun:1` vs `oven/bun:1.1.x`)?
- Do we want to support `--bun` as a short alias?
