package dbaccesslist

import (
	"mongo-admin-backend/api/presenter"
	"mongo-admin-backend/entity"
)

//Reader interface
type Reader interface {
	Get() (*[]entity.DBAccessRequest, error)
	GetOne(id string) (*entity.DBAccessRequest, error)
}

//Writer user writer
type Writer interface {
	Create(e *entity.DBAccessRequest) error
	Delete(e *entity.DBAccessRequest) error
	Update(e *entity.DBAccessRequest) error
	UpdateRequestStatus(u, p, id, status string) (bool, error)
	CreateRequest(u, p string, model *entity.DBAccessRequest) (*presenter.DBAccessRequest, *presenter.ErrorDetail, error)
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetAllDBAccessList() (*[]entity.DBAccessRequest, error)
	CreateDBAccessList(ent *entity.DBAccessRequest) error
	GetOneDBAccessRequest(id string) (*entity.DBAccessRequest, error)
	CreateDBAccessRequest(u, p string, ent *entity.DBAccessRequest) (*presenter.DBAccessRequest, *presenter.ErrorDetail, error)
	UpdateDBAccessRequestStatus(u, p, id, status string) (bool, error)
}
