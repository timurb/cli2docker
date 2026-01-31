## Why

The shim currently runs containers with minimal isolation. We want safer-by-default execution for packaged CLIs without forcing users to hand-edit the generated script.

## What Changes

- Enable hardening defaults in shim execution: drop Linux capabilities and set `no-new-privileges`.
- Enable container root filesystem `--read-only` by default, with an explicit opt-out flag.
- Add opt-out flags for the hardening defaults (caps drop, no-new-privileges, read-only).
- Emit a warning to stderr that read-only mode is experimental when the shim runs.

## Non-goals

- Changing the build workflow or Dockerfile generation.
- Introducing sandboxing beyond Docker’s existing isolation (e.g., seccomp profiles, gVisor).
- Adding new mount behaviors beyond explicit user-specified mounts.

## Capabilities

### New Capabilities
- (none)

### Modified Capabilities
- `shim`: default runtime hardening and explicit opt-out flags.

## Impact

- Shim script generation logic, CLI flags, and README usage documentation.
