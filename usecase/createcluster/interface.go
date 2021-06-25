package createcluster

import (
	"mongo-admin-backend/api/presenter"
	"mongo-admin-backend/entity"
)

type Reader interface {
	Get() *[]entity.AdminRequest
	GetOne(id string) *entity.AdminRequest
}

type Writer interface {
	Create(e *entity.UserRequest) error
	Delete(e *entity.UserRequest) error
	Update(e *entity.UserRequest) error
	UpdateRequestStatus(u, p, id, status string) (bool error)
	CreateRequest(u, p string, model *presenter.ClusterDetails)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetAllCreateClusterList() (*[]entity.AdminRequest, error)
	CreateNewClusterRequest(ent *entity.AdminRequest) error
	GetOneClusterCreateRequest(id string) *presenter.ClusterDetails
	SendClusterRequest(u, p string, ent *presenter.ClusterDetails) (*presenter.ClusterDetails, *presenter.ErrorDetail, error)
	UpdateAdminRequestStatus(u, p, id, status string) (bool, error)
}
