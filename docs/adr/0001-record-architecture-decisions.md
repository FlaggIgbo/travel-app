# 0001. Record Architecture Decisions

## Metadata
- **Status**: Accepted
- **Date**: 2026-11-24
- **Deciders**: Engineering Team
- **Consulted**: Core Contributors
- **Informed**: All Team Members

---

## Context & Problem Statement
As Triangle Travel evolves with route graph algorithms, caching layers, database migrations, and AI integrations (MCP servers, evaluation benchmarks), we need a centralized, version-controlled mechanism to record key architectural decisions, rationale, trade-offs, and outcomes.

## Decision Drivers
- Need for transparent historical context behind architectural trade-offs.
- Seamless developer onboarding and architectural alignment.
- Low friction: Markdown files stored alongside code in Git.

## Considered Options
- **Option 1**: Lightweight Architecture Decision Records (ADRs) stored in `docs/adr/`.
- **Option 2**: External wiki / Google Docs.
- **Option 3**: Git commit messages / pull request descriptions only.

## Decision Outcome
Chosen Option: **Option 1 (ADRs in `docs/adr/`)**, because storing ADRs directly in version control keeps decisions synchronized with code changes, supports code reviews during PRs, and requires no external tooling.

## Consequences
### Positive Impacts
- Clear historical record of decisions and technical trade-offs.
- Encourages deliberation and structured evaluation of alternatives prior to major refactors.

### Negative Impacts / Trade-offs
- Requires discipline to keep records updated when architectural assumptions change.
