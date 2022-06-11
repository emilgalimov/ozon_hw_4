package storageService

type Storage struct {
	repo repository
}

func NewStorage(repo repository) *Storage {
	return &Storage{
		repo: repo,
	}
}

func (s Storage) Reserve(orderID int, productID []int) error {
	return nil
}

func (s Storage) Unreserve(orderID int, productID []int) error {
	return nil
}
