## Context

`cli2docker shim` already supports `--mount-home` (validated under `$HOME`) and `--mount-cwd`. Users frequently need app-scoped config/cache/state directories; today they must manually pass specific paths, and `--mount-home` always targets `/home/node` even when the image runs as a different user. The change adds standard XDG-style homedir mounts, aligns mount targets with the container user's home directory, and uses a build-time label as the source of truth for runtime user.

## Goals / Non-Goals

**Goals:**
- Provide explicit shim flags to mount XDG-style app directories under `$HOME` with read-only default.
- Align `--mount-home` target paths with the container user's home directory.
- Always label the runtime user in `build` and require that label in `shim` for mount resolution.
- Reuse existing validation and mount-building patterns to keep behavior consistent.
- Document that enabling these mounts may create host directories by default.
- Allow the caller to override the XDG app subpath and select which XDG dirs to mount.

**Non-Goals:**
- Support macOS/Windows standard directories in this change.
- Mount arbitrary homedir subtrees beyond XDG defaults.
- Change build behavior or global hardening defaults.

## Decisions

- **User selection for mount resolution.** Use `io.cli2docker.user` label as the source of truth; require it and fail if missing. Use `/root` for `root`, `/home/<user>` for others.
  - *Alternative:* Derive from image-configured user (rejected to avoid ambiguity and missing metadata).
- **XDG mounts plus aligned `--mount-home`.** Add XDG Base Directory defaults and optional `$XDG_*_HOME` overrides (kept within `$HOME`), and apply the same effective-user home mapping to `--mount-home` and XDG mounts.
- **Read-only by default with explicit opt-in.** Follow the `--mount-home` pattern: XDG mounts are `:ro` unless an explicit `--mount-xdg-rw` (or equivalent) flag is provided.
  - *Alternative:* Default to read-write (rejected to preserve shim hardening posture).
- **Explicit XDG app subpath with sane default.** Add a parameter to override the XDG subpath (e.g., `--mount-xdg-app <name>`). If omitted, default to the app name derived from image labels (prefer `io.cli2docker.bin`, else `io.cli2docker.package`), and fail with a clear error if no label is available.
  - *Alternative:* Always require the app subpath (rejected for ergonomics).
- **Selectable XDG dirs with default all.** Add a parameter to choose which XDG dirs to mount (e.g., `--mount-xdg-dirs config,cache,data,state`), defaulting to all when omitted.
  - *Alternative:* Separate per-dir flags (rejected to avoid flag explosion).
- **Bind mounts via `-v` for consistency.** Keep using the existing `-v` style mount generation for shims, and document that host paths may be created if missing.
  - *Alternative:* Use `--mount type=bind` and fail if missing (rejected to avoid behavioral divergence and added flags).

## Risks / Trade-offs

- **Directory creation side effect** → Users may see empty XDG dirs created even if the app never uses them. Mitigation: document this in CLI params and keep mounts opt-in.
- **App name mismatch** → Defaults could be wrong for some tools. Mitigation: `--mount-xdg-app` override and clear help text.
- **Missing user label** → Shims for older images will fail. Mitigation: rebuild images; change is acceptable for single-user workflow.
- **XDG env override outside `$HOME`** → Rejected with validation error, which could surprise users. Mitigation: clear error message referencing `$HOME` constraint.

## Migration Plan

- Add new flags to shim command (`--mount-xdg*`) and update help text.
- Add build label for runtime user and require it in shim.
- Implement path resolution/validation using existing helpers.
- Add tests covering path derivation, read-only default, and read-write opt-in.
- No user data migration; change is additive and backward compatible.

## Open Questions

- Should we include legacy dotfile mounts (e.g., `~/.<app>`) as optional flags?
- Do we want a single “mount all XDG dirs” convenience flag or keep per-dir flags only?
