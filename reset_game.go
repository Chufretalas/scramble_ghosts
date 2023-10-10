package main

func (g *Game) ResetGame() {
	g.TimerSystem.Clear()
	CanShoot = true
	bulletsToRemove = make([]int, 0)
	g.PBullets = make([]*PBullet, 0)
	g.EHBullets = make([]*EHommingBullet, 0)
	g.Enemies = make([]*Enemy, 0)
	g.Player.X = SCREENWIDTH/2 - 15
	g.Player.Y = SCREENHEIGHT - 30
	g.Score = 0
	g.ShowDWWL = false
	g.ShowDWWR = false
	g.ShouldSpawnEnemy = true
	g.Arcshot.Reset()
	g.DWL.Reset()
	g.DWR.Reset()
	g.Diff.Reset()
}
