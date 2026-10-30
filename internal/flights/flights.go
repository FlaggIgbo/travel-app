package flights

import (
	"strings"
	"triangle_travel/internal/db"
)

// SearchInput holds user search inputs
type SearchInput struct {
	Start     string `json:"start"`
	End       string `json:"end"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Cabin     string `json:"cabin"`
	Alliance  string `json:"alliance"`
}

// SearchResult holds the travel options found
type SearchResult struct {
	DriveThenFly map[string]float64 `json:"driveThenFly"`
	FlyThenFly   map[string]float64 `json:"flyThenFly"`
	AvgPrice     float64            `json:"avgPrice"`
}

// FindTriangleTravel finds stopover cities for drive/train or connecting flights
func FindTriangleTravel(database *db.DB, input SearchInput) (*SearchResult, error) {
	// Clean up user inputs
	input.Start = strings.ToUpper(strings.TrimSpace(input.Start))
	input.End = strings.ToUpper(strings.TrimSpace(input.End))
	if input.Alliance == "" {
		input.Alliance = "None"
	}

	result := &SearchResult{
		DriveThenFly: make(map[string]float64),
		FlyThenFly:   make(map[string]float64),
		AvgPrice:     -1,
	}

	// 1. Find nearby airports you can drive or take a train to (between 55 and 300 miles)
	distances, err := database.GetDistancesFrom(input.End)
	if err != nil {
		return result, err
	}
	for airport, distance := range distances {
		if distance >= 55 && distance <= 300 {
			result.DriveThenFly[airport] = distance
		}
	}

	// 2. Find direct flight destinations from the end city
	routes, err := database.GetRoutesFromWithFallback(input.End, input.Alliance)
	if err != nil {
		return result, err
	}
	for _, airport := range routes {
		result.FlyThenFly[airport] = 0
	}

	return result, nil
}
