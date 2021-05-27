package groups

import (
	"encoding/json"
	"fmt"
	"mongo-admin-backend/api/presenter"
	"mongo-admin-backend/config"
	"mongo-admin-backend/entity"
	"mongo-admin-backend/pkg/digest"
)

// Service  interface.
type Service struct {
	entity.AtlasParams
	GroupId   string
	ProcessId string
	Params    map[string]string
}

// NewService create new use case.
func NewService(publicKey, privateKey, groupId, processId string, qParams map[string]string) *Service {
	atlasParams := entity.AtlasParams{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}
	return &Service{
		AtlasParams: atlasParams,
		GroupId:     groupId,
		ProcessId:   processId,
		Params:      qParams,
	}
}

// Get Groups.
func (s *Service) Get() (presenter.Measurements, presenter.ErrorDetail, error) {
	return s.digestProcessMeasurements(s.PublicKey, s.PrivateKey, s.GroupId, s.ProcessId, s.Params)
}

// Get Groups.
func (s *Service) digestProcessMeasurements(u, p, g, proc string, params map[string]string) (presenter.Measurements, presenter.ErrorDetail, error) {
	digestor := digest.NewDigestor(u, p)
	path := config.ENVIRONMENT.GROUPS_PATH + "/" + g + "/" +
		config.PATH_PROCESS + "/" + proc + "/" +
		config.PATH_MEASUREMENT + "?pretty=true"

	granStr, isGranPresent := params[config.PARAM_GRANULARITY]
	prdStr, isPeriodPresent := params[config.PARAM_PERIOD]
	stDate, isStpresent := params[config.PARAM_ST_DATE]
	enDate, isEndPresent := params[config.PARAM_END_DATE]
	measurementName, isMPresent := params[config.PARAM_MEASUREMENT]

	if isGranPresent {
		path = path + "&" + config.PARAM_GRANULARITY + "=" + granStr
	}
	if isPeriodPresent {
		path = path + "&" + config.PARAM_PERIOD + "=" + prdStr
	}
	if isStpresent {
		path = path + "&" + config.PARAM_ST_DATE + "=" + stDate
	}

	if isEndPresent {
		path = path + "&" + config.PARAM_END_DATE + "=" + enDate
	}
	if isMPresent {
		path = path + "&" + config.PARAM_MEASUREMENT + "=" + measurementName
	}

	result, err := digestor.Digest(config.ENVIRONMENT.BASE_PATH, path, "GET", nil)
	fmt.Println(err)
	var mongoMeasurements presenter.Measurements
	var errDetails presenter.ErrorDetail
	var val map[string]interface{}

	_ = json.Unmarshal(result, &val)
	fmt.Println(val)
	jsonEncoded, _ := json.Marshal(val)

	_ = json.Unmarshal(jsonEncoded, &mongoMeasurements)
	_ = json.Unmarshal(jsonEncoded, &errDetails)
	if err == nil {
		return mongoMeasurements, presenter.ErrorDetail{}, err
	} else {
		return mongoMeasurements, errDetails, err
	}

}
