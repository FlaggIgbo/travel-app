package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"triangle_travel/internal/db"
)

func TestAPIRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Connect to database
	database, err := db.New(filepath.Join("..", ".."))
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer database.Close()

	// Setup router and endpoints
	apiHandler := &API{DB: database}
	router := gin.New()
	router.POST("/api/search", apiHandler.Search)
	router.GET("/api/cities", apiHandler.Cities)

	// 1. Test GET /api/cities returns status 200
	req, _ := http.NewRequest("GET", "/api/cities", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 for /api/cities, got %d", w.Code)
	}

	// 2. Test POST /api/search rejects same-city search (e.g. JFK and LGA)
	sameCityBody, _ := json.Marshal(SearchRequest{Start: "JFK", End: "LGA"})
	req, _ = http.NewRequest("POST", "/api/search", bytes.NewBuffer(sameCityBody))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for same-city search, got %d", w.Code)
	}

	// 3. Test POST /api/search succeeds for different cities (JFK to SFO)
	validBody, _ := json.Marshal(SearchRequest{
		Start:     "JFK",
		End:       "SFO",
		StartDate: "2025-03-01",
		EndDate:   "2025-03-15",
	})
	req, _ = http.NewRequest("POST", "/api/search", bytes.NewBuffer(validBody))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 for valid search, got %d", w.Code)
	}
}
