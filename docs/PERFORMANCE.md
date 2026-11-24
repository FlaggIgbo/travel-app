# Performance & Latency Bottlenecks

Use this document to track throughput and latency bottlenecks identified through k6 load testing, profiling (`pprof`), or database inspection, along with the optimizations implemented and verification results.

---

## Summary Log

| ID | Date | Subsystem / Endpoint | Baseline (P95 / RPS) | Bottleneck Root Cause | Fix / Optimization | Post-Fix (P95 / RPS) | Status |
|----|------|----------------------|----------------------|-----------------------|--------------------|----------------------|--------|
| *PB-01* | *YYYY-MM-DD* | *e.g., GET /api/routes* | *e.g., 420ms @ 50 RPS* | *e.g., Missing index on route destination & sequential airport scans* | *e.g., Added composite index & in-memory LRU cache* | *e.g., 38ms @ 400 RPS* | *Fixed* |
| *PB-02* | *YYYY-MM-DD* | *[Endpoint / Component]* | *[Baseline]* | *[Root Cause]* | *[Fix]* | *[Post-Fix]* | *Investigating* |

---

## Bottleneck Investigation Template

### [PB-XXXX] Title / Short Description

#### 1. Baseline Profile
- **Target Endpoint / Function**: `[e.g. GET /api/search?from=JFK&to=SFO]`
- **Test Scenario**: `[e.g. 50 VUs ramp-up over 2m using load-tests/k6-test.js]`
- **Baseline Throughput**: `[e.g. 85 req/s]`
- **Baseline Latency**:
  - P50: `[e.g. 180ms]`
  - P95: `[e.g. 620ms]`
  - P99: `[e.g. 1100ms]`
- **Error Rate**: `[e.g. 2.4% timeouts]`

#### 2. Root Cause Analysis
- **Observed Behavior**: `[CPU saturation, memory leak, DB connection starvation, lock contention]`
- **Diagnostic Tool Used**: `[e.g., go tool pprof, sqlite3 EXPLAIN QUERY PLAN, k6 telemetry]`
- **Root Cause Details**:
  <!-- Provide succinct code snippet, query plan, or profile finding -->
  ```text
  [TODO: Insert diagnostic excerpt, query plan, or flamegraph notes here]
  ```

#### 3. Optimization Applied
- **Changes Implemented**:
  - `[TODO: Description of code refactor, index addition, caching strategy, concurrency tweak]`
- **PR / Commit**: `[TODO: Link to commit or PR]`

#### 4. Post-Optimization Verification
- **Post-Fix Throughput**: `[e.g. 450 req/s (+430%)]`
- **Post-Fix Latency**:
  - P50: `[e.g. 22ms]`
  - P95: `[e.g. 45ms]`
  - P99: `[e.g. 85ms]`
- **Error Rate**: `[e.g. 0.0%]`
- **Regression Checks Passed**: `[Yes/No]`
