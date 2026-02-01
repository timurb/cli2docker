## 1. Image user resolution

- [ ] 1.1 Extend image inspection to read the configured user (`Config.User`) alongside labels.
- [ ] 1.2 Add a helper to map image user to container home (`/root` for `root`, `/home/<user>` otherwise, `/home/node` when empty).

## 2. Shim mount integration

- [ ] 2.1 Thread the computed container home into home mount resolution (adjust `resolveHomeMount` usage/signature or add a wrapper).
- [ ] 2.2 Update shim exec line generation to use the computed container home for `--mount-home`.

## 3. Tests

- [ ] 3.1 Add tests for Bun user mounts resolving to `/home/bun`.
- [ ] 3.2 Add a test for empty image user falling back to `/home/node`.
