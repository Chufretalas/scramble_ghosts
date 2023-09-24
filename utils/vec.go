package utils

import "math"

type Vec struct {
	X float64
	Y float64
}

func (v *Vec) ToUnit() {
	v.EscDiv(v.GetMod())
}

func (v *Vec) LimitMod(threshold float64) {
	if v.GetMod() > threshold {
		v.ToUnit()
		v.EscMult(threshold)
	}
}

// v * n multiply vector by scalar
func (v *Vec) EscMult(n float64) {
	v.X *= n
	v.Y *= n
}

// v / n divide vector by scalar
func (v *Vec) EscDiv(n float64) {
	v.X = v.X / n
	v.Y = v.Y / n
}

func (v *Vec) Add(otherVec Vec) {
	v.X += otherVec.X
	v.Y += otherVec.Y
}

func (v Vec) GetMod() float64 {
	return math.Sqrt(math.Pow(float64(v.X), 2) + math.Pow(float64(v.Y), 2))
}
