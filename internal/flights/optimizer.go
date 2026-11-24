package flights

import (
	"errors"
	"triangle_travel/internal/db"
)

// ErrNotImplemented is returned by optimizer stubs that are pending implementation.
var ErrNotImplemented = errors.New("not implemented: pending agentic optimizer milestone")

// TripRequest defines the strict JSON schema for structured constraint satisfaction
// and multi-city award/positioning optimization.
//
// TODO(step-2): Add JSON schema validation tags and a deterministic validator so LLM-parsed
// intents are strictly checked before executing graph queries.
type TripRequest struct {
	Origin              string         `json:"origin"`
	TargetDestinations  []string       `json:"targetDestinations"`
	StartDate           string         `json:"startDate"`
	EndDate             string         `json:"endDate"`
	FlexibilityDays     int            `json:"flexibilityDays"`
	PreferredAlliances  []string       `json:"preferredAlliances"`  // e.g., ["ONE_WORLD", "STAR_ALLIANCE"]
	PointBalances       map[string]int `json:"pointBalances"`       // e.g., {"CHASE_UR": 120000, "BILT": 50000, "AVIOS": 35000}
	TargetCabin         string         `json:"targetCabin"`         // e.g., "BUSINESS", "ECONOMY"
	MaxPositioningLegs  int            `json:"maxPositioningLegs"`  // e.g., 1
	MaxPointsBudget     int            `json:"maxPointsBudget"`     // e.g., 120000
	MaxCashBudgetUSD    float64        `json:"maxCashBudgetUsd"`    // Max out-of-pocket cash for positioning + taxes/fees
	AllowGroundBridging bool           `json:"allowGroundBridging"` // Allow regional rail / one-way car rental open-jaws
}

// ItineraryLeg represents a single segment (award flight, cash positioning flight, or ground bridge)
// within a candidate multi-city itinerary.
type ItineraryLeg struct {
	FromIATA       string  `json:"fromIata"`
	ToIATA         string  `json:"toIata"`
	Date           string  `json:"date"`
	LegType        string  `json:"legType"`        // "AWARD_FLIGHT", "CASH_POSITIONING", "GROUND_BRIDGE"
	CarrierOrMode  string  `json:"carrierOrMode"`  // Airline IATA code or ground mode ("RAIL", "CAR_RENTAL")
	Alliance       string  `json:"alliance"`       // Alliance network if flight
	LoyaltyProgram string  `json:"loyaltyProgram"` // Program used if award flight (e.g., "AVIOS", "AADVANTAGE")
	Cabin          string  `json:"cabin"`
	PointsCost     int     `json:"pointsCost"`
	CashCostUSD    float64 `json:"cashCostUsd"`
	TransitMinutes int     `json:"transitMinutes"`
}

// CandidateItinerary represents a complete triangle or open-jaw travel plan
// (e.g., Origin -> Hub A -> Destination -> Hub B -> Origin) with deterministic score metrics.
type CandidateItinerary struct {
	Legs                []ItineraryLeg `json:"legs"`
	TotalPoints         int            `json:"totalPoints"`
	TotalCashUSD        float64        `json:"totalCashUsd"`
	EffectiveCPP        float64        `json:"effectiveCpp"`        // Cents-per-point vs. baseline round-trip cash price
	PositioningLegCount int            `json:"positioningLegCount"` // Number of cash positioning flights used
	GroundBridgeCount   int            `json:"groundBridgeCount"`   // Number of rail/car open-jaw bridges
	Score               float64        `json:"score"`               // Deterministic composite rank score
	TradeoffSummary     string         `json:"tradeoffSummary"`     // Optional LLM-synthesized explanation of trade-offs
}

// EvaluateTriangleArbitrage searches the deterministic route graph for open-jaw and triangle combinations
// (`Origin -> Hub A -> Destination -> Hub B -> Origin`) where mixing cash positioning flights or ground bridges
// with partner award redemptions beats standard round-trip pricing.
func EvaluateTriangleArbitrage(repo db.RouteRepository, req TripRequest) ([]CandidateItinerary, error) {
	// TODO(triangle-arbitrage):
	// 1. Expand candidate Hub A and Hub B nodes reachable from req.Origin and req.TargetDestinations.
	// 2. If req.AllowGroundBridging is true, include multi-modal ground bridges (repo.GetGroundBridgesFrom)
	//    between nearby regional hubs.
	// 3. Prune paths that exceed req.MaxPositioningLegs, req.MaxPointsBudget, or req.MaxCashBudgetUSD.
	// 4. Filter and price award segments using FilterByLoyaltyAndAlliance.
	// 5. Rank surviving itineraries deterministically with ScoreCandidateItineraries.
	return nil, ErrNotImplemented
}

// FilterByLoyaltyAndAlliance applies transfer currency rules (e.g., Chase, Bilt, Avios, AAdvantage)
// and alliance direct-flight preferences to prune low-value redemptions or undesirable layovers.
func FilterByLoyaltyAndAlliance(repo db.RouteRepository, candidates []CandidateItinerary, req TripRequest) ([]CandidateItinerary, error) {
	// TODO(loyalty-routing):
	// 1. Look up transfer ratios and partner programs via repo.GetTransferPartners for each currency in req.PointBalances.
	// 2. Filter out candidate itineraries whose required points exceed available transferable balances.
	// 3. Filter out redemptions below the partner program's baseline cents-per-point (CPP) threshold.
	return candidates, ErrNotImplemented
}

// ScoreCandidateItineraries deterministically scores and sorts candidate triangle itineraries
// based on cents-per-point (CPP) arbitrage value, total travel friction, cabin class, and positioning overhead.
func ScoreCandidateItineraries(candidates []CandidateItinerary, req TripRequest) []CandidateItinerary {
	// TODO(step-2): Implement deterministic multi-objective scoring:
	//   score = (w_cpp * EffectiveCPP) - (w_pos * PositioningLegCount) - (w_transit * TotalTransitHours) + (w_cabin * CabinBonus)
	// Keep all math and ranking 100% deterministic (never delegated to an LLM).
	return candidates
}
