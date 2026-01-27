## Why

`docs/feature-01-cli.md` описывает базовый CLI‑контур (`cli2docker`, `build`, `shim`), но эти требования не зафиксированы в OpenSpec. Нужно перенести их в спецификацию, чтобы OpenSpec отражал текущее поведение CLI.

## What Changes

- Преобразовать требования из `feature-01-cli` в OpenSpec‑спеку.
- Создать новый capability `cli` в `openspec/specs/cli/spec.md`.

## Capabilities

### New Capabilities
- `cli`: базовый CLI‑контур с корневой командой и сабкомандами.

### Modified Capabilities
None.

## Non-goals

- Не менять поведение `cli2docker`.
- Не затрагивать спецификации `build` и `shim` (они синхронизируются отдельно).
- Не править `docs/feature-01-cli.md` на этом шаге.

## Impact

- `openspec/changes/sync-feature-01-cli/specs/cli/spec.md`.
- `openspec/specs/cli/spec.md` (после sync).
