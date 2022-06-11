package notify

import "time"

type SuccessNotification struct {
	OrderID   int
	Message   string
	CreatedAt time.Time
}
