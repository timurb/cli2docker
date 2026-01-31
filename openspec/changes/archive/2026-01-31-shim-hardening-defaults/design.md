## Context

The `shim` command emits a shell script that wraps `docker run` with minimal isolation. We want safer-by-default execution while keeping explicit opt-out flags for compatibility. This change touches CLI flags, shim generation, and user-facing documentation, but does not alter build behavior.

## Goals / Non-Goals

**Goals:**
- Enable default hardening for shim execution: drop Linux capabilities and set `no-new-privileges`.
- Enable `--read-only` by default with an opt-out flag.
- Emit a warning to stderr that read-only mode is experimental.
- Keep hardening explicitly reversible via flags.

**Non-Goals:**
- Modifying Dockerfile generation or build flow.
- Introducing new sandboxing mechanisms (seccomp profiles, gVisor).
- Automatically adding new mounts beyond what the user requests.

## Decisions

1. **Default hardening flags in the shim runtime.**
   - Add `--cap-drop=ALL` and `--security-opt=no-new-privileges` to the generated `docker run` command by default.
   - *Alternatives considered:* opt-in hardening flags only (rejected: defaults remain unsafe).

2. **Opt-out flag naming.**
   - Use `--no-drop-caps` to disable `--cap-drop=ALL`.
   - Use `--no-new-privileges` to disable `no-new-privileges`.
   - Use `--no-read-only` to disable `--read-only`.
   - *Rationale:* uniform negative flags signal default-on behavior, explicit opt-out keeps script readable.

3. **Read-only default with explicit warning.**
   - Add `--read-only` by default and print a warning to stderr that this experimental behavior is enabled.
   - *Rationale:* security by default with explicit user visibility; stderr preserves stdout for scripting.

4. **Mount semantics are unchanged.**
   - Read-only rootfs does not block explicit RW bind mounts (e.g., `-v ~/.config:~/.config:rw` remains writable).
   - *Rationale:* rely on Docker’s documented behavior; avoid introducing implicit mounts.

## Risks / Trade-offs

- **[Compatibility regressions]** Some CLIs expect to write outside mounted volumes (e.g., `$HOME`, `/var`) and will fail under read-only defaults. → Mitigation: `--read-write` opt-out and explicit docs.
- **[Flag confusion]** Negative flags may be misread or inverted. → Mitigation: keep names direct, document in README and `--help`.
- **[Silent failures]** If warnings are missed, users may misattribute failures. → Mitigation: stderr warning on every shim run when read-only default is active.

## Migration Plan

- Add flags to the shim CLI and update generated script logic.
- Update README and `shim` help text to describe defaults and opt-outs.
- No runtime migrations; behavior change is immediate upon upgrade and can be reverted per invocation.

## Open Questions

- Keep naming uniform across all opt-out flags in this change. If we need mixed naming styles, create a separate change.
- The warning SHOULD mention the exact opt-out flag (`--no-read-only`) for discoverability.
