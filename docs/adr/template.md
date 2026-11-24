# [ADR-XXXX] Title of Architectural Decision

## Metadata
- **Status**: [ Proposed | Accepted | Rejected | Deprecated | Superseded by ADR-XXXX ]
- **Date**: YYYY-MM-DD
- **Deciders**: [Names / Handles]
- **Consulted**: [Stakeholders / Reviewers]
- **Informed**: [Team members]

---

## Context & Problem Statement
<!-- 
Describe the context, problem, forces, and constraints driving this decision.
What problem are we solving? What are the technical or business constraints?
-->
[TODO: Explain the context and problem statement]

## Decision Drivers
<!-- 
Key requirements and non-functional priorities influencing this choice (e.g., latency, maintainability, cost).
-->
- [TODO: Driver 1, e.g., Sub-50ms graph route discovery response time]
- [TODO: Driver 2, e.g., Offline/local development ergonomics]
- [TODO: Driver 3, e.g., Ease of decoupling external provider APIs]

## Considered Options
<!-- 
List alternative approaches evaluated.
-->
- **Option 1**: [Description, pros, cons]
- **Option 2**: [Description, pros, cons]
- **Option 3**: [Description, pros, cons]

## Decision Outcome
<!-- 
State the chosen solution clearly, along with justification.
-->
Chosen Option: **[Option X]**, because:
[TODO: Succinct rationale for why this option was chosen over alternatives]

## Consequences
<!-- 
What becomes easier or harder as a result of this decision?
-->
### Positive Impacts
- [TODO: e.g., Simpler operational overhead]
- [TODO: e.g., Predictable query performance]

### Negative Impacts / Trade-offs
- [TODO: e.g., Need to implement custom caching or migration layer]

### Neutral / Operational
- [TODO: e.g., Requires developer onboarding documentation]

## Validation & Benchmarks
<!-- 
How do we verify this decision in practice? (e.g. k6 load test thresholds, unit test coverage, profiling)
-->
- [TODO: Validation criteria, e.g., Load test validates P95 < 200ms under 50 VUs]

## References & Related Artifacts
- [TODO: Links to PRs, issues, RFCs, or benchmarks]
