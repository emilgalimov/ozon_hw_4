package saga

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/notifyService"
	"log"
	"strconv"
)

type notifier struct {
	notifyService   *notifyService.NotifyService
	producer        sarama.SyncProducer
	numOfTries      int
	repeatTopicName string
}

func NewNotifier(
	notifyService *notifyService.NotifyService,
	producer sarama.SyncProducer,
	numOfTries int,
	repeatTopicName string,
) *notifier {
	return &notifier{
		notifyService:   notifyService,
		producer:        producer,
		numOfTries:      numOfTries,
		repeatTopicName: repeatTopicName,
	}
}

func (s *notifier) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *notifier) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *notifier) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			confirmMessage := &RepeatableOrderToConfirmMessage{}

			err := json.Unmarshal(message.Value, confirmMessage)
			if err != nil {
				break
			}

			if confirmMessage.CurrentTry >= s.numOfTries {
				break
			}

			withdrawErr := s.notifyService.Notify(confirmMessage.OrderID)
			if withdrawErr != nil {
				confirmMessage.CurrentTry++
				repeatMessage, _ := json.Marshal(confirmMessage)
				log.Printf("Notify ERROR %v", confirmMessage.OrderID)
				_, _, _ = s.producer.SendMessage(&sarama.ProducerMessage{
					Topic: s.repeatTopicName,
					Key:   sarama.StringEncoder(strconv.Itoa(confirmMessage.OrderID)),
					Value: sarama.ByteEncoder(repeatMessage),
				})
				log.Printf("Notify REPEAT %v", confirmMessage.OrderID)
				break
			}

			log.Printf("Notify SUCCESS %v", confirmMessage.OrderID)

		case <-session.Context().Done():
			return nil
		}
	}
}
