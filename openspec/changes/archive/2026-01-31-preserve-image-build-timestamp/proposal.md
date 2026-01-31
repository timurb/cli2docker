## Why

Builds from unversioned packages (for example `eslint` without a tag) are hard to trace to the exact upstream version. Capturing a stable build timestamp in image metadata provides the minimal provenance needed to later map the build to the package version that was resolved at that time.

## What Changes

- Capture a single UTC build timestamp at the start of `cli2docker build` and use it consistently for the build.
- Embed the timestamp as an image label, and include it in the generated Dockerfile when `--print-dockerfile` is used.
- Keep the existing CLI interface unchanged (no new flags).

## Non-goals

- Determining or labeling the resolved package version for unversioned builds.
- Implementing reproducible builds or `SOURCE_DATE_EPOCH` support.
- Changing Docker image `Created` metadata or registry timestamps.

## Capabilities

### New Capabilities
- (none)

### Modified Capabilities
- `build`: add a requirement to record a build timestamp label for images and print-only Dockerfiles.

## Impact

- `cli2docker build` workflow and Dockerfile generation (labels).
- Build-related tests and fixtures.
