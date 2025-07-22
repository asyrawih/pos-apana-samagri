package main

import (
	"pos-apana-samagri/internal/config"
	"pos-apana-samagri/internal/models"
	"pos-apana-samagri/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Transaction{})
}

func main() {
	r := gin.Default()
	config := config.LoadConfig()
	port := config.Server.Port

	logger.Info("Server starting on port %s", zap.String("port", port))

	db, err := config.Database.Connect(&gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return
	}

	migrate(db)

	if err := r.Run(":" + port); err != nil {
		logger.Error("Failed to start server", zap.Error(err))
		return
	}
}
