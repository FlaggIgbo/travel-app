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
│   ├── schema.sql       # Table definitions (cities, distances, routes)
│   └── seed_data.sql    # Airport and route data
├── internal/
│   ├── api/             # HTTP route handlers (/api/search, /api/cities)
│   ├── db/              # Database helper functions
│   ├── flights/         # Triangle travel search logic
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

## TODO

- **Award Seat Availability Poller & Multi-Program Transfer Sniper**: A background daemon (similar to [Seats.aero](https://seats.aero/)) that monitors Oneworld / Star Alliance / SkyTeam saver award space across airline mileage programs and credit card transfer partners for routes like JFK → LHR and sends instant alerts when partner seats open.
  - **Dynamic Polling & Token Bucket Rate Limiting**: Distributed rate-limiter (Redis token bucket) with rotating user-agent worker pools to query partner endpoints without IP bans.
  - **Priority Queues (`Redis ZSET`)**: Score flight searches by departure proximity (e.g., polling every 15 mins for flights 3 weeks out vs. daily for 6 months out).
  - **Transfer Valuation Calculator Engine**: In-memory graph of point currencies, transfer ratios, and baseline cash valuations (cents-per-point) to flag when transferring credit card points to airline miles beats booking with cash.
