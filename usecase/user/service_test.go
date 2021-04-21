package user
//
//import (
//	"github.com/stretchr/testify/assert"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"mongo-admin-backend/entity"
//	"testing"
//	"time"
//)
//
//
//func newFixtureUser() *entity.User {
//	return &entity.User{
//		ID:        primitive.NewObjectID(),
//		Email:     "ozzy@metalgods.net",
//		Password:  "123456",
//		Name: "Ozzy",
//		CreatedAt: time.Now(),
//	}
//}
//
//func Test_Create(t *testing.T) {
//	repo := newInmem()
//	m := NewService(repo)
//	u := newFixtureUser()
//	_, err := m.CreateUser(u.Email, u.Password, u.Name, u.AtlasParams.PublicKey, u.AtlasParams.PrivateKey)
//	assert.Nil(t, err)
//	assert.False(t, u.CreatedAt.IsZero())
//	assert.True(t, u.UpdatedAt.IsZero())
//}
//
////func Test_Get(t *testing.T) {
////	repo := newInmem()
////	m := NewService(repo)
////	u1 := newFixtureUser()
////	u2 := newFixtureUser()
////	u2.Name = "Lemmy"
////
////	user, _ := m.CreateUser(u1.Email, u1.Password, u1.Name, u1.AtlasParams.PublicKey, u1.AtlasParams.PrivateKey)
////	_, _ = m.CreateUser(u2.Email, u2.Password, u2.Name, u2.AtlasParams.PublicKey, u2.AtlasParams.PrivateKey)
////
////
////	t.Run("get", func(t *testing.T) {
////		saved, err := m.GetUser(user.Email)
////		assert.Nil(t, err)
////		assert.Equal(t, u1.Name, saved.Name)
////	})
////}
////
////	err = m.DeleteUser(u1.ID)
////	assert.Equal(t, entity.ErrNotFound, err)
////	id, _ := m.CreateUser(u3.Email, u3.Password, u3.FirstName, u3.LastName)
////func Test_Update(t *testing.T) {
////	repo := newInmem()
////	m := NewService(repo)
////	u := newFixtureUser()
////	id, err := m.CreateUser(u.Email, u.Password, u.FirstName, u.LastName)
////	assert.Nil(t, err)
////	saved, _ := m.GetUser(id)
////	saved.FirstName = "Dio"
////	saved.Books = append(saved.Books, entity.NewID())
////	assert.Nil(t, m.UpdateUser(saved))
////	updated, err := m.GetUser(id)
////	assert.Nil(t, err)
////	assert.Equal(t, "Dio", updated.FirstName)
////	assert.False(t, updated.UpdatedAt.IsZero())
////	assert.Equal(t, 1, len(updated.Books))
////}
////
////func TestDelete(t *testing.T) {
////	repo := newInmem()
////	m := NewService(repo)
////	u1 := newFixtureUser()
////	u2 := newFixtureUser()
////	u2ID, _ := m.CreateUser(u2.Email, u2.Password, u2.FirstName, u2.LastName)
////
////
////	err = m.DeleteUser(u2ID)
////	assert.Nil(t, err)
////	_, err = m.GetUser(u2ID)
////	assert.Equal(t, entity.ErrNotFound, err)
////
////	u3 := newFixtureUser()
////	saved, _ := m.GetUser(id)
////	saved.Books = []entity.ID{entity.NewID()}
////	_ = m.UpdateUser(saved)
////	err = m.DeleteUser(id)
////	assert.Equal(t, entity.ErrCannotBeDeleted, err)
////}