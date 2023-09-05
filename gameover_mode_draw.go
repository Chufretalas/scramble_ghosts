package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) GameoverModeDraw(screen *ebiten.Image) {
	if GotHighscore {
		screen.DrawImage(GameoverImageHS, &ebiten.DrawImageOptions{})
		textSize := text.BoundString(MyEpicGamerFont, fmt.Sprintf("New Highscore: %v", g.FinalScore))
		text.Draw(screen, fmt.Sprintf("New Highscore: %v", g.FinalScore), MyEpicGamerFont, ScreenWidth/2-textSize.Size().X/2, ScreenHeight/2-30, color.White)
	} else {
		screen.DrawImage(GameoverImage, &ebiten.DrawImageOptions{})
		textSize := text.BoundString(MyEpicGamerFont, fmt.Sprintf("Score: %v", g.FinalScore))
		text.Draw(screen, fmt.Sprintf("Score: %v", g.FinalScore), MyEpicGamerFont, ScreenWidth/2-textSize.Size().X/2, ScreenHeight/2-70, color.White)
		hsTextSize := text.BoundString(MyEpicGamerFont, fmt.Sprintf("Current Highscore: %v", UInfo.Highscore))
		text.Draw(screen, fmt.Sprintf("Current Highscore: %v", UInfo.Highscore), MyEpicGamerFont, ScreenWidth/2-hsTextSize.Size().X/2, ScreenHeight/2-5, color.White)
	}
}
