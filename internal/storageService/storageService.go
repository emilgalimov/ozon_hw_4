package storageService

import (
	"errors"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/model/storage"
	"reflect"
)

type Storage struct {
	repo repository
}

func NewStorage(repo repository) *Storage {
	return &Storage{
		repo: repo,
	}
}

func (s *Storage) Reserve(orderID int, productIDs []int) error {
	products := s.repo.GetUnreservedByOrderID(orderID)
	equal := equalProductSets(productIDs, products)
	if !equal {
		return errors.New("товары для резерва не совпадают со складом")
	}
	s.repo.ReserveByOrderID(orderID)
	return nil
}

func (s *Storage) Unreserve(orderID int) {
	s.repo.UnreserveByOrderID(orderID)
}

func equalProductSets(productIDs []int, products []storage.StoreItem) bool {
	if len(productIDs) != len(products) {
		return false
	}

	//инициализируется явно для корректной работы reflect.DeepEqual
	dbProductIds := []int{}

	for _, product := range products {
		dbProductIds = append(dbProductIds, product.ProductID)
	}

	return reflect.DeepEqual(productIDs, dbProductIds)
}
