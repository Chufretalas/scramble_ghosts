package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// TODO: remove deprecated thigns here
func (g *Game) GameoverModeDraw(screen *ebiten.Image) {
	if GotHighscore {
		screen.DrawImage(Sprites.GameoverHS.Img, &ebiten.DrawImageOptions{})
		textSize := text.BoundString(MyEpicGamerFont, fmt.Sprintf("New Highscore: %v", g.FinalScore))
		text.Draw(screen, fmt.Sprintf("New Highscore: %v", g.FinalScore), MyEpicGamerFont, SCREENWIDTH/2-textSize.Size().X/2, SCREENHEIGHT/2-30, color.White)
	} else {
		screen.DrawImage(Sprites.Gameover.Img, &ebiten.DrawImageOptions{})
		textSize := text.BoundString(MyEpicGamerFont, fmt.Sprintf("Score: %v", g.FinalScore))
		text.Draw(screen, fmt.Sprintf("Score: %v", g.FinalScore), MyEpicGamerFont, SCREENWIDTH/2-textSize.Size().X/2, SCREENHEIGHT/2-70, color.White)
		hsTextSize := text.BoundString(MyEpicGamerFont, fmt.Sprintf("Current Highscore: %v", UInfo.Highscore))
		text.Draw(screen, fmt.Sprintf("Current Highscore: %v", UInfo.Highscore), MyEpicGamerFont, SCREENWIDTH/2-hsTextSize.Size().X/2, SCREENHEIGHT/2-5, color.White)
	}
}
