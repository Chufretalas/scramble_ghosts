package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.design/x/clipboard"
)

func (g *Game) EditModeUpdate() {

	if editSelected == "" {
		if inpututil.IsKeyJustPressed(ebiten.KeyN) {
			editSelected = "name"
			g.EditText = UInfo.Name
		} else if inpututil.IsKeyJustPressed(ebiten.KeyU) {
			editSelected = "ld_url"
			g.EditText = UInfo.LD_URL
		} else if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			editSelected = "ld_pass"
			g.EditText = UInfo.LD_Pass
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			editSelected = ""
			g.Mode = "title"
			UpdateUserInfo()
			LDConnection = "waiting..."
			go CheckLDConnection()
			return
		}
	} else {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			switch editSelected {
			case "name":
				if g.EditText != UInfo.Name {
					UInfo.Highscore = 0
				}
				UInfo.Name = g.EditText
			case "ld_url":
				UInfo.LD_URL = g.EditText
			case "ld_pass":
				UInfo.LD_Pass = g.EditText
			}
			editSelected = ""
			g.EditText = ""
			g.EditRunes = make([]rune, 0)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			editSelected = ""
			g.EditText = ""
			g.EditRunes = make([]rune, 0)
		}
		if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyV) && canUseClipboard {
			g.EditText += string(clipboard.Read(clipboard.FmtText))
		} else {
			g.EditRunes = ebiten.AppendInputChars(g.EditRunes[:0])
			g.EditText += string(g.EditRunes)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
			if len(g.EditText) > 0 {
				g.EditText = g.EditText[:len(g.EditText)-1]
			}
		}
		EditTextBlink += 1
		if EditTextBlink > 180 {
			EditTextBlink = 0
		}
	}
}
