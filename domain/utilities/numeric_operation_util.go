package utilities

import (
	"math"
)

func ToRadians(degrees float64) float64 {
	return float64(degrees * math.Pi / 180)
}

func ToDegrees(radians float64) float64 {
	return radians * (180.0 / math.Pi)
}
