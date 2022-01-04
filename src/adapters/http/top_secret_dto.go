package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/ronaltrianat/operation-quasar-fire/src/core/domain"
)

var validate = validator.New()

type Position struct {
	X float64 `json:"x" validate:"required"`
	Y float64 `json:"y" validate:"required"`
}

type SatelliteRequest struct {
	Name     string   `json:"name" validate:"required"`
	Distance float64  `json:"distance" validate:"required"`
	Message  []string `json:"message" validate:"required"`
	Position Position `json:"position" validate:"required"`
}

func (r SatelliteRequest) ToDomainData(name string) *domain.Satellite {
	satellite := &domain.Satellite{
		Name:     name,
		Distance: r.Distance,
		Message:  r.Message,
		Position: domain.Position(r.Position),
	}
	return satellite
}

func (r SatelliteRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

type TopSecretRequest struct {
	Satellites []SatelliteRequest `json:"satellites" validate:"required,len=3,dive,required"`
}

func (r TopSecretRequest) ToDomainData() *domain.SatellitesData {
	satellitesData := &domain.SatellitesData{}
	for _, v := range r.Satellites {
		satellitesData.Satellites = append(satellitesData.Satellites, domain.Satellite{
			Name:     v.Name,
			Distance: v.Distance,
			Message:  v.Message,
			Position: domain.Position(v.Position),
		})
	}
	return satellitesData
}

func (r TopSecretRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

type GeneralResponse struct {
	Success                  bool   `json:"success"`
	MessageResponse          string `json:"messageResponse"`
	TechnicalMessageResponse string `json:"technicalMessageResponse"`
}

type TopSecretResponse struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
	GeneralResponse
}
