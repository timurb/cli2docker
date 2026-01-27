# ADR 0003: Shim печатается в stdout

## Status
Implemented

## Context
В требованиях есть “OS level alias or shim”. Нужно решить: ставить shim автоматически или печатать его в stdout.

## Decision
Команда `cli2docker shim` печатает shim‑скрипт в stdout. Пользователь сам решает, куда сохранить и как назвать файл.

## Consequences
- Максимальный контроль у пользователя.
- Требуется явное перенаправление и `chmod +x`.

## Alternatives
- Автоматическая установка в `~/.local/bin`.
