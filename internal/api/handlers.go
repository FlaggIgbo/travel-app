package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"triangle_travel/internal/db"
	"triangle_travel/internal/flights"
)

// API holds the database connection for API routes
type API struct {
	DB *db.DB
}

// SearchRequest holds the search form data sent by the client
type SearchRequest struct {
	Start     string `json:"start" binding:"required"`
	End       string `json:"end" binding:"required"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Cabin     string `json:"cabin"`
	Alliance  string `json:"alliance"`
}

// Search handles POST /api/search
func (a *API) Search(c *gin.Context) {
	var req SearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide start and end airports"})
		return
	}

	// Make sure start and end airports are in different cities
	sameCity, err := a.DB.SameCity(req.Start, req.End)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if sameCity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Start and end must be different cities"})
		return
	}

	// Search for triangle travel options
	input := flights.SearchInput{
		Start:     req.Start,
		End:       req.End,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Cabin:     req.Cabin,
		Alliance:  req.Alliance,
	}

	result, err := flights.FindTriangleTravel(a.DB, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Search failed"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// Cities handles GET /api/cities
func (a *API) Cities(c *gin.Context) {
	// Query all unique city codes
	rows, err := a.DB.Query("SELECT DISTINCT city_code FROM iata_cities ORDER BY city_code")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not load cities"})
		return
	}
	defer rows.Close()

	var cities []string
	for rows.Next() {
		var cityCode string
		if err := rows.Scan(&cityCode); err == nil {
			cities = append(cities, cityCode)
		}
	}

	c.JSON(http.StatusOK, cities)
}

// OptimizeRequest holds either a natural-language prompt or a structured TripRequest
// for the agentic multi-city award and positioning optimizer.
type OptimizeRequest struct {
	Prompt      string               `json:"prompt"`
	TripRequest *flights.TripRequest `json:"tripRequest,omitempty"`
}

// OptimizeItinerary handles POST /api/optimize
//
// TODO(step-2 & step-3): Wire this handler to internal/agent.Orchestrator.RunBoundedOptimization
// (when Prompt is provided) or directly to internal/flights.EvaluateTriangleArbitrage
// (when a structured TripRequest is provided).
func (a *API) OptimizeItinerary(c *gin.Context) {
	var req OptimizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid optimization request payload"})
		return
	}

	c.JSON(http.StatusNotImplemented, gin.H{
		"error":   "Agentic multi-city award & positioning optimizer is not yet wired up",
		"todo":    "See internal/flights/optimizer.go and internal/agent/orchestrator.go",
		"request": req,
	})
}
