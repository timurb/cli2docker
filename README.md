# cli2docker

**Disclaimer:** this code was mostly written using ChatGPT Codex 5.2

Quickly package Node.js CLI tools into Docker images without hand-writing a Dockerfile each time.

## Goals
- Provide a simple script-first workflow to build minimal images for Node.js CLIs.
- Keep usage fast and repeatable for multiple tools.
- Avoid boilerplate Dockerfiles for trivial packaging needs.

## Non-goals (for now)
- Building multi-arch images or advanced build pipelines.
- Managing releases, registries, or image signing.

## Status
MVP implementation is available in Go.

## Repository layout
- `openspec/specs/openspec.yaml` - product/spec description for the project.
- `openspec/specs/overview.md` - project overview and design decisions.
- `openspec/specs/<capability>/feature.md` - capability overview/constraints (black+white box).
- `openspec/specs/<capability>/spec.md` - normative requirements and scenarios.
- `main.go` - Go CLI implementation.
- `.gitignore` - common ignores for scripts and tooling.
- `.editorconfig` - basic formatting rules.
- `AGENTS.md` - guidance for future contributors/agents.

## Docs
- Overview: `openspec/specs/overview.md`
- CLI: `openspec/specs/cli/feature.md` / `openspec/specs/cli/spec.md`
- Build: `openspec/specs/build/feature.md` / `openspec/specs/build/spec.md`
- Shim: `openspec/specs/shim/feature.md` / `openspec/specs/shim/spec.md`

## Requirements
- Docker installed and running.
- A published npm package providing a CLI binary.

## Quick start
Build an image for a Node.js CLI and tag it:

```bash
go run . build --package eslint --bin eslint --image acme/eslint
```

Derive the image name from the package with a prefix:

```bash
go run . build --package eslint --image-prefix cli/
```

Run it directly with Docker:

```bash
docker run --rm -it acme/eslint:latest --version
```

Print a shim and install it manually:

```bash
go run . shim --image acme/eslint:latest > ~/.local/bin/eslint
chmod +x ~/.local/bin/eslint
```

## Notes
- The default base image is `node:20-alpine`. Use `--base` to override it.
- By default the image drops to the `node` user for runtime. Use `--no-user` for images without that user.
- `--image-prefix` defaults to `cli/`.
- If `--image` is omitted, it is derived from `--package`
- Shim scripts are printed to stdout. Redirect to a file on your `PATH`.

## Provenance
Built images include labels for origin metadata:
`io.cli2docker.package`, `io.cli2docker.package-version` (only when explicit), and `io.cli2docker.bin`.
Inspect labels with:
`docker image inspect --format '{{json .Config.Labels}}' <image>`

Shim output includes comment lines when the labels are present:
```
# io.cli2docker.package=eslint
# io.cli2docker.package-version=1.2.3
# io.cli2docker.bin=eslint
```

## Build

```bash
go build -o cli2docker .
./cli2docker build --package eslint --bin eslint --image acme/eslint
```
