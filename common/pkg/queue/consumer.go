package queue

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	done    chan error
	down    chan struct{}
	config  *ConsumerConfig
}

type ConsumerConfig struct {
	AmqpURI      *string
	Exchange     *string
	ExchangeType *string
	QueueName    *string
	RoutingKey   *string
	Tag          *string
}

func NewConsumer(consumerConfig *ConsumerConfig) *Consumer {
	return &Consumer{
		done:   make(chan error),
		config: consumerConfig,
	}
}

func (c *Consumer) RunConsumer(handler func(<-chan amqp.Delivery, chan<- error, <-chan struct{})) error {
	var err error
	config := amqp.Config{Properties: amqp.NewConnectionProperties()}
	config.Properties.SetClientConnectionName("consumer" + time.Now().Format(time.DateTime))
	log.Printf("dialing %q", c.config.AmqpURI)
	c.conn, err = amqp.DialConfig(*c.config.AmqpURI, config)
	if err != nil {
		return fmt.Errorf("Dial: %s", err)
	}

	c.channel, err = c.conn.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}

	if err = c.channel.ExchangeDeclare(
		*c.config.Exchange,     // name of the exchange
		*c.config.ExchangeType, // type
		true,                   // durable
		false,                  // delete when complete
		false,                  // internal
		false,                  // noWait
		nil,                    // arguments
	); err != nil {
		return fmt.Errorf("Exchange Declare: %s", err)
	}

	queue, err := c.channel.QueueDeclare(
		*c.config.QueueName, // name of the queue
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // noWait
		nil,                 // arguments
	)

	if err != nil {
		return fmt.Errorf("Queue Declare: %s", err)
	}

	if err = c.channel.QueueBind(
		queue.Name,           // name of the queue
		*c.config.RoutingKey, // bindingKey
		*c.config.Exchange,   // sourceExchange
		false,                // noWait
		nil,                  // arguments
	); err != nil {
		return fmt.Errorf("Queue Bind: %s", err)
	}

	deliveries, err := c.channel.Consume(
		queue.Name,    // name
		*c.config.Tag, // consumerTag,
		false,         // autoAck
		false,         // exclusive
		false,         // noLocal
		false,         // noWait
		nil,           // arguments
	)
	if err != nil {
		return fmt.Errorf("Queue Consume: %s", err)
	}

	go handler(deliveries, c.done, c.down)
	return nil
}

func (c *Consumer) Shutdown() error {
	c.down <- struct{}{}
	if err := <-c.done; err != nil {
		log.Println("down handler err:", err)
	}
	// will close() the deliveries channel
	if err := c.channel.Cancel(*c.config.Tag, true); err != nil {
		return err
	}

	if err := c.conn.Close(); err != nil {
		return err
	}

	// wait for handle() to exit
	return nil
}
