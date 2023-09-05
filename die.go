package main

func (g *Game) Die() {
	g.Mode = "gameover"
	g.FinalScore = g.Score
	g.ResetGame()
	g.StartedTheTimers = false
	if g.FinalScore > UInfo.Highscore {
		UInfo.Highscore = g.FinalScore
		GotHighscore = true
		go UpdateUserInfo()
		go SendScore(g.FinalScore)
	} else {
		GotHighscore = false
	}
}
