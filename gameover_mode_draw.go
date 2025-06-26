package main

import (
	"fmt"
	"image/color"

	u "github.com/Chufretalas/scramble_ghosts/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) GameoverModeDraw(screen *ebiten.Image) {
	if GotHighscore {
		screen.DrawImage(Sprites.GameoverHS.Img, &ebiten.DrawImageOptions{})
		textWidth := u.GetTextWidth(MyEpicGamerFont, fmt.Sprintf("New Highscore: %v", g.FinalScore))
		text.Draw(screen, fmt.Sprintf("New Highscore: %v", g.FinalScore), MyEpicGamerFont, SCREENWIDTH/2-textWidth/2, SCREENHEIGHT/2-30, color.White)
	} else {
		screen.DrawImage(Sprites.Gameover.Img, &ebiten.DrawImageOptions{})
		textSize := u.GetTextWidth(MyEpicGamerFont, fmt.Sprintf("Score: %v", g.FinalScore))
		text.Draw(screen, fmt.Sprintf("Score: %v", g.FinalScore), MyEpicGamerFont, SCREENWIDTH/2-textSize/2, SCREENHEIGHT/2-70, color.White)
		hsTextSize := u.GetTextWidth(MyEpicGamerFont, fmt.Sprintf("Current Highscore: %v", UInfo.Highscore))
		text.Draw(screen, fmt.Sprintf("Current Highscore: %v", UInfo.Highscore), MyEpicGamerFont, SCREENWIDTH/2-hsTextSize/2, SCREENHEIGHT/2-5, color.White)
	}
}
