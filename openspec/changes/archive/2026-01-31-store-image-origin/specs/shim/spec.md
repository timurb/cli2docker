## ADDED Requirements

### Requirement: Shim includes origin metadata
The system SHALL include the image origin metadata as comments in the generated shim output when the image labels are present. The metadata SHALL use the same label keys as the image (`io.cli2docker.package`, `io.cli2docker.package-version`, `io.cli2docker.bin`).

#### Scenario: Shim includes package and bin
- **WHEN** `cli2docker shim --image <image>` is run and the image labels include `io.cli2docker.package=eslint` and `io.cli2docker.bin=eslint`
- **THEN** the shim output includes comment lines with `io.cli2docker.package=eslint` and `io.cli2docker.bin=eslint`

#### Scenario: Shim includes explicit version
- **WHEN** `cli2docker shim --image <image>` is run and the image labels include `io.cli2docker.package=@acme/eslint`, `io.cli2docker.package-version=1.2.3`, and `io.cli2docker.bin=eslint-cli`
- **THEN** the shim output includes comment lines with `io.cli2docker.package=@acme/eslint`, `io.cli2docker.package-version=1.2.3`, and `io.cli2docker.bin=eslint-cli`
