package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"

	"github.com/olteffe/balancex/balance/config"
)

// NewRabbitMQConn Initialize new RabbitMQ connection
func NewRabbitMQConn(cfg *config.Config) (*amqp.Connection, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RabbitMQ.User,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)
	return amqp.Dial(connAddr)
}
