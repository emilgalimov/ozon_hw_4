package pay

import "time"

type Payments struct {
	OrderID   int
	Sum       float32
	CreatedAt time.Time
}
