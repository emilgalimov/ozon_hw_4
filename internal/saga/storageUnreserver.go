package saga

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/storageService"
	"log"
)

type StorageUnreserver struct {
	storage  *storageService.Storage
	producer sarama.SyncProducer
}

func NewStorageUnreserver(
	storage *storageService.Storage,
	producer sarama.SyncProducer,
) *StorageUnreserver {
	return &StorageUnreserver{
		storage:  storage,
		producer: producer,
	}
}

func (s *StorageUnreserver) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *StorageUnreserver) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *StorageUnreserver) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			confirmMessage := &OrderToConfirmMessage{}

			err := json.Unmarshal(message.Value, confirmMessage)
			if err != nil {
				return err
			}
			s.storage.Unreserve(confirmMessage.OrderID)
			log.Printf("Unreserve SUCCESS ID %v", confirmMessage.OrderID)
		case <-session.Context().Done():
			return nil
		}
	}
}
