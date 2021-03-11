package entrypoint

import (
	"co.edu.meli/luisillera/prueba-tecnica/domain/dto"
	"co.edu.meli/luisillera/prueba-tecnica/domain/usecase"
)

type TopSecretResource struct {
	ExtractInformation usecase.ExtractInformationUsecase
	RequestExtractor   usecase.RequestExtractorUsecase
}

func (r *TopSecretResource) HandleRequest(request dto.SatelliteRequest) (dto.Response, error) {
	spaceship, messages := r.RequestExtractor.Process(request)
	if result, err := r.ExtractInformation.Process(spaceship, messages); err == nil {
		return dto.Response{
			Position: dto.Position{
				X: result.Position.East,
				Y: result.Position.North,
			},
			Message: result.Message,
		}, nil
	} else {
		return dto.Response{}, err
	}
}
