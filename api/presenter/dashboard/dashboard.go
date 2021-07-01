package dashboard

import (
	"fmt"
	"time"
)

type ServiceCallResult struct {
	Results []ServiceCallResponse
}

type ServiceCallResponse struct {
	ServiceCall  Service
	Name         string
	Responsebody interface{}
	Error        error
}
type Process_Response func([]byte) (interface{}, error)

type Service struct {
	Name             string
	Host             string
	Url              string
	Method           string
	RequestBody      interface{}
	Process_Response Process_Response
}

type ErrorDetail struct {
	Detail     string   `json:"detail"`
	Error      int      `json:"error"`
	ErrorCode  string   `json:"errorCode"`
	Parameters []string `json:"parameters"`
	Reason     string   `json:"reason"`
}

type ClusterDetails struct {
	Backupenabled         bool      `json:"backupEnabled"`
	Clustertype           string    `json:"clusterType"`
	Createdate            time.Time `json:"createDate"`
	Disksizegb            float64   `json:"diskSizeGB"`
	Groupid               string    `json:"groupId"`
	ID                    string    `json:"id"`
	Mongodbmajorversion   string    `json:"mongoDBMajorVersion"`
	Mongodbversion        string    `json:"mongoDBVersion"`
	Name                  string    `json:"name"`
	Numshards             int       `json:"numShards"`
	Paused                bool      `json:"paused"`
	Pitenabled            bool      `json:"pitEnabled"`
	Providerbackupenabled bool      `json:"providerBackupEnabled"`
	Providersettings      struct {
		Providername string `json:"providerName"`
		Autoscaling  struct {
			Compute struct {
				Maxinstancesize interface{} `json:"maxInstanceSize"`
				Mininstancesize interface{} `json:"minInstanceSize"`
			} `json:"compute"`
		} `json:"autoScaling"`
		Backingprovidername string `json:"backingProviderName"`
		Regionname          string `json:"regionName"`
		Instancesizename    string `json:"instanceSizeName"`
	} `json:"providerSettings"`
	Replicationfactor    int    `json:"replicationFactor"`
	Versionreleasesystem string `json:"versionReleaseSystem"`
}

type ProcessDetails struct {
	Created        time.Time `json:"created"`
	Groupid        string    `json:"groupId"`
	Hostname       string    `json:"hostname"`
	ID             string    `json:"id"`
	Lastping       time.Time `json:"lastPing"`
	Port           int       `json:"port"`
	Replicasetname string    `json:"replicaSetName"`
	Typename       string    `json:"typeName"`
	Useralias      string    `json:"userAlias"`
	Version        string    `json:"version"`
}

//Implement the parsing functions as Callback functions

func process_first(response []byte) (interface{}, error) {
	fmt.Println("This  function is a callback function")
	return nil, nil
}
