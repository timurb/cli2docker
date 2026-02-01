## Why

Users need a copy-pasteable shim run command right after `cli2docker build` without hunting through logs or re-deriving image names. Printing it by default makes the build flow self-contained and reduces user error.

## What Changes

- `cli2docker build` will print the shim command to stdout on the normal build path.
- `--print-dockerfile` remains a separate output mode that prints only the Dockerfile to stdout.

## Capabilities

### New Capabilities

<!-- None -->

### Modified Capabilities

- `build`: Standard build output now includes the shim command on stdout.

## Impact

- CLI output changes for standard builds; downstream scripts that parse stdout may need updates.
- Tests covering build output and CLI UX will need updates to reflect the new stdout content.

## Non-goals

- Changing shim behavior or flags.
- Changing Dockerfile output behavior for `--print-dockerfile`.
- Adding new subcommands or output formats.

## Risks / Tradeoffs

- Writing the shim command to stdout risks contaminating machine-consumed output; this is accepted to keep the command copyable without extra flags.
