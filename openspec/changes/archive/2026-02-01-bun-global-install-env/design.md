## Context

`cli2docker build` generates a Dockerfile for Bun with `RUN bun add -g <pkg>` before dropping to the `bun` user. Bun installs global packages under `$HOME/.bun` by default; when run as root this lands in `/root`, making the CLI unavailable to the runtime user. We need a deterministic, non-home location for Bun global installs while keeping npm behavior unchanged for now.

## Goals / Non-Goals

**Goals:**
- Ensure Bun global installs land in a system path accessible at runtime.
- Keep the runtime binary on a standard PATH location.
- Limit scope to Dockerfile generation for Bun builds.

**Non-Goals:**
- Multi-stage builds or image size optimizations.
- Changing npm global install behavior unless a matching runtime issue is confirmed.
- Adding support for additional package managers.

## Decisions

- **Set Bun global install environment variables in the Dockerfile.**  
  Use `BUN_INSTALL_GLOBAL_DIR=/usr/local/bun/global` and `BUN_INSTALL_BIN=/usr/local/bin` before `bun add -g`. This keeps global payloads out of user home and places binaries on a standard path.

- **Keep install running as root.**  
  Installing as root into system paths avoids per-user homes and preserves current Dockerfile flow. The runtime user only needs read/execute access.

- **Do not set Bun cache overrides.**  
  Cache affects build-time only and is out of scope; we can revisit in a later multi-stage-build change.

- **Leave npm unchanged for now.**  
  Only adjust npm if we reproduce the same runtime access issue; otherwise keep the current behavior to avoid unintended changes.

## Risks / Trade-offs

- **Packages attempt to write into their install directory at runtime** → Such tools will fail under the unprivileged user. Mitigation: document behavior and recommend `--no-user` or a custom base if a tool requires write access.
- **PATH assumptions differ across base images** → `/usr/local/bin` may not be on PATH for all images. Mitigation: select a standard path used by typical base images; if a custom base lacks it, users can override `--base`/`--user` or future enhancements can add explicit PATH.

## Migration Plan

- Regenerate images to pick up the new Dockerfile env vars.
- No runtime migration needed beyond rebuilding images.

## Open Questions

- If npm shows the same runtime access issue, should we adopt `NPM_CONFIG_PREFIX=/usr/local` in a follow-on change?
