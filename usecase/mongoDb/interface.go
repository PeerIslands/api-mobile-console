package mongoDb

import "mongo-admin-backend/entity"

//Reader interface
type Reader interface {
	Get(publicKey, privateKey string) ([]entity.MongoGroups, error)
}


//UseCase interface
type UseCase interface {
	Reader
}