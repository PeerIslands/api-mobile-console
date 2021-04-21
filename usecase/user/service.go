package user

import (
	"errors"
	"mongo-admin-backend/entity"
	"time"
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

//CreateUser Create an user
func (s *Service) CreateUser(email, password, name, publicKey, privateKey string) (entity.User, error) {
	u, err := s.GetUser(email)
	if err != nil {
		return entity.User{}, err
	}

	if u.Name != "" {
		return entity.User{}, errors.New("user already exists")
	}

	e, err := entity.NewUser(email, password, name, publicKey, privateKey)
	if err != nil {
		return entity.User{}, err
	}


	user, err := s.repo.Create(e)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

// GetUser Get an user.
func (s *Service) GetUser(email string) (*entity.User, error) {
	return s.repo.Get(email)
}


// GetUser Get an user.
func (s *Service) GetAtlasParams(email string) (*entity.User, error) {
	return s.repo.Get(email)
}


//DeleteUser Delete an user
func (s *Service) DeleteUser(email string) error {
	u, err := s.GetUser(email)
	if u == nil {
		return errors.New("User not found")
	}
	if err != nil {
		return err
	}

	return s.repo.Delete(email)
}

//UpdateUser Update an user
func (s *Service) UpdateUser(e *entity.User) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
