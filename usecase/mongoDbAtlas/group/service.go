package group

import (
	"encoding/json"
	"mongo-admin-backend/api/presenter"
	"mongo-admin-backend/config"
	"mongo-admin-backend/entity"
	"mongo-admin-backend/pkg/digest"
)

// Service  interface.
type Service struct {
	entity.AtlasParams
}

// NewService create new use case.
func NewService(publicKey, privateKey string) *Service {
	return &Service{
		entity.AtlasParams{
			PublicKey: publicKey,
			PrivateKey: privateKey,
		},
	}
}

// Get Groups.
func (s *Service) Get() ([]presenter.MongoGroups, error) {
	return s.digestGroups(s.PublicKey, s.PrivateKey)
}

// Get Groups.
func (s *Service) digestGroups(u, p string) ([]presenter.MongoGroups, error) {
	digestor := digest.NewDigestor(u,p)
	result := digestor.Digest(config.ENVIRONMENT.BASE_PATH, config.ENVIRONMENT.GROUPS_PATH, "GET", nil)

	var val map[string]interface{}

	_ = json.Unmarshal(result, &val)
	var mongoGroups []presenter.MongoGroups

	jsonEncoded, _ := json.Marshal(val["results"])

	_ = json.Unmarshal(jsonEncoded, &mongoGroups)

	return mongoGroups, nil
}
