## 1. Spec Parsing & Defaults

- [x] 1.1 Add helper(s) to parse `github:owner/repo[#ref]` and return owner, repo, ref, with validation errors.
- [x] 1.2 Extend `imageFromPackage` to derive `owner/repo` (strip `#ref`) for `github:` specs.
- [x] 1.3 Extend `packageBaseName` to derive `repo` (strip `#ref`) for `github:` specs.
- [x] 1.4 Apply validation only when defaults are needed (when `--image`/`--bin` omitted).

## 2. Labels & Metadata

- [x] 2.1 Extend package/version parsing for labels to handle `github:` specs and store `#ref` in `io.cli2docker.package-version`.
- [x] 2.2 Update Dockerfile label generation to use `github:owner/repo` as `io.cli2docker.package` when applicable.

## 3. Tests

- [x] 3.1 Add unit tests for default derivation with `github:owner/repo` and `github:owner/repo#ref` for both image and bin.
- [x] 3.2 Add tests for invalid github shorthand when defaults are required.
- [x] 3.3 Add tests for label generation with `github:` specs including `#ref`.

## 4. Docs

- [x] 4.1 Update README with supported `github:` shorthand, examples, and limitations.
