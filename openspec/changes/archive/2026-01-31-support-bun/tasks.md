## 1. CLI flags and defaults

- [x] 1.1 Add `PackageManager` to `buildFlags` and wire `--package-manager` flag with `npm`/`bun` help text.
- [x] 1.2 Validate `--package-manager` early in `build` RunE and return a clear error on invalid values.
- [x] 1.3 Apply runtime-aware defaults for `--base` and `--user` when omitted (`node:20-alpine`/`node` for npm, `oven/bun:1`/`bun` for Bun).

## 2. Dockerfile generation

- [x] 2.1 Branch Dockerfile content by package manager (npm env + `npm install -g`, bun uses `bun add -g`).
- [x] 2.2 Keep labels and entrypoint unchanged across package managers; honor `--no-user`.

## 3. Tests

- [x] 3.1 Add tests for package manager parsing/defaults and invalid values.
- [x] 3.2 Add Dockerfile content tests covering npm vs Bun install lines and default base/user behavior.

## 4. Documentation

- [x] 4.1 Update README usage to include `--package-manager` and a Bun example.
