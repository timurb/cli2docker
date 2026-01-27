---
name: phase-transition
description: Auto-write phase artifacts on explicit phase change command (no chat preview).
license: MIT
compatibility: Works with OpenSpec workflows.
metadata:
  author: local
  version: "1.0"
---

When the user explicitly says "переходим на фазу <X>" (or "next phase"), generate and write the artifacts for that phase directly to files without previewing content in chat.

Behavior:
- Treat this as an explicit transition signal.
- Write artifacts immediately to their target files.
- After writing, report the file paths and a brief summary of what was written.
- Accept edits on request; do not re-preview unless asked.
