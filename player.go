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
}
