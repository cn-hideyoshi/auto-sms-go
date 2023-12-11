package amqp

import (
	model "blog.hideyoshi.top/common/pkg/db/model/msg"
	"blog.hideyoshi.top/common/pkg/queue"
	"blog.hideyoshi.top/common/utils"
	"blog.hideyoshi.top/msg/config"
	"context"
	"encoding/json"
	"github.com/alibabacloud-go/tea/tea"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type SmsQueue struct {
}

type SmsBody struct {
	Info  model.MsgGroup        `json:"info"`
	Users []*model.MsgGroupUser `json:"users"`
}

func (s *SmsQueue) Push(body string) {
	queueConfig := &queue.ProducerConfig{
		AmqpURI:      utils.BuildAmqpUri(config.C.Amqp),
		Exchange:     tea.String("hello"),
		ExchangeType: tea.String("topic"),
		QueueName:    tea.String("hello-1"),
		RoutingKey:   tea.String("h"),
	}
	producer := queue.NewProducer(queueConfig)
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	err := producer.Publish(timeout, &body)
	if err != nil {
		log.Println(err)
	}
}

func (s *SmsQueue) Consumer() {
	queueConfig := &queue.ConsumerConfig{
		AmqpURI:      utils.BuildAmqpUri(config.C.Amqp),
		Exchange:     tea.String("hello"),
		ExchangeType: tea.String("topic"),
		QueueName:    tea.String("hello-1"),
		RoutingKey:   tea.String("h"),
		Tag:          tea.String(time.Now().Format(time.DateTime)),
	}
	consumer := queue.NewConsumer(queueConfig)

	log.Println("sms consumer start...")
	err := consumer.RunConsumer(s.handler)
	if err != nil {
		log.Fatalln(err)
	}
}

func (s *SmsQueue) handler(deliveries <-chan amqp.Delivery, errCh chan<- error, chanDown <-chan struct{}) {
consumerRunning:
	for {
		select {
		case queueData := <-deliveries:
			s.SendSms(queueData, errCh)
		case <-chanDown:
			break consumerRunning
		}
	}
	errCh <- nil
}

func (s *SmsQueue) SendSms(queueData amqp.Delivery, errCh chan<- error) {
	var err error

	body := &SmsBody{}

	err = json.Unmarshal(queueData.Body, body)
	if err != nil {
		errCh <- err
	}
	err = queueData.Ack(true)
	if err != nil {
		errCh <- err
	}
}
