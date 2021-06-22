package accesslist

import (
	"mongo-admin-backend/api/presenter"
	"mongo-admin-backend/entity"
)

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

func (s *Service) GetAllNetworkAccessList() (*[]entity.NetworkAccessList, error) {
	return s.repo.Get()
}

func (s *Service) GetOneNetworkAccessRequest(id string) (*entity.NetworkAccessList, error) {
	return s.repo.GetOne(id)
}

func (s *Service) CreateNetworkAccessRequest(u, p string, ent *entity.NetworkAccessList) (*[]presenter.NetworkAccessList, *presenter.ErrorDetail, error) {
	model := ent.ConvertToAPIRequest()
	return s.repo.CreateRequest(u, p, model)
}

func (s *Service) CreateNetworkAccessList(ent *entity.NetworkAccessList) error {
	return s.repo.Create(ent)
}

func (s *Service) UpdateNetworkRequestStatus(u, p, id, status string) (bool, error) {
	return s.repo.UpdateRequestStatus(u, p, id, status)
}
