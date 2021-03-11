package infrastructure

import (
	"co.edu.meli/luisillera/prueba-tecnica/domain/model"
	"co.edu.meli/luisillera/prueba-tecnica/domain/model/sat"
	"co.edu.meli/luisillera/prueba-tecnica/domain/utilities"
	"fmt"
)

type SatelliteProvider struct {
	Kenobi     sat.Kenobi
	Skywalker  sat.Skywalker
	Sato       sat.Sato
	calculator utilities.DistanceCalculatorUtil
}

func (s *SatelliteProvider) Initialize(kenobi model.Point, skywalker model.Point, sato model.Point) {
	s.calculator = utilities.DistanceCalculatorUtil{}
	s.buildKenobi(kenobi, skywalker, sato)
	s.buildSkywalker(kenobi, skywalker, sato)
	s.buildSato(kenobi, skywalker, sato)
	fmt.Println()
}

func (s *SatelliteProvider) buildKenobi(kenobi model.Point, skywalker model.Point, sato model.Point) {
	s.Kenobi = sat.Kenobi{
		Coordinates: kenobi,
		Skywalker:   s.calculator.Process(kenobi, skywalker),
		Sato:        s.calculator.Process(kenobi, sato),
	}
	fmt.Printf("\nKENOBI %s", s.Kenobi)
}

func (s *SatelliteProvider) buildSkywalker(kenobi model.Point, skywalker model.Point, sato model.Point) {
	s.Skywalker = sat.Skywalker{
		Coordinates: skywalker,
		Sato:        s.calculator.Process(skywalker, sato),
		Kenobi:      s.calculator.Process(kenobi, skywalker),
	}
	fmt.Printf("\nSKYWALKER %s", s.Skywalker)
}

func (s *SatelliteProvider) buildSato(kenobi model.Point, skywalker model.Point, sato model.Point) {
	s.Sato = sat.Sato{
		Coordinates: sato,
		Skywalker:   s.calculator.Process(skywalker, sato),
		Kenobi:      s.calculator.Process(kenobi, sato),
	}
	fmt.Printf("\nSATO %s", s.Sato)
}
