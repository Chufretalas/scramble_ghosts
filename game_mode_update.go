package main

import (
	"github.com/Chufretalas/scramble_ghosts/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// 1 should return the Update function and 0 should continue
func (g *Game) GameModeUpdate() int {

	g.TimerSystem.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		InvincibleMode = !InvincibleMode
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		showDebug = !showDebug
		NewRandomEnemy(ScreenWidth, ScreenHeight, 10)
	}

	g.Player.Move(9, 0.8)

	// fire bullets
	if CanShoot {
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

	// move enemies and check for collisions
	for _, enemy := range g.Enemies {
		if enemy.Alive {
			enemy.Move()
			if enemy.X+enemy.Width < 0 || enemy.X > ScreenWidth || enemy.Y > ScreenHeight || enemy.Y < -(enemy.Height*2) {
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
			if utils.IsColliding(enemy.X, enemy.Y, enemy.Width, enemy.Height, g.Player.X, g.Player.Y, PlayerBaseSize*g.Player.SizeMult, PlayerBaseSize*g.Player.SizeMult) && enemy.Alive {
				enemy.Hit = true
				if !InvincibleMode {
					g.Mode = "gameover"
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
