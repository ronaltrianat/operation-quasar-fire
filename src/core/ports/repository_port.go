package ports

import "github.com/ronaltrianat/operation-quasar-fire/src/core/domain"

type RepositoryPort interface {
	Get(id string) (*domain.Satellite, error)
	Save(data *domain.Satellite) error
}
