package agent

import (
	"context"
	"errors"
	"time"
	"triangle_travel/internal/db"
	"triangle_travel/internal/flights"
)

// ErrAgentNotImplemented is returned by agent stubs pending LLM client integration.
var ErrAgentNotImplemented = errors.New("not implemented: pending bounded agent loop milestone")

// LoopBudget enforces strict cost, token, step, and latency bounds on the iterative search agent.
type LoopBudget struct {
	MaxSteps       int           `json:"maxSteps"`       // Maximum tool-calling iterations (e.g., 4)
	MaxTokens      int           `json:"maxTokens"`      // Strict prompt + completion token ceiling
	StepTimeout    time.Duration `json:"stepTimeout"`    // Per-tool / per-API call timeout
	EnableFallback bool          `json:"enableFallback"` // Fall back to cached SQLite routes when live APIs time out
}

// DefaultLoopBudget returns sensible production defaults for bounded agent execution.
func DefaultLoopBudget() LoopBudget {
	return LoopBudget{
		MaxSteps:       4,
		MaxTokens:      4096,
		StepTimeout:    5 * time.Second,
		EnableFallback: true,
	}
}

// OptimizeResponse bundles the structured TripRequest, deterministically ranked itineraries,
// and the LLM-synthesized trade-off comparison.
type OptimizeResponse struct {
	ParsedRequest   flights.TripRequest          `json:"parsedRequest"`
	Itineraries     []flights.CandidateItinerary `json:"itineraries"`
	SynthesisReport string                       `json:"synthesisReport"`
	StepsUsed       int                          `json:"stepsUsed"`
	TokensUsed      int                          `json:"tokensUsed"`
	FallbackUsed    bool                         `json:"fallbackUsed"`
}

// Orchestrator coordinates deterministic route-graph search with bounded LLM calls
// restricted strictly to intent parsing and trade-off synthesis.
type Orchestrator struct {
	Repo   db.RouteRepository
	Budget LoopBudget
}

// ParseTripIntent converts a natural-language multi-city prompt
// (e.g., "NYC to two European cities in Business, max 1 positioning flight, under 120k points")
// into a strictly validated flights.TripRequest struct.
func (o *Orchestrator) ParseTripIntent(ctx context.Context, naturalLanguagePrompt string) (*flights.TripRequest, error) {
	// TODO(deterministic-plus-agentic):
	// 1. Invoke LLM with structured JSON output schema for flights.TripRequest.
	// 2. Validate IATA airport/city codes, date ranges, cabin enums, and non-negative budgets.
	// 3. Do NOT let the LLM suggest routes or calculate points—only parse user constraints.
	return nil, ErrAgentNotImplemented
}

// RunBoundedOptimization executes the state machine:
//  1. Parse natural-language prompt -> validated TripRequest (LLM).
//  2. Query route graph & evaluate triangle/open-jaw arbitrage (Deterministic).
//  3. If live availability is sparse or times out, relax flexibility window or fall back to cached schedule data.
//  4. Synthesize clear trade-off explanations across the top ranked itineraries (LLM).
func (o *Orchestrator) RunBoundedOptimization(ctx context.Context, prompt string) (*OptimizeResponse, error) {
	// TODO(bounded-agent-loops):
	// Implement bounded retry/relaxation loop respecting o.Budget.MaxSteps, o.Budget.MaxTokens,
	// and o.Budget.StepTimeout, recording telemetry (StepsUsed, TokensUsed, FallbackUsed) for eval benchmarks.
	return nil, ErrAgentNotImplemented
}

// SynthesizeTradeoffs generates a concise human-readable comparison of the top-ranked deterministic
// itineraries (comparing positioning inconvenience vs. points saved and ground transit time).
func (o *Orchestrator) SynthesizeTradeoffs(ctx context.Context, req flights.TripRequest, ranked []flights.CandidateItinerary) (string, error) {
	// TODO(deterministic-plus-agentic):
	// Pass only the top N deterministically scored itineraries to the LLM so it explains
	// trade-offs grounded strictly in the computed numbers.
	return "", ErrAgentNotImplemented
}
