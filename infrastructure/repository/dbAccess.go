package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"mongo-admin-backend/api/presenter"
	"mongo-admin-backend/entity"
	"mongo-admin-backend/infrastructure/database"
	"mongo-admin-backend/pkg/contextWrapper"
	"mongo-admin-backend/pkg/digest"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// mongoDB repo.
type DatabaseAccessMongoDB struct {
	client     *mongo.Client
	collection *mongo.Collection
	apiURL     string
}

func NewDatabaseAccessMongoDB(client *mongo.Client, apiUrl string) *DatabaseAccessMongoDB {
	collection := client.Database("mongoDbAdmin").Collection("databaseAccessRequest")
	if apiUrl != "" {
		return &DatabaseAccessMongoDB{
			client:     client,
			collection: collection,
			apiURL:     apiUrl,
		}
	} else {
		return &DatabaseAccessMongoDB{
			client:     client,
			collection: collection,
		}
	}
}

func (r *DatabaseAccessMongoDB) Get() (*[]entity.DBAccessRequest, error) {
	return getAllDBRequest(contextWrapper.Ctx, r.collection)
}

func (r *DatabaseAccessMongoDB) GetOne(id string) (*entity.DBAccessRequest, error) {
	return getDBRequest(contextWrapper.Ctx, r.collection, id)
}

func (r *DatabaseAccessMongoDB) CreateRequest(u, p string, model *entity.DBAccessRequest) (*presenter.DBAccessRequest, *presenter.ErrorDetail, error) {
	return createDBRequest(contextWrapper.Ctx, r.collection, u, p, r.apiURL, model)
}

// Update an user.
func (r *DatabaseAccessMongoDB) Update(e *entity.DBAccessRequest) error {
	fmt.Println(e)

	return nil
}

// Delete an user.
func (r *DatabaseAccessMongoDB) Delete(e *entity.DBAccessRequest) error {
	fmt.Println(e)

	return nil
}

func createDBRequest(ctx context.Context, collection *mongo.Collection, u, p, baseURL string, model *entity.DBAccessRequest) (*presenter.DBAccessRequest, *presenter.ErrorDetail, error) {
	groupId := model.Groupid
	// fmt.Println(groupId)
	// fmt.Println(model)
	// fmt.Println(u)
	// fmt.Println(p)

	digestor := digest.NewDigestor(u, p)
	authDB, _, _ := database.DBCredentialsVal.GetDBCredentials()
	model.Databasename = authDB
	path := "/groups" + "/" + groupId + "/" + "databaseUsers"
	//reqBody, err1 := model.CreateJSON()
	//fmt.Println(reqBody)
	fmt.Println(authDB)
	reqBody, err1 := model.CreateJSON()
	var request = make(map[string]interface{})
	json.Unmarshal(reqBody, &request)
	fmt.Println(request)
	result, err2 := digestor.Digest(baseURL, path, "POST", reqBody)
	var dbaccessList presenter.DBAccessRequest
	var errDetails presenter.ErrorDetail
	var val map[string]interface{}

	_ = json.Unmarshal(result, &val)

	jsonEncoded, _ := json.Marshal(val)
	_ = json.Unmarshal(jsonEncoded, &dbaccessList)
	_ = json.Unmarshal(jsonEncoded, &errDetails)

	if err2 == nil {
		return &dbaccessList, &presenter.ErrorDetail{}, err1
	} else {
		fmt.Println("The output is ")
		fmt.Println(dbaccessList)
		return &dbaccessList, &errDetails, err1
	}
}

func getAllDBRequest(ctx context.Context, collection *mongo.Collection) (*[]entity.DBAccessRequest, error) {

	cursor, _ := collection.Find(ctx, bson.M{})

	var aList []entity.DBAccessRequest

	for cursor.Next(ctx) {
		var elem entity.DBAccessRequest
		err := cursor.Decode(&elem)
		if err != nil {
			fmt.Println("Decoding failed")
		}
		aList = append(aList, elem)
	}

	return &aList, nil
}

func getDBRequest(ctx context.Context, collection *mongo.Collection, id string) (*entity.DBAccessRequest, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("Invalid Request ID")
	}
	singleResult := collection.FindOne(ctx, bson.M{"_id": objectId})
	var dbAccesslist entity.DBAccessRequest
	singleResult.Decode(&dbAccesslist)
	return &dbAccesslist, nil
}

func (r *DatabaseAccessMongoDB) Create(e *entity.DBAccessRequest) error {
	return createOneDBAccessRequest(contextWrapper.Ctx, r.collection, e)
}

func createOneDBAccessRequest(ctx context.Context, collection *mongo.Collection, ent *entity.DBAccessRequest) error {
	result, err := collection.InsertOne(ctx, ent)
	if err != nil {
		return err
	} else {
		if result.InsertedID == nil {
			return errors.New("Cannot insert")
		} else {
			return nil
		}
	}
}
