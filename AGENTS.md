# AGENTS.md

This repo is a small, script-first toolkit to package Node.js CLI tools into Docker images without writing Dockerfiles by hand.

## Conventions
- Keep tooling minimal and dependency-light; prefer stdlib for Python and POSIX shell for Bash.
- Provide clear CLI help text and examples for any scripts you add.
- Update `README.md` and `specs/openspec.yaml` when scope or usage changes.

## Repository notes
- There are no required tests yet.
- Avoid adding heavy build systems unless they unlock clear product value.

## Suggested checks (if applicable)
- Shell: `shellcheck` on bash scripts.
- Python: `python -m pytest` if tests are introduced.
