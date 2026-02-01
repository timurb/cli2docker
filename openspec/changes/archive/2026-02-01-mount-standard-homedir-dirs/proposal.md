## Why

Users running CLI tools via `cli2docker shim` often need app-scoped config/cache/state persisted in their home directory. Today they must manually pass a `--mount-home` path and know the exact directory layout, and `--mount-home` always targets `/home/node` even when the image runs as a different user (e.g., Bun or root). Adding standard homedir mounts reduces friction, and aligning mount targets with the container user fixes incorrect home mounts.

## What Changes

- Add shim flags to mount app-scoped standard directories under the container user's home directory (e.g., `.config/<app>`, `.cache/<app>`, `.local/share/<app>`, `.local/state/<app>`), defaulting to read-only unless explicitly requested read-write.
- Provide consistent validation and error messaging for these derived mounts, similar to existing `--mount-home` validation.
- Resolve `--mount-home` into the container user's home directory instead of hard-coding `/home/node`.
- Always set an image label for the runtime user during `build`, and require it for `shim` mount resolution.
- Document in the CLI parameters that enabling these mounts can create the host directories by default even if the app does not use them.

## Non-goals

- No automatic mounting of arbitrary home subtrees or hidden directories beyond the app-scoped standard dirs.
- No changes to `cli2docker build` behavior.

## Capabilities

### New Capabilities
- None.

### Modified Capabilities
- `shim`: Add requirements for standard homedir mounts and their flags/behavior; require user label for mount resolution.
- `build`: Add requirement to label the runtime user on built images.

## Impact

- `cli2docker shim` CLI surface (new flags, derived mount paths, mount-home target changes, and user label requirement).
- `cli2docker build` label set (runtime user label).
- Shim script generation logic and validation.
- Tests for mount path derivation, defaults (read-only), and read-write opt-ins.
