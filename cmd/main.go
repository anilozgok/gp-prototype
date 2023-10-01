package main

import (
	"github.com/anilozgok/gp-prototype/internal/config"
	"github.com/anilozgok/gp-prototype/internal/handlers"
	"github.com/anilozgok/gp-prototype/internal/log"
	"github.com/anilozgok/gp-prototype/internal/rabbit"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Get("./configs")
	if err != nil {
		log.Logger().Fatal("failed to read configs", zap.Error(err))
	}

	rabbitClient, err := rabbit.New(cfg)
	if err != nil {
		log.Logger().Fatal("failed to create rabbit client", zap.Error(err))
	}
	defer rabbitClient.CloseConnection()

	handler := handlers.New(rabbitClient)

	app := fiber.New()
	router := app.Group("/api/v1")

	router.Get("/health", handler.HealthCheck)
	router.Post("/publish", handler.PublishMessage)

	if err = app.Listen(":8080"); err != nil {
		log.Logger().Fatal("failed to start server", zap.Error(err))
	}
}
