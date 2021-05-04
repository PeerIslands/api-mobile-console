package groups

import (
	"mongo-admin-backend/api/presenter"
)

//Reader interface
type Reader interface {
	Get(publicKey, privateKey, groupId string) ([]presenter.Process, error)
}


//UseCase interface
type UseCase interface {
	Reader
}