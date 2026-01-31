## 1. Build image metadata

- [x] 1.1 Add label keys for package, package-version (explicit only), and bin in build flow
- [x] 1.2 Parse `--package` to split name and explicit version for labels
- [x] 1.3 Add tests for image label generation for versioned and unversioned packages
- [x] 1.4 Add Dockerfile label test for unversioned package

## 2. Shim provenance comments

- [x] 2.1 Read image labels for package/bin/version when generating shim
- [x] 2.2 Emit shim header comments with origin metadata when labels exist
- [x] 2.3 Add tests covering shim output with and without package-version label
- [x] 2.4 Add shim script output test covering origin comment insertion

## 3. Documentation touchpoints

- [x] 3.1 Document how to inspect labels for origin metadata
- [x] 3.2 Document shim header comment format for provenance
