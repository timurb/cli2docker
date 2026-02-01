## Context

- `cli2docker build` already writes status lines and `docker build` output to stdout.
- `--print-dockerfile` is a separate mode that writes only the Dockerfile to stdout.
- `cli2docker shim` requires a concrete image reference; `build` already resolves that via image+tag.

## Goals / Non-Goals

**Goals:**
- Print a runnable shim-generation command as part of the standard `build` flow.
- Use the resolved image reference (including tag) so the command matches the built image.
- Keep `--print-dockerfile` output isolated to the Dockerfile only.

**Non-Goals:**
- Changing shim behavior or flags.
- Introducing new output formats or structured/machine-readable output.
- Changing Dockerfile output behavior for `--print-dockerfile`.

## Decisions

1. **Stdout is the destination.**
   - **Why:** Explicit requirement; the command should be copy-pasteable without flags.
   - **Alternatives:** stderr or an opt-in flag; rejected to avoid extra UX steps.

2. **Emit only after a successful build.**
   - **Why:** Avoid printing a command that points to a failed build.
   - **Alternatives:** Print before build or regardless of outcome; rejected to prevent stale guidance.

3. **Command format is a single line: `cli2docker shim --image <image>`.**
   - **Why:** Minimal, directly runnable, and does not assume install path or file location.
   - **Alternatives:** Include redirects or `chmod` guidance; rejected because paths are user-specific.

4. **Non-shim build output goes to stderr.**
   - **Why:** Keeps stdout dedicated to the shim command for copy-paste and scripting.
   - **Alternatives:** Keep status/build output on stdout; rejected due to requirement for stdout-only shim line.

## Risks / Trade-offs

- **[Stdout contamination]** The shim line can break consumers that parse stdout. → **Mitigation:** Accept the trade-off; keep the shim command as a standalone line emitted last.
- **[Discoverability in noisy logs]** The shim line may be buried in docker build output. → **Mitigation:** Emit it after build completion so it is the final line.

## Migration Plan

- No migration required. Update specs and tests to reflect new stdout output.

## Open Questions

- None.
