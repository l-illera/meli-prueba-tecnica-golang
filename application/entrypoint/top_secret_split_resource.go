package entrypoint

import (
	"co.edu.meli/luisillera/prueba-tecnica/domain/dto"
	"co.edu.meli/luisillera/prueba-tecnica/domain/usecase"
	"co.edu.meli/luisillera/prueba-tecnica/infrastructure"
)

type TopSecretSplitResource struct {
	MessageProvider    infrastructure.MessageProvider
	ExtractInformation usecase.ExtractInformationUsecase
	RequestExtractor   usecase.RequestExtractorUsecase
}

func (t *TopSecretSplitResource) LoadMessage(satellite dto.Satellite) (dto.Response, error) {
	t.MessageProvider.AddMessage(satellite)
	return t.LocateMessage()
}

func (t *TopSecretSplitResource) LocateMessage() (dto.Response, error) {
	spaceship, messages := t.RequestExtractor.Process(t.MessageProvider.GetMessages())
	if result, err := t.ExtractInformation.Process(spaceship, messages); err == nil {
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
