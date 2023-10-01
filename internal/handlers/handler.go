package handlers

import (
	"fmt"
	"github.com/anilozgok/gp-prototype/internal/log"
	"github.com/anilozgok/gp-prototype/internal/messages"
	"github.com/anilozgok/gp-prototype/internal/rabbit"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handler struct {
	rabbitClient *rabbit.RabbitClient
}

func New(client *rabbit.RabbitClient) *Handler {
	return &Handler{
		rabbitClient: client,
	}
}

func (h *Handler) PublishMessage(ctx *fiber.Ctx) error {
	if err := h.rabbitClient.OpenChannel(); err != nil {
		log.Logger().Fatal("failed to open channel", zap.Error(err))
		ctx.Status(500).SendString("failed to open channel")
		return err
	}
	defer h.rabbitClient.CloseChannel()

	if err := h.rabbitClient.DeclareQueue("test"); err != nil {
		log.Logger().Fatal("failed to declare queue", zap.Error(err))
		ctx.Status(500).SendString("failed to declare queue")
		return err
	}

	message := messages.Message{Message: "test message"}

	if err := h.rabbitClient.PublishMessage("test", message); err != nil {
		log.Logger().Fatal("failed to publish message", zap.Error(err))
		ctx.Status(500).SendString("failed to publish message")
		return err
	}

	return nil
}

func (h *Handler) HealthCheck(ctx *fiber.Ctx) error {
	ctx.Status(200).SendString(fmt.Sprintf("Service %s is healthy.", ctx.Route().Path))
	return nil
}
