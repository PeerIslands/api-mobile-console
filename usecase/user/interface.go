package user

import "mongo-admin-backend/entity"

//Reader interface
type Reader interface {
	Get(email string) (*entity.User, error)
}

//Writer user writer
type Writer interface {
	Create(e *entity.User) (entity.User, error)
	Update(e *entity.User) error
	Delete(email string) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetUser(email string) (*entity.User, error)
	CreateUser(email, password, name, publicKey, privateKey string) (entity.User, error)
	UpdateUser(e *entity.User) error
	DeleteUser(email string) error
}