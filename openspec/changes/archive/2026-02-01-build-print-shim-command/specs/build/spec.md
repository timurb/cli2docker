## ADDED Requirements

### Requirement: Build prints shim command

The system SHALL print a shim command to stdout after a successful build. The shim command SHALL reference the resolved image reference used for the build, including any resolved tag. All other build output SHALL be written to stderr.

#### Scenario: Shim command uses resolved image
- **WHEN** `build` is executed with `--image acme/eslint` and `--tag v1` and `--print-dockerfile` is not set
- **THEN** stdout includes `cli2docker shim --image acme/eslint:v1`

#### Scenario: Build output separation
- **WHEN** `build` is executed successfully and `--print-dockerfile` is not set
- **THEN** stdout contains only the shim command line
- **AND** stderr includes build status output
