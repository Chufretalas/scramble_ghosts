package main

import (
	"math/rand"
)

type EnemyType int64

const (
	Linear EnemyType = iota
	CurveL
	CurveR
)

const curveAcc = 0.05

func NewEnemy(X, Y, VX, VY float32) *Enemy {
	return &Enemy{X: X, Y: Y, VX: VX, VY: VY, Width: EnemyW, Height: EnemyH, Hit: false, Alive: true, Type: Linear, Score: 20}
}

func NewRandomEnemy(screenWidth, screenHeight, VY float32) *Enemy {

	x := float32(rand.Int31n(ScreenWidth - EnemyW))

	var eType EnemyType
	var vx float32
	var score int
	if v := rand.Int63n(6); v >= 4 {
		if x+EnemyH/2 < screenWidth/2 {
			eType = CurveL
		} else {
			eType = CurveR
		}
		vx = 0
		score = 50
	} else {
		eType = Linear
		vx = rand.Float32() - 0.5
		score = 20
	}

	return &Enemy{
		X:      x,
		Y:      -EnemyH - 5,
		VX:     vx,
		VY:     VY,
		Width:  EnemyW,
		Height: EnemyH,
		Hit:    false,
		Alive:  true,
		Type:   eType,
		Score:  score,
	}
}

type Enemy struct {
	X, Y          float32
	VX, VY        float32
	Width, Height float32
	Hit           bool
	Alive         bool
	Type          EnemyType
	Score         int
}

func (e *Enemy) Move() {
	maxVX := e.VY * 1.2
	switch e.Type {
	case CurveL:
		if e.VX <= maxVX {
			e.VX += curveAcc
		} else {
			e.VX = maxVX
		}
	case CurveR:
		if e.VX >= -maxVX {
			e.VX -= curveAcc
		} else {
			e.VX = -maxVX
		}
	}
	e.X += e.VX
	e.Y += e.VY
}
