package tests

import (
	"testing"

	"github.com/ronaltrianat/operation-quasar-fire/src/config/apperrors"
	"github.com/ronaltrianat/operation-quasar-fire/src/core/domain"
	"github.com/ronaltrianat/operation-quasar-fire/src/core/usecases"
	"github.com/stretchr/testify/assert"
)

func TestTriangulatePositionHappyPath(t *testing.T) {
	trilaterationUseCase := usecases.NewTrilaterationUseCase()
	data := domain.SatellitesData{
		Satellites: []domain.Satellite{
			{
				Distance: 500.0,
				Position: domain.Position{X: 100, Y: -100},
			},
			{
				Distance: 728.0110,
				Position: domain.Position{X: 500, Y: 100},
			},
			{
				Distance: 583.0952,
				Position: domain.Position{X: -500, Y: -200},
			},
		},
	}

	position, err := trilaterationUseCase.TriangulatePosition(&data)

	assert.Equal(t, err, nil)
	assert.Equal(t, position, domain.Position{X: -200, Y: 300})
}

func TestTriangulatePositionErrInvalidNumberSatellites(t *testing.T) {
	trilaterationUseCase := usecases.NewTrilaterationUseCase()
	data := domain.SatellitesData{}

	position, err := trilaterationUseCase.TriangulatePosition(&data)

	assert.Equal(t, err, apperrors.ErrInvalidNumberSatellites)
	assert.Equal(t, position, domain.Position{})

}
