package presenter

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
