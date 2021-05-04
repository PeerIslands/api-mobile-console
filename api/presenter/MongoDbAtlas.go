package presenter

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

type Process struct {
	Created  time.Time `json:"created"`
	GroupId  string    `json:"groupId"`
	Hostname string    `json:"hostname"`
	Id       string    `json:"id"`
	LastPing time.Time `json:"lastPing"`
	Links    []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
	Port           int    `json:"port"`
	ReplicaSetName string `json:"replicaSetName"`
	TypeName       string `json:"typeName"`
	UserAlias      string `json:"userAlias"`
	Version        string `json:"version"`
}
