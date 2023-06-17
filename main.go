package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/Chufretalas/scramble_ghosts/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/solarlune/ebitick"
)

const (
	screenWidth  = 600
	screenHeight = 350
	bV           = 2
)

var (
	bulletsToRemove []int
)

type Enemy struct {
	x, y          float32
	width, height float32
	hit           bool
	alive         bool
}

type Bullet struct {
	x, y          float32
	width, height float32
}

type Game struct {
	enemies     []*Enemy
	bullets     []*Bullet
	player      Player
	TimerSystem *ebitick.TimerSystem
}

func (g *Game) Update() error {
	g.TimerSystem.Update()

	// Player movement
	g.player.Move(10, 0.35)
	//end player movement

	// fire bullets
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.bullets = append(g.bullets, &Bullet{g.player.x + g.player.width/2, g.player.y, 10, 10})
	}

	// move bullets
	for i, bullet := range g.bullets {
		bullet.y -= bV
		if bullet.y+10 < 0 {
			bulletsToRemove = append(bulletsToRemove, i)
		}
	}

	//check for collisions
	for _, enemy := range g.enemies {
		if enemy.alive {
			for bullet_index, bullet := range g.bullets { //TODO: remove the -5 magic number once the bullets stop beign a weird circle
				if utils.IsColliding(bullet.x-5, bullet.y, bullet.width, bullet.height, enemy.x, enemy.y, enemy.width, enemy.height) {
					// enemy.hit = true
					enemy.alive = false
					bulletsToRemove = append(bulletsToRemove, bullet_index)
					break
				}
			}
			if utils.IsColliding(enemy.x, enemy.y, enemy.width, enemy.height, g.player.x, g.player.y, g.player.width, g.player.height) && enemy.alive {
				enemy.hit = true
			} else {
				enemy.hit = false
			}
		}
	}

	if len(bulletsToRemove) != 0 {
		newBullets := make([]*Bullet, 0, len(g.bullets)-len(bulletsToRemove))
		for i, bullet := range g.bullets {
			if !utils.InSlice(bulletsToRemove, i) {
				newBullets = append(newBullets, bullet)
			}
		}
		g.bullets = newBullets
		bulletsToRemove = make([]int, 0)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v\nBullets: %v", ebiten.ActualFPS(), len(g.bullets)))
	vector.DrawFilledRect(screen, g.player.x, g.player.y, g.player.width, g.player.height, color.White, true)
	var enemyColor color.Color
	for _, enemy := range g.enemies {
		if enemy.alive {
			if enemy.hit {
				enemyColor = color.RGBA{100, 200, 100, 255}
			} else {
				enemyColor = color.RGBA{255, 0, 0, 255}
			}
		} else {
			enemyColor = color.RGBA{200, 0, 0, 255}
		}
		vector.DrawFilledRect(screen,
			enemy.x,
			enemy.y,
			enemy.width,
			enemy.height,
			enemyColor,
			true)
	}

	for _, bullet := range g.bullets {
		vector.DrawFilledCircle(screen, bullet.x, bullet.y, bullet.width, color.RGBA{255, 0, 0, 255}, true)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func RepeatTimer(g *Game) {
	g.TimerSystem.After(time.Second, func() {
		RepeatTimer(g)
	})
	fmt.Println("opa")
}

func main() {
	game := &Game{
		enemies:     make([]*Enemy, 0),
		bullets:     make([]*Bullet, 0),
		player:      Player{x: 0, y: 0, width: 30, height: 30},
		TimerSystem: ebitick.NewTimerSystem(),
	}
	game.TimerSystem.After(time.Second, func() {
		RepeatTimer(game)
	})
	game.enemies = append(game.enemies, &Enemy{x: screenWidth / 2, y: screenHeight / 2, width: 30, height: 30, hit: false, alive: true})
	game.enemies = append(game.enemies, &Enemy{x: screenWidth/2 + 50, y: screenHeight / 2, width: 30, height: 30, hit: false, alive: true})
	game.enemies = append(game.enemies, &Enemy{x: screenWidth / 2, y: screenHeight/2 + 50, width: 30, height: 30, hit: false, alive: true})
	bulletsToRemove = make([]int, 0)
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Scramble Ghosts 👻")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
