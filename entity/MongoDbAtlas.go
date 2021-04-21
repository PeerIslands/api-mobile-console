package entity

import (
	"time"
)

type MongoGroups struct {
	ClusterCount int       `json:"clusterCount"`
	Created      time.Time `json:"created"`
	Id           string    `json:"id"`
	Links        []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
	Name  string `json:"name"`
	OrgId string `json:"orgId"`
}

////NewUser create a new user
//func NewGroup(email, password, name, publicKey, privateKey string) (*User, error) {
//	var a = &AtlasParams{
//		PublicKey: publicKey,
//		PrivateKey: privateKey,
//	}
//	var u = &User{
//		Email: email,
//		Name:  name,
//		AtlasParams: *a,
//		Password: password,
//	}
//	u.Encrypt()
//	err := u.Validate()
//	if err != nil {
//		return nil, err
//	}
//	return u, nil
//}

