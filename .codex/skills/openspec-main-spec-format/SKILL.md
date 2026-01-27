---
name: openspec-main-spec-format
description: Enforce the main OpenSpec spec format with overview, scope, interfaces, and requirements.
license: MIT
compatibility: OpenSpec projects
metadata:
  author: local
  version: "1.0"
---

When creating or updating capability docs under `openspec/specs/<capability>/`, use a split-file format:

1) `spec.md` (normative tests only)
   - Purpose (1–3 sentences)
   - Requirements/Scenarios (normative)
   - Links to sibling docs (markdown links)

2) `feature.md` (black + white box in one file)
   - Overview / Context / Problems addressed
   - Goals / Non-Goals (or In scope / Out of scope)
   - Usage workflows / scenarios beyond tests
   - Success criteria (if applicable)
   - Architecture description
   - Constraints
   - Entities
   - Interfaces (Inputs/Outputs)
   - Invariants (if applicable)
   - Events / Triggers
   - Links to sibling docs (markdown links)

Cross-links are recommended in all files:
- Link to sibling docs within the same capability
- Link across capabilities when concepts are shared
Use markdown links (e.g., `[feature](feature.md)`), not backticked filenames.

Keep `spec.md` lean; requirements and scenarios are the normative tests.
