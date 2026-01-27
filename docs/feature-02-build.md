# Фича 2: Команда build

## Status
Implemented

## Summary
Собирает Docker‑образ для npm‑CLI на основе временного Dockerfile.

## In scope
- Сборка образа из npm‑пакета.
- Генерация Dockerfile с `ENTRYPOINT`.
- Базовые настройки образа (base image, user, cache).

## Out of scope
- Multi‑arch сборка.
- Публикация в registry/подпись образов.
- Генерация Dockerfile для произвольных приложений.

## Interface
Обязательные флаги:
- `--package` — npm пакет;
- `--bin` — имя CLI бинарника;
- `--image` — имя образа.

Опциональные флаги:
- `--tag` — тег (по умолчанию `latest`);
- `--base` — базовый образ (по умолчанию `node:20-alpine`);
- `--user` — runtime‑пользователь (по умолчанию `node`);
- `--no-user` — не добавлять `USER` в образ;
- `--no-cache` — отключить build cache.

## Behavior
- `newBuildCmd` регистрирует флаги и required‑поля.
- `buildWithOptions` формирует ссылку на образ, создаёт temp‑директорию, пишет Dockerfile и запускает `docker build`.
- Dockerfile: `FROM <base>`, `npm install -g <package>`, `ENTRYPOINT ["<bin>"]`, опционально `USER <user>`.

## Constraints
- Требуется работающий Docker daemon.
- npm‑пакет должен публиковать указанный бинарник.

## Errors
- Нет `docker` в `PATH`.
- Ошибка `docker build` (сеть/реестр/пакет/daemon).
- Неверные флаги или пустые обязательные параметры.

## Security
- По умолчанию используется `node:20-alpine` и `USER node`.
- `--no-user` отключает понижение привилегий.

## Examples
```bash
cli2docker build --package eslint --bin eslint --image acme/eslint
```

```bash
cli2docker build --package eslint --bin eslint --image acme/eslint --tag v1
```

```bash
cli2docker build --package eslint --bin eslint --image acme/eslint --no-user
```

## ADR
- нет
