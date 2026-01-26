# node2docker

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
- `specs/openspec.yaml` - product/spec description for the project.
- `main.go` - Go CLI implementation.
- `.gitignore` - common ignores for scripts and tooling.
- `.editorconfig` - basic formatting rules.
- `AGENTS.md` - guidance for future contributors/agents.

## Requirements
- Docker installed and running.
- A published npm package providing a CLI binary.

## Quick start
Build an image for a Node.js CLI and tag it:

```bash
go run . build --package eslint --bin eslint --image acme/eslint
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
- Shim scripts are printed to stdout. Redirect to a file on your `PATH`.

## Build

```bash
go build -o node2docker-go .
./node2docker-go build --package eslint --bin eslint --image acme/eslint
```
