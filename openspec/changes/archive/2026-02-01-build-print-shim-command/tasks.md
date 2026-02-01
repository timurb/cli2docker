## 1. Spec and API alignment

- [x] 1.1 Confirm stdout-only shim command behavior matches specs (no output in `--print-dockerfile` path)
- [x] 1.2 Decide whether to prefix the shim line (resolve open question in design)

## 2. Build flow changes

- [x] 2.1 Emit shim command after successful build using resolved image ref
- [x] 2.2 Ensure shim line is the final stdout line in the standard build path

## 3. Tests and verification

- [x] 3.1 Update/extend build output tests to assert shim command line
- [x] 3.2 Add test coverage for tag resolution in shim command output
