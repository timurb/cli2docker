## 1. CLI surface and flow control

- [x] 1.1 Add `--print-dockerfile` to `build` flags and `buildFlags`.
- [x] 1.2 Update `runBuild` to apply defaults, then short-circuit to print-only mode without calling Docker.
- [x] 1.3 Ensure Docker availability is checked only when a build is executed.

## 2. Dockerfile rendering pipeline

- [x] 2.1 Introduce `renderDockerfile(opts)` to return the Dockerfile content as a string.
- [x] 2.2 Update `writeDockerfile` to use `renderDockerfile`.
- [x] 2.3 Use `renderDockerfile` for print-only output to stdout.

## 3. Tests and docs

- [x] 3.1 Add tests for print-only mode (stdout contains Dockerfile, no `docker` check, no build run).
- [x] 3.2 Add test coverage for stderr warnings in print-only mode when defaults are derived.
- [x] 3.3 Update README/build usage to document `--print-dockerfile`.
