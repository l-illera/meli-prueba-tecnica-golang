package infrastructure

import (
	"co.edu.meli/luisillera/prueba-tecnica/domain/dto"
	"strings"
)

type MessageProvider struct {
	Messages []dto.Satellite
}

func (m *MessageProvider) AddMessage(satellite dto.Satellite) {
	exists := false
	for i := 0; i < len(m.Messages); i++ {
		if strings.ToLower(m.Messages[i].Name) == strings.ToLower(satellite.Name) {
			m.Messages[i] = satellite
			exists = true
			break
		}
	}
	if !exists {
		m.Messages = append(m.Messages, satellite)
	}
}

func (m *MessageProvider) GetMessages() dto.SatelliteRequest {
	return dto.SatelliteRequest{m.Messages}
}

func (m *MessageProvider) Initialize() {
	m.Messages = []dto.Satellite{}
}
