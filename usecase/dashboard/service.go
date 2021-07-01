package dashboard

//Service  interface
type Service struct {
	repo Repository
}

//NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetDashboardData(u, p, groupId string) (*map[string]interface{}, error) {
	return s.repo.Get(u, p, groupId)
}
