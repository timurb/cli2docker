# Project status

## Decisions
- [Implemented] ADR-0001: opt-in монтирование текущего каталога в shim (`--mount-cwd`, `/work`).
- [Implemented] ADR-0002: политика монтирования путей из `$HOME` (только внутри `$HOME`, `:ro` по умолчанию, `--mount-home-rw`).
- [Implemented] ADR-0003: shim печатается в stdout.
- [Implemented] ADR-0005: Go CLI парсинг через `spf13/cobra`.
- [Superseded] ADR-0006: прототипы в bash/python/go для сравнения.
- [Pending] build: дефолтные `--image` и `--bin` вычисляются из `--package` (без scope), с предупреждением при авто-значениях.
