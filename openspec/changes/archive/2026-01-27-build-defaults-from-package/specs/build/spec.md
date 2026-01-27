## ADDED Requirements

### Requirement: Default image from package

If `--image` is omitted, the system SHALL derive the image name from `--package`.

#### Scenario: Plain package name
- **WHEN** `--package` is `eslint` and `--image` is omitted
- **THEN** image name is `eslint`
- **AND** a warning is emitted to stderr

#### Scenario: Scoped package name
- **WHEN** `--package` is `@acme/eslint` and `--image` is omitted
- **THEN** image name is `acme/eslint`
- **AND** a warning is emitted to stderr

### Requirement: Default bin from package

If `--bin` is omitted, the system SHALL derive the binary name from `--package` by using the package name without scope.

#### Scenario: Plain package name
- **WHEN** `--package` is `eslint` and `--bin` is omitted
- **THEN** binary name is `eslint`
- **AND** a warning is emitted to stderr

#### Scenario: Scoped package name
- **WHEN** `--package` is `@acme/eslint` and `--bin` is omitted
- **THEN** binary name is `eslint`
- **AND** a warning is emitted to stderr

### Requirement: Explicit values override defaults

If `--image` or `--bin` is provided explicitly, the system SHALL use the explicit value.

#### Scenario: Explicit image and bin
- **WHEN** `--package` is `eslint` and `--image` is `acme/eslint` and `--bin` is `eslint-cli`
- **THEN** image name is `acme/eslint`
- **AND** binary name is `eslint-cli`
