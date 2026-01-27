# ADR 0005: Go CLI парсинг через spf13/cobra

## Status
Implemented

## Context
Go CLI использует сабкоманды (`build`, `shim`) и требует удобного парсинга флагов и help/usage.

## Decision
Использовать `github.com/spf13/cobra` для парсинга CLI.

## Consequences
- Появляется внешняя зависимость.
- Поведение help/usage соответствует типичному Go CLI с сабкомандами.

## Alternatives
- Стандартный `flag` (быстро, но неудобно для сабкоманд).
- Другие библиотеки (например, `urfave/cli/v2`).
