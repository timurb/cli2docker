## 1. CLI flags and defaults

- [x] 1.1 Add shim flags: `--no-drop-caps`, `--no-new-privileges`, `--no-read-only`.
- [x] 1.2 Wire defaults so generated shim includes `--cap-drop=ALL`, `--security-opt=no-new-privileges`, and `--read-only` unless opt-outs are set.

## 2. Shim generation behavior

- [x] 2.1 Emit stderr warning that read-only mode is experimental and can be disabled with `--no-read-only`.
- [x] 2.2 Ensure opt-out flags remove only their corresponding hardening option.

## 3. Documentation and tests

- [x] 3.1 Update README and `shim --help` output to document defaults and opt-out flags.
- [x] 3.2 Add tests covering default hardening flags, each opt-out flag, and the stderr warning.
