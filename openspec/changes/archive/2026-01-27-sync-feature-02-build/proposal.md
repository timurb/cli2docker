## Why

Текущая документация `docs/feature-02-build.md` описывает поведение команды `build`, но OpenSpec‑спека покрывает лишь часть требований (дефолты по `--package`). Нужно синхронизировать требования в OpenSpec, чтобы спецификация была источником истины.

## What Changes

- Преобразовать содержание `docs/feature-02-build.md` в требования/сценарии OpenSpec.
- Обновить существующую спецификацию `build` в `openspec/specs/build/spec.md` через delta‑spec.

## Capabilities

### New Capabilities
None.

### Modified Capabilities
- `build`: добавить требования по интерфейсу флагов, поведению, ограничениям и ошибкам из `feature-02-build`.

## Non-goals

- Не менять реализацию `cli2docker build`.
- Не синхронизировать другие фичи или ADR.
- Не править `docs/feature-02-build.md` на этом шаге.

## Impact

- `openspec/changes/sync-feature-02-build/specs/build/spec.md` (delta‑spec).
- `openspec/specs/build/spec.md` (после sync).
