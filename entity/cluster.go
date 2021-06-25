package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminRequest struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Groupid        string             `json:"groupId" bson:"groupId"`
	Requester      string             `json:"requester" bson:"requester"`
	Requesttype    string             `json:"requestType" bson:"requestType"`
	Status         string             `json:"status" bson:"status"`
	Requesteddate  string             `json:"requestedDate" bson:"requestedDate"`
	Requestdetails interface{}        `json:"requestDetails" bson:"requestDetails"`
	V              int                `json:"__v" bson:"__v"`
}

type ClusterDetails struct {
	Name                  string             `json:"name"`
	Disksizegb            int                `json:"diskSizeGB"`
	Numshards             int                `json:"numShards"`
	Providersettings      Providersettings   `json:"providerSettings"`
	Clustertype           string             `json:"clusterType"`
	Replicationfactor     int                `json:"replicationFactor"`
	Replicationspecs      []Replicationspecs `json:"replicationSpecs"`
	Backupenabled         bool               `json:"backupEnabled"`
	Providerbackupenabled bool               `json:"providerBackupEnabled"`
	Autoscaling           Autoscaling        `json:"autoScaling"`
}
type Providersettings struct {
	Providername     string `json:"providerName"`
	Instancesizename string `json:"instanceSizeName"`
	Regionname       string `json:"regionName"`
}
type RegionSpecs struct {
	Analyticsnodes int `json:"analyticsNodes"`
	Electablenodes int `json:"electableNodes"`
	Priority       int `json:"priority"`
	Readonlynodes  int `json:"readOnlyNodes"`
}

type Replicationspecs struct {
	Numshards     int                    `json:"numShards"`
	Regionsconfig map[string]RegionSpecs `json:"regionsConfig"`
	Zonename      string                 `json:"zoneName"`
}
type Autoscaling struct {
	Diskgbenabled bool `json:"diskGBEnabled"`
}
