package usecase

import (
	"co.edu.meli/luisillera/prueba-tecnica/domain/model"
	"co.edu.meli/luisillera/prueba-tecnica/domain/utilities"
	"co.edu.meli/luisillera/prueba-tecnica/infrastructure"
	"errors"
	"fmt"
)

type CalculatePositionUsecase struct {
	SatelliteProvider infrastructure.SatelliteProvider
}

func (a *CalculatePositionUsecase) Process(SpaceshipReference model.SpaceshipReference) (model.Point, error) {
	positionAlpha, alphaErr := a.findAlpha(SpaceshipReference)
	if alphaErr != nil {
		return model.Point{}, alphaErr
	}
	positionBravo, bravoErr := a.findBravo(SpaceshipReference)
	if bravoErr != nil {
		return model.Point{}, bravoErr
	}
	fmt.Println()
	if a.assertPositions(positionAlpha, positionBravo) {
		return positionAlpha, nil
	} else {
		return model.Point{}, errors.New("Final Positions Doesn't Match.")
	}
}

func (a *CalculatePositionUsecase) assertPositions(alpha model.Point, bravo model.Point) bool {
	north := alpha.North == bravo.North
	east := alpha.East == bravo.East
	return north && east
}

func (a *CalculatePositionUsecase) findAlpha(reference model.SpaceshipReference) (model.Point, error) {
	kenobi := a.SatelliteProvider.Kenobi
	beta := utilities.BetaAngle(kenobi.Skywalker, reference.ToSkywalker, reference.ToKenobi)
	azimuth := utilities.Azimuth(kenobi.Coordinates, a.SatelliteProvider.Skywalker.Coordinates) - beta
	north := utilities.North(azimuth, reference.ToKenobi, kenobi.Coordinates)
	east := utilities.East(azimuth, reference.ToKenobi, kenobi.Coordinates)
	return model.Point{
		East:  east,
		North: north,
	}, nil
}

func (a *CalculatePositionUsecase) findBravo(reference model.SpaceshipReference) (model.Point, error) {
	sato := a.SatelliteProvider.Sato
	beta := utilities.BetaAngle(sato.Skywalker, reference.ToSkywalker, reference.ToSato)
	azimut := utilities.Azimuth(sato.Coordinates, a.SatelliteProvider.Skywalker.Coordinates) + beta
	north := utilities.North(azimut, reference.ToSato, sato.Coordinates)
	east := utilities.East(azimut, reference.ToSato, sato.Coordinates)
	return model.Point{
		East:  east,
		North: north,
	}, nil
}
