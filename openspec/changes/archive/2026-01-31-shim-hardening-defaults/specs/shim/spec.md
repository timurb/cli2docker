## ADDED Requirements

### Requirement: Shim hardening defaults
The system SHALL include `--cap-drop=ALL`, `--security-opt=no-new-privileges`, and `--read-only` in the shim `docker run` command by default.

#### Scenario: Default hardening flags
- **WHEN** `cli2docker shim --image <image>` is run without opt-out flags
- **THEN** the shim includes `--cap-drop=ALL`, `--security-opt=no-new-privileges`, and `--read-only`

### Requirement: Opt-out flags for hardening
The system SHALL provide explicit opt-out flags that remove the default hardening flags from the shim.

#### Scenario: Disable cap drop
- **WHEN** `cli2docker shim --image <image> --no-drop-caps` is run
- **THEN** the shim does not include `--cap-drop=ALL`

#### Scenario: Disable no-new-privileges
- **WHEN** `cli2docker shim --image <image> --no-new-privileges` is run
- **THEN** the shim does not include `--security-opt=no-new-privileges`

#### Scenario: Disable read-only rootfs
- **WHEN** `cli2docker shim --image <image> --no-read-only` is run
- **THEN** the shim does not include `--read-only`

### Requirement: Read-only warning
The system SHALL emit a warning to stderr stating that read-only mode is experimental and indicating the opt-out flag.

#### Scenario: Warning on default read-only
- **WHEN** `cli2docker shim --image <image>` is run
- **THEN** stderr contains a warning that read-only mode is experimental and can be disabled with `--no-read-only`
