package main

import "github.com/gin-gonic/gin"

func main() {
    // Create a simple Gin router
    r := gin.Default()
    // Basic health check endpoint
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello, world!")
    })
    // Run server on localhost:8080
    r.Run("localhost:8080")
}
