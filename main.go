package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/Chufretalas/scramble_ghosts/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 600
	screenHeight = 350
	bV           = 2
)

var (
	toRemove []int
)

type Enemy struct {
	x, y          float32
	width, height float32
	hit           bool
}

type Player struct {
	x, y          float32
	width, height float32
	v             float32
}

type Bullet struct {
	x, y float32
}

type Game struct {
	enemies []*Enemy
	bullets []*Bullet
	player  Player
}

func (g *Game) Update() error {

	// Player movement
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.x += g.player.v
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.x -= g.player.v
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.player.y += g.player.v
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.player.y -= g.player.v
	}

	// fire bullets
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.bullets = append(g.bullets, &Bullet{g.player.x + g.player.width/2, g.player.y})
	}

	//check for player collisions
	for _, enemy := range g.enemies {
		if utils.IsColliding(enemy.x, enemy.y, enemy.width, enemy.height, g.player.x, g.player.y, g.player.width, g.player.height) {
			enemy.hit = true
		} else {
			enemy.hit = false
		}
	}

	// Remove bullets outr of frame
	for i, bullet := range g.bullets {
		bullet.y -= bV
		if bullet.y+10 < 0 {
			toRemove = append(toRemove, i)
		}
	}

	if len(toRemove) != 0 {
		newBullets := make([]*Bullet, 0, len(g.bullets)-len(toRemove))
		for i, bullet := range g.bullets {
			if !utils.InSlice(toRemove, i) {
				newBullets = append(newBullets, bullet)
			}
		}
		g.bullets = newBullets
		toRemove = make([]int, 0)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v\nBullets: %v", ebiten.ActualFPS(), len(g.bullets)))
	vector.DrawFilledRect(screen, g.player.x, g.player.y, g.player.width, g.player.height, color.White, true)
	var enemyColor color.Color
	for _, enemy := range g.enemies {
		if enemy.hit {
			enemyColor = color.RGBA{100, 200, 100, 255}
		} else {
			enemyColor = color.RGBA{255, 0, 0, 255}
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
		vector.DrawFilledCircle(screen, bullet.x, bullet.y, 10, color.RGBA64{255, 0, 0, 255}, true)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{
		enemies: make([]*Enemy, 0),
		bullets: make([]*Bullet, 0),
		player:  Player{0, 0, 30, 30, 10},
	}
	game.enemies = append(game.enemies, &Enemy{x: screenWidth / 2, y: screenHeight / 2, width: 30, height: 30, hit: false})
	toRemove = make([]int, 0)
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Scramble Ghosts 👻")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
