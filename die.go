package main

func (g *Game) Die() {
	g.Mode = "gameover"
	if g.Score > UInfo.Highscore {
		UInfo.Highscore = g.Score
		GotHighscore = true
		go SaveHighscore()
		go SendScore()
	} else {
		GotHighscore = false
	}
}
