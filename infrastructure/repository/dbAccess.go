package repository

import (
	"context"
	"encoding/json"
	"errors"
	"mongo-admin-backend/config"

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
	collection := client.Database("mongoDbAdmin").Collection("mongoadminrequests")
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

func (r *DatabaseAccessMongoDB) UpdateRequestStatus(u, p, id, status string) (bool, error) {
	return updateOneRequestStatus(contextWrapper.Ctx, r.collection, id, status)
}

// Update an user.
func (r *DatabaseAccessMongoDB) Update(e *entity.DBAccessRequest) error {
	return nil
}

// Delete an user.
func (r *DatabaseAccessMongoDB) Delete(e *entity.DBAccessRequest) error {
	return nil
}

func createDBRequest(ctx context.Context, collection *mongo.Collection, u, p, baseURL string, model *entity.DBAccessRequest) (*presenter.DBAccessRequest, *presenter.ErrorDetail, error) {
	groupId := model.Groupid
	digestor := digest.NewDigestor(u, p)
	authDB, _, _ := database.DBCredentialsVal.GetDBCredentials()
	model.Databasename = authDB
	path := "/groups" + "/" + groupId + "/" + "databaseUsers"
	reqBody, err1 := model.CreateJSON()
	var request = make(map[string]interface{})
	json.Unmarshal(reqBody, &request)
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
		return &dbaccessList, &errDetails, err2
	}
}

func getAllDBRequest(ctx context.Context, collection *mongo.Collection) (*[]entity.DBAccessRequest, error) {

	cursor, _ := collection.Find(ctx, bson.M{"requestType": config.STR_REQ_TYPE_DB_ACCESS, "status": config.STR_REQ_STATUS_OPEN})

	var aList []entity.DBAccessRequest

	for cursor.Next(ctx) {
		var elem entity.UserRequest
		err := cursor.Decode(&elem)
		if err != nil {
		}
		aList = append(aList, elem.GetDBUserRequest())
	}

	return &aList, nil
}

func getDBRequest(ctx context.Context, collection *mongo.Collection, id string) (*entity.DBAccessRequest, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("Invalid Request ID")
	}
	singleResult := collection.FindOne(ctx, bson.M{"_id": objectId})
	var dbAccesslist entity.UserRequest
	singleResult.Decode(&dbAccesslist)
	returnModel := dbAccesslist.GetDBUserRequest()
	return &returnModel, nil
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

func updateOneRequestStatus(ctx context.Context, collection *mongo.Collection, id, status string) (bool, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, errors.New("Invalid ID string")
	}
	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": objectId},
		bson.D{{"$set", bson.D{{"status", status}}}})
	if result.ModifiedCount > 0 {
		return true, nil
	} else {
		return false, errors.New("Failed to update status")
	}
}
