package tests

import (
	"testing"

	"github.com/ronaltrianat/operation-quasar-fire/src/config/apperrors"
	"github.com/ronaltrianat/operation-quasar-fire/src/core/domain"
	"github.com/ronaltrianat/operation-quasar-fire/src/core/usecases"
	"github.com/stretchr/testify/assert"
)

func TestCalculateMessageHappyPath(t *testing.T) {
	messageUseCase := usecases.NewMessageUseCase()
	data := domain.SatellitesData{
		Satellites: []domain.Satellite{
			{
				Message: []string{"este", "", "", "mensaje", ""},
			},
			{
				Message: []string{"", "es", "", "", "secreto"},
			},
			{
				Message: []string{"este", "", "un", "", ""},
			},
		},
	}

	message, err := messageUseCase.CalculateMessage(&data)

	assert.Equal(t, err, nil)
	assert.Equal(t, message, "este es un mensaje secreto")
}

func TestCalculateMessageErrInvalidNumberSatellites(t *testing.T) {
	messageUseCase := usecases.NewMessageUseCase()
	data := domain.SatellitesData{}
	message, err := messageUseCase.CalculateMessage(&data)

	assert.Equal(t, err, apperrors.ErrInvalidNumberSatellites)
	assert.Equal(t, message, "")
}
