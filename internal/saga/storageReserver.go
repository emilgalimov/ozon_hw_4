package saga

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/storageService"
	"log"
	"strconv"
)

type StorageReserver struct {
	storage          *storageService.Storage
	producer         sarama.SyncProducer
	successTopicName string
}

func NewStorageReserver(
	storage *storageService.Storage,
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
			reserveError := s.storage.Reserve(confirmMessage.OrderID, confirmMessage.ProductIds)
			if reserveError != nil {
				log.Printf("Reserve ERROR ID = %v", confirmMessage.OrderID)
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
			log.Printf("Reserve SUCCESS ID = %v", confirmMessage.OrderID)
		case <-session.Context().Done():
			return nil
		}
	}
}
