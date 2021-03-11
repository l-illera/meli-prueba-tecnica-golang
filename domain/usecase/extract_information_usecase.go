package usecase

import (
	"co.edu.meli/luisillera/prueba-tecnica/domain/model"
	"errors"
	"fmt"
)

type ExtractInformationUsecase struct {
	CalculatePosition CalculatePositionUsecase
	MessageBuilder    MessageBuilderUsecase
}

func (u *ExtractInformationUsecase) Process(reference model.SpaceshipReference, messages [][]string) (model.Result, error) {
	result := model.Result{}
	result.Message = u.findMessage(messages)
	result.Position = u.findPosition(reference)
	if result.Position.North == 0 {
		return model.Result{}, errors.New("Position can't be found.")
	}
	return result, nil
}

func (u *ExtractInformationUsecase) findPosition(reference model.SpaceshipReference) model.Point {
	if point, err := u.CalculatePosition.Process(reference); err == nil {
		return point
	} else {
		fmt.Printf("ExtractInformationUsecase.Process.CalculatePosition Err: [%s]", err.Error())
		return model.Point{}
	}
}

func (u *ExtractInformationUsecase) findMessage(messages [][]string) string {
	if msg, err := u.MessageBuilder.Process(messages); err == nil {
		return msg
	} else {
		fmt.Printf("ExtractInformationUsecase.Process.MessageBuilder Err: [%s]", err.Error())
		return ""
	}
}
