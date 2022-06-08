package storage

type Storage struct {
	repo repository
}

func NewStorage(repo repository) *Storage {
	return &Storage{
		repo: repo,
	}
}

func (s Storage) WriteOff(orderID int, productID []int) bool {
	return true
}
