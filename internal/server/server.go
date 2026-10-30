package server

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"triangle_travel/internal/api"
	"triangle_travel/internal/db"
)

// Run starts the Triangle Travel web server
func Run() {
	// Optional flags for host and port
	host := flag.String("host", "0.0.0.0", "Server host")
	port := flag.String("port", "8080", "Server port")
	flag.Parse()

	// Check if cloud host provided a PORT environment variable
	if envPort := os.Getenv("PORT"); envPort != "" {
		*port = envPort
	}

	// Connect to database
	database, err := db.New(".")
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
	defer database.Close()

	// Create Gin router
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Setup API endpoints
	apiHandler := &api.API{DB: database}
	router.POST("/api/search", apiHandler.Search)
	router.GET("/api/cities", apiHandler.Cities)

	// Serve static frontend files if built
	buildFolder := filepath.Join(".", "build")
	if _, err := os.Stat(buildFolder); err == nil {
		router.Use(static.Serve("/", static.LocalFile(buildFolder, true)))
		router.NoRoute(func(c *gin.Context) {
			c.File(filepath.Join(buildFolder, "index.html"))
		})
	} else {
		router.GET("/", func(c *gin.Context) {
			c.String(200, "Triangle Travel API is running. Run 'npm run build' to see the frontend.")
		})
	}

	// Start listening
	address := fmt.Sprintf("%s:%s", *host, *port)
	fmt.Printf("Server running at http://%s\n", address)
	if err := router.Run(address); err != nil {
		log.Fatal("Server failed to run:", err)
	}
}
