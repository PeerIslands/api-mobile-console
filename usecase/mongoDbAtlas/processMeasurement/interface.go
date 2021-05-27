package groups

import (
	"mongo-admin-backend/api/presenter"
)

//Reader interface
type Reader interface {
	Get(publicKey, privateKey, groupId, processId string, qParams map[string]string) (presenter.Measurements, error)
}

//UseCase interface
type UseCase interface {
	Reader
}
