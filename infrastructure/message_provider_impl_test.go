package infrastructure

import (
	"co.edu.meli/luisillera/prueba-tecnica/domain/dto"
	"testing"
)

var provider MessageProvider

func TestMessageProvider_AddMessage(t *testing.T) {
	provider = MessageProvider{}
	provider.Initialize()
	var satellite = dto.Satellite{
		Name:     "Kenobi",
		Distance: 485.25,
		Message:  []string{"este", "", "un", "mensaje", ""},
	}
	type fields struct {
		messages []dto.Satellite
	}
	type args struct {
		satellite dto.Satellite
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wants  []dto.Satellite
	}{
		{"Satellite Not Exists",
			fields{messages: []dto.Satellite{}},
			args{satellite: satellite},
			[]dto.Satellite{
				satellite,
			}},
		{"Satellite Exists Yet",
			fields{messages: []dto.Satellite{satellite}},
			args{satellite: dto.Satellite{
				Name:     "Kenobi",
				Distance: 1234.25,
				Message:  []string{"este", "es", "", "mensaje", "secreto"},
			}},
			[]dto.Satellite{
				dto.Satellite{
					Name:     "Kenobi",
					Distance: 1234.25,
					Message:  []string{"este", "es", "", "mensaje", "secreto"},
				},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider.AddMessage(tt.args.satellite)
			if len(provider.Messages) != len(tt.wants) {
				t.Errorf("Expected [%d] Got [%d]", len(tt.wants), len(provider.Messages))
			}
			for v := range tt.wants {
				SatelliteComparator(t, tt.wants[v], provider.Messages[v])
			}
		})
	}
}

func SatelliteComparator(t *testing.T, expected dto.Satellite, actual dto.Satellite) {
	if expected.Name != actual.Name {
		t.Errorf("Name Didn't match Expected [%s] Got [%s]", expected.Name, actual.Name)
	}
	if expected.Distance != actual.Distance {
		t.Errorf("Distante didn't match: Expected [%.1f] Got [%.1f]", expected.Distance, actual.Distance)
	}
	if len(expected.Message) != len(actual.Message) {
		t.Errorf("Message Didn't match: Expected [%d] Got [%d]", len(expected.Message), len(actual.Message))
	}
	for v := range expected.Message {
		if expected.Message[v] != actual.Message[v] {
			t.Errorf("Message Part Didn't match: Expected [%s] Got [%s]", expected.Message[v], actual.Message[v])
		}
	}
}
