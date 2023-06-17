package main

func NewEnemy(X, Y, VX, VY float32) *Enemy {
	return &Enemy{X: X, Y: Y, VX: VX, VY: VY, Width: EnemyW, Height: EnemyH, Hit: false, Alive: true}
}

type Enemy struct {
	X, Y          float32
	VX, VY        float32
	Width, Height float32
	Hit           bool
	Alive         bool
}

func (e *Enemy) Move() {
	e.X += e.VX
	e.Y += e.VY
}
