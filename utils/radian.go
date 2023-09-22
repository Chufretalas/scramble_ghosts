package utils

import "math"

func Deg2Rad(deg float32) float64 {
	return float64((deg * math.Pi) / 180)
}
