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
	GroupId string
}

// NewService create new use case.
func NewService(publicKey, privateKey, groupId string) *Service {
	atlasParams := entity.AtlasParams{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}
	return &Service{
		AtlasParams: atlasParams,
		GroupId:     groupId,
	}
}

// Get Groups.
func (s *Service) Get() ([]presenter.Process, error) {
	return s.digestGroups(s.PublicKey, s.PrivateKey, s.GroupId)
}

// Get Groups.
func (s *Service) digestGroups(u, p, g string) ([]presenter.Process, error) {
	digestor := digest.NewDigestor(u, p)
	result, err := digestor.Digest(config.ENVIRONMENT.BASE_PATH, config.ENVIRONMENT.GROUPS_PATH+"/"+g+
		"/processes?pretty=true", "GET", nil)
	fmt.Println(err)

	var mongoProcess []presenter.Process
	var val map[string]interface{}

	_ = json.Unmarshal(result, &val)

	jsonEncoded, _ := json.Marshal(val["results"])

	_ = json.Unmarshal(jsonEncoded, &mongoProcess)

	return mongoProcess, nil
}
