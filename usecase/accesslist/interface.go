package accesslist

import (
	"mongo-admin-backend/api/presenter"
	"mongo-admin-backend/entity"
)

//Reader interface
type Reader interface {
	Get() (*[]entity.NetworkAccessList, error)
	GetOne(id string) (*entity.NetworkAccessList, error)
}

//Writer user writer
type Writer interface {
	Create(e *entity.NetworkAccessList) error
	Delete(e *entity.NetworkAccessList) error
	Update(e *entity.NetworkAccessList) error
	CreateRequest(u, p string, model *presenter.NetworkAccessList) (*[]presenter.NetworkAccessList, *presenter.ErrorDetail, error)
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetAllNetworkAccessList() (*[]entity.NetworkAccessList, error)
	CreateNetworkAccessList(ent *entity.NetworkAccessList) error
	GetOneNetworkAccessRequest(id string) (*entity.NetworkAccessList, error)
	CreateNetworkAccessRequest(u, p string, ent *entity.NetworkAccessList) (*[]presenter.NetworkAccessList, *presenter.ErrorDetail, error)
}
