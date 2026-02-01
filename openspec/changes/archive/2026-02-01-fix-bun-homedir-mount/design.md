## Context

`cli2docker shim --mount-home` currently resolves the container path using a hard-coded `/home/node`, which breaks images that run as a different user (notably Bun with user `bun` and home `/home/bun`). The shim already inspects image metadata via `docker image inspect` to read labels, so the runtime user is available without new external dependencies.

## Goals / Non-Goals

**Goals:**
- Resolve the mount target path using the image's configured runtime user so Bun home mounts land under `/home/bun`.
- Preserve current behavior for Node images and existing validation rules for `--mount-home`.

**Non-Goals:**
- Add new flags or configuration options for mount behavior.
- Change build defaults or runtime user selection.
- Validate or enforce home directory contents inside the image.

## Decisions

- **Derive container user from image config (`docker image inspect` / `Config.User`).**
  - *Alternative:* Add a new build-time label for user and read it in the shim. Rejected to avoid new labels and backfilling older images.
  - *Alternative:* Infer from base image or package manager (bun/npm). Rejected as brittle and incorrect for custom images.

- **Map user to home path as `/root` for `root`, `/home/<user>` otherwise; fallback to `/home/node` when user is empty.**
  - *Alternative:* Read `/etc/passwd` from the image to get the actual home directory. Rejected due to higher complexity and extra image access.

## Risks / Trade-offs

- **Images with non-standard home directories or numeric users may resolve incorrectly** → Accept for now; document a follow-up option to inspect `/etc/passwd` or add an override flag if this becomes an issue.
- **Extra image metadata access** → Reuse existing `docker image inspect` call to avoid additional commands.
