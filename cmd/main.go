package main

import (
	"github.com/anilozgok/gp-prototype/internal/config"
	"github.com/anilozgok/gp-prototype/internal/log"
	"github.com/anilozgok/gp-prototype/internal/rabbit"
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

	err = rabbitClient.DeclareQueue("test")
	if err != nil {
		log.Logger().Fatal("failed to declare queue", zap.Error(err))
	}

	err = rabbitClient.PublishMessage("test", "test message")
	if err != nil {
		log.Logger().Fatal("failed to publish message", zap.Error(err))
	}

}
