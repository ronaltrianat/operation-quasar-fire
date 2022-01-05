package repositories

import (
	"encoding/json"

	"github.com/ronaltrianat/operation-quasar-fire/src/config/apperrors"
	"github.com/ronaltrianat/operation-quasar-fire/src/core/domain"
)

type inmemoryDB struct {
	db map[string][]byte
}

func NewInMemoryDB() *inmemoryDB {
	return &inmemoryDB{
		db: map[string][]byte{},
	}
}

func (repo *inmemoryDB) Get(id string) (*domain.Satellite, error) {
	if value, ok := repo.db[id]; ok {
		satellite := &domain.Satellite{}
		err := json.Unmarshal(value, &satellite)
		if err != nil {
			return &domain.Satellite{}, err
		}

		return satellite, nil
	}

	return &domain.Satellite{}, apperrors.ErrRecordNotFoundInDB
}

func (repo *inmemoryDB) Save(data *domain.Satellite) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	repo.db[data.Name] = bytes
	return nil
}
