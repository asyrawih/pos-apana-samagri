package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"pos-apana-samagri/internal/config"
	"pos-apana-samagri/internal/handlers"
	"pos-apana-samagri/internal/models"
	"pos-apana-samagri/pkg/logger"
)

func migrate(db *gorm.DB) {
	// AutoMigrate all models - constraints are defined in the model struct tags
	modelsToMigrate := []interface{}{
		&models.Transaction{},
		&models.Customer{},
		&models.Inventory{},
		&models.User{},
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
	}
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

	transactionsHandler := handlers.NewTransactionHandler()

	r.GET("/transactions", transactionsHandler.ListTransactions)
	r.GET("/transactions/:id", transactionsHandler.GetTransaction)
	r.POST("/transactions", transactionsHandler.CreateTransaction)

	if err := r.Run(":" + port); err != nil {
		logger.Error("Failed to start server", zap.Error(err))
		return
	}
}
