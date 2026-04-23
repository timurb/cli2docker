## Context

`cli2docker shim` currently emits a POSIX `sh` script and assembles the `docker run` command line in Go. Today that assembly assumes Unix-style host paths:

- `--mount-cwd` renders `-v "${PWD}:/work" -w /work`
- `--mount-home` and `--mount-xdg` render `-v <host>:<container>[:mode]` using `filepath`-derived host paths

That shape is acceptable on Unix hosts, but it breaks down on Windows. Windows host paths contain drive letters and backslashes, while Git Bash/MSYS may also rewrite Unix-looking arguments before passing them to native executables such as `docker.exe`. The change scope is limited to Windows POSIX environments where a POSIX `sh` shim is still a valid contract: `git-bash` and `WSL`.

## Goals / Non-Goals

**Goals:**
- Keep a single POSIX `sh` shim output and make its mount-related `docker run` arguments valid in `git-bash` and `WSL`.
- Preserve current container-side paths and semantics for `--mount-cwd`, `--mount-home`, and `--mount-xdg`.
- Preserve existing read-only/read-write behavior for home and XDG mounts.
- Make the Windows behavior explicit enough to cover with deterministic tests.

**Non-Goals:**
- Generate native `PowerShell` or `cmd.exe` shims.
- Change `build` behavior, package-manager behavior, or image layout.
- Solve Docker Desktop installation, file-sharing, or shell bootstrap issues outside the generated command line.

## Decisions

1. **Keep the generated shim as POSIX `sh` and explicitly scope Windows support to POSIX environments.**
   - The generated script remains `#!/usr/bin/env sh`.
   - `git-bash` and `WSL` are in scope because they can execute the existing shim contract directly.
   - *Alternative considered:* Generate separate native Windows shim variants. Rejected because it expands the command contract and is explicitly out of scope.

2. **Render bind mounts with `--mount type=bind,...` instead of `-v`.**
   - Windows drive letters make `-v <host>:<container>[:opts]` ambiguous and brittle.
   - `--mount` separates `src=` and `dst=` fields, so `C:/Users/...` can be passed without relying on colon parsing.
   - Use `bind-create-src` to preserve the current practical behavior of `-v`, which tolerates missing source directories for home/XDG paths.
   - *Alternative considered:* Keep `-v` and escape or transform drive letters specially. Rejected as fragile and harder to reason about across Git Bash and WSL.

3. **Normalize host paths differently for runtime cwd mounts versus generation-time home/XDG mounts.**
   - `--mount-cwd` is runtime-derived, so the shim script should resolve the host cwd immediately before `docker run`.
   - In `WSL` and other non-MSYS POSIX environments, the runtime cwd path can be passed through unchanged.
   - In `git-bash`/MSYS, the runtime cwd path should be converted to Docker-friendly Windows form using `cygpath -m`, producing `C:/...`.
   - `--mount-home` and `--mount-xdg` are already resolved during shim generation in Go, so those host paths should be normalized in Go before rendering the final `--mount` argument.
   - *Alternative considered:* Recompute all host paths in the shell script at runtime. Rejected because it broadens current behavior and needlessly moves option resolution out of the Go code path.

4. **Disable MSYS automatic argument conversion for the final `docker run` invocation.**
   - Container-side paths such as `dst=/work`, `dst=/home/<user>/...`, and `-w /work` must reach Docker unchanged.
   - The shim should invoke Docker with `MSYS2_ARG_CONV_EXCL='*'` so Git Bash/MSYS does not rewrite Docker arguments that merely look like Unix paths.
   - This environment variable is ignored outside MSYS, so the same script remains valid in `WSL`.
   - *Alternative considered:* Rely on MSYS automatic conversion. Rejected because it can rewrite container-side paths and mixed bind-mount arguments in ways the shim does not control.

5. **Keep container-side targets unchanged.**
   - `/work` remains the mount target and workdir for `--mount-cwd`.
   - Home and XDG destinations remain POSIX container paths derived from the image user label.
   - The design problem is host-path serialization, not Linux container layout.
   - *Alternative considered:* Change container-side targets specifically for Windows hosts. Rejected because the images remain Linux containers and the current target paths are already coherent inside the container.

## Risks / Trade-offs

- **[Git Bash conversion dependency]** The design depends on `cygpath`/MSYS behavior in `git-bash`. → Mitigation: keep the runtime conversion helper minimal and cover the expected output in tests.
- **[`--mount` behavior drift]** Switching from `-v` to `--mount` can change corner-case behavior if source creation is not preserved. → Mitigation: include `bind-create-src` where current behavior depends on Docker creating missing directories.
- **[Environment-specific divergence]** `git-bash` and `WSL` need different host-path handling. → Mitigation: isolate environment detection and path normalization into explicit helpers instead of scattering formatting logic.
- **[Generated shim portability]** Home and XDG host paths are still captured from the machine that generates the shim. → Mitigation: keep this as the current contract and document the intended local-generation workflow rather than broadening scope in this change.

## Migration Plan

- Update shim command generation to build bind mounts through a dedicated renderer instead of inline `-v` strings.
- Add a small runtime helper in the emitted shell script for cwd path normalization in `git-bash`.
- Update tests to assert Windows-specific mount rendering, MSYS conversion guard behavior, and unchanged Unix behavior.
- Update README to document supported Windows shell environments (`git-bash`, `WSL`) and explicitly exclude native `PowerShell`/`cmd.exe`.
- No data migration is required; rollback is a code revert to the previous shim rendering logic.

## Open Questions

- None at design time. If implementation reveals that `MSYS2_ARG_CONV_EXCL='*'` is broader than necessary, narrowing it to specific argument prefixes can be handled without changing the overall design.
