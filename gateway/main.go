package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Service struct {
	Name string
	URL  string
}

var services = map[string]Service{
	"game":    {Name: "Game Service", URL: "http://localhost:8081"},
	"auth":    {Name: "Auth Service", URL: "http://localhost:8082"},
	"payment": {Name: "Payment Service", URL: "http://localhost:8083"},
}

func main() {
	// Load configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../configs") // Adjust this path as needed
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	r := gin.Default()

	// Middleware for all routes
	r.Use(corsMiddleware())
	r.Use(loggingMiddleware())

	// Routes
	r.Any("/game/*path", createReverseProxy("game"))
	r.Any("/auth/*path", createReverseProxy("auth"))
	r.Any("/payment/*path", createReverseProxy("payment"))

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	port := viper.GetString("gateway.port")
	fmt.Printf("API Gateway is running on port %s\n", port)
	r.Run(":" + port)
}

func createReverseProxy(serviceKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		service, exists := services[serviceKey]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
			return
		}

		remote, err := url.Parse(service.URL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing service URL"})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = c.Param("path")
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the request
		fmt.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL.Path)

		// Process request
		c.Next()

		// Log the response status
		fmt.Printf("Response Status: %d\n", c.Writer.Status())
	}
}
