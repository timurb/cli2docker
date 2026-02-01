## MODIFIED Requirements

### Requirement: Image origin labels

The system SHALL set image labels on the built image for the originating package and bin. The system SHALL set `io.cli2docker.package`, `io.cli2docker.bin`, `io.cli2docker.build-timestamp`, and `io.cli2docker.user` for every build, and SHALL set `io.cli2docker.package-version` only when an explicit version is present in `--package`. For `github:` shorthand with `#ref`, the `io.cli2docker.package-version` label SHALL contain the git ref value. The system SHALL set `io.cli2docker.build-timestamp` for every build to a single RFC3339 UTC timestamp captured at the start of `cli2docker build`. The `io.cli2docker.user` label SHALL be set to the effective runtime user: `root` when `--no-user` is set, otherwise the value of `--user`.

#### Scenario: Package without explicit version
- **WHEN** `--package` is `eslint` and `cli2docker build` completes
- **THEN** the image labels include `io.cli2docker.package=eslint`, `io.cli2docker.bin=eslint`, and `io.cli2docker.user=<user>`
- **AND** the image label `io.cli2docker.package-version` is not set
- **AND** the image label `io.cli2docker.build-timestamp` is set to an RFC3339 UTC timestamp

#### Scenario: No user build
- **WHEN** `--no-user` is set and `cli2docker build` completes
- **THEN** the image label `io.cli2docker.user` is set to `root`
