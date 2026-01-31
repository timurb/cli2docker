## 1. Core derivation logic

- [x] 1.1 Add helper to strip version/tag suffix from package specs (last `@` when not leading scope).
- [x] 1.2 Apply version stripping in `packageBaseName` and `imageFromPackage` while leaving install package unchanged.
- [x] 1.3 Confirm explicit `--image`, `--bin`, and `--tag` flows remain unaffected.

## 2. Tests and verification

- [x] 2.1 Extend unit tests for `packageBaseName` to include versioned specs (plain and scoped).
- [x] 2.2 Extend unit tests for `imageFromPackage` to include versioned specs (plain and scoped).
- [x] 2.3 Run `go test ./...` and verify scenarios align with the spec updates.
- [x] 2.4 Add test covering warning emission when deriving defaults from versioned packages.
