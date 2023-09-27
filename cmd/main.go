package main

import (
	"fmt"
	"github.com/anilozgok/gp-prototype/internal/config"
	"github.com/anilozgok/gp-prototype/internal/log"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var (
	QUEUE_NAME = "gp_prototype_queue"
)

func main() {

	cfg, err := config.Get("./configs")
	if err != nil {
		log.Logger().Fatal("failed to read configs", zap.Error(err))
	}

	connStr := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		cfg.Secrets.RabbitMqCredentials.Username,
		cfg.Secrets.RabbitMqCredentials.Password,
		cfg.RabbitMq.Host,
		cfg.RabbitMq.Port,
	)

	conn, err := amqp.Dial(connStr)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	log.Logger().Info("Successfully Connected to our RabbitMQ Instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		QUEUE_NAME,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		QUEUE_NAME,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("gp prototype rabbit mq message"),
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	log.Logger().Info("Successfully Published Message to Queue")
}
