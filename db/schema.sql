-- Triangle Travel SQLite Schema

-- City codes to airport IATA codes
CREATE TABLE IF NOT EXISTS iata_cities (
    city_code TEXT NOT NULL,
    airport_code TEXT NOT NULL,
    PRIMARY KEY (city_code, airport_code)
);

CREATE INDEX IF NOT EXISTS idx_iata_cities_city ON iata_cities(city_code);

-- Distances between airports in miles
CREATE TABLE IF NOT EXISTS distances (
    from_iata TEXT NOT NULL,
    to_iata TEXT NOT NULL,
    distance_miles REAL NOT NULL,
    PRIMARY KEY (from_iata, to_iata)
);

CREATE INDEX IF NOT EXISTS idx_distances_from ON distances(from_iata);
CREATE INDEX IF NOT EXISTS idx_distances_to ON distances(to_iata);

-- City routes by alliance
CREATE TABLE IF NOT EXISTS city_routes (
    city_iata TEXT NOT NULL,
    alliance TEXT NOT NULL,
    route_to TEXT NOT NULL,
    PRIMARY KEY (city_iata, alliance, route_to)
);

CREATE INDEX IF NOT EXISTS idx_city_routes_city ON city_routes(city_iata);
CREATE INDEX IF NOT EXISTS idx_city_routes_alliance ON city_routes(city_iata, alliance);

-- TODO(loyalty-routing): Transfer partner rules and baseline cents-per-point (CPP) valuations
-- CREATE TABLE IF NOT EXISTS loyalty_transfer_partners (
--     source_currency TEXT NOT NULL,   -- e.g. 'CHASE_UR', 'BILT', 'AMEX_MR'
--     partner_program TEXT NOT NULL,   -- e.g. 'AVIOS', 'AADVANTAGE', 'UNITED_MILEAGEPLUS'
--     alliance TEXT NOT NULL,          -- 'ONE_WORLD', 'STAR_ALLIANCE', 'SKY_TEAM', 'NONE'
--     transfer_ratio REAL NOT NULL DEFAULT 1.0,
--     baseline_cpp REAL NOT NULL DEFAULT 1.5,
--     PRIMARY KEY (source_currency, partner_program)
-- );

-- TODO(multi-modal-bridging): Regional rail & one-way car rental links between nearby hubs
-- CREATE TABLE IF NOT EXISTS ground_bridges (
--     from_iata TEXT NOT NULL,
--     to_iata TEXT NOT NULL,
--     mode TEXT NOT NULL,              -- 'RAIL', 'CAR_RENTAL', 'BUS'
--     distance_miles REAL NOT NULL,
--     est_transit_mins INTEGER NOT NULL,
--     est_cash_cost_usd REAL NOT NULL,
--     PRIMARY KEY (from_iata, to_iata, mode)
-- );

-- TODO(award-cache): Cached saver award inventory & cash positioning fares for arbitrage scoring
-- CREATE TABLE IF NOT EXISTS cached_award_fares (
--     from_iata TEXT NOT NULL,
--     to_iata TEXT NOT NULL,
--     flight_date TEXT NOT NULL,
--     cabin TEXT NOT NULL,             -- 'ECONOMY', 'PREMIUM_ECONOMY', 'BUSINESS', 'FIRST'
--     program TEXT NOT NULL,           -- e.g. 'AVIOS', 'AEROPLAN'
--     points_cost INTEGER NOT NULL,
--     taxes_fees_usd REAL NOT NULL,
--     cash_equivalent_usd REAL NOT NULL,
--     seats_available INTEGER NOT NULL DEFAULT 1,
--     updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
--     PRIMARY KEY (from_iata, to_iata, flight_date, cabin, program)
-- );
