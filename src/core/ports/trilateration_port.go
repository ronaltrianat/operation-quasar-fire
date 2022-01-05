package ports

import "github.com/ronaltrianat/operation-quasar-fire/src/core/domain"

type TrilaterationPort interface {
	TriangulatePosition(data *domain.SatellitesData) (domain.Position, error)
}
