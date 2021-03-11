package entrypoint

import (
	"co.edu.meli/luisillera/prueba-tecnica/domain/dto"
	"co.edu.meli/luisillera/prueba-tecnica/domain/usecase"
	"co.edu.meli/luisillera/prueba-tecnica/infrastructure"
)

type TopSecretSplitResource struct {
	messageProvider    infrastructure.MessageProvider
	extractInformation usecase.ExtractInformationUsecase
	requestExtractor   usecase.RequestExtractorUsecase
}

func (t *TopSecretSplitResource) Initialize(provider infrastructure.MessageProvider) {
	t.messageProvider = provider
}

func (t *TopSecretSplitResource) LoadMessage(satellite dto.Satellite) (dto.Response, error) {
	t.messageProvider.AddMessage(satellite)
	return t.LocateMessage()
}

func (t *TopSecretSplitResource) LocateMessage() (dto.Response, error) {
	spaceship, messages := t.requestExtractor.Process(t.messageProvider.GetMessages())
	if result, err := t.extractInformation.Process(spaceship, messages); err == nil {
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
