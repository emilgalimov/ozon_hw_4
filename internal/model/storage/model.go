package storage

type StoreItem struct {
	OrderID    int  `json:"order_id"`
	ProductID  int  `json:"product_id"`
	IsReserved bool `json:"is_reserved"`
}
