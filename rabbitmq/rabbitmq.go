package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	QueueName  string
}

func NewRabbitMQ(queueName, connectionString string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to RabbitMQ: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("Failed to open a channel: %s", err)
	}

	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("Failed to declare a queue: %s", err)
	}

	return &RabbitMQ{
		Connection: conn,
		Channel:    ch,
		QueueName:  queueName,
	}, nil
}

func (r *RabbitMQ) PublishMessage(message []byte) error {
	err := r.Channel.Publish(
		"",
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	return err
}

func (r *RabbitMQ) ConsumeMessages() (<-chan amqp.Delivery, error) {
	msgs, err := r.Channel.Consume(
		r.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	return msgs, err
}

func (r *RabbitMQ) Close() {
	r.Channel.Close()
	r.Connection.Close()
}
