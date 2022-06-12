package saga

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/payService"
	"log"
	"strconv"
)

type orderPayer struct {
	payService       *payService.PayService
	producer         sarama.SyncProducer
	successTopicName string
	errorTopicName   string
}

func NewOrderPayer(
	payService *payService.PayService,
	producer sarama.SyncProducer,
	successTopicName string,
	errorTopicName string,
) *orderPayer {
	return &orderPayer{
		payService:       payService,
		producer:         producer,
		successTopicName: successTopicName,
		errorTopicName:   errorTopicName,
	}
}

func (s *orderPayer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *orderPayer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *orderPayer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			confirmMessage := &OrderToConfirmMessage{}

			err := json.Unmarshal(message.Value, confirmMessage)
			if err != nil {
				return err
			}

			withdrawErr := s.payService.WithdrawPayment(confirmMessage.OrderID)
			if withdrawErr != nil {
				_, _, err := s.producer.SendMessage(&sarama.ProducerMessage{
					Topic: s.errorTopicName,
					Key:   sarama.StringEncoder(strconv.Itoa(confirmMessage.OrderID)),
					Value: sarama.ByteEncoder(message.Value),
				})
				if err != nil {
					return err
				}
				log.Printf("Pay ERROR ID %v", confirmMessage.OrderID)
				continue
			}

			_, _, _ = s.producer.SendMessage(&sarama.ProducerMessage{
				Topic: s.successTopicName,
				Key:   sarama.StringEncoder(strconv.Itoa(confirmMessage.OrderID)),
				Value: sarama.ByteEncoder(message.Value),
			})
			if err != nil {
				return err
			}
			log.Printf("Pay SUCCESS ID %v", confirmMessage.OrderID)

		case <-session.Context().Done():
			return nil
		}
	}
}
