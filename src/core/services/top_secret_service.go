package services

import (
	"github.com/ronaltrianat/operation-quasar-fire/src/config/apperrors"
	"github.com/ronaltrianat/operation-quasar-fire/src/core/domain"
	"github.com/ronaltrianat/operation-quasar-fire/src/core/ports"
)

type topSecretService struct {
	trilaterationUseCase ports.TrilaterationPort
	messageUseCase       ports.MessagePort
	repository           ports.RepositoryPort
}

func NewTopSecretService(trilateration ports.TrilaterationPort, message ports.MessagePort,
	repository ports.RepositoryPort) *topSecretService {
	return &topSecretService{
		trilaterationUseCase: trilateration,
		messageUseCase:       message,
		repository:           repository,
	}
}

func (service *topSecretService) TopSecret(data *domain.SatellitesData) (*domain.TopSecret, error) {
	if data == nil || len(data.Satellites) < 3 {
		return nil, apperrors.ErrInvalidNumberSatellites
	}
	return service.generateTopSecret(data)
}

func (service *topSecretService) SaveTopSecretSplit(data *domain.Satellite) error {
	if err := service.repository.Save(data); err != nil {
		return err
	}
	return nil
}

func (service *topSecretService) RetrieveTopSecretSplit(satellites []string) (*domain.TopSecret, error) {
	satellitesList := []domain.Satellite{}
	for _, v := range satellites {
		data, err := service.repository.Get(v)
		if err != nil {
			return nil, err
		}
		satellitesList = append(satellitesList, *data)
	}

	satellitesData := domain.SatellitesData{
		Satellites: satellitesList,
	}

	return service.generateTopSecret(&satellitesData)
}

func (service *topSecretService) generateTopSecret(data *domain.SatellitesData) (*domain.TopSecret, error) {
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
