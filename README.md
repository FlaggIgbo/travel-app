# Triangle Travel

A travel planning web app that helps you find a third city to explore on a round-trip itinerary.

For example, when flying from New York (JFK) to San Francisco (SFO), the app shows:
1. **Drive or train then fly**: Cities within 55–300 miles of your destination (like Oakland or San Jose).
2. **Fly then fly**: Direct connecting flight routes from your destination.

---

## Requirements

- **Go** (version 1.21 or newer)
- **Node.js** (version 18 or newer)

---

## Quick Start (3 Steps)

### 1. Seed the database
Creates the SQLite database at `db/data.sqlite3` with airport codes, distances, and routes:
```bash
go run ./cmd/seed
```

### 2. Build the frontend
Installs frontend dependencies and builds the static site into `build/`:
```bash
npm install
npm run build
```

### 3. Start the server
Runs the web server on **http://localhost:8080**:
```bash
go run .
```

Open your browser and visit **http://localhost:8080**.

---

## Development Mode (Hot Reload)

To work on frontend and backend simultaneously:

**Terminal 1 (Backend API):**
```bash
go run .
```

**Terminal 2 (Frontend Dev Server):**
```bash
npm run dev
```
Open **http://localhost:5173**. Vite automatically proxies `/api` requests to port 8080.

---

## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/api/search` | Search for triangle travel stopovers |
| `GET` | `/api/cities` | List supported airport city codes |

### Example Search Request
```json
{
  "start": "JFK",
  "end": "SFO",
  "startDate": "2025-03-01",
  "endDate": "2025-03-15",
  "alliance": "STAR_ALLIANCE"
}
```

---

## Project Structure

```text
triangle-travel/
├── cmd/seed/            # Script to create and seed the SQLite database
├── db/
│   ├── schema.sql       # Table definitions (cities, distances, routes, loyalty, transit)
│   └── seed_data.sql    # Airport and route data
├── internal/
│   ├── agent/           # Bounded LLM orchestration (intent parsing & trade-off synthesis)
│   ├── api/             # HTTP route handlers (/api/search, /api/cities, /api/optimize)
│   ├── db/              # Database helper functions & RouteRepository abstraction
│   ├── evals/           # Automated evaluation harness & multi-city benchmark checks
│   ├── flights/         # Triangle travel search, constraint engine & itinerary scorer
│   ├── mcp/             # Model Context Protocol (MCP) server & tool definitions
│   └── server/          # Web server setup
├── src/                 # Svelte frontend
│   └── routes/          # Web pages (+page.svelte)
├── main.go              # Application entry point
├── package.json         # Frontend configuration
└── go.mod               # Go dependencies
```

---

## Testing

Run backend tests:
```bash
go test ./...
```

Run frontend type and lint checks:
```bash
npm run check
```

---

## Roadmap & TODO: Agentic Multi-City Award & Positioning Optimizer

Turn `triangle-travel` from a route-lookup scraper into an **agentic multi-city award and positioning optimizer**.

FlightsFrom shows *which* direct routes and timetables exist, but it does not solve the hard combinatorial and preference-driven problem of triangle travel: optimizing alliance transfer partners, positioning flights, cabin upgrades, and ground transit around real constraints. Rather than building another CRUD wrapper around an LLM, the goal is a **deterministic-plus-agentic pipeline** with strict schema validation, tool orchestration, and an evaluation harness.

### 1. Product Scope (Beyond FlightsFrom)

Treat route graphs as an input layer and focus on **constraint satisfaction and points optimization**:

