## MODIFIED Requirements

### Requirement: Mount home path

If `--mount-home` is set, the system SHALL mount a path from `$HOME` to the container user's home directory read-only by default. The container user's home directory SHALL be `/root` when the configured image user is `root`, otherwise `/home/<user>`. If the image user is empty, the system SHALL default to `/home/node`.

#### Scenario: Mount home read-only for bun
- **WHEN** `--mount-home` is set and `--mount-home-rw` is not set and the image user is `bun`
- **THEN** the shim includes a `:ro` mount targeting `/home/bun/<relative>` for the resolved path
