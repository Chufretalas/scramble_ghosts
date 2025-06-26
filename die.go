package main

import (
	u "github.com/Chufretalas/scramble_ghosts/utils"
)

func (g *Game) Die() {
	g.Mode = "gameover"
	g.FinalScore = g.Score
	g.ResetGame()
	if g.FinalScore > UInfo.Highscore {
		UInfo.Highscore = g.FinalScore
		GotHighscore = true
		if !u.IsWASM() {
			go UpdateUserInfo()
			if LDConnection == "ok" {
				go SendScore(g.FinalScore)
			}
		}
	} else {
		GotHighscore = false
	}
}
