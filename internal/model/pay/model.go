package pay

import "time"

type Payment struct {
	OrderID   int
	Sum       float32
	CreatedAt time.Time
}
