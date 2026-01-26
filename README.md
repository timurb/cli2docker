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
Bootstrapped repo structure. Implementation and usage will be added next.

## Repository layout
- `specs/openspec.yaml` - product/spec description for the project.
- `.gitignore` - common ignores for scripts and tooling.
- `.editorconfig` - basic formatting rules.
- `AGENTS.md` - guidance for future contributors/agents.

## Planned workflow (placeholder)
1. Provide a CLI name and entrypoint.
2. Run a script to generate a Docker build context and image.
3. Run the resulting container like any standard CLI tool.
