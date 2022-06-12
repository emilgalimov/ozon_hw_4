package saga

type OrderToConfirmMessage struct {
	OrderID    int   `json:"order_id"`
	ProductIds []int `json:"product_ids"`
}

type RepeatableOrderToConfirmMessage struct {
	OrderToConfirmMessage

	CurrentTry int `json:"current_try"`
}
