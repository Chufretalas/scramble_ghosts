package main

import (
	"math"
	"math/rand"
)

type ACSide int

const (
	LEFT ACSide = iota
	RIGHT
)

type Arcshot struct {
	X, Y       float64
	ShotsFired float64 // how many shots has he fired in one sweep
	rad        float64
	State      string // "idle" | "firing", this just controls the sprite that is being shown
	Active     bool
	Side       ACSide
}

func (a *Arcshot) Move() {
	switch a.Side {
	case LEFT:
		if a.X > SCREENWIDTH+150 {
			a.Reset()
			return
		}
		a.X += 5

	case RIGHT:
		if a.X < -150 {
			a.Reset()
			return
		}
		a.X -= 5
	}

	a.rad += 0.07
	middlePos := a.X + 75 - SCREENWIDTH/2
	a.Y = (-0.7*math.Pow(middlePos/(SCREENWIDTH/2), 2)+1)*300 - 250 + math.Sin(a.rad)*40
}

func (a *Arcshot) Reset() {
	a.rad = 0
	a.ShotsFired = 0
	a.State = "idle"
	a.Active = false
	if rand.Int31n(2) == 1 {
		a.Side = LEFT
		a.X = -150
	} else {
		a.Side = RIGHT
		a.X = SCREENWIDTH + 150
	}
}
