package main

func (g *Game) ResetGame() {
	bulletsToRemove = make([]int, 0)
	g.Bullets = make([]*Bullet, 0)
	g.Enemies = make([]*Enemy, 0)
	g.Player.X = ScreenWidth/2 - 15
	g.Player.Y = ScreenHeight - 30
	g.Score = 0
	g.ShowDWWL = false
	g.ShowDWWR = false
	g.DWL.Reset()
	g.DWR.Reset()
}
