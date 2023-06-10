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
	pAcc         = 0.4 // player acceleration
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
	vx, vy        float32
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
	walkedx := false
	walkedy := false

	// Player movement
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if g.player.vx < 10 {
			if g.player.vx < 0 {
				g.player.vx = 0
			}
			g.player.vx += pAcc
		}
		walkedx = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if g.player.vx > -10 {
			if g.player.vx > 0 {
				g.player.vx = 0
			}
			g.player.vx -= pAcc
		}
		walkedx = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if g.player.vy < 10 {
			if g.player.vy < 0 {
				g.player.vy = 0
			}
			g.player.vy += pAcc
		}
		walkedy = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if g.player.vy > -10 {
			if g.player.vy > 0 {
				g.player.vy = 0
			}
			g.player.vy -= pAcc
		}
		walkedy = true
	}

	g.player.x += g.player.vx
	g.player.y += g.player.vy

	if !walkedx {
		g.player.vx = 0
	}
	if !walkedy {
		g.player.vy = 0
	}
	//end player movement

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
		vector.DrawFilledCircle(screen, bullet.x, bullet.y, 10, color.RGBA{255, 0, 0, 255}, true)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{
		enemies: make([]*Enemy, 0),
		bullets: make([]*Bullet, 0),
		player:  Player{x: 0, y: 0, width: 30, height: 30},
	}
	game.enemies = append(game.enemies, &Enemy{x: screenWidth / 2, y: screenHeight / 2, width: 30, height: 30, hit: false})
	game.enemies = append(game.enemies, &Enemy{x: screenWidth/2 + 50, y: screenHeight / 2, width: 30, height: 30, hit: false})
	game.enemies = append(game.enemies, &Enemy{x: screenWidth / 2, y: screenHeight/2 + 50, width: 30, height: 30, hit: false})
	toRemove = make([]int, 0)
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Scramble Ghosts 👻")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
