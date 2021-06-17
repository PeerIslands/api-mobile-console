package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mongo-admin-backend/api/presenter"
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
	collection := client.Database("mongoDbAdmin").Collection("userDBAccess")
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

func createRequest(ctx context.Context, collection *mongo.Collection, u, p string, baseURL string, model *presenter.NetworkAccessList) (*[]presenter.NetworkAccessList, *presenter.ErrorDetail, error) {
	groupId := model.Groupid
	// fmt.Println(groupId)
	// fmt.Println(model)
	// fmt.Println(u)
	// fmt.Println(p)

	digestor := digest.NewDigestor(u, p)
	path := "/groups" + "/" + groupId + "/" + "accessList"
	fmt.Println(path)
	fmt.Println(model)
	reqBody, err1 := model.CreateJSON()
	fmt.Println(reqBody)
	result, err2 := digestor.Digest(baseURL, path, "POST", reqBody)
	var nwaccessList []presenter.NetworkAccessList
	var errDetails presenter.ErrorDetail
	var val map[string]interface{}

	_ = json.Unmarshal(result, &val)

	jsonEncoded, _ := json.Marshal(val["results"])
	fmt.Println(val["results"])
	_ = json.Unmarshal(jsonEncoded, &nwaccessList)
	_ = json.Unmarshal(jsonEncoded, &errDetails)

	if err2 == nil {
		return &nwaccessList, &presenter.ErrorDetail{}, err1
	} else {
		fmt.Println("The output is ")
		fmt.Println(nwaccessList)
		return &nwaccessList, &errDetails, err1
	}
}

func getAllRequest(ctx context.Context, collection *mongo.Collection) (*[]entity.NetworkAccessList, error) {

	cursor, _ := collection.Find(ctx, bson.M{})

	var aList []entity.NetworkAccessList

	for cursor.Next(ctx) {
		fmt.Println("the cursot next")
		var elem entity.NetworkAccessList
		err := cursor.Decode(&elem)
		if err != nil {
			fmt.Println("Decoding failed")
		}
		aList = append(aList, elem)
	}

	return &aList, nil
}

func getRequest(ctx context.Context, collection *mongo.Collection, id string) (*entity.NetworkAccessList, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("Invalid Request ID")
	}
	singleResult := collection.FindOne(ctx, bson.M{"_id": objectId})
	var nwAccesslist entity.NetworkAccessList
	singleResult.Decode(&nwAccesslist)
	return &nwAccesslist, nil
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

// Send the request to MongoDB Api

// Update an user.
func (r *NetworkAccessMongoDB) Update(e *entity.NetworkAccessList) error {
	fmt.Println(e)

	return nil
}

// Delete an user.
func (r *NetworkAccessMongoDB) Delete(e *entity.NetworkAccessList) error {
	fmt.Println(e)

	return nil
}
