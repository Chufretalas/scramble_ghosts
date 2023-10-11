package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) GameModeDraw(screen *ebiten.Image) {

	if showDebug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v\nBullets: %v\nEnemies: %v\n%v", ebiten.ActualFPS(), len(g.PBullets), len(g.Enemies), g.Diff))
	}

	for _, bullet := range g.PBullets {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(bullet.X, bullet.Y)
		screen.DrawImage(PlayerBulletImage, op)
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
		enemyOp.GeoM.Translate(enemy.X, enemy.Y)
		screen.DrawImage(enemy.GetSprite(), enemyOp)
		enemyOp.ColorScale.Reset()
		enemyOp.GeoM.Reset()
	}

	for _, bullet := range g.EHBullets {
		bulletOp := &ebiten.DrawImageOptions{}
		bulletOp.GeoM.Translate(bullet.X, bullet.Y)
		if bullet.Size == 30 {
			screen.DrawImage(EnemyBullet30Image, bulletOp)
		} else {
			screen.DrawImage(EnemyBullet50Image, bulletOp)
		}
	}

	// Drawing arcshot
	arcshotOP := &ebiten.DrawImageOptions{}
	arcshotOP.GeoM.Translate(g.Arcshot.X, g.Arcshot.Y+100)
	switch g.Arcshot.State {
	case "idle":
		screen.DrawImage(ArcshotSheet.SubImage(image.Rect(0, 0, 150, 200)).(*ebiten.Image), arcshotOP)
	case "firing":
		screen.DrawImage(ArcshotSheet.SubImage(image.Rect(150, 0, 300, 200)).(*ebiten.Image), arcshotOP)
	}
	// End Arcshot

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
	text.Draw(screen, fmt.Sprintf("Score: %v", g.Score), MyEpicGamerFont, 15, 40, color.White)

	levelOp := &ebiten.DrawImageOptions{}
	levelOp.GeoM.Scale(0.7, 0.7)
	levelOp.GeoM.Translate(10, SCREENHEIGHT-15)
	text.DrawWithOptions(screen, fmt.Sprintf("Lvl: %v", g.Diff.Level), MyEpicGamerFont, levelOp)
}
