package notifyService

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/model/notify"
	"math/rand"
	"time"
)

type NotifyService struct {
	repo repository
}

func NewNotifyService(repo repository) *NotifyService {
	return &NotifyService{
		repo: repo,
	}
}

func (s NotifyService) Notify(orderID int) error {
	message := fmt.Sprintf("Ваш заказ под номером %v выдан", orderID)
	err := externalSerivceSendingMock(message)
	if err != nil {
		return err
	}

	s.repo.Save(notify.SuccessNotification{
		OrderID:   orderID,
		Message:   message,
		CreatedAt: time.Now(),
	})
	return nil
}

func externalSerivceSendingMock(message string) error {
	if rand.Intn(2) == 1 {
		return errors.New("notification sending error")
	}
	logrus.Info(message)
	return nil
}
