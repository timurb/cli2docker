---
name: openspec-main-spec-format
description: Enforce the main OpenSpec spec format with overview, scope, interfaces, and requirements.
license: MIT
compatibility: OpenSpec projects
metadata:
  author: local
  version: "1.0"
---

When creating or updating files under `openspec/specs/*/spec.md`, use this structure:

1) Overview
    This section describes context in which the developed solution operates, its blackbox description and intent of its creation.
    It is recommended to describe the following as paragraph parts or as subsections:
    - Context
    - Problems addressed
    - Goals and intents
2) Constraints
3) Goals / Non-Goals (or In scope / Out of scope)
4) Architecture
    This section describes glass box description of the developed solution.
    Depending on the application nature it can hold the following subsections (non exhaustive)
    - Entities
    - Interfaces (or Inputs and Outputs)
    - Events / Triggers
5) Additional sections (recommended but optional):
    - Invariants
    - Success criteria
    - Usage orkflows
6) Requirements (test cases with scenarios)

Keep sections concise; requirements and scenarios remain the normative tests.
