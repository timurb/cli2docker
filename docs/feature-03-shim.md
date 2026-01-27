# Фича 3: Команда shim

## Status
Implemented

## Summary
Генерирует shim‑скрипт в stdout для запуска Docker‑образа как CLI.

## In scope
- Печать shim в stdout.
- Поддержка монтирования `$PWD` и одного пути из `$HOME`.
- Read‑only по умолчанию для home‑пути.

## Out of scope
- Автоматическая установка shim в `PATH`.
- Множественные монтирования из `$HOME`.
- Произвольные пути вне `$HOME`.

## Interface
Обязательные флаги:
- `--image` — ссылка на образ.

Опциональные флаги:
- `--mount-cwd` — монтировать текущий каталог в `/work` и установить `-w /work`.
- `--mount-home` — смонтировать каталог из `$HOME` в `$HOME` контейнера (read‑only).
- `--mount-home-rw` — смонтировать `--mount-home` в режиме read‑write.

## Behavior
- Скрипт печатается в stdout и не имеет собственных флагов.
- `buildShimExecLine` формирует строку `docker run`.
- `resolveHomeMount` нормализует путь и строит `/home/node/<relative>`.
- `expandHomePath` поддерживает `~` и `~/`.

## Constraints
- `--mount-home` принимает только пути внутри `$HOME`.
- `--mount-home-rw` требует `--mount-home`.
- Допускается только один путь из `$HOME`.
- Точка монтирования CWD фиксирована: `/work`.
- Точка монтирования home‑пути фиксирована: `/home/node/<relative>`.

## Errors
- Нет `docker` в `PATH`.
- `--mount-home-rw` без `--mount-home`.
- Пустой `--mount-home`.
- Путь `--mount-home` вне `$HOME`.

## Security
- Монтирования выключены по умолчанию.
- `--mount-home` по умолчанию read‑only.

## Examples
Shim без монтирований:
```bash
cli2docker shim --image acme/eslint:latest > ~/.local/bin/eslint
chmod +x ~/.local/bin/eslint
eslint --version
```

Shim с текущим каталогом:
```bash
cli2docker shim --image acme/eslint:latest --mount-cwd > ~/.local/bin/eslint
chmod +x ~/.local/bin/eslint
eslint .
```

Shim с конфигом из `$HOME` (read‑only):
```bash
cli2docker shim --image acme/eslint:latest --mount-home ~/.codex > ~/.local/bin/eslint
chmod +x ~/.local/bin/eslint
eslint .
```

Shim с конфигом из `$HOME` (read‑write):
```bash
cli2docker shim --image acme/eslint:latest --mount-home ~/.codex --mount-home-rw > ~/.local/bin/eslint
chmod +x ~/.local/bin/eslint
eslint .
```

Ошибка при попытке монтировать путь вне `$HOME`:
```bash
cli2docker shim --image acme/eslint:latest --mount-home /etc
```

## ADR
- `docs/adr/0001-shim-mount-cwd.md`
- `docs/adr/0002-home-mount-policy.md`
- `docs/adr/0003-shim-stdout.md`
