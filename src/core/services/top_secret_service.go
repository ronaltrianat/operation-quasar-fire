package services

import (
	"github.com/ronaltrianat/operation-quasar-fire/src/core/domain"
	"github.com/ronaltrianat/operation-quasar-fire/src/core/ports"
)

type topSecretService struct {
	trilaterationUseCase ports.TrilaterationPort
	messageUseCase       ports.MessagePort
}

func NewTopSecretService(trilateration ports.TrilaterationPort, message ports.MessagePort) *topSecretService {
	return &topSecretService{
		trilaterationUseCase: trilateration,
		messageUseCase:       message,
	}
}

func (service *topSecretService) TopSecret(data *domain.SatellitesData) (*domain.TopSecret, error) {
	message, err := service.messageUseCase.CalculateMessage(data)
	if err != nil {
		return nil, err
	}

	position, err := service.trilaterationUseCase.TriangulatePosition(data)
	if err != nil {
		return nil, err
	}

	response := &domain.TopSecret{
		Position: position,
		Message:  message,
	}

	return response, nil
}
