package main

import (
	"fmt"

	u "github.com/Chufretalas/scramble_ghosts/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var EditTextBlink = 0

func (g *Game) EditModeDraw(screen *ebiten.Image) {
	nameOptionOp := &ebiten.DrawImageOptions{}
	nameOptionOp.GeoM.Scale(0.6, 0.6)
	nameOptionOp.GeoM.Translate(2, 24)
	text.DrawWithOptions(screen, fmt.Sprintf("(N) Name: %v", UInfo.Name), MyEpicGamerFont, nameOptionOp)

	urlOptionOp := &ebiten.DrawImageOptions{}
	urlOptionOp.GeoM.Scale(0.6, 0.6)
	urlOptionOp.GeoM.Translate(2, 50)
	text.DrawWithOptions(screen, fmt.Sprintf("(U) Leaderboard URL: %v", UInfo.LD_URL), MyEpicGamerFont, urlOptionOp)

	passOptionOp := &ebiten.DrawImageOptions{}
	passOptionOp.GeoM.Scale(0.6, 0.6)
	passOptionOp.GeoM.Translate(2, 76)
	text.DrawWithOptions(screen, fmt.Sprintf("(P) Leaderboard password: %v", UInfo.LD_Pass), MyEpicGamerFont, passOptionOp)

	if editSelected == "" {
		instructionsOp := &ebiten.DrawImageOptions{}
		instructionsText := "press 'N', 'U', 'P' to edit or 'ESC' to save and exit"
		instructionsTextWidth := u.GetTextWidth(MyEpicGamerFont, instructionsText)
		instructionsOp.GeoM.Translate(float64(SCREENWIDTH/2-instructionsTextWidth/2), SCREENHEIGHT-200)
		text.DrawWithOptions(screen, instructionsText, MyEpicGamerFont, instructionsOp)
	} else {
		t := g.EditText
		if EditTextBlink%60 < 30 {
			t += "_"
		}
		editTextOp := &ebiten.DrawImageOptions{}
		editTextWidth := u.GetTextWidth(MyEpicGamerFont, g.EditText)
		editTextOp.GeoM.Translate(float64(SCREENWIDTH/2-editTextWidth/2), SCREENHEIGHT-200)
		editTextOp.GeoM.Scale(0.8, 0.8)
		text.DrawWithOptions(screen, t, MyEpicGamerFont, editTextOp)

		instructionsOp := &ebiten.DrawImageOptions{}
		instructionsText := "press 'ESC' to cancel or 'ENTER' to confirm the changes"
		instructionsTextWidth := u.GetTextWidth(MyEpicGamerFont, instructionsText)
		instructionsOp.GeoM.Translate(float64(SCREENWIDTH/2-instructionsTextWidth/2), SCREENHEIGHT-20)
		text.DrawWithOptions(screen, instructionsText, MyEpicGamerFont, instructionsOp)
	}
}
