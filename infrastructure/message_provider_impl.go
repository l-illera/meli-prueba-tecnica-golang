package infrastructure

import (
	"co.edu.meli/luisillera/prueba-tecnica/domain/dto"
	"strings"
)

type MessageProvider struct {
	messages []dto.Satellite
}

func (m *MessageProvider) AddMessage(satellite dto.Satellite) {
	exists := false
	for i := 0; i < len(m.messages); i++ {
		if strings.ToLower(m.messages[i].Name) == strings.ToLower(satellite.Name) {
			m.messages[i] = satellite
			exists = true
			break
		}
	}
	if !exists {
		m.messages = append(m.messages, satellite)
	}
}

func (m *MessageProvider) GetMessages() dto.SatelliteRequest {
	return dto.SatelliteRequest{m.messages}
}

func (m *MessageProvider) Initialize() {
	m.messages = []dto.Satellite{}
}
