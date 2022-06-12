package storageService

import "gitlab.ozon.dev/emilgalimov/homework-4/internal/model/storage"

type repository interface {
	GetUnreservedByOrderID(orderID int) []storage.StoreItem
	ReserveByOrderID(orderID int)
	UnreserveByOrderID(orderID int)
}
