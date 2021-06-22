package repository

import (
	"context"
	"encoding/json"
	"errors"
	"mongo-admin-backend/api/presenter"
	"mongo-admin-backend/config"
	"mongo-admin-backend/entity"
	"mongo-admin-backend/pkg/contextWrapper"
	"mongo-admin-backend/pkg/digest"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// mongoDB repo.
type NetworkAccessMongoDB struct {
	client     *mongo.Client
	collection *mongo.Collection
	apiURL     string
}

func NewNetworkAccessMongoDB(client *mongo.Client, apiUrl string) *NetworkAccessMongoDB {
	collection := client.Database("mongoDbAdmin").Collection("mongoadminrequests")
	if apiUrl != "" {
		return &NetworkAccessMongoDB{
			client:     client,
			collection: collection,
			apiURL:     apiUrl,
		}
	} else {
		return &NetworkAccessMongoDB{
			client:     client,
			collection: collection,
		}
	}
}

func (r *NetworkAccessMongoDB) Get() (*[]entity.NetworkAccessList, error) {
	return getAllRequest(contextWrapper.Ctx, r.collection)
}

func (r *NetworkAccessMongoDB) GetOne(id string) (*entity.NetworkAccessList, error) {
	return getRequest(contextWrapper.Ctx, r.collection, id)
}

func (r *NetworkAccessMongoDB) CreateRequest(u, p string, model *presenter.NetworkAccessList) (*[]presenter.NetworkAccessList, *presenter.ErrorDetail, error) {
	return createRequest(contextWrapper.Ctx, r.collection, u, p, r.apiURL, model)
}
func (r *NetworkAccessMongoDB) UpdateRequestStatus(u, p, id, status string) (bool, error) {
	return updateOneRequestStatus(contextWrapper.Ctx, r.collection, id, status)
}

func createRequest(ctx context.Context, collection *mongo.Collection, u, p string, baseURL string, model *presenter.NetworkAccessList) (*[]presenter.NetworkAccessList, *presenter.ErrorDetail, error) {
	groupId := model.Groupid

	digestor := digest.NewDigestor(u, p)
	path := "/groups" + "/" + groupId + "/" + "accessList"

	reqBody, err1 := model.CreateJSON()
	result, err2 := digestor.Digest(baseURL, path, "POST", reqBody)
	var nwaccessList []presenter.NetworkAccessList
	var errDetails presenter.ErrorDetail
	var val map[string]interface{}

	_ = json.Unmarshal(result, &val)

	jsonEncoded, _ := json.Marshal(val["results"])
	_ = json.Unmarshal(jsonEncoded, &nwaccessList)
	_ = json.Unmarshal(jsonEncoded, &errDetails)

	if err2 == nil {
		return &nwaccessList, &presenter.ErrorDetail{}, err1
	} else {
		return &nwaccessList, &errDetails, err1
	}
}

func getAllRequest(ctx context.Context, collection *mongo.Collection) (*[]entity.NetworkAccessList, error) {

	cursor, _ := collection.Find(ctx, bson.M{"requestType": config.STR_REQ_TYPE_NETWORK_ACCESS, "status": config.STR_REQ_STATUS_OPEN})

	var aList []entity.NetworkAccessList
	var err error
	for cursor.Next(ctx) {
		var elem entity.UserRequest
		err = cursor.Decode(&elem)
		aList = append(aList, elem.GetNetworkEntity())
	}
	return &aList, err
}

func getRequest(ctx context.Context, collection *mongo.Collection, id string) (*entity.NetworkAccessList, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("Invalid Request ID")
	}
	singleResult := collection.FindOne(ctx, bson.M{"_id": objectId})
	var nwAccesslist entity.UserRequest
	singleResult.Decode(&nwAccesslist)
	var returnNwAccesslist entity.NetworkAccessList
	returnNwAccesslist = nwAccesslist.GetNetworkEntity()
	return &returnNwAccesslist, nil
}

func (r *NetworkAccessMongoDB) Create(e *entity.NetworkAccessList) error {
	return createOneAccessRequest(contextWrapper.Ctx, r.collection, e)
}

func createOneAccessRequest(ctx context.Context, collection *mongo.Collection, ent *entity.NetworkAccessList) error {
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

// Update an user.
func (r *NetworkAccessMongoDB) Update(e *entity.NetworkAccessList) error {
	return nil
}

// Delete an user.
func (r *NetworkAccessMongoDB) Delete(e *entity.NetworkAccessList) error {
	return nil
}
