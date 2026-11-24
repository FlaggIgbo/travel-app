package mcp

import (
	"context"
	"errors"
	"triangle_travel/internal/db"
)

// ErrMCPNotImplemented is returned by MCP tool handler stubs pending protocol wiring.
var ErrMCPNotImplemented = errors.New("not implemented: pending MCP server milestone")

// ToolDefinition describes an MCP-exposed tool name, description, and JSON input schema.
type ToolDefinition struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	InputSchema string `json:"inputSchema"`
}

// Server exposes triangle-travel's deterministic route-graph queries, award-chart calculators,
// and itinerary scorer as Model Context Protocol (MCP) tools so any AI client (Claude Code,
// local Ollama models, or the web UI) can invoke them.
type Server struct {
	Repo db.RouteRepository
}

// NewServer creates a new MCP tool server backed by a RouteRepository.
func NewServer(repo db.RouteRepository) *Server {
	return &Server{Repo: repo}
}

// ListTools returns the MCP tool definitions registered by triangle-travel.
func (s *Server) ListTools() []ToolDefinition {
	// TODO(step-3-mcp): Finalize JSON schemas for each MCP tool and wire up stdio/SSE transport.
	return []ToolDefinition{
		{
			Name:        "query_route_graph",
			Description: "Query direct alliance flight routes and multi-modal ground bridges (rail/car) from an airport or city.",
			InputSchema: `{"type":"object","properties":{"origin":{"type":"string"},"alliance":{"type":"string"},"maxGroundMiles":{"type":"number"}},"required":["origin"]}`,
		},
		{
			Name:        "calculate_transfer_arbitrage",
			Description: "Evaluate credit card / loyalty transfer partner ratios and cents-per-point (CPP) value for a route.",
			InputSchema: `{"type":"object","properties":{"pointBalances":{"type":"object"},"targetCabin":{"type":"string"},"cashPriceUsd":{"type":"number"}},"required":["pointBalances"]}`,
		},
		{
			Name:        "score_triangle_itineraries",
			Description: "Deterministically rank candidate triangle and open-jaw itineraries against a structured TripRequest.",
			InputSchema: `{"type":"object","properties":{"tripRequest":{"type":"object"}},"required":["tripRequest"]}`,
		},
	}
}

// CallTool dispatches an incoming MCP tool invocation to the corresponding deterministic engine function.
func (s *Server) CallTool(ctx context.Context, name string, argumentsJSON []byte) ([]byte, error) {
	// TODO(step-3-mcp): Unmarshal argumentsJSON, invoke flights.FindTriangleTravel or
	// flights.EvaluateTriangleArbitrage / flights.ScoreCandidateItineraries, and marshal the result.
	return nil, ErrMCPNotImplemented
}
