## Why

`docs/feature-03-shim.md` описывает поведение команды `shim`, но в OpenSpec нет соответствующей спеки. Нужно перенести требования в OpenSpec и сделать его источником истины.

## What Changes

- Преобразовать `feature-03-shim` в OpenSpec требования/сценарии.
- Создать capability `shim` в `openspec/specs/shim/spec.md`.

## Capabilities

### New Capabilities
- `shim`: генерация shim‑скрипта и правила монтирования.

### Modified Capabilities
None.

## Non-goals

- Не менять реализацию `cli2docker shim`.
- Не конвертировать другие фичи или ADR.
- Не править `docs/feature-03-shim.md` на этом шаге.

## Impact

- `openspec/changes/sync-feature-03-shim/specs/shim/spec.md`.
- `openspec/specs/shim/spec.md` (после sync).
