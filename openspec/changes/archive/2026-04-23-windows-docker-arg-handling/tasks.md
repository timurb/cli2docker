## 1. Mount rendering foundation

- [x] 1.1 Replace inline `-v` mount formatting in shim generation with shared helpers that emit bind mounts for cwd, home, and XDG paths.
- [x] 1.2 Normalize generation-time Windows host paths for `--mount-home` and `--mount-xdg` into Docker-compatible source paths while preserving existing `$HOME` validation.

## 2. Runtime shim behavior

- [x] 2.1 Update the generated POSIX shim script to normalize runtime cwd paths for `git-bash` while leaving `WSL` and Unix-like environments unchanged.
- [x] 2.2 Guard the `docker run` invocation from MSYS argument rewriting so container-side paths such as `/work` and `/home/<user>` reach Docker unchanged.

## 3. Verification and documentation

- [x] 3.1 Extend shim tests to cover cwd, home, and XDG mount rendering for Unix, `git-bash`, and `WSL`-compatible behavior.
- [x] 3.2 Update README and relevant CLI help text to document supported Windows environments (`git-bash`, `WSL`) and explicitly exclude native `PowerShell` and `cmd.exe`.
