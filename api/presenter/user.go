package presenter

import "go.mongodb.org/mongo-driver/bson/primitive"

//User data
type User struct {
	ID      primitive.ObjectID `json:"id"`
	Email   string             `json:"email"`
	Name    string             `json:"name"`
	IsAdmin bool               `json:"is_admin"`
}

type Token struct {
	Token string `json:"token"`
}
