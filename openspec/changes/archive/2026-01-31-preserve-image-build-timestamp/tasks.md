## 1. Build timestamp plumbing

- [x] 1.1 Add a new build timestamp label constant (`io.cli2docker.build-timestamp`) alongside existing label keys.
- [x] 1.2 Capture a single UTC RFC3339 timestamp at the start of `cli2docker build` and pass it through the build configuration.

## 2. Dockerfile and label generation

- [x] 2.1 Include the build timestamp label in the generated Dockerfile `LABEL` instruction.
- [x] 2.2 Ensure the same captured timestamp is used for image labels during `docker build` and for `--print-dockerfile`.

## 3. Tests and documentation

- [x] 3.1 Update build tests to assert the new label is present and matches RFC3339 UTC formatting.
- [x] 3.2 Update README label examples to include `io.cli2docker.build-timestamp`.
