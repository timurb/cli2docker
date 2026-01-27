## 1. Build defaults from package

- [x] 1.1 Add helper functions to derive image/bin from package (`packageBaseName`, `imageFromPackage`).
- [x] 1.2 Remove `--image` and `--bin` from required flags in `newBuildCmd`.
- [x] 1.3 Apply defaults in `RunE` before `buildWithOptions`, emit warnings to stderr when defaults are used.

## 2. Tests

- [x] 2.1 Add unit tests for package name derivation helpers (plain and scoped cases).

## 3. Verify

- [x] 3.1 Run `go test ./...`.
