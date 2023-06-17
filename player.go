package main

import "github.com/hajimehoshi/ebiten/v2"

var (
	walkedx bool
	walkedy bool
)

type Player struct {
	X, Y          float32
	Width, Height float32
	VX, VY        float32
}

func (p *Player) Move(maxV, acc float32) {
	walkedx = false
	walkedy = false
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if p.VX < maxV {
			if p.VX < 0 {
				p.VX = 0
			}
			p.VX += acc
		}
		walkedx = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if p.VX > -maxV {
			if p.VX > 0 {
				p.VX = 0
			}
			p.VX -= acc
		}
		walkedx = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if p.VY < maxV {
			if p.VY < 0 {
				p.VY = 0
			}
			p.VY += acc
		}
		walkedy = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if p.VY > -maxV {
			if p.VY > 0 {
				p.VY = 0
			}
			p.VY -= acc
		}
		walkedy = true
	}

	p.X += p.VX
	p.Y += p.VY

	//Check bounds
	if p.X+p.Width > ScreenWidth {
		p.X = ScreenWidth - p.Width
	}
	if p.X < 0 {
		p.X = 0
	}

	if p.Y+p.Height > ScreenHeight {
		p.Y = ScreenHeight - p.Height
	}
	if p.Y < 0 {
		p.Y = 0
	}
	// End check bounds

	if !walkedx {
		p.VX = 0
	}
	if !walkedy {
		p.VY = 0
	}
}
