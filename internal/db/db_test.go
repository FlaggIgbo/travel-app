package db

import (
	"path/filepath"
	"testing"
)

func TestSameCity(t *testing.T) {
	// Open test database
	database, err := New(filepath.Join("..", ".."))
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer database.Close()

	// JFK and LGA are both in New York
	same, err := database.SameCity("JFK", "LGA")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !same {
		t.Errorf("Expected JFK and LGA to be in the same city")
	}

	// JFK and SFO are in different cities
	diff, err := database.SameCity("JFK", "SFO")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if diff {
		t.Errorf("Expected JFK and SFO to be in different cities")
	}
}
