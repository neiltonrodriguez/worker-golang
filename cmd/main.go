package main

import (
	"context"
	"net/http"
	"worker-api/config"
	"worker-api/internal/domain"
	"worker-api/internal/router"
	"worker-api/internal/worker"
	"worker-api/pkg/common"
	"worker-api/pkg/database"
	"worker-api/pkg/email"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	log.Info().Msg("Server started")

	// Load configuration
	if err := config.GlobalConfig.LoadVariables(); err != nil {
		log.Error().Msgf("Failed to load configuration: %v", err)
		return
	}

	// Connect to database
	dbConfig := database.Config{
		Host:     config.GlobalConfig.Database.Host,
		Port:     config.GlobalConfig.Database.Port,
		Username: config.GlobalConfig.Database.Username,
		Password: config.GlobalConfig.Database.Password,
		Database: config.GlobalConfig.Database.Database,
	}

	// Connect to database
	db := database.Connect(dbConfig)

	// Auto migrate the schema
	db.AutoMigrate(&domain.Order{})

	// Initialize SQS
	config.InitSQS(config.GlobalConfig.AWS.QueueURL)

	// Initialize Email Service
	emailConfig := email.Config{
		Host:     config.GlobalConfig.SMTP.Host,
		Port:     config.GlobalConfig.SMTP.Port,
		Username: config.GlobalConfig.SMTP.Username,
		Password: config.GlobalConfig.SMTP.Password,
		From:     config.GlobalConfig.SMTP.From,
	}
	emailService := email.NewEmailService(emailConfig)

	// Start the worker in a goroutine
	orderWorker := worker.NewOrderWorker(db, emailService)
	ctx := context.Background()
	go orderWorker.Start(ctx)

	// Initialize Echo framework
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	common.NewLogger()
	e.Use(common.LoggingMiddleware)

	// Register routes
	router.RegisterRoutes(e, db)

	// Start server
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Error().Msg("Error message: " + err.Error())
	}
}
