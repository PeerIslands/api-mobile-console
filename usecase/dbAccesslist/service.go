package dbaccesslist

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

func (s *Service) GetAllDBAccessList() (*[]entity.DBAccessRequest, error) {
	return s.repo.Get()
}

func (s *Service) GetOneDBAccessRequest(id string) (*entity.DBAccessRequest, error) {
	return s.repo.GetOne(id)
}

func (s *Service) CreateDBAccessRequest(u, p string, ent *entity.DBAccessRequest) (*presenter.DBAccessRequest, *presenter.ErrorDetail, error) {
	return s.repo.CreateRequest(u, p, ent)
}

func (s *Service) CreateDBAccessList(ent *entity.DBAccessRequest) error {
	return s.repo.Create(ent)
}
