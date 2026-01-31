## MODIFIED Requirements

### Requirement: Default image from package

If `--image` is omitted, the system SHALL derive the image name from `--package`.

#### Scenario: Plain package name
- **WHEN** `--package` is `eslint` and `--image` is omitted
- **THEN** image name is `cli/eslint`
- **AND** a warning is emitted to stderr

#### Scenario: Scoped package name
- **WHEN** `--package` is `@acme/eslint` and `--image` is omitted
- **THEN** image name is `cli/acme/eslint`
- **AND** a warning is emitted to stderr

### Requirement: Explicit values override defaults

If `--image` is provided explicitly, the system SHALL use the explicit value.

#### Scenario: Explicit image overrides prefix
- **WHEN** `--image` is `acme/eslint` and `--image-prefix` is set
- **THEN** image name is `acme/eslint`
- **AND** a warning is emitted that the prefix was ignored

## ADDED Requirements

### Requirement: Image prefix flag

The system SHALL accept `--image-prefix` to prefix the derived image name when `--image` is omitted.

#### Scenario: Prefix applied to derived image
- **WHEN** `--package` is `eslint`, `--image` is omitted, and `--image-prefix` is `cli/`
- **THEN** image name is `cli/eslint`

#### Scenario: Prefix applied to scoped derived image
- **WHEN** `--package` is `@acme/eslint`, `--image` is omitted, and `--image-prefix` is `cli/`
- **THEN** image name is `cli/acme/eslint`

#### Scenario: Default prefix
- **WHEN** `--package` is `eslint`, `--image` is omitted, and `--image-prefix` is not set
- **THEN** image name is `cli/eslint`
