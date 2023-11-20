package main

import (
	"math/rand"
	"slices"
	"time"

	"github.com/Chufretalas/scramble_ghosts/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// 1 should return the Update function and 0 should continue
func (g *Game) GameModeUpdate() int {

	g.TimerSystem.Update()

	//arcshot stuff
	if g.Arcshot.Active {
		g.Arcshot.Move()
		if g.Arcshot.ShotsFired < float64(g.Diff.ArcshotShots) && rand.Int31n(200) == 10 {
			g.Arcshot.ShotsFired++

			// X: Archsot.X + Arcshot.Width/2 - bullet.Size = Archsot.X + 75 - 50
			origin := utils.Vec{X: g.Arcshot.X + 50, Y: g.Arcshot.Y + 220}
			vel := utils.Vec{X: g.Player.X + PlayerBaseSize/2 - origin.X, Y: g.Player.Y + PlayerBaseSize/2 - origin.Y}
			vel.ToUnit().EscMult(10)
			g.EHBullets = append(g.EHBullets, &EHommingBullet{X: origin.X, Y: origin.Y, Vel: vel, Strength: 0.5, Size: 50, Alive: true})
			g.Arcshot.State = "firing"
			g.TimerSystem.After(time.Second, func() { g.Arcshot.State = "idle" })
		}
	} else {
		g.TimerSystem.After(g.Diff.ArcshotDelay, func() {
			g.Arcshot.Active = true
		})
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.Arcshot.Reset()
	}

	// increase difficulty
	if g.Diff.ShouldIncrease {
		g.Diff.ShouldIncrease = false
		g.Diff.Increase()
		g.TimerSystem.After(DIFF_INCREASE_DELAY, func() {
			g.Diff.ShouldIncrease = true
		})
	}

	// spawn enemies
	if g.ShouldSpawnEnemy {
		g.ShouldSpawnEnemy = false
		for i := 0; i < g.Diff.EnemiesPerSpawn; i++ {
			g.Enemies = append(g.Enemies, NewRandomEnemy(6))
		}
		g.TimerSystem.After(g.Diff.EnemySpawnDelay, func() {
			g.ShouldSpawnEnemy = true
		})
	}

	// debug stuff
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		InvincibleMode = !InvincibleMode
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		showDebug = !showDebug
	}

	g.Player.Move(10, 0.8)

	// fire bullets
	if CanShoot && !InvincibleMode {
		var angle float32
		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			angle = 120
		} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			angle = 60
		} else {
			angle = 90
		}
		g.PBullets = append(g.PBullets, &PBullet{X: g.Player.X + PlayerBaseSize/2 - PlayerBulletSize/2, Y: g.Player.Y, Rad: utils.Deg2Rad(angle), Speed: 6, Alive: true})
		CanShoot = false
		g.TimerSystem.After(ShotDelay, func() { CanShoot = true })
	}

	// move bullets
	for _, bullet := range g.EHBullets {
		bullet.Move(g.Player.X+PlayerBaseSize/2, g.Player.Y+PlayerBaseSize/2)
		if bullet.Y+bullet.Size < 0 || bullet.Y > SCREENHEIGHT || bullet.X+bullet.Size < 0 || bullet.X > SCREENWIDTH {
			bullet.Alive = false
		}
		if utils.IsColliding(bullet.X, bullet.Y, bullet.Size, bullet.Size, g.Player.X+6, g.Player.Y+6, PlayerBaseSize-12, PlayerBaseSize-12) {
			bullet.Alive = false
			if !InvincibleMode {
				g.Die()
				return 0
			}
		}
	}

	for _, bullet := range g.PBullets {
		bullet.Move()
		if bullet.Y+26 < 0 || bullet.Y > SCREENHEIGHT || bullet.X+26 < 0 || bullet.X > SCREENWIDTH {
			bullet.Alive = false
		}
	}
	// end move bullets

	// spawn DWs, move them and check for collision with the player
	if !g.DWL.Active && !g.DWL.IsSpawning {
		if g.Player.X+PlayerBaseSize/2 < SCREENWIDTH*0.35 {
			if n := rand.Int31n(g.Diff.DWSpawnChance); n == 10 {
				g.SpawnDeathWall("left")
			}
		}
	} else if g.DWL.Active {
		g.DWL.Move(g.Diff.DWSpeedMult)
		if g.Player.X < g.DWL.X+DWWidth-DWSafeZone {
			if !InvincibleMode {
				g.Die()
				return 0
			}
		}
	}

	if !g.DWR.Active && !g.DWR.IsSpawning {
		if g.Player.X+PlayerBaseSize/2 > SCREENWIDTH*0.65 {
			if n := rand.Int31n(450); n == 10 {
				g.SpawnDeathWall("right")
			}
		}
	} else if g.DWR.Active {
		g.DWR.Move(g.Diff.DWSpeedMult)
		if float64(g.Player.X)+PlayerBaseSize > g.DWR.X+DWSafeZone {
			if !InvincibleMode {
				g.Die()
				return 0
			}
		}
	}

	// move enemies and check for collisions
	for _, enemy := range g.Enemies {
		if enemy.Alive {
			enemy.Move(g.Diff.EnemySpeedMult)
			if g.DWL.Active {
				if enemy.X < g.DWL.X+DWWidth-DWSafeZone {
					enemy.Alive = false
					continue
				}
			}
			if g.DWR.Active {
				if enemy.X+EnemyW > g.DWR.X+DWSafeZone {
					enemy.Alive = false
					continue
				}
			}
			if enemy.X+enemy.Width < 0 || enemy.X > SCREENWIDTH || enemy.Y > SCREENHEIGHT || enemy.Y < -(enemy.Height*2) {
				enemy.Alive = false
				continue
			}
			for _, bullet := range g.PBullets {
				if !bullet.Alive {
					continue
				}

				if utils.IsColliding(bullet.X, bullet.Y, PlayerBulletSize, PlayerBulletSize, enemy.X, enemy.Y, enemy.Width, enemy.Height) {
					// enemy.hit = true
					enemy.Alive = false
					bullet.Alive = false
					g.Score += enemy.Score
					break
				}
			}
			if utils.IsColliding(enemy.X, enemy.Y, enemy.Width, enemy.Height, g.Player.X+6, g.Player.Y+6, PlayerBaseSize-12, PlayerBaseSize-12) && enemy.Alive {
				enemy.Hit = true
				if !InvincibleMode {
					g.Die()
					return 0
				}
			} else {
				enemy.Hit = false
			}
		}
	}

	// Remove enemies
	g.Enemies = slices.DeleteFunc(g.Enemies, func(e *Enemy) bool {
		return !e.Alive
	})

	// Remove bullets
	g.PBullets = slices.DeleteFunc(g.PBullets, func(b *PBullet) bool {
		return !b.Alive
	})

	g.EHBullets = slices.DeleteFunc(g.EHBullets, func(b *EHommingBullet) bool {
		return !b.Alive
	})

	return 0
}
