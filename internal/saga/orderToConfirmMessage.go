package saga

type OrderToConfirmMessage struct {
	OrderID    int   `json:"order_id"`
	ProductIds []int `json:"product_ids"`
}
