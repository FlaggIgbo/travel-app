package evals

import (
	"context"
	"time"
	"triangle_travel/internal/flights"
)

// BenchmarkCase defines a single multi-city evaluation prompt and the ground-truth constraints
// that any returned itinerary must satisfy.
type BenchmarkCase struct {
	ID                     string   `json:"id"`
	Prompt                 string   `json:"prompt"`
	ExpectedOrigin         string   `json:"expectedOrigin"`
	ExpectedMinCities      int      `json:"expectedMinCities"`
	ExpectedCabin          string   `json:"expectedCabin"`
	MaxPositioningFlights  int      `json:"maxPositioningFlights"`
	MaxPointsBudget        int      `json:"maxPointsBudget"`
	RequiredAlliances      []string `json:"requiredAlliances,omitempty"`
	AllowGroundBridge      bool     `json:"allowGroundBridge"`
}

// EvalMetrics summarizes model performance across the benchmark dataset for README documentation.
type EvalMetrics struct {
	ModelName          string        `json:"modelName"`
	TotalCases         int           `json:"totalCases"`
	ConstraintPassRate float64       `json:"constraintPassRate"` // 0.0 - 1.0
	P50Latency         time.Duration `json:"p50Latency"`
	P95Latency         time.Duration `json:"p95Latency"`
	AvgTokensPerSearch int           `json:"avgTokensPerSearch"`
	AvgCostUSD         float64       `json:"avgCostUsd"`
	FallbackRate       float64       `json:"fallbackRate"`
}

// SampleBenchmarkDataset provides starter prompts for the 25-30 complex multi-city eval suite.
//
// TODO(step-4-evals): Expand dataset to 25-30 diverse multi-city prompts covering open-jaw rail bridges,
// alliance transfer bonuses, and sparse saver availability fallbacks.
func SampleBenchmarkDataset() []BenchmarkCase {
	return []BenchmarkCase{
		{
			ID:                    "nyc-eu-biz-120k",
			Prompt:                "NYC to two European cities in Business, max 1 positioning flight, under 120k points using Chase or Bilt",
			ExpectedOrigin:        "NYC",
			ExpectedMinCities:     2,
			ExpectedCabin:         "BUSINESS",
			MaxPositioningFlights: 1,
			MaxPointsBudget:       120000,
			AllowGroundBridge:     true,
		},
		{
			ID:                    "sfo-asia-oneworld-openjaw",
			Prompt:                "SFO to Tokyo and Seoul in Oneworld Business or First, open to bullet train between regional hubs, under 160k Avios/AAdvantage",
			ExpectedOrigin:        "SFO",
			ExpectedMinCities:     2,
			ExpectedCabin:         "BUSINESS",
			MaxPositioningFlights: 1,
			MaxPointsBudget:       160000,
			RequiredAlliances:     []string{"ONE_WORLD"},
			AllowGroundBridge:     true,
		},
	}
}

// AssertConstraintsSatisfied runs deterministic assertions verifying that a candidate itinerary
// honors every hard constraint in the benchmark case (points ceiling, positioning flight limit, cabin, and valid graph edges).
func AssertConstraintsSatisfied(tc BenchmarkCase, req flights.TripRequest, itinerary flights.CandidateItinerary) (bool, []string) {
	// TODO(step-4-evals): Implement automated assertions checking:
	// - itinerary.TotalPoints <= tc.MaxPointsBudget
	// - itinerary.PositioningLegCount <= tc.MaxPositioningFlights
	// - All award legs match tc.ExpectedCabin and valid route-graph edges
	// - Return (passed, failureReasons)
	return false, []string{"TODO: constraint assertion harness not yet implemented"}
}

// RunEvalSuite executes the benchmark dataset against a target model configuration and computes
// latency, token cost, and constraint-pass rate metrics.
func RunEvalSuite(ctx context.Context, modelName string, dataset []BenchmarkCase) (*EvalMetrics, error) {
	// TODO(step-4-evals): Wire up Orchestrator.RunBoundedOptimization across all cases in dataset,
	// evaluate with AssertConstraintsSatisfied, and output markdown table rows for README.md.
	return &EvalMetrics{
		ModelName:  modelName,
		TotalCases: len(dataset),
	}, nil
}
