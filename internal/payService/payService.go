package payService

type PayService struct {
	repo repository
}

func NewPayService(repo repository) *PayService {
	return &PayService{
		repo: repo,
	}
}

func (s PayService) WithdrawPayment(orderID int) error {
	return nil
}
