package main

import (
	"fmt"
	"github.com/Chufretalas/scramble_ghosts/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

const (
	screenWidth  = 600
	screenHeight = 350
	pWidth       = 20
	pHeight      = 20
	bV           = 2
)

var (
	toRemove []int
)

type Enemy struct {
	x, y float64
}

type Player struct {
	x, y float64
	v    float64
}

type Bullet struct {
	x, y float64
}

type Game struct {
	enemies []Enemy
	bullets []Bullet
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
		g.bullets = append(g.bullets, Bullet{g.player.x + pWidth/2, g.player.y})
	}

	// Remove bullets outr of frame
	for i := range g.bullets {
		g.bullets[i].y -= bV
		if g.bullets[i].y+10 < 0 {
			toRemove = append(toRemove, i)
		}
	}

	if len(toRemove) != 0 {
		newBullets := make([]Bullet, 0, len(g.bullets)-len(toRemove))
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
	vector.DrawFilledRect(screen, float32(g.player.x), float32(g.player.y), pWidth, pHeight, color.White, true)
	for _, bullet := range g.bullets {
		vector.DrawFilledCircle(screen, float32(bullet.x), float32(bullet.y), 10, color.RGBA{255, 0, 0, 255}, true)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{
		enemies: make([]Enemy, 0),
		bullets: make([]Bullet, 0),
		player:  Player{0, 0, 10},
	}
	toRemove = make([]int, 0)
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Scramble Ghosts 👻")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
