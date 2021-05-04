package group

import (
	"mongo-admin-backend/api/presenter"
)

//Reader interface
type Reader interface {
	Get(publicKey, privateKey string) ([]presenter.MongoGroups, error)
}


//UseCase interface
type UseCase interface {
	Reader
}