package rabbit

import (
	"fmt"
	"github.com/anilozgok/gp-prototype/internal/config"
	"github.com/anilozgok/gp-prototype/internal/log"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

//TODO: we cannot re-open a channel after it is closed. we need to find a way to declare queue and publish message before channel is closed

type RabbitClient struct {
	Ch *amqp.Channel
}

func New(cfg *config.Config) (*RabbitClient, error) {

	connStr := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		cfg.Secrets.RabbitMqCredentials.Username,
		cfg.Secrets.RabbitMqCredentials.Password,
		cfg.RabbitMq.Host,
		cfg.RabbitMq.Port,
	)

	conn, err := amqp.Dial(connStr)
	if err != nil {
		log.Logger().Fatal("failed to connect to rabbitmq", zap.Error(err))
		return nil, err
	}
	defer conn.Close()
	log.Logger().Info("successfully connected to rabbitMq")

	ch, err := conn.Channel()
	if err != nil {
		log.Logger().Fatal("failed to create channel", zap.Error(err))
		return nil, err
	}
	defer ch.Close()
	log.Logger().Info("successfully created a channel")
	return &RabbitClient{Ch: ch}, nil
}

func (r *RabbitClient) DeclareQueue(name string) error {

	q, err := r.Ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Logger().Fatal("failed to declare a queue", zap.Error(err))
		return err
	}
	log.Logger().Info("successfully declared a queue", zap.String("queue", q.Name))

	return nil
}

func (r *RabbitClient) PublishMessage(name string, message string) error {
	err := r.Ch.Publish(
		"",
		name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Logger().Fatal("failed to publish a message", zap.Error(err))
		return err
	}
	log.Logger().Info("successfully published a message", zap.String("queue", name))

	return nil
}
