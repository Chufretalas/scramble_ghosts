package main

import "github.com/hajimehoshi/ebiten/v2"

var (
	walkedx bool
	walkedy bool
)

type Player struct {
	x, y          float32
	width, height float32
	vx, vy        float32
}

func (p *Player) Move(maxV, acc float32) {
	walkedx = false
	walkedy = false
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if p.vx < maxV {
			if p.vx < 0 {
				p.vx = 0
			}
			p.vx += acc
		}
		walkedx = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if p.vx > -maxV {
			if p.vx > 0 {
				p.vx = 0
			}
			p.vx -= acc
		}
		walkedx = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if p.vy < maxV {
			if p.vy < 0 {
				p.vy = 0
			}
			p.vy += acc
		}
		walkedy = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if p.vy > -maxV {
			if p.vy > 0 {
				p.vy = 0
			}
			p.vy -= acc
		}
		walkedy = true
	}

	p.x += p.vx
	p.y += p.vy

	if !walkedx {
		p.vx = 0
	}
	if !walkedy {
		p.vy = 0
	}
}
