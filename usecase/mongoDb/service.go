package mongoDb

import (
	"encoding/json"
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
func (s *Service) Get() ([]entity.MongoGroups, error) {
	return s.digestGroups(s.PublicKey, s.PrivateKey)
}

// Get Groups.
func (s *Service) digestGroups(u, p string) ([]entity.MongoGroups, error) {
	digestor := digest.NewDigestor(u,p)
	result := digestor.Digest(config.BASE_PATH, config.GROUPS_PATH, "GET", nil)

	var val map[string]interface{}

	_ = json.Unmarshal(result, &val)
	var mongoGroups []entity.MongoGroups

	jsonEncoded, _ := json.Marshal(val["results"])

	_ = json.Unmarshal(jsonEncoded, &mongoGroups)

	return mongoGroups, nil
}
