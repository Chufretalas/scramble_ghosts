package main

import (
	"fmt"

	u "github.com/Chufretalas/scramble_ghosts/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) TitleDraw(screen *ebiten.Image) {
	titleOp := &ebiten.DrawImageOptions{}
	titleOp.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(Sprites.Title.Img, titleOp)

	versionOp := &ebiten.DrawImageOptions{}
	versionOp.GeoM.Scale(0.5, 0.5)
	versionOp.GeoM.Translate(10, SCREENHEIGHT-30)
	text.DrawWithOptions(screen, fmt.Sprintf("Version: %v", VERSION), MyEpicGamerFont, versionOp)

	if !u.IsWASM() {
		if LDConnection == "ok" {
			ldButtonOP := &ebiten.DrawImageOptions{}
			if x, y := ebiten.CursorPosition(); x <= 350 && y <= 200 {
				screen.DrawImage(Sprites.LDButtonActive.Img, ldButtonOP)
			} else {
				screen.DrawImage(Sprites.LDButton.Img, ldButtonOP)
			}
		}

		ldConnectionOp := &ebiten.DrawImageOptions{}
		ldConnectionOp.GeoM.Scale(0.5, 0.5)
		ldConnectionOp.GeoM.Translate(10, SCREENHEIGHT-5)
		text.DrawWithOptions(screen, fmt.Sprintf("Leaderboard connection: %v", LDConnection), MyEpicGamerFont, ldConnectionOp)

		userNameOp := &ebiten.DrawImageOptions{}
		userNameOp.GeoM.Scale(0.5, 0.5)
		userNameOp.GeoM.Translate(10, SCREENHEIGHT-55)
		text.DrawWithOptions(screen, fmt.Sprintf("User Name: %v", UInfo.Name), MyEpicGamerFont, userNameOp)

		editInfoMsg := "press i to edit your leaderboard information"
		editInfoMsgOp := &ebiten.DrawImageOptions{}
		editInfoMsgOp.GeoM.Scale(0.5, 0.5)
		editInfoMsgOp.GeoM.Translate(
			float64((SCREENWIDTH - u.GetTextWidth(MyEpicGamerFont, editInfoMsg)/2 - 15)),
			float64(SCREENHEIGHT-u.GetTextHeight(MyEpicGamerFont, editInfoMsg)/2+10))
		text.DrawWithOptions(screen, editInfoMsg, MyEpicGamerFont, editInfoMsgOp)
	}

}
