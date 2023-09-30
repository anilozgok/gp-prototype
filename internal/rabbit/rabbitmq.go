package rabbit

import (
	"fmt"
	"github.com/anilozgok/gp-prototype/internal/config"
	"github.com/anilozgok/gp-prototype/internal/log"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type RabbitClient struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
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
		return nil, err
	}
	log.Logger().Info("successfully created rabbit client")

	return &RabbitClient{Conn: conn}, nil
}

func (r *RabbitClient) CloseConnection() {
	r.Conn.Close()
}

func (r *RabbitClient) OpenChannel() error {
	ch, err := r.Conn.Channel()
	if err != nil {
		log.Logger().Fatal("failed to open a channel", zap.Error(err))
		return err
	}
	r.Ch = ch
	log.Logger().Info("successfully opened a channel")

	return nil
}

func (r *RabbitClient) CloseChannel() {
	r.Ch.Close()
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
		return err
	}
	log.Logger().Info("successfully published a message", zap.String("queue", name))

	return nil
}