- [ ] **Triangle & Open-Jaw Arbitrage** (`internal/flights/optimizer.go`): Automatically evaluate `Origin -> Hub A -> Destination -> Hub B -> Origin` (or ground-leg open-jaws) to find sweet spots where mixing cash positioning flights with partner award redemptions beats round-trip pricing.
- [ ] **Loyalty & Alliance Routing** (`internal/flights/optimizer.go`, `db/schema.sql`): Encode rules for transfer currencies (e.g., Chase, Bilt, Avios, AAdvantage) and alliance networks (like *oneworld* direct-flight preferences) so the engine filters out low-value redemptions or undesirable layovers before surfacing options.
- [ ] **Multi-Modal Ground Bridging** (`internal/flights/flights.go`, `db/schema.sql`): Factor in regional rail or one-way car rentals between nearby airports (e.g., flying into one regional hub and departing from another) when it unlocks better award availability.
- [ ] **Award Seat Availability Poller & Multi-Program Transfer Sniper**: Background daemon (similar to [Seats.aero](https://seats.aero/)) monitoring *oneworld* / Star Alliance / SkyTeam saver award space across airline mileage programs and credit card transfer partners:
  - **Dynamic Polling & Token Bucket Rate Limiting**: Distributed rate-limiter (Redis token bucket) with rotating worker pools to query partner endpoints without IP bans.
  - **Priority Queues (`Redis ZSET`)**: Score flight searches by departure proximity (e.g., polling every 15 mins for flights 3 weeks out vs. daily for 6 months out).
  - **Transfer Valuation Calculator Engine**: In-memory graph of point currencies, transfer ratios, and baseline cash valuations (cents-per-point) to flag when transferring credit card points to airline miles beats booking with cash.

### 2. Architecture Upgrades for AI Engineering

| Pattern | What to Build in `triangle-travel` | Engineering Skill Gained |
| --- | --- | --- |
| **Deterministic Graph + LLM Orchestration** | Keep route discovery deterministic (graph search / SQLite or Postgres queries over cached schedule data), and use an LLM agent strictly for **intent parsing** (turning vague trip constraints into structured JSON) and **trade-off synthesis**. | Knowing when *not* to use an LLM; preventing hallucinations in math/routing logic. |
| **Tool-Calling & MCP Server** | Expose route-graph queries, award-chart calculators, and live flight/hotel search APIs as **Model Context Protocol (MCP)** tools (`internal/mcp/server.go`). | Building reusable tool interfaces that any AI client (Claude Code, local Ollama models, or the web UI) can invoke. |
| **Bounded Agent Loops & Fallbacks** | Implement an iterative search agent (`internal/agent/orchestrator.go`) with a strict token/step budget, structured output validation, and fallback handling when a route API times out or returns sparse availability. | Production resilience, state machines, and cost/latency control. |
| **Automated Evaluation (Evals)** | Create a benchmark dataset of 25–30 complex multi-city prompts (`internal/evals/harness.go`, e.g., *"NYC to two European cities in Business, max 1 positioning flight, under 120k points"*) and write automated assertions checking constraint satisfaction. | Regression testing non-deterministic systems. |

### 3. Concrete Implementation Milestones

1. [ ] **Decouple the Data Ingestion** (`internal/db/db.go`): Abstract the FlightsFrom scraping/data layer behind a clean `RouteRepository` interface so cached OpenFlights/airline timetable datasets or official flight search APIs can be swapped in without breaking core logic.
2. [ ] **Add a Structured Constraint Engine** (`internal/flights/optimizer.go`): Define a strict JSON schema for a `TripRequest` (dates, flexibility window, alliance preferences, point balances, cabin class) and build a deterministic scorer (`ScoreCandidateItineraries`) that ranks candidate triangle itineraries.
3. [ ] **Wrap It in an MCP Server + Lightweight UI** (`internal/mcp/server.go`, `src/routes/+page.svelte`): Expose the scorer and route finder as MCP tools first for immediate terminal/editor usage while planning trips, then wire up a visual itinerary comparison view.
4. [ ] **Document the Benchmarks in the README** (`internal/evals/harness.go`): Track latency, token cost per trip search, and constraint-pass rate across different models (comparing frontier models against local open-weight models).

### 4. Evaluation Benchmarks (TODO: Populate via `internal/evals`)

| Model | Prompt Dataset Size | Constraint Pass Rate (%) | P50 / P95 Latency (ms) | Avg Cost / Search ($) | Fallback Rate (%) |
| --- | --- | --- | --- | --- | --- |
| *Frontier Model (e.g. Claude / Gemini Pro)* | 30 | `TODO` | `TODO` | `TODO` | `TODO` |
| *Fast Cloud Model (e.g. Haiku / Flash)* | 30 | `TODO` | `TODO` | `TODO` | `TODO` |
| *Local Open-Weight (Ollama)* | 30 | `TODO` | `TODO` | `$0.00` | `TODO` |

