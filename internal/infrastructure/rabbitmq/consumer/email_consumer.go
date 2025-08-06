package consumer

import (
	"encoding/json"
	"log"

	"github.com/CosmeticsShiraz/Backend/bootstrap"
	"github.com/CosmeticsShiraz/Backend/internal/domain/communication"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/rabbitmq"
)

type EmailConsumer struct {
	constants    *bootstrap.RabbitMQConstants
	rabbitMQ     *rabbitmq.RabbitMQ
	emailService communication.EmailService
}

func NewEmailConsumer(
	constants *bootstrap.RabbitMQConstants,
	rabbitMQ *rabbitmq.RabbitMQ,
	emailService communication.EmailService,
) *EmailConsumer {
	return &EmailConsumer{
		constants:    constants,
		rabbitMQ:     rabbitMQ,
		emailService: emailService,
	}
}

func (consumer *EmailConsumer) Start() error {
	return consumer.rabbitMQ.ConsumeMessages(consumer.constants.Events.NotificationsEmail, consumer.handleMessage)
}

func (consumer *EmailConsumer) handleMessage(body []byte) error {
	var msg struct {
		ToEmail      string      `json:"toEmail"`
		Subject      string      `json:"subject"`
		TemplateFile string      `json:"templateFile"`
		Data         interface{} `json:"data"`
	}
	if err := json.Unmarshal(body, &msg); err != nil {
		log.Printf("Failed to unmarshal email notification message: %v", err)
	}

	if err := consumer.emailService.SendEmail(msg.ToEmail, msg.Subject, msg.TemplateFile, msg.Data); err != nil {
		log.Printf("Failed to send email: %v", err)
	}
	return nil
}
