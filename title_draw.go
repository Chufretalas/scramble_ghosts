package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) TitleDraw(screen *ebiten.Image) {
	titleOp := &ebiten.DrawImageOptions{}
	titleOp.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(TitleImage, titleOp)

	if LDConnection == "ok" {
		ldButtonOP := &ebiten.DrawImageOptions{}
		if x, y := ebiten.CursorPosition(); x <= 350 && y <= 200 {
			screen.DrawImage(LDButtonActiveImage, ldButtonOP)
		} else {
			screen.DrawImage(LDButtonImage, ldButtonOP)
		}
	}

	ldConnectionOp := &ebiten.DrawImageOptions{}
	ldConnectionOp.GeoM.Scale(0.5, 0.5)
	ldConnectionOp.GeoM.Translate(10, SCREENHEIGHT-5)
	text.DrawWithOptions(screen, fmt.Sprintf("Leaderboard connection: %v", LDConnection), MyEpicGamerFont, ldConnectionOp)

	versionOp := &ebiten.DrawImageOptions{}
	versionOp.GeoM.Scale(0.5, 0.5)
	versionOp.GeoM.Translate(10, SCREENHEIGHT-30)
	text.DrawWithOptions(screen, fmt.Sprintf("Version: %v", VERSION), MyEpicGamerFont, versionOp)

	userNameOp := &ebiten.DrawImageOptions{}
	userNameOp.GeoM.Scale(0.5, 0.5)
	userNameOp.GeoM.Translate(10, SCREENHEIGHT-55)
	text.DrawWithOptions(screen, fmt.Sprintf("User Name: %v", UInfo.Name), MyEpicGamerFont, userNameOp)

	editInfoMsg := "press i to edit your leaderboard information"
	editInfoMsgOp := &ebiten.DrawImageOptions{}
	editInfoMsgOp.GeoM.Scale(0.5, 0.5)
	editInfoMsgOp.GeoM.Translate(
		float64((SCREENWIDTH - text.BoundString(MyEpicGamerFont, editInfoMsg).Dx()/2 - 15)),
		float64(SCREENHEIGHT-text.BoundString(MyEpicGamerFont, editInfoMsg).Dy()/2+10))
	text.DrawWithOptions(screen, editInfoMsg, MyEpicGamerFont, editInfoMsgOp)

}
