package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) GameoverModeUpdate() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.Mode = "game"
		g.ResetGame()
		return
	}
}
