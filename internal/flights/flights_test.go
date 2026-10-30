package flights

import (
	"path/filepath"
	"testing"
	"triangle_travel/internal/db"
)

func TestFindTriangleTravel(t *testing.T) {
	// Open test database
	database, err := db.New(filepath.Join("..", ".."))
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer database.Close()

	// Test search from JFK to SFO
	input := SearchInput{
		Start:     "JFK",
		End:       "SFO",
		StartDate: "2025-03-01",
		EndDate:   "2025-03-15",
	}

	result, err := FindTriangleTravel(database, input)
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}

	// SFO has nearby airports like OAK and SJC
	if len(result.DriveThenFly) == 0 {
		t.Errorf("Expected nearby drive/train airports near SFO")
	}

	// SFO has connecting flight routes
	if len(result.FlyThenFly) == 0 {
		t.Errorf("Expected fly-then-fly routes from SFO")
	}
}
