## ADDED Requirements

### Requirement: GitHub shorthand defaults validation

When `--package` uses the `github:` shorthand, the system SHALL validate the
shorthand only when defaults are required for `--image` or `--bin`. If the
shorthand cannot be parsed into `<owner>/<repo>` (optionally with `#<ref>`), the
command SHALL fail with an error indicating an invalid `github:` package spec.

#### Scenario: Invalid github shorthand with defaults
- **WHEN** `--package` is `github:acme` and `--image` is omitted
- **THEN** the command fails with an error indicating an invalid `github:` package spec

#### Scenario: Explicit image and bin bypass github validation
- **WHEN** `--package` is `github:acme` and `--image` is `acme/eslint` and `--bin` is `eslint`
- **THEN** image name is `acme/eslint`
- **AND** binary name is `eslint`

## MODIFIED Requirements

### Requirement: Default image from package

If `--image` is omitted, the system SHALL derive the image name from `--package`
after removing any version/tag suffix. For `github:` shorthand, the system SHALL
drop the `github:` prefix, remove any `#<ref>` suffix, and use the remaining
`<owner>/<repo>` as the derived image name.

#### Scenario: Plain package name with version tag
- **WHEN** `--package` is `eslint@latest` and `--image` is omitted
- **THEN** image name is `cli/eslint`
- **AND** a warning is emitted to stderr

#### Scenario: Scoped package name with version
- **WHEN** `--package` is `@acme/eslint@1.2.3` and `--image` is omitted
- **THEN** image name is `cli/acme/eslint`
- **AND** a warning is emitted to stderr

#### Scenario: GitHub shorthand without ref
- **WHEN** `--package` is `github:acme/eslint` and `--image` is omitted
- **THEN** image name is `cli/acme/eslint`

#### Scenario: GitHub shorthand with ref
- **WHEN** `--package` is `github:acme/eslint#v1.2.3` and `--image` is omitted
- **THEN** image name is `cli/acme/eslint`
- **AND** a warning is emitted to stderr

### Requirement: Default bin from package

If `--bin` is omitted, the system SHALL derive the binary name from `--package`
by removing any version/tag suffix and then using the package name without
scope. For `github:` shorthand, the system SHALL remove any `#<ref>` suffix and
use the repository name as the derived binary.

#### Scenario: Plain package name with version
- **WHEN** `--package` is `eslint@latest` and `--bin` is omitted
- **THEN** binary name is `eslint`
- **AND** a warning is emitted to stderr

#### Scenario: Scoped package name with version
- **WHEN** `--package` is `@acme/eslint@1.2.3` and `--bin` is omitted
- **THEN** binary name is `eslint`
- **AND** a warning is emitted to stderr

#### Scenario: GitHub shorthand without ref
- **WHEN** `--package` is `github:acme/eslint` and `--bin` is omitted
- **THEN** binary name is `eslint`

#### Scenario: GitHub shorthand with ref
- **WHEN** `--package` is `github:acme/eslint#v1.2.3` and `--bin` is omitted
- **THEN** binary name is `eslint`
- **AND** a warning is emitted to stderr

### Requirement: Image origin labels

The system SHALL set image labels on the built image for the originating package
and bin. The system SHALL set `io.cli2docker.package` and `io.cli2docker.bin`
for every build, and SHALL set `io.cli2docker.package-version` only when an
explicit version is present in `--package`. For `github:` shorthand with `#ref`,
the `io.cli2docker.package-version` label SHALL contain the git ref value.

#### Scenario: Package without explicit version
- **WHEN** `--package` is `eslint` and `cli2docker build` completes
- **THEN** the image labels include `io.cli2docker.package=eslint` and `io.cli2docker.bin=eslint`
- **AND** the image label `io.cli2docker.package-version` is not set

#### Scenario: Package with explicit version
- **WHEN** `--package` is `@acme/eslint@1.2.3`, `--bin` is `eslint-cli`, and `cli2docker build` completes
- **THEN** the image labels include `io.cli2docker.package=@acme/eslint`, `io.cli2docker.package-version=1.2.3`, and `io.cli2docker.bin=eslint-cli`

#### Scenario: GitHub package with explicit ref
- **WHEN** `--package` is `github:acme/eslint#v1.2.3`, `--bin` is `eslint`, and `cli2docker build` completes
- **THEN** the image labels include `io.cli2docker.package=github:acme/eslint`, `io.cli2docker.package-version=v1.2.3`, and `io.cli2docker.bin=eslint`

#### Scenario: Dockerfile contains origin labels
- **WHEN** `cli2docker build` generates a Dockerfile for `--package` `eslint` with derived `--bin`
- **THEN** the Dockerfile includes a `LABEL` instruction containing `io.cli2docker.package=eslint` and `io.cli2docker.bin=eslint`
- **AND** the Dockerfile does not include `io.cli2docker.package-version`
