package main

func (g *Game) Die() {
	g.Mode = "gameover"
	g.FinalScore = g.Score
	g.ResetGame()
	if g.FinalScore > UInfo.Highscore {
		UInfo.Highscore = g.FinalScore
		GotHighscore = true
		go UpdateUserInfo()
		if LDConnection == "ok" {
			go SendScore(g.FinalScore)
		}
	} else {
		GotHighscore = false
	}
}
