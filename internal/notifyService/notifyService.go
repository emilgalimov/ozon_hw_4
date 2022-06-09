package notifyService

type NotifyService struct {
	repo repository
}

func NewNotifyService(repo repository) *NotifyService {
	return &NotifyService{
		repo: repo,
	}
}

func (s NotifyService) Notify(orderID int) error {
	return nil
}
