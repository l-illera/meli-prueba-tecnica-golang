package utilities

import (
	"co.edu.meli/luisillera/prueba-tecnica/domain/model"
	"fmt"
	"math"
	"strconv"
)

func North(azimut float64, spaceshipDistance float64, satellite model.Point) float64 {
	return findPosition(satellite.North, math.Cos(ToRadians(azimut)), spaceshipDistance)
}

func East(azimut float64, spaceshipDistance float64, satellite model.Point) float64 {
	return findPosition(satellite.East, math.Sin(ToRadians(azimut)), spaceshipDistance)
}

func findPosition(initial float64, inclination float64, distance float64) float64 {
	gross := initial + inclination*distance
	round := fmt.Sprintf("%.1f", gross)
	net, _ := strconv.ParseFloat(round, 64)
	return net
}
