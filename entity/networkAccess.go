package entity

import (
	"encoding/json"
	"mongo-admin-backend/api/presenter"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NetworkAccessList struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Cidrblock       string             `json:"cidrBlock" bson:"cidrBlock"`
	Comment         string             `json:"comment" bson:"comment"`
	Groupid         string             `json:"groupId" bson:"groupId"`
	Ipaddress       string             `json:"ipAddress,omitempty" bson:"ipAddress,omitempty"`
	Deleteafterdate time.Time          `json:"deleteAfterDate,omitempty" bson:"deleteAfterDate,omitempty"`
	CreatedAt       time.Time          `bson:"created_at,omitempty"`
	UpdatedAt       time.Time          `bson:"updated_at,omitempty"`
	CreatedBy       string             `bson:"createdBy,omitempty"`
	Status          string             `bson:"status"`
}

func (r *NetworkAccessList) convertFromAPIRequest(entity *presenter.NetworkAccessList) {
	r.Cidrblock = entity.Cidrblock
	r.Comment = entity.Comment
	r.Groupid = entity.Groupid
	r.Ipaddress = entity.Ipaddress
	r.Deleteafterdate = entity.Deleteafterdate
	r.Status = "Done"
}

func (r *NetworkAccessList) ConvertToAPIRequest() *presenter.NetworkAccessList {
	var model presenter.NetworkAccessList
	model.Cidrblock = r.Cidrblock
	model.Comment = r.Comment
	model.Ipaddress = r.Ipaddress
	model.Deleteafterdate = r.Deleteafterdate
	model.Groupid = r.Groupid
	return &model
}

//Database user access request
type DBAccessRequest struct {
	Id           string    `json:"id" bson:"_id"`
	Groupid      string    `json:"groupId" bson:"groupId"`
	Databasename string    `json:"databaseName" bson:"databaseName"`
	Password     string    `json:"password" bson:"password"`
	Roles        []Roles   `json:"roles" bson:"roles"`
	Scopes       []Scopes  `json:"scopes" bson:"scopes"`
	Username     string    `json:"username" bson:"username"`
	CreatedAt    time.Time `bson:"created_at,omitempty"`
	UpdatedAt    time.Time `bson:"updated_at,omitempty"`
	CreatedBy    string    `bson:"createdBy,omitempty"`
	Status       string    `bson:"status,omitempty"`
}
type Roles struct {
	Databasename string `bson:"databaseName" json:"databaseName"`
	Rolename     string `bson:"roleName" json:"roleName"`
}
type Scopes struct {
	Name string `bson:"name" json:"name"`
	Type string `bson:"type" json:"type"`
}

func (model *DBAccessRequest) CreateJSON() ([]byte, error) {
	var request = make(map[string]interface{})

	if model.Databasename != "" {
		request["databaseName"] = model.Databasename
	}
	if model.Groupid != "" {
		request["groupId"] = model.Groupid
	}

	if model.Password != "" {
		request["password"] = model.Password
	}
	if model.Username != "" {
		request["username"] = model.Username
	}
	if model.Roles != nil {
		request["roles"] = model.Roles
	}

	j, err := json.Marshal(request)
	return j, err

}
