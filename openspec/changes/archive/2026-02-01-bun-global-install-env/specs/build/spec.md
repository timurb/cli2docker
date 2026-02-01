## MODIFIED Requirements

### Requirement: Dockerfile content

The generated Dockerfile SHALL include: `FROM <base>`, an install command based on the selected package manager, `ENTRYPOINT ["<bin>"]`, and `USER <user>` unless `--no-user` is set. When `--package-manager` is `bun`, the Dockerfile SHALL set Bun global install environment variables to ensure global packages are installed outside the user home directory.

#### Scenario: Npm install
- **WHEN** `--package-manager` is omitted
- **THEN** the Dockerfile includes `npm install -g <package>`

#### Scenario: Bun install
- **WHEN** `--package-manager` is `bun`
- **THEN** the Dockerfile includes `bun add -g <package>`

#### Scenario: Bun global install env
- **WHEN** `--package-manager` is `bun`
- **THEN** the Dockerfile includes `BUN_INSTALL_GLOBAL_DIR=/usr/local/bun/global`
- **AND** the Dockerfile includes `BUN_INSTALL_BIN=/usr/local/bin`

#### Scenario: Default user
- **WHEN** `--no-user` is not set
- **THEN** the Dockerfile includes `USER <user>`

#### Scenario: No user
- **WHEN** `--no-user` is set
- **THEN** the Dockerfile does not include a `USER` instruction
