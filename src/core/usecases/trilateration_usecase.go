package usecases

import (
	"math"

	"github.com/ronaltrianat/operation-quasar-fire/src/config/apperrors"
	"github.com/ronaltrianat/operation-quasar-fire/src/core/domain"
)

type TrilaterationUseCase struct{}

func NewTrilaterationUseCase() *TrilaterationUseCase {
	return &TrilaterationUseCase{}
}

const (
	// SatelliteOne docs
	SatelliteOne = iota
	// SatelliteTwo docs
	SatelliteTwo
	// SatelliteThree docs
	SatelliteThree
)

// GetLocation docs.
//
// metodo de eliminacion de 3 ecuaciones y 2 incognitas
func (useCase *TrilaterationUseCase) TriangulatePosition(data *domain.SatellitesData) (domain.Position, error) {
	if data == nil || len(data.Satellites) < 3 {
		return domain.Position{}, apperrors.ErrInvalidNumberSatellites
	}

	positionY := useCase.getPositionY(data)

	var positionX float64

	if !math.IsNaN(positionY) {
		positionX = useCase.getPositionX(data, positionY)
	}

	response := domain.Position{X: positionX, Y: positionY}

	return response, nil
}

// GetPositionY docs
//
// Formula para calcular la coordenada 'Y' del portacarga imperial,
// teniendo las coordenadas y distancias de los satelites.
//
// =========================================================================================================
// =|      2(x1 - x3)(d2² - d1² + x1² - x2² + y1² - y2²) - 2(x1 - x2)(d3² - d1² + x1² - x3² + y1² - y3²)  |=
// =| Y = ----------------------------------------------------------------------------------------------- |=
// =|                           4((x1 - x3)(y1 - y2) - (x1 - x2)(y1 - y3))                                |=
// =========================================================================================================
//
func (useCase TrilaterationUseCase) getPositionY(data *domain.SatellitesData) float64 {
	// dividendo de la ecuacion: 2(x1 - x3)(d2² - d1² + x1² - x2² + y1² - y2²) - 2(x1 - x2)(d3² - d1² + x1² - x3² + y1² - y3²)

	// 2(x1 - x3)(d2² - d1² + x1² - x2² + y1² - y2²)
	minuend := 2 * (data.Satellites[SatelliteOne].Position.X - data.Satellites[SatelliteThree].Position.X) *
		(math.Pow(data.Satellites[SatelliteTwo].Distance, 2) -
			math.Pow(data.Satellites[SatelliteOne].Distance, 2) +
			math.Pow(data.Satellites[SatelliteOne].Position.X, 2) -
			math.Pow(data.Satellites[SatelliteTwo].Position.X, 2) +
			math.Pow(data.Satellites[SatelliteOne].Position.Y, 2) -
			math.Pow(data.Satellites[SatelliteTwo].Position.Y, 2))

	// 2(x1 - x2)(d3² - d1² + x1² - x3² + y1² - y3²)
	subtracting := 2 * (data.Satellites[SatelliteOne].Position.X - data.Satellites[SatelliteTwo].Position.X) *
		(math.Pow(data.Satellites[SatelliteThree].Distance, 2) -
			math.Pow(data.Satellites[SatelliteOne].Distance, 2) +
			math.Pow(data.Satellites[SatelliteOne].Position.X, 2) -
			math.Pow(data.Satellites[SatelliteThree].Position.X, 2) +
			math.Pow(data.Satellites[SatelliteOne].Position.Y, 2) -
			math.Pow(data.Satellites[SatelliteThree].Position.Y, 2))

	// divisor de la ecuacion
	// 4((x1 - x3)(y1 - y2) - (x1 - x2)(y1 - y3))
	divider := 4 * ((data.Satellites[SatelliteOne].Position.X-data.Satellites[SatelliteThree].Position.X)*
		(data.Satellites[SatelliteOne].Position.Y-data.Satellites[SatelliteTwo].Position.Y) -
		(data.Satellites[SatelliteOne].Position.X-data.Satellites[SatelliteTwo].Position.X)*
			(data.Satellites[SatelliteOne].Position.Y-data.Satellites[SatelliteThree].Position.Y))

	positionY := (minuend - subtracting) / divider
	roundY := math.Round(positionY*100) / 100
	return roundY
}

// GetPositionX docs
//
// Formula para calcular la coordenada 'X' del portacarga imperial,
// teniendo las coordenadas, distancias de los satelites y el valor de la coordenada 'Y'.
//
// ==============================================================
// =|       d2² - d1² + x1² - x2² + y1² - y2² - 2Y(y1 - y2)    |=
// =|  X = -------------------------------------------------   |=
// =|                        2(x1 - x2)                        |=
// =|                                                          |=
// =|       d3² - d1² + x1² - x3² + y1² - y3² - 2Y(y1 - y3)    |=
// =|  X = -------------------------------------------------   |=
// =|                        2(x1 - x3)                        |=
// ==============================================================
//
func (useCase TrilaterationUseCase) getPositionX(data *domain.SatellitesData, positionY float64) float64 {
	dividend := math.Pow(data.Satellites[SatelliteTwo].Distance, 2) -
		math.Pow(data.Satellites[SatelliteOne].Distance, 2) +
		math.Pow(data.Satellites[SatelliteOne].Position.X, 2) -
		math.Pow(data.Satellites[SatelliteTwo].Position.X, 2) +
		math.Pow(data.Satellites[SatelliteOne].Position.Y, 2) -
		math.Pow(data.Satellites[SatelliteTwo].Position.Y, 2) -
		2*positionY*(data.Satellites[SatelliteOne].Position.Y-data.Satellites[SatelliteTwo].Position.Y)

	divider := 2 * (data.Satellites[SatelliteOne].Position.X - data.Satellites[SatelliteTwo].Position.X)
	roundX := math.Round((dividend/divider)*100) / 100
	return roundX
}
