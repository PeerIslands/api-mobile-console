package presenter

import (
	"encoding/json"
	"time"
)

type NetworkAccessListResponse struct {
	Links      []Links             `json:"links"`
	Results    []NetworkAccessList `json:"results"`
	Totalcount int                 `json:"totalCount"`
}
type Links struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}
type NetworkAccessList struct {
	Cidrblock       string    `json:"cidrBlock,omitempty"`
	Comment         string    `json:"comment,omitempty"`
	Groupid         string    `json:"groupId,omitempty"`
	Ipaddress       string    `json:"ipAddress,omitempty"`
	Links           []Links   `json:"links,omitempty"`
	Deleteafterdate time.Time `json:"deleteAfterDate,omitempty"`
}

func (model *NetworkAccessList) CreateJSON() ([]byte, error) {
	type RequestNWAccess struct {
		cidrBlock string
		ipAddress string
		groupId   string
		comment   string
	}
	var request = make(map[string]string)
	if model.Cidrblock != "" {
		request["cidrBlock"] = model.Cidrblock
	}

	if model.Ipaddress != "" {
		request["ipAddress"] = model.Ipaddress
	}
	if model.Comment != "" {
		request["comment"] = model.Comment
	}
	if model.Groupid != "" {
		request["groupId"] = model.Groupid
	}
	requestArr := [](map[string]string){request}

	j, err := json.Marshal(requestArr)
	return j, err

}

//Database user access request
type DBAccessRequest struct {
	Awsiamtype   string        `json:"awsIAMType"`
	Databasename string        `json:"databaseName"`
	Groupid      string        `json:"groupId"`
	Labels       []interface{} `json:"labels"`
	Ldapauthtype string        `json:"ldapAuthType"`
	Links        []Links       `json:"links"`
	Roles        []Roles       `json:"roles"`
	Scopes       []interface{} `json:"scopes"`
	Username     string        `json:"username"`
	X509Type     string        `json:"x509Type"`
	Password     string        `json:"password"`
}

type Roles struct {
	Databasename string `json:"databaseName"`
	Rolename     string `json:"roleName"`
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
