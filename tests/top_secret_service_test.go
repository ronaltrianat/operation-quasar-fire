package tests

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ronaltrianat/operation-quasar-fire/src/config/apperrors"
	"github.com/ronaltrianat/operation-quasar-fire/src/core/domain"
	"github.com/ronaltrianat/operation-quasar-fire/src/core/services"
	mock_ports "github.com/ronaltrianat/operation-quasar-fire/tests/mocks/mockups"
	"github.com/stretchr/testify/assert"
)

type mocks struct {
	trilaterationUseCase *mock_ports.MockTrilaterationPort
	messageUseCase       *mock_ports.MockMessagePort
	repository           *mock_ports.MockRepositoryPort
}

func TestTopSecretHappyPath(t *testing.T) {

	m := mocks{
		trilaterationUseCase: mock_ports.NewMockTrilaterationPort(gomock.NewController(t)),
		messageUseCase:       mock_ports.NewMockMessagePort(gomock.NewController(t)),
		repository:           mock_ports.NewMockRepositoryPort(gomock.NewController(t)),
	}

	topSecretService := services.NewTopSecretService(m.trilaterationUseCase, m.messageUseCase, m.repository)

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

	position := domain.Position{X: -200, Y: 300}
	m.messageUseCase.EXPECT().CalculateMessage(&data).Return("este es un mensaje secreto", nil)
	m.trilaterationUseCase.EXPECT().TriangulatePosition(&data).Return(position, nil)

	response, err := topSecretService.TopSecret(&data)

	assert.Nil(t, err)
	assert.Equal(t, response, &domain.TopSecret{
		Position: position,
		Message:  "este es un mensaje secreto",
	})
}

func TestTopSecretErrInvalidNumberSatellites(t *testing.T) {

	m := mocks{
		trilaterationUseCase: mock_ports.NewMockTrilaterationPort(gomock.NewController(t)),
		messageUseCase:       mock_ports.NewMockMessagePort(gomock.NewController(t)),
		repository:           mock_ports.NewMockRepositoryPort(gomock.NewController(t)),
	}

	topSecretService := services.NewTopSecretService(m.trilaterationUseCase, m.messageUseCase, m.repository)

	data := domain.SatellitesData{}
	_, err := topSecretService.TopSecret(&data)

	assert.NotNil(t, err)
	assert.Equal(t, err, apperrors.ErrInvalidNumberSatellites)
}
