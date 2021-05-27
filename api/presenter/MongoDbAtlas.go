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

// type Measurements struct {
// 	End         string `json:"end"`
// 	Granularity string `json:"granularity"`
// 	GroupId     string `json:"groupId"`
// 	HostId      string `json:"hostId"`
// 	Links       []struct {
// 		Href string `json:"href"`
// 		Rel  string `json:"rel"`
// 	} `json:"links"`
// 	Measurements []struct {
// 		DataPoints []struct {
// 			Timestamp string   `json:"timestamp"`
// 			Value     *float64 `json:"value"`
// 		} `json:"dataPoints"`
// 		Name  string `json:"name"`
// 		Units string `json:"units"`
// 	} `json:"measurements"`
// 	ProcessId string `json:"processId"`
// 	Start     string `json:"start"`
// }

type Measurements struct {
	End         time.Time `json:"end"`
	Granularity string    `json:"granularity"`
	GroupId     string    `json:"groupId"`
	HostId      string    `json:"hostId"`
	Links       []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
	Measurements []struct {
		DataPoints []struct {
			Timestamp time.Time `json:"timestamp"`
			Value     *float64  `json:"value"`
		} `json:"dataPoints"`
		Name  string `json:"name"`
		Units string `json:"units"`
	} `json:"measurements"`
	ProcessId string    `json:"processId"`
	Start     time.Time `json:"start"`
}

type ErrorDetail struct {
	Detail     string   `json:"detail"`
	Error      int      `json:"error"`
	ErrorCode  string   `json:"errorCode"`
	Parameters []string `json:"parameters"`
	Reason     string   `json:"reason"`
}
