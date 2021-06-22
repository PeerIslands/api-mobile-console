package repository

import (
	"context"
	"mongo-admin-backend/entity"
	"mongo-admin-backend/pkg/contextWrapper"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserMongoDB mongoDB repo.
type UserMongoDB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

//NewUserMongoDB create new repository
func NewUserMongoDB(client *mongo.Client) *UserMongoDB {
	collection := client.Database("mongoDbAdmin").Collection("users")

	return &UserMongoDB{
		client:     client,
		collection: collection,
	}
}

// Create an user.
func (r *UserMongoDB) Create(e *entity.User) (entity.User, error) {
	now := time.Now()
	e.CreatedAt = now
	e.UpdatedAt = now
	result, databaseErr := r.collection.InsertOne(contextWrapper.Ctx, e)
	if databaseErr != nil {
		return entity.User{}, databaseErr
	}
	e.ID = result.InsertedID.(primitive.ObjectID)
	return *e, nil
}

// Get an user.
func (r *UserMongoDB) Get(email string) (*entity.User, error) {
	return getUser(contextWrapper.Ctx, email, r.collection)
}

func getUser(ctx context.Context, email string, collection *mongo.Collection) (*entity.User, error) {
	var u entity.User

	_ = collection.FindOne(ctx, bson.M{
		"email": email,
	}).Decode(&u)

	return &u, nil
}

// Update an user.
func (r *UserMongoDB) Update(e *entity.User) error {
	return nil
}

// Delete an user.
func (r *UserMongoDB) Delete(email string) error {
	_, err := r.collection.DeleteOne(contextWrapper.Ctx, bson.M{
		"email": email,
	}, nil)

	return err
}

//type Reader interface {
//	Get(email string) (*entity.User, error)
//}
//
////Writer user writer
//type Writer interface {
//	Create(e *entity.User) (entity.User, error)
//	Update(e *entity.User) error
//	Delete(email string) error
//}
