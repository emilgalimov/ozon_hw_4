package payService

import "gitlab.ozon.dev/emilgalimov/homework-4/internal/model/pay"

type repository interface {
	Save(pay pay.Payment)
}
