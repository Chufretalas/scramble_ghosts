package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	walkedx bool
	walkedy bool
)

type Player struct {
	X, Y   float32
	VX, VY float32
}

func (p *Player) GetSprite() *ebiten.Image {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		return PlayerSheet.SubImage(image.Rect(40, 0, 80, 40)).(*ebiten.Image)
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		return PlayerSheet.SubImage(image.Rect(80, 0, 120, 40)).(*ebiten.Image)
	}

	return PlayerSheet.SubImage(image.Rect(0, 0, 40, 40)).(*ebiten.Image)
}

func (p *Player) Move(maxV, acc float32) {
	walkedx = false
	walkedy = false
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if p.VX < maxV {
			if p.VX < 0 && p.VX < -acc*StoppingMult {
				p.VX += acc * StoppingMult
			} else if p.VX < 0 {
				p.VX = 0
			} else {
				p.VX += acc
			}
		}
		walkedx = !walkedx
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if p.VX > -maxV {
			if p.VX > 0 && p.VX > acc*StoppingMult {
				p.VX -= acc * StoppingMult
			} else if p.VX > 0 {
				p.VX = 0
			} else {
				p.VX -= acc
			}
		}
		walkedx = !walkedx
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if p.VY < maxV {
			if p.VY < 0 && p.VY < -acc*StoppingMult {
				p.VY += acc * StoppingMult
			} else if p.VY < 0 {
				p.VY = 0
			} else {
				p.VY += acc
			}
		}
		walkedy = !walkedy
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if p.VY > -maxV {
			if p.VY > 0 && p.VY > acc*StoppingMult {
				p.VY += acc * StoppingMult
			} else if p.VY > 0 {
				p.VY = 0
			} else {
				p.VY -= acc
			}
		}
		walkedy = !walkedy
	}

	// Actually walk
	if walkedx {
		p.X += p.VX
	} else {
		p.VX = 0
	}

	if walkedy {
		p.Y += p.VY
	} else {
		p.VY = 0
	}
	// End actually walk

	//Check bounds
	if p.X+PlayerBaseSize > SCREENWIDTH {
		p.X = SCREENWIDTH - PlayerBaseSize
	}
	if p.X < 0 {
		p.X = 0
	}

	if p.Y+PlayerBaseSize > SCREENHEIGHT {
		p.Y = SCREENHEIGHT - PlayerBaseSize
	}
	if p.Y < 0 {
		p.Y = 0
	}
	// End check bounds
}
