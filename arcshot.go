package main

import (
	"math"
)

type Arcshot struct {
	X, Y  float64
	rad   float64
	State string // "idle" | "firing", this just controls the sprite that is being shown
}

func (a *Arcshot) Move() {
	a.rad += 0.1
	a.X += 5
	middlePos := a.X + 75 - SCREENWIDTH/2
	a.Y = (-0.7*math.Pow(middlePos/(SCREENWIDTH/2), 2)+1)*300 - 250 + math.Sin(a.rad)*40
}

func (a *Arcshot) Reset() {
	a.rad = 0
	a.X = -200
	a.State = "idle"
}
