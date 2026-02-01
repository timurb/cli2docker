## 1. Build labels

- [x] 1.1 Add `io.cli2docker.user` to image labels in the Dockerfile render.
- [x] 1.2 Set `io.cli2docker.user` to `root` when `--no-user` is set, otherwise to `--user`.
- [x] 1.3 Update build label tests to assert `io.cli2docker.user`.

## 2. Shim user resolution

- [x] 2.1 Require `io.cli2docker.user` label when generating shims and surface a clear error if missing.
- [x] 2.2 Use `io.cli2docker.user` to resolve container home for `--mount-home` and XDG mounts.

## 3. XDG mounts

- [x] 3.1 Implement XDG mount path derivation and validation under the labeled home.
- [x] 3.2 Add tests for mount targets with labeled users (`bun`, `root`).
