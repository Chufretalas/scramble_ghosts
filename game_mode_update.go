package main

import (
	"math/rand"

	"github.com/Chufretalas/scramble_ghosts/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// 1 should return the Update function and 0 should continue
func (g *Game) GameModeUpdate() int {

	g.TimerSystem.Update()

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
			g.Enemies = append(g.Enemies, NewRandomEnemy(SCREENWIDTH, SCREENHEIGHT, 6))
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
		NewRandomEnemy(SCREENWIDTH, SCREENHEIGHT, 10)
	}

	g.Player.Move(9, 0.8)

	// fire bullets
	if CanShoot && !InvincibleMode {
		g.Bullets = append(g.Bullets, &Bullet{g.Player.X + PlayerBaseSize*g.Player.SizeMult/2 - BulletBaseSize/2, g.Player.Y, 1})
		CanShoot = false
		g.TimerSystem.After(ShotDelay, func() { CanShoot = true })
	}

	// move bullets
	for i, bullet := range g.Bullets {
		bullet.Y -= bV
		if bullet.Y+10 < 0 {
			bulletsToRemove = append(bulletsToRemove, i)
		}
	}

	// spawn DWs, move them and check for collision with the player
	if !g.DWL.Active && !g.DWL.IsSpawning {
		if g.Player.X+PlayerBaseSize*g.Player.SizeMult/2 < SCREENWIDTH*0.35 {
			if n := rand.Int31n(g.Diff.DWSpawnChance); n == 10 {
				g.SpawnDeathWall("left")
			}
		}
	} else if g.DWL.Active {
		g.DWL.Move(g.Diff.DWSpeedMult)
		if g.Player.X < float32(g.DWL.X)+DWWidth-DWSafeZone {
			if !InvincibleMode {
				g.Die()
				return 0
			}
		}
	}

	if !g.DWR.Active && !g.DWR.IsSpawning {
		if g.Player.X+PlayerBaseSize*g.Player.SizeMult/2 > SCREENWIDTH*0.65 {
			if n := rand.Int31n(450); n == 10 {
				g.SpawnDeathWall("right")
			}
		}
	} else if g.DWR.Active {
		g.DWR.Move(g.Diff.DWSpeedMult)
		if g.Player.X+PlayerBaseSize*g.Player.SizeMult > float32(g.DWR.X)+DWSafeZone {
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
				if enemy.X < float32(g.DWL.X)+DWWidth-DWSafeZone {
					enemy.Alive = false
					continue
				}
			}
			if g.DWR.Active {
				if enemy.X+EnemyW > float32(g.DWR.X)+DWSafeZone {
					enemy.Alive = false
					continue
				}
			}
			if enemy.X+enemy.Width < 0 || enemy.X > SCREENWIDTH || enemy.Y > SCREENHEIGHT || enemy.Y < -(enemy.Height*2) {
				enemy.Alive = false
				continue
			}
			for bullet_index, bullet := range g.Bullets {
				if utils.IsColliding(bullet.X, bullet.Y, BulletBaseSize*bullet.SizeMult, BulletBaseSize*bullet.SizeMult, enemy.X, enemy.Y, enemy.Width, enemy.Height) {
					// enemy.hit = true
					enemy.Alive = false
					bulletsToRemove = append(bulletsToRemove, bullet_index)
					g.Score += enemy.Score
					break
				}
			}
			if utils.IsColliding(enemy.X, enemy.Y, enemy.Width, enemy.Height, g.Player.X+6, g.Player.Y+6, PlayerBaseSize*g.Player.SizeMult-12, PlayerBaseSize*g.Player.SizeMult-12) && enemy.Alive {
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
	new_enemies := make([]*Enemy, 0, len(g.Enemies))
	for _, enemy := range g.Enemies {
		if enemy.Alive {
			new_enemies = append(new_enemies, enemy)
		}
	}
	g.Enemies = new_enemies

	// Remove bullets
	// TODO: maybe use the same strategy to remove the bullets as is used for the enemies
	if len(bulletsToRemove) != 0 {
		bulletsToRemove = utils.RemoveDups(bulletsToRemove)
		newBullets := make([]*Bullet, 0, len(g.Bullets)-len(bulletsToRemove))
		for i, bullet := range g.Bullets {
			if !utils.InSlice(bulletsToRemove, i) {
				newBullets = append(newBullets, bullet)
			}
		}
		g.Bullets = newBullets
		bulletsToRemove = make([]int, 0)
	}
	return 0
}
