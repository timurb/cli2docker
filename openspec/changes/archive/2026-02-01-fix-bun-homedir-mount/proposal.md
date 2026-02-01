## Why

`cli2docker shim --mount-home` always targets `/home/node`, which is wrong for Bun images that run as user `bun` with home `/home/bun`. This breaks home mounts for Bun-based images and blocks config access that should work.

## What Changes

- Resolve `--mount-home` into the container user's home directory instead of hard-coding `/home/node`.
- Ensure Bun images mount into `/home/bun/<relative>` while keeping existing behavior for Node images.

## Non-goals

- Change `--mount-home` validation rules (paths must remain within `$HOME`).
- Change build defaults or runtime user selection.
- Add new flags or options for mount behavior.

## Capabilities

### New Capabilities
- none

### Modified Capabilities
- `shim`: mount-home target path must match the container user's home directory rather than `/home/node`.

## Impact

Shim mount path logic (e.g., `resolveHomeMount`), shim-related tests, and the shim spec requirement for home mounts.
