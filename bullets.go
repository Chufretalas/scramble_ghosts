package main

import (
	"math"

	"github.com/Chufretalas/scramble_ghosts/utils"
)

// Player bullet
type PBullet struct {
	X, Y  float32 // current position
	Rad   float64 // launch angle of the bullet in Radians
	Speed float32 // module of the movement vector
}

func (b *PBullet) Move() {
	b.X += float32(math.Cos(b.Rad)) * b.Speed
	b.Y -= float32(math.Sin(b.Rad)) * b.Speed
}

type EHommingBullet struct {
	X, Y     float32 // current position
	Vel      utils.Vec
	Strength float64
}

// Aways pass the player's CENTER!
func (b *EHommingBullet) Move(playerX, playerY float32) {
	acc := utils.Vec{X: float64(playerX - b.X - 15), Y: float64(playerY - b.Y - 15)}
	mod := acc.GetMod()

	if mod > 100 {
		b.X += float32(b.Vel.X)
		b.Y += float32(b.Vel.Y)
		return
	}

	acc.ToUnit()
	acc.EscMult(b.Strength)

	b.Vel.Add(acc)

	//  1   <-> 100
	// mult <-> mod
	mult := mod / 100

	// graph this in desmos if you are confused -ax^{2}+1
	b.Vel.LimitMod(5*(-0.6*math.Pow(mult, 2)+1) + 5)

	b.X += float32(b.Vel.X)
	b.Y += float32(b.Vel.Y)
}
