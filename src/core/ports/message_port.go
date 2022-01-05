package ports

import "github.com/ronaltrianat/operation-quasar-fire/src/core/domain"

type MessagePort interface {
	CalculateMessage(data *domain.SatellitesData) (string, error)
}
