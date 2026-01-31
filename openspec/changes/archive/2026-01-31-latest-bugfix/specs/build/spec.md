## MODIFIED Requirements

### Requirement: Default image from package

If `--image` is omitted, the system SHALL derive the image name from `--package` after removing any version/tag suffix.

#### Scenario: Plain package name with version tag
- **WHEN** `--package` is `eslint@latest` and `--image` is omitted
- **THEN** image name is `cli/eslint`
- **AND** a warning is emitted to stderr

#### Scenario: Scoped package name with version
- **WHEN** `--package` is `@acme/eslint@1.2.3` and `--image` is omitted
- **THEN** image name is `cli/acme/eslint`
- **AND** a warning is emitted to stderr

### Requirement: Default bin from package

If `--bin` is omitted, the system SHALL derive the binary name from `--package` by removing any version/tag suffix and then using the package name without scope.

#### Scenario: Plain package name with version
- **WHEN** `--package` is `eslint@latest` and `--bin` is omitted
- **THEN** binary name is `eslint`
- **AND** a warning is emitted to stderr

#### Scenario: Scoped package name with version
- **WHEN** `--package` is `@acme/eslint@1.2.3` and `--bin` is omitted
- **THEN** binary name is `eslint`
- **AND** a warning is emitted to stderr
