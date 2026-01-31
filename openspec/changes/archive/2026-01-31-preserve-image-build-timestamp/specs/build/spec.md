## MODIFIED Requirements

### Requirement: Image origin labels

The system SHALL set image labels on the built image for the originating package and bin. The system SHALL set `io.cli2docker.package` and `io.cli2docker.bin` for every build, and SHALL set `io.cli2docker.package-version` only when an explicit version is present in `--package`. The system SHALL set `io.cli2docker.build-timestamp` for every build to a single RFC3339 UTC timestamp captured at the start of `cli2docker build`.

#### Scenario: Package without explicit version
- **WHEN** `--package` is `eslint` and `cli2docker build` completes
- **THEN** the image labels include `io.cli2docker.package=eslint` and `io.cli2docker.bin=eslint`
- **AND** the image label `io.cli2docker.package-version` is not set
- **AND** the image label `io.cli2docker.build-timestamp` is set to an RFC3339 UTC timestamp

#### Scenario: Package with explicit version
- **WHEN** `--package` is `@acme/eslint@1.2.3`, `--bin` is `eslint-cli`, and `cli2docker build` completes
- **THEN** the image labels include `io.cli2docker.package=@acme/eslint`, `io.cli2docker.package-version=1.2.3`, and `io.cli2docker.bin=eslint-cli`
- **AND** the image label `io.cli2docker.build-timestamp` is set to an RFC3339 UTC timestamp

#### Scenario: Dockerfile contains origin labels
- **WHEN** `cli2docker build` generates a Dockerfile for `--package` `eslint` with derived `--bin`
- **THEN** the Dockerfile includes a `LABEL` instruction containing `io.cli2docker.package=eslint` and `io.cli2docker.bin=eslint`
- **AND** the Dockerfile does not include `io.cli2docker.package-version`
- **AND** the Dockerfile includes `io.cli2docker.build-timestamp` set to an RFC3339 UTC timestamp
