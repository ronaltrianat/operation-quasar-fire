package ports

import "github.com/ronaltrianat/operation-quasar-fire/src/core/domain"

type TopSecretPort interface {
	TopSecret(data *domain.SatellitesData) (*domain.TopSecret, error)
	SaveTopSecretSplit(data *domain.Satellite) error
	RetrieveTopSecretSplit(satellites []string) (*domain.TopSecret, error)
}
