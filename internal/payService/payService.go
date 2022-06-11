package payService

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/model/pay"
	"math/rand"
	"time"
)

type PayService struct {
	repo repository
}

func NewPayService(repo repository) *PayService {
	return &PayService{
		repo: repo,
	}
}

func (p *PayService) WithdrawPayment(orderID int) error {
	sum := externalOrderServiceMockGetOrderSum(orderID)
	err := externalPaymentServiceMock(orderID, sum)
	if err != nil {
		return err
	}
	p.repo.Save(pay.Payment{
		OrderID:   orderID,
		Sum:       sum,
		CreatedAt: time.Now(),
	})

	return nil
}

func externalOrderServiceMockGetOrderSum(orderID int) float32 {
	return rand.Float32()
}

func externalPaymentServiceMock(orderID int, sum float32) error {
	if rand.Intn(2) == 1 {
		return errors.New("payment error")
	}
	logrus.Info(fmt.Sprintf("прошла оплата по заказу %v на сумму %v", orderID, sum))
	return nil
}
