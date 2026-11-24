import http from 'k6/http';
import { check, sleep, group } from 'k6';
import { Counter, Rate, Trend } from 'k6/metrics';

// Custom Metrics
const searchLatencyTrend = new Trend('route_search_duration_ms');
const errorRate = new Rate('failed_requests_rate');
const searchSuccessCounter = new Counter('successful_route_searches');

// Load Test Configuration Options
export const options = {
  scenarios: {
    // 1. Standard Load / Smoke Test
    standard_load: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: '30s', target: 20 }, // Ramp up to 20 users
        { duration: '1m', target: 20 },  // Steady load
        { duration: '20s', target: 0 },  // Ramp down
      ],
      gracefulRampDown: '10s',
    },
    // 2. Stress / Spike Test (Uncomment when conducting spike/capacity tests)
    /*
    stress_test: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: '15s', target: 100 }, // Spike to 100 users
        { duration: '30s', target: 100 }, // Hold high load
        { duration: '15s', target: 0 },   // Cool down
      ],
    },
    */
    // 3. Chaos / Fault Injection Scenario (Uncomment when testing resilience)
    /*
    chaos_scenario: {
      executor: 'constant-vus',
      vus: 10,
      duration: '1m',
    },
    */
  },
  thresholds: {
    http_req_duration: ['p(95)<500', 'p(99)<1000'], // 95% of requests under 500ms
    failed_requests_rate: ['rate<0.02'],            // Error rate under 2%
    route_search_duration_ms: ['p(95)<400'],        // Search API P95 under 400ms
  },
};

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8080';

// Setup function: runs once before virtual users start
export function setup() {
  console.log(`Setting up k6 load test targeting ${BASE_URL}...`);
  // TODO: Verify server health or pre-seed test database state if necessary
  const healthRes = http.get(`${BASE_URL}/api/cities`);
  check(healthRes, {
    'setup health check responded 200': (r) => r.status === 200,
  });
  return { startTime: new Date().toISOString() };
}

export default function (data) {
  group('City and Airport Directory', () => {
    testFetchCities();
  });

  group('Route Search Pipeline', () => {
    testSearchDirectRoutes();
    testSearchTriangleRoutes();
  });

  group('Chaos & Resilience Stubs', () => {
    // Chaos tests simulate dropped connections, latency spikes, or malformed payloads
    testChaosMalformedInputs();
    testChaosSimulatedRateLimiting();
  });

  sleep(1); // Think time between iterations
}

// ---------------------------------------------------------------------------
// Load Test Function Stubs
// ---------------------------------------------------------------------------

export function testFetchCities() {
  // TODO: Add query parameters or pagination when supported
  const res = http.get(`${BASE_URL}/api/cities`);
  
  const success = check(res, {
    'cities status is 200': (r) => r.status === 200,
    'cities response has items': (r) => {
      try {
        const json = JSON.parse(r.body);
        return Array.isArray(json) && json.length > 0;
      } catch (e) {
        return false;
      }
    },
  });

  errorRate.add(!success);
}

export function testSearchDirectRoutes() {
  const payload = JSON.stringify({
    from: 'JFK',
    to: 'SFO',
    // TODO: Add dates, passenger count, or cabin class filters as schema expands
  });

  const params = {
    headers: { 'Content-Type': 'application/json' },
  };

  const start = Date.now();
  const res = http.post(`${BASE_URL}/api/search`, payload, params);
  searchLatencyTrend.add(Date.now() - start);

  const success = check(res, {
    'search status is 200': (r) => r.status === 200,
    'search returns routes': (r) => r.body && r.body.length > 0,
  });

  if (success) searchSuccessCounter.add(1);
  errorRate.add(!success);
}

export function testSearchTriangleRoutes() {
  // TODO: Test multi-city triangle route discovery (e.g., JFK -> SFO with intermediate explore city)
  const payload = JSON.stringify({
    from: 'ORD',
    to: 'LAX',
    maxDriveMiles: 150,
  });

  const params = {
    headers: { 'Content-Type': 'application/json' },
  };

  const start = Date.now();
  const res = http.post(`${BASE_URL}/api/search`, payload, params);
  searchLatencyTrend.add(Date.now() - start);

  const success = check(res, {
    'triangle search status is 200': (r) => r.status === 200,
  });

  errorRate.add(!success);
}

// ---------------------------------------------------------------------------
// Chaos & Fault Injection Function Stubs
// ---------------------------------------------------------------------------

export function testChaosMalformedInputs() {
  // TODO: Fuzz invalid IATA codes, negative distances, or oversized payloads to test error handling
  const invalidPayload = JSON.stringify({
    from: 'INVALID_AIRPORT_CODE',
    to: '???',
  });

  const res = http.post(`${BASE_URL}/api/search`, invalidPayload, {
    headers: { 'Content-Type': 'application/json' },
  });

  // Chaos expectation: Should return graceful client error (400), never a 500 panic or connection reset
  check(res, {
    'gracefully handles invalid airport code': (r) => r.status === 400 || r.status === 422,
  });
}

export function testChaosSimulatedRateLimiting() {
  // TODO: Rapid burst to verify rate-limiter triggers HTTP 429 without service degradation
  /*
  for (let i = 0; i < 20; i++) {
    http.get(`${BASE_URL}/api/cities`);
  }
  */
}

// Teardown function: runs once after all VUs complete
export function teardown(data) {
  console.log(`k6 load test completed. Started at: ${data.startTime}`);
  // TODO: Clean up ephemeral state, parse final metrics, or alert on anomalies
}
