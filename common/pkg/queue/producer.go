package queue

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type Producer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	done    chan struct{}
	config  *ProducerConfig
}

type ProducerConfig struct {
	AmqpURI      *string
	Exchange     *string
	ExchangeType *string
	QueueName    *string
	RoutingKey   *string
}

func NewProducer(config *ProducerConfig) *Producer {
	producer := &Producer{
		done:   make(chan struct{}),
		config: config,
	}
	return producer
}

func (p Producer) Publish(ctx context.Context, body *string) (err error) {
	config := amqp.Config{
		Vhost:      "/",
		Properties: amqp.NewConnectionProperties(),
	}

	config.Properties.SetClientConnectionName("producer-" + time.Now().Format(time.DateTime))

	p.conn, err = amqp.DialConfig(*p.config.AmqpURI, config)
	if err != nil {
		return err
	}
	defer p.conn.Close()
	p.channel, err = p.conn.Channel()
	if err != nil {
		return err
	}
	defer p.channel.Close()
	if err = p.channel.ExchangeDeclare(
		*p.config.Exchange,     // name
		*p.config.ExchangeType, // type
		true,                   // durable
		false,                  // auto-delete
		false,                  // internal
		false,                  // noWait
		nil,                    // arguments
	); err != nil {
		return err
	}

	d, err := p.channel.QueueDeclare(
		*p.config.QueueName, // name of the queue
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // noWait
		nil,                 // arguments
	)

	if err := p.channel.QueueBind(
		d.Name,
		*p.config.RoutingKey,
		*p.config.Exchange,
		false,
		nil,
	); err != nil {
		return err
	}

	if err := p.channel.Confirm(false); err != nil {
		return err
	}

	if err = p.channel.PublishWithContext(
		ctx,
		*p.config.Exchange,
		*p.config.RoutingKey,
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(*body),
			DeliveryMode:    amqp.Persistent,
			Priority:        0,
		},
	); err != nil {
		return err
	}
	return nil
}
