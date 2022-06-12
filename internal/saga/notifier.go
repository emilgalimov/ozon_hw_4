package saga

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/notifyService"
	"log"
)

type notifier struct {
	notifyService *notifyService.NotifyService
	producer      sarama.SyncProducer
	numOfTries    int
}

func NewNotifier(
	notifyService *notifyService.NotifyService,
	producer sarama.SyncProducer,
	numOfTries int,
) *notifier {
	return &notifier{
		notifyService: notifyService,
		producer:      producer,
		numOfTries:    numOfTries,
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

			confirmMessage := &OrderToConfirmMessage{}

			err := json.Unmarshal(message.Value, confirmMessage)
			if err != nil {
				return err
			}
			var withdrawErr error
			//TODO переделать в повтор транзакции
			for i := 0; i < 10; i++ {
				withdrawErr = s.notifyService.Notify(confirmMessage.OrderID)
				if withdrawErr == nil {
					log.Printf("Notify SUCCESS %v", confirmMessage.OrderID)
					return nil
				}
				log.Printf("Notify ERROR %v", confirmMessage.OrderID)
			}
			return withdrawErr

		case <-session.Context().Done():
			return nil
		}
	}
}
