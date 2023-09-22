package main

import "math"

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
