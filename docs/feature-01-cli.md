# Фича 1: CLI скрипт

## Status
Implemented

## Summary
CLI‑обёртка с подкомандами `build` и `shim`, общими help/usage и единым способом парсинга флагов.

## In scope
- Корневая команда `cli2docker`.
- Регистрация подкоманд `build` и `shim`.
- Парсинг флагов и help/usage.

## Out of scope
- Логика сборки образов (это фича 2).
- Генерация shim (это фича 3).

## Interface
- `cli2docker`
- `cli2docker build ...`
- `cli2docker shim ...`

## Behavior
- Входная точка: `main()` вызывает `newRootCmd().Execute()`.
- Подкоманды регистрируются в `newRootCmd()`.
- Парсинг флагов выполняет `spf13/cobra`.

## Constraints
- Подкоманды `build` и `shim` требуют наличия `docker` в `PATH`.

## Errors
- Неверные флаги/аргументы → ненулевой код выхода и сообщение в stderr.

## Security
- Отдельных security‑настроек нет; безопасность определяется фичами 2/3.

## Examples
```bash
cli2docker --help
cli2docker build --help
cli2docker shim --help
```

## ADR
- `docs/adr/0005-go-cli-cobra.md`
