package usecase

import (
	"co.edu.meli/luisillera/prueba-tecnica/domain/dto"
	"co.edu.meli/luisillera/prueba-tecnica/domain/model"
	"errors"
	"strings"
)

type RequestExtractorUsecase struct {
}

func (u *RequestExtractorUsecase) Process(request dto.SatelliteRequest) (model.SpaceshipReference, [][]string) {
	messages := make([][]string, 3)
	spaceship := model.SpaceshipReference{}
	messages[0], spaceship.ToKenobi, _ = u.findSatellite(request.Satellites, "kenobi")
	messages[1], spaceship.ToSkywalker, _ = u.findSatellite(request.Satellites, "skywalker")
	messages[2], spaceship.ToSato, _ = u.findSatellite(request.Satellites, "Sato")
	return spaceship, messages
}

func (u *RequestExtractorUsecase) findSatellite(satellites []dto.Satellite, name string) ([]string, float64, error) {
	for i := 0; i < len(satellites); i++ {
		satellite := satellites[i]
		if strings.ToLower(satellite.Name) == strings.ToLower(name) {
			return satellite.Message, satellite.Distance, nil
		}
	}
	return nil, 0, errors.New("Satellite [" + name + "] not found in the request")
}
