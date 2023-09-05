package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) GameModeDraw(screen *ebiten.Image) {
	if showDebug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v\nBullets: %v\nEnemies: %v\nDWL X:%v", ebiten.ActualFPS(), len(g.Bullets), len(g.Enemies), g.DWL.X))
	}

	for _, bullet := range g.Bullets {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(bullet.X), float64(bullet.Y))
		screen.DrawImage(BulletImage, op)
	}

	// draw player
	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(float64(g.Player.X), float64(g.Player.Y))

	screen.DrawImage(g.Player.GetSprite(), playerOp)

	// draw enemies
	enemyOp := &ebiten.DrawImageOptions{}
	for _, enemy := range g.Enemies {
		if enemy.Hit {
			enemyOp.ColorScale.SetB(255)
			enemyOp.ColorScale.SetG(100)
		}
		enemyOp.GeoM.Translate(float64(enemy.X), float64(enemy.Y))
		screen.DrawImage(enemy.GetSprite(), enemyOp)
		enemyOp.ColorScale.Reset()
		enemyOp.GeoM.Reset()
	}

	// Death Walls warnings
	if g.ShowDWWL {
		screen.DrawImage(DWWLImage, nil)
	}

	if g.ShowDWWR {
		DWWROp := &ebiten.DrawImageOptions{}
		DWWROp.GeoM.Translate(1200, 0)
		screen.DrawImage(DWWRImage, DWWROp)
	}

	// Death Walls ☠️
	if g.DWL.Active {
		dwlOp := &ebiten.DrawImageOptions{}
		dwlOp.GeoM.Translate(g.DWL.X, 0)
		screen.DrawImage(g.DWL.Image, dwlOp)
	}

	if g.DWR.Active {
		dwrOp := &ebiten.DrawImageOptions{}
		dwrOp.GeoM.Translate(g.DWR.X, 0)
		screen.DrawImage(g.DWR.Image, dwrOp)
	}

	// Score 🏆
	text.Draw(screen, fmt.Sprintf("Score: %v", g.Score), MyEpicGamerFont, 20, 40, color.White)
}
