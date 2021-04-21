package user

import (
	"errors"
	"fmt"
	"mongo-admin-backend/entity"
)


//inmem in memory repo
type inmem struct {
	m map[string]*entity.User
}

//newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*entity.User{}
	return &inmem{
		m: m,
	}
}

//Create an user
func (r *inmem) Create(e *entity.User) (entity.User, error) {
	r.m[e.Email] = e
	return *e, nil
}

//Get an user
func (r *inmem) Get(email string) (*entity.User, error) {
	if r.m[email] == nil {
		return nil, errors.New("Not found")
	}
	return r.m[email], nil
}

//Update an user
func (r *inmem) Update(e *entity.User) error {
	_, err := r.Get(e.Email)
	if err != nil {
		return err
	}
	r.m[e.Email] = e
	return nil
}

//Delete an user
func (r *inmem) Delete(email string) error {
	if r.m[email] == nil {
		return fmt.Errorf("not found")
	}
	r.m[email] = nil
	return nil
}
