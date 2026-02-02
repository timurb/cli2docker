## 1. Shim metadata output

- [x] 1.1 Extend `originCommentLines` to emit `io.cli2docker.build-timestamp` when present on the image labels.
- [x] 1.2 Ensure shim script output includes the timestamp comment in the correct header section.

## 2. Tests

- [x] 2.1 Update `TestOriginCommentLines` to cover build timestamp emission.
- [x] 2.2 Update `TestBuildShimScriptIncludesOriginComments` to assert the build timestamp comment appears when labeled.
