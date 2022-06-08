package saga

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/storage"
	"log"
	"strconv"
)

type StorageReserver struct {
	storage          *storage.Storage
	producer         sarama.SyncProducer
	successTopicName string
}

func NewStorageReserver(
	storage *storage.Storage,
	producer sarama.SyncProducer,
	successTopicName string,
) *StorageReserver {
	return &StorageReserver{
		storage:          storage,
		producer:         producer,
		successTopicName: successTopicName,
	}
}

func (s *StorageReserver) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *StorageReserver) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *StorageReserver) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			confirmMessage := &OrderToConfirmMessage{}

			err := json.Unmarshal(message.Value, confirmMessage)
			if err != nil {
				return err
			}
			isSend := s.storage.WriteOff(confirmMessage.OrderID, confirmMessage.ProductIds)
			if isSend {
				par, off, err := s.producer.SendMessage(&sarama.ProducerMessage{
					Topic: s.successTopicName,
					Key:   sarama.StringEncoder(strconv.Itoa(confirmMessage.OrderID)),
					Value: sarama.ByteEncoder(message.Value),
				})
				if err != nil {
					return err
				}
				log.Printf("Message SUCCESS %v -> %v; %v", par, off, err)

			} else {
				log.Printf("Message ERROR: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
			}
		case <-session.Context().Done():
			return nil
		}
	}
}
