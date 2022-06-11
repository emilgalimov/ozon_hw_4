package notifyService

import (
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/model/notify"
)

type repository interface {
	Save(notify.SuccessNotification)
}
