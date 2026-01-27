## Context

`docs/feature-03-shim.md` описывает `cli2docker shim`, но OpenSpec пока не фиксирует эти требования. Нужно создать новую spec‑capability `shim` без изменения реализации.

## Goals / Non-Goals

**Goals:**
- Создать OpenSpec spec для `shim` на основе `feature-03-shim`.
- Зафиксировать правила монтирования и поведение флагов.

**Non-Goals:**
- Не менять поведение `cli2docker shim`.
- Не обновлять другие фичи/ADR.
- Не править документацию до архивирования change.

## Decisions

### Decision 1: New capability `shim`
Создаём `openspec/changes/sync-feature-03-shim/specs/shim/spec.md` и синкаем в `openspec/specs/shim/spec.md`.

### Decision 2: Явно фиксируем требование docker
Требование `docker` в PATH фиксируется в `shim` (а не в `cli`) как часть поведения подкоманды.

## Risks / Trade-offs

- **Risk:** Часть поведения скрыта в коде (например, формат stdout) → **Mitigation:** держим требования минимальными и привязанными к текущей документации.
- **Risk:** Спека и legacy docs разойдутся → **Mitigation:** удалить `docs/feature-03-shim.md` после архивации.
