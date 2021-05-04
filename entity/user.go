package entity

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	crypto2 "mongo-admin-backend/pkg/crypto"
	"time"
)

type User struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Key         string             `json:"key" bson:"key"`
	Email       string             `json:"email" bson:"email"`
	Name        string             `json:"name" bson:"name"`
	Password    string             `json:"password" bson:"password"`
	AtlasParams AtlasParams        `json:"atlas_params" bson:"atlas_params"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

//NewUser create a new user
func NewUser(email, password, name, publicKey, privateKey string) (*User, error) {
	var a = &AtlasParams{
		PublicKey: publicKey,
		PrivateKey: privateKey,
	}
	var u = &User{
		Email: email,
		Name:  name,
		AtlasParams: *a,
		Password: password,
	}
	u.Encrypt()
	err := u.Validate()
	if err != nil {
		return nil, err
	}
	return u, nil
}

type AtlasParams struct {
	PublicKey  string `json:"public_key" bson:"public_key"`
	PrivateKey string `json:"private_key" bson:"private_key"`
}

func (u *User) Encrypt() {
	u.Key = crypto2.GenerateKey()
	u.AtlasParams.PrivateKey = crypto2.Encrypt(u.AtlasParams.PrivateKey, u.Key)
	u.AtlasParams.PublicKey = crypto2.Encrypt(u.AtlasParams.PublicKey, u.Key)
	u.Password = crypto2.GetMD5Hash(u.Password)
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name can't be empty")
	} else if u.Email == "" {
		return errors.New("email can't be empty")
	} else if u.Password == "" {
		return errors.New("password can't be empty")
	} else if u.AtlasParams.PrivateKey == "" {
		return errors.New("privateKey can't be empty")
	} else if u.AtlasParams.PublicKey == "" {
		return errors.New("publicKey can't be empty")
	}

	return nil
}

func (u *User) ValidatePassword(p string) error {
	if u.Password == crypto2.GetMD5Hash(p) {
		return nil
	}

	return errors.New("email or password is wrong")
}