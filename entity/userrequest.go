package entity

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRequest struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Groupid        string             `json:"groupId" bson:"groupId"`
	Requester      string             `json:"requester" bson:"requester"`
	Requesttype    string             `json:"requestType" bson:"requestType"`
	Status         string             `json:"status" bson:"status"`
	Requesteddate  string             `json:"requestedDate" bson:"requestedDate"`
	Requestdetails Requestdetails     `json:"requestDetails" bson:"requestDetails"`
	V              int                `json:"__v" bson:"__v"`
}

type Requestdetails struct {
	Ipcidr   string  `json:"ipcidr" bson:"ipcidr"`
	Comment  string  `json:"comment" bson:"comment"`
	Groupid  string  `json:"groupId" bson:"groupId"`
	Roles    []Roles `json:"roles" bson:"roles"`
	UserName string  `json:"userName" bson:"userName"`
	Password string  `json:"password" bson:"password"`
}

func (model *UserRequest) GetNetworkEntity() NetworkAccessList {
	var returnVal NetworkAccessList
	var err error
	layout := "2006-01-02 15:04:05"
	returnVal.ID = model.ID
	returnVal.Cidrblock = model.Requestdetails.Ipcidr
	returnVal.Comment = model.Requestdetails.Comment
	returnVal.Groupid = model.Requestdetails.Groupid
	returnVal.CreatedBy = model.Requester
	returnVal.CreatedAt, err = time.Parse(layout, model.Requesteddate)
	if err != nil {
		fmt.Println("Error is formatting the date" + err.Error())
	}
	returnVal.Status = model.Status
	return returnVal
}

func (model *UserRequest) GetDBUserRequest() DBAccessRequest {
	var returnObj DBAccessRequest
	var err error
	layout := "2006-01-02 15:04:05"
	returnObj.ID = model.ID
	returnObj.Username = model.Requestdetails.UserName
	returnObj.Password = model.Requestdetails.Password
	returnObj.Groupid = model.Requestdetails.Groupid
	returnObj.Roles = model.Requestdetails.Roles
	returnObj.Status = model.Status
	returnObj.CreatedAt, err = time.Parse(layout, model.Requesteddate)

	if err != nil {
		fmt.Println("Error in formatting the date" + err.Error())
	}
	returnObj.CreatedBy = model.Requester
	return returnObj
}
