package utils

import "math"

type Vec struct {
	X float64
	Y float64
}

func (v *Vec) ToUnit() *Vec {
	v.EscDiv(v.GetMod())
	return v
}

func (v *Vec) LimitMod(threshold float64) *Vec {
	if v.GetMod() > threshold {
		v.ToUnit()
		v.EscMult(threshold)
	}
	return v
}

// v * n multiply vector by scalar
func (v *Vec) EscMult(n float64) *Vec {
	v.X *= n
	v.Y *= n
	return v
}

// v / n divide vector by scalar
func (v *Vec) EscDiv(n float64) *Vec {
	v.X = v.X / n
	v.Y = v.Y / n
	return v
}

func (v *Vec) Add(otherVec Vec) *Vec {
	v.X += otherVec.X
	v.Y += otherVec.Y
	return v
}

func (v Vec) GetMod() float64 {
	return math.Sqrt(math.Pow(float64(v.X), 2) + math.Pow(float64(v.Y), 2))
}
