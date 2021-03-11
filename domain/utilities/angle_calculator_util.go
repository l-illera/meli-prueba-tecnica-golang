package utilities

import (
	"co.edu.meli/luisillera/prueba-tecnica/domain/model"
	"math"
)

func BetaAngle(distanceAB float64, distanceBC float64, distanceAC float64) float64 {
	op1 := (0.5 * math.Pow(distanceAB, 2)) - (0.5 * math.Pow(distanceBC, 2)) + (0.5 * math.Pow(distanceAC, 2))
	op2 := distanceAB * distanceAC
	acos := math.Acos(op1 / op2)
	return ToDegrees(acos)
}

func fi(First model.Point, Second model.Point) float64 {
	deltaX := Second.East - First.East
	deltaY := Second.North - First.North
	atan := math.Atan(deltaX / deltaY)
	return ToDegrees(atan)
}

func Azimuth(first model.Point, second model.Point) float64 {
	alpha := fi(first, second)
	return AdjustAzimuth(first, second, alpha)
}

func AdjustAzimuth(first model.Point, second model.Point, alpha float64) float64 {
	if first.East < second.East && first.North > second.North {
		return 180 - alpha
	} else if first.East > second.East && first.North > second.North {
		return 180 + alpha
	} else if first.East > second.East && first.North < second.North {
		return 360 - alpha
	} else {
		return alpha
	}
}
