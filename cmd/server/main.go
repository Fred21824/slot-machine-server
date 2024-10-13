package main

import (
	"fmt"
	"net/http"

	"slot-machine-server/internal/api"
	"slot-machine-server/internal/cache"
	"slot-machine-server/internal/db"
	"slot-machine-server/internal/logger"
	"slot-machine-server/internal/websocket"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	logger.Init()
	defer logger.Log.Sync()

	// Load configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Error("Error reading config file", zap.Error(err))
		return
	}

	// Initialize database connection
	if err := db.Initialize(); err != nil {
		logger.Error("Failed to initialize database", zap.Error(err))
		return
	}

	// Initialize Redis
	cache.InitRedis()

	// Setup routes
	router := api.SetupRoutes()

	// Add WebSocket handler
	router.HandleFunc("/ws", websocket.HandleWebSocket)

	// Start server
	port := viper.GetString("server.port")
	logger.Info("Starting server", zap.String("port", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		logger.Error("Failed to start server", zap.Error(err))
	}
}
