package main

import (
	"math"

	"github.com/Chufretalas/scramble_ghosts/utils"
)

// Player bullet
type PBullet struct {
	X, Y  float64 // current position
	Rad   float64 // launch angle of the bullet in Radians
	Speed float64 // module of the movement vector
	Alive bool    // checks if the bullet should be removed
}

func (b *PBullet) Move() {
	b.X += math.Cos(b.Rad) * b.Speed
	b.Y -= math.Sin(b.Rad) * b.Speed
}

const HOMMING_RANGE = 180

type EHommingBullet struct {
	X, Y     float64 // current position
	Vel      utils.Vec
	Strength float64
	Size     float64 // 30 or 50
	Alive    bool    // checks if the bullet should be removed
}

// Aways pass the player's CENTER!
func (b *EHommingBullet) Move(playerX, playerY float64) {

	acc := utils.Vec{X: playerX - b.X - float64(b.Size)/2, Y: playerY - b.Y - float64(b.Size)/2}
	mod := acc.GetMod()

	if mod > HOMMING_RANGE {
		b.X += b.Vel.X
		b.Y += b.Vel.Y
		return
	}

	acc.ToUnit()
	acc.EscMult(b.Strength)

	b.Vel.Add(acc)

	//  1   <-> 100
	// mult <-> mod
	mult := mod / HOMMING_RANGE

	// graph this in desmos if you are confused -ax^{2}+1
	b.Vel.LimitMod(5*(-0.6*math.Pow(mult, 2)+1) + 5)

	b.X += b.Vel.X
	b.Y += b.Vel.Y
}
