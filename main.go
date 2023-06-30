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
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/solarlune/ebitick"
	"golang.org/x/image/font"
)

const (
	ScreenWidth    = 750
	ScreenHeight   = 380
	bV             = 4
	EnemyW         = 30
	EnemyH         = 30
	EnemySpawnTime = time.Millisecond * 100
	StoppingMult   = 4
)

var (
	bulletsToRemove []int
	ShotDelay       time.Duration
	CanShoot        bool
	MyEpicGamerFont font.Face
	showDebug       bool
	titleImage      *ebiten.Image
)

type Bullet struct {
	X, Y          float32
	Width, Height float32
}

type Game struct {
	Enemies     []*Enemy
	Bullets     []*Bullet
	Player      Player
	TimerSystem *ebitick.TimerSystem
	Score       int
	Mode        string
}

func (g *Game) Update() error {

	if g.Mode == "title" {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.Mode = "game"
		}
	} else {

		g.TimerSystem.Update()

		if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			showDebug = !showDebug
			NewRandomEnemy(ScreenWidth, ScreenHeight, 10)
		}

		g.Player.Move(8, 0.75)

		// fire bullets
		if CanShoot {
			g.Bullets = append(g.Bullets, &Bullet{g.Player.X + g.Player.Width/2, g.Player.Y, 10, 10})
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
				for bullet_index, bullet := range g.Bullets { //TODO: remove the -5 magic number once the bullets stop beign a weird circle
					if utils.IsColliding(bullet.X-5, bullet.Y, bullet.Width, bullet.Height, enemy.X, enemy.Y, enemy.Width, enemy.Height) {
						// enemy.hit = true
						enemy.Alive = false
						bulletsToRemove = append(bulletsToRemove, bullet_index)
						g.Score += enemy.Score
						break
					}
				}
				if utils.IsColliding(enemy.X, enemy.Y, enemy.Width, enemy.Height, g.Player.X, g.Player.Y, g.Player.Width, g.Player.Height) && enemy.Alive {
					enemy.Hit = true
					g.Mode = "title" //TODO: change here to add the game over screen
					g.ResetGame()
					return nil
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
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.Mode == "title" {
		titleOp := &ebiten.DrawImageOptions{}
		screen.DrawImage(titleImage, titleOp)

	} else {
		if showDebug {
			ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v\nBullets: %v\nEnemies: %v", ebiten.ActualFPS(), len(g.Bullets), len(g.Enemies)))
		}
		vector.DrawFilledRect(screen, g.Player.X, g.Player.Y, g.Player.Width, g.Player.Height, color.White, true)
		var enemyColor color.Color
		for _, enemy := range g.Enemies {
			if enemy.Alive {
				if enemy.Hit {
					enemyColor = color.RGBA{100, 200, 100, 255}
				} else {
					switch enemy.Type {
					case Linear:
						enemyColor = color.RGBA{200, 0, 0, 255}
					case CurveL:
						enemyColor = color.RGBA{255, 100, 100, 255}
					case CurveR:
						enemyColor = color.RGBA{50, 50, 200, 255}
					}
				}
			} else {
				enemyColor = color.RGBA{200, 0, 0, 255}
			}
			vector.DrawFilledRect(screen,
				enemy.X,
				enemy.Y,
				enemy.Width,
				enemy.Height,
				enemyColor,
				true)
		}

		for _, bullet := range g.Bullets {
			vector.DrawFilledCircle(screen, bullet.X, bullet.Y, bullet.Width, color.RGBA{255, 0, 0, 255}, true)
		}

		text.Draw(screen, fmt.Sprintf("Score: %v", g.Score), MyEpicGamerFont, 8, 20, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func SpawnEnemies(g *Game) {
	g.TimerSystem.After(EnemySpawnTime, func() {
		SpawnEnemies(g)
	})
	g.Enemies = append(g.Enemies, NewRandomEnemy(ScreenWidth, ScreenHeight, 3))
}

func main() {

	showDebug = false

	LoadFont()

	var tIError error

	titleImage, _, tIError = ebitenutil.NewImageFromFile("./assets/title_screen_750-380.png")

	if tIError != nil {
		log.Fatal("Title image did not load" + tIError.Error())
	}

	game := &Game{
		Enemies:     make([]*Enemy, 0),
		Bullets:     make([]*Bullet, 0),
		Player:      Player{X: ScreenWidth/2 - 15, Y: ScreenHeight - 30, Width: 30, Height: 30},
		TimerSystem: ebitick.NewTimerSystem(),
		Score:       0,
		Mode:        "title",
	}
	game.TimerSystem.After(EnemySpawnTime, func() {
		SpawnEnemies(game)
	})
	ShotDelay = time.Millisecond * 200
	CanShoot = true
	bulletsToRemove = make([]int, 0)
	ebiten.SetWindowSize(ScreenWidth*2, ScreenHeight*2)
	ebiten.SetWindowTitle("Scramble Ghosts 👻")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
