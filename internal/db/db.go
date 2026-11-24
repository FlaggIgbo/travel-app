package db

import (
	"database/sql"
	"path/filepath"

	_ "modernc.org/sqlite"
)

// DB wraps our SQLite database connection
type DB struct {
	*sql.DB
}

// RouteRepository abstracts route graph and distance lookups so data ingestion
// (cached SQLite, OpenFlights/airline timetable datasets, or live flight APIs)
// is decoupled from core triangle optimization logic.
//
// TODO(step-1): Migrate internal/flights and internal/api to depend on RouteRepository
// instead of concrete *db.DB, and add an external schedule/API adapter implementation.
type RouteRepository interface {
	GetCity(airportCode string) (string, error)
	SameCity(airport1, airport2 string) (bool, error)
	GetAirportsForCity(cityCode string) ([]string, error)
	GetDistancesFrom(fromAirport string) (map[string]float64, error)
	GetRoutesFrom(cityOrAirport, alliance string) ([]string, error)
	GetRoutesFromWithFallback(cityOrAirport, alliance string) ([]string, error)
	GetGroundBridgesFrom(fromAirport string, maxTransitMinutes int) ([]GroundBridge, error)
	GetTransferPartners(currency string) ([]TransferPartnerRule, error)
}

// GroundBridge represents a regional rail or one-way car rental link between nearby airports.
type GroundBridge struct {
	FromIATA       string  `json:"fromIata"`
	ToIATA         string  `json:"toIata"`
	Mode           string  `json:"mode"` // e.g., "RAIL", "CAR_RENTAL", "BUS"
	DistanceMiles  float64 `json:"distanceMiles"`
	EstTransitMins int     `json:"estTransitMins"`
	EstCashCostUSD float64 `json:"estCashCostUsd"`
}

// TransferPartnerRule encodes credit card or loyalty point transfer ratios and alliance mappings.
type TransferPartnerRule struct {
	SourceCurrency  string  `json:"sourceCurrency"`  // e.g., "CHASE_UR", "BILT", "AMEX_MR"
	PartnerProgram  string  `json:"partnerProgram"`  // e.g., "AVIOS", "AADVANTAGE", "AEROPLAN"
	Alliance        string  `json:"alliance"`        // e.g., "ONE_WORLD", "STAR_ALLIANCE", "SKY_TEAM"
	TransferRatio   float64 `json:"transferRatio"`   // e.g., 1.0 for 1:1 transfers
	BaselineCPP     float64 `json:"baselineCpp"`     // Baseline cents-per-point threshold
}

// Ensure *DB satisfies RouteRepository at compile time.
var _ RouteRepository = (*DB)(nil)

// New opens a connection to the SQLite database
func New(dataDir string) (*DB, error) {
	dbPath := filepath.Join(dataDir, "db", "data.sqlite3")
	database, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	// Check if the connection works
	err = database.Ping()
	if err != nil {
		database.Close()
		return nil, err
	}

	return &DB{database}, nil
}

// GetCity returns the city code for an airport code
func (d *DB) GetCity(airportCode string) (string, error) {
	var cityCode string
	query := "SELECT city_code FROM iata_cities WHERE airport_code = ? LIMIT 1"
	err := d.QueryRow(query, airportCode).Scan(&cityCode)
	if err == sql.ErrNoRows {
		// If not found, use airport code itself
		return airportCode, nil
	}
	if err != nil {
		return "", err
	}
	return cityCode, nil
}

// SameCity checks if two airports are in the same city
func (d *DB) SameCity(airport1, airport2 string) (bool, error) {
	if airport1 == airport2 {
		return true, nil
	}

	city1, err := d.GetCity(airport1)
	if err != nil {
		return false, err
	}

	city2, err := d.GetCity(airport2)
	if err != nil {
		return false, err
	}

	return city1 == city2, nil
}

// GetAirportsForCity returns all airport codes in a city
func (d *DB) GetAirportsForCity(cityCode string) ([]string, error) {
	query := "SELECT airport_code FROM iata_cities WHERE city_code = ? ORDER BY airport_code"
	rows, err := d.Query(query, cityCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var airports []string
	for rows.Next() {
		var code string
		if err := rows.Scan(&code); err == nil {
			airports = append(airports, code)
		}
	}

	// If no airports found, return cityCode itself
	if len(airports) == 0 {
		airports = append(airports, cityCode)
	}

	return airports, nil
}

// GetDistancesFrom returns a map of destination -> distance in miles
func (d *DB) GetDistancesFrom(fromAirport string) (map[string]float64, error) {
	query := "SELECT to_iata, distance_miles FROM distances WHERE from_iata = ?"
	rows, err := d.Query(query, fromAirport)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	distances := make(map[string]float64)
	for rows.Next() {
		var toAirport string
		var distance float64
		if err := rows.Scan(&toAirport, &distance); err == nil {
			distances[toAirport] = distance
		}
	}

	return distances, nil
}

// GetRoutesFrom returns reachable destinations for a city and alliance
func (d *DB) GetRoutesFrom(cityOrAirport, alliance string) ([]string, error) {
	query := "SELECT route_to FROM city_routes WHERE city_iata = ? AND alliance = ? ORDER BY route_to"
	rows, err := d.Query(query, cityOrAirport, alliance)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []string
	for rows.Next() {
		var destination string
		if err := rows.Scan(&destination); err == nil {
			routes = append(routes, destination)
		}
	}

	return routes, nil
}

// GetRoutesFromWithFallback checks city first, then each airport in that city
func (d *DB) GetRoutesFromWithFallback(cityOrAirport, alliance string) ([]string, error) {
	// Try direct lookup first
	routes, err := d.GetRoutesFrom(cityOrAirport, alliance)
	if err != nil || len(routes) > 0 {
		return routes, err
	}

	// Otherwise, find all airports in the city
	airports, err := d.GetAirportsForCity(cityOrAirport)
	if err != nil {
		return nil, err
	}

	// Collect unique routes
	seen := make(map[string]bool)
	var allRoutes []string

	for _, airport := range airports {
		airportRoutes, _ := d.GetRoutesFrom(airport, alliance)
		for _, destination := range airportRoutes {
			if !seen[destination] {
				seen[destination] = true
				allRoutes = append(allRoutes, destination)
			}
		}
	}

	return allRoutes, nil
}

// GetGroundBridgesFrom returns multi-modal ground transit links (regional rail, one-way car rental)
// from an airport that unlock alternative departure hubs for open-jaw award redemptions.
func (d *DB) GetGroundBridgesFrom(fromAirport string, maxTransitMinutes int) ([]GroundBridge, error) {
	// TODO(multi-modal-bridging): Query ground_bridges table (or compute from distances + rail schedules)
	// instead of relying solely on straight-line 55-300 mile radius checks.
	return nil, nil
}

// GetTransferPartners returns loyalty transfer partner rules and alliance networks for a given point currency.
func (d *DB) GetTransferPartners(currency string) ([]TransferPartnerRule, error) {
	// TODO(loyalty-routing): Query loyalty_transfer_partners table to filter out low-value redemptions
	// and match transfer currencies (Chase, Bilt, Avios, AAdvantage) with alliance direct routes.
	return nil, nil
}
