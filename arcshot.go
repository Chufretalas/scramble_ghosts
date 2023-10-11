package main

import (
	"math"
)

type Arcshot struct {
	X, Y       float64
	ShotsFired float64 // how many shots has he fired in one sweep
	rad        float64
	State      string // "idle" | "firing", this just controls the sprite that is being shown
	Active     bool
}

func (a *Arcshot) Move() {
	if a.X > SCREENWIDTH+150 {
		a.Reset()
		return
	}

	a.rad += 0.07
	a.X += 5
	middlePos := a.X + 75 - SCREENWIDTH/2
	a.Y = (-0.7*math.Pow(middlePos/(SCREENWIDTH/2), 2)+1)*300 - 250 + math.Sin(a.rad)*40
}

func (a *Arcshot) Reset() {
	a.rad = 0
	a.X = -150
	a.ShotsFired = 0
	a.State = "idle"
	a.Active = false
}
