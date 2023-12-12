package amqp

import (
	model "blog.hideyoshi.top/common/pkg/db/model/msg"
	"blog.hideyoshi.top/common/pkg/queue"
	"blog.hideyoshi.top/common/utils"
	"blog.hideyoshi.top/msg/config"
	"blog.hideyoshi.top/msg/pkg/third/sms/sms_aliyun"
	"context"
	"encoding/json"
	"github.com/alibabacloud-go/tea/tea"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"strconv"
	"time"
)

type SmsQueue struct {
}

type SmsBody struct {
	Info     *model.MsgGroup       `json:"info"`
	Users    []*model.MsgGroupUser `json:"users"`
	Template *model.MsgTemplate    `json:"template"`
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
			err := s.SendSms(queueData, errCh)
			if err != nil {
				log.Println("sendSms err:", err)
			}
		case <-chanDown:
			break consumerRunning
		}
	}
	errCh <- nil
}

func (s *SmsQueue) SendSms(queueData amqp.Delivery, errCh chan<- error) error {
	var err error

	sms := &SmsBody{}
	err = json.Unmarshal(queueData.Body, sms)
	if err != nil {
		return err
	}
	arrLen := len(sms.Users)
	phoneNumberArr := make([]string, arrLen)
	signArr := make([]string, arrLen)
	paramArr := make([]map[string]string, arrLen)
	for i, user := range sms.Users {
		if user != nil {
			phoneNumberArr[i] = user.PhoneNo
			signArr[i] = sms.Template.TemplateSign
			paramArr[i] = map[string]string{
				"name": user.UserName,
			}
		}
	}
	phoneNumberJson, _ := json.Marshal(phoneNumberArr)
	signNameJson, _ := json.Marshal(signArr)
	paramJson, _ := json.Marshal(paramArr)
	yun := sms_aliyun.AliYun{}
	id := config.C.Sms.AccessKeyId
	secret := config.C.Sms.AccessKeySecret
	yun.SetAccessKeyId(id)
	yun.SetAccessKeySecret(secret)
	log.Println(yun)
	args := &sms_aliyun.AliYunBatchArgs{
		PhoneNumberJson:   tea.String(string(phoneNumberJson)),
		SignNameJson:      tea.String(string(signNameJson)),
		TemplateParamJson: tea.String(string(paramJson)),
		TemplateCode:      tea.String(sms.Template.TemplateCode),
		OutId:             tea.String(strconv.FormatInt(sms.Info.GroupId, 10)),
	}
	err = yun.SendBatchSms(args)
	if err != nil {
		return err
	}
	err = queueData.Ack(false)
	if err != nil {
		return err
	}
	return nil
}
