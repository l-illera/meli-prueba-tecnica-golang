package utilities

import (
	"co.edu.meli/luisillera/prueba-tecnica/domain/model"
	"math"
)

type DistanceCalculatorUtil struct{}

func (d *DistanceCalculatorUtil) Process(first model.Point, second model.Point) float64 {
	deltaX := d.findDelta(second.East, first.East)
	deltaY := d.findDelta(second.North, first.North)
	distance := math.Sqrt(deltaX + deltaY)
	return distance
}

func (d *DistanceCalculatorUtil) findDelta(x1 float64, x2 float64) float64 {
	return math.Pow(x1-x2, 2)
}
