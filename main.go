package main

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/solarlune/ebitick"
	"golang.org/x/image/font"

	_ "github.com/silbinarywolf/preferdiscretegpu"
)

const (
	ScreenWidth    = 1920
	ScreenHeight   = 1080
	bV             = 6
	BulletBaseSize = 30
	PlayerBaseSize = 40
	EnemyW         = 50
	EnemyH         = 50
	EnemySpawnTime = time.Millisecond * 50
	StoppingMult   = 4
)

var (
	bulletsToRemove []int
	ShotDelay       time.Duration
	CanShoot        bool
	MyEpicGamerFont font.Face
	showDebug       bool
	titleImage      *ebiten.Image
	gameoverImage   *ebiten.Image
	bulletImage     *ebiten.Image
	playerImage     *ebiten.Image
	InvincibleMode  bool
)

type Bullet struct {
	X, Y     float32
	SizeMult float32
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

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("ahahaha")
	}

	switch g.Mode {
	case "title":
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.Mode = "game"
		}
	case "game":
		g.GameModeUpdate()
	case "gameover":
		g.GameoverModeUpdate()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.Mode == "title" {
		titleOp := &ebiten.DrawImageOptions{}
		titleOp.GeoM.Scale(0.5, 0.5)
		screen.DrawImage(titleImage, titleOp)
	} else if g.Mode == "gameover" {
		screen.DrawImage(gameoverImage, &ebiten.DrawImageOptions{})
		textSize := text.BoundString(MyEpicGamerFont, fmt.Sprintf("Score: %v", g.Score))
		text.Draw(screen, fmt.Sprintf("Score: %v", g.Score), MyEpicGamerFont, ScreenWidth/2-textSize.Size().X/2, ScreenHeight/2-30, color.White)
	} else {
		if showDebug {
			ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v\nBullets: %v\nEnemies: %v", ebiten.ActualFPS(), len(g.Bullets), len(g.Enemies)))
		}

		for _, bullet := range g.Bullets {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(bullet.X), float64(bullet.Y))
			screen.DrawImage(bulletImage, op)
		}

		// draw player
		playerOp := &ebiten.DrawImageOptions{}
		playerOp.GeoM.Translate(float64(g.Player.X), float64(g.Player.Y))

		screen.DrawImage(playerImage, playerOp)

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

		text.Draw(screen, fmt.Sprintf("Score: %v", g.Score), MyEpicGamerFont, 20, 40, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func SpawnEnemies(g *Game) {
	g.TimerSystem.After(EnemySpawnTime, func() {
		SpawnEnemies(g)
	})
	g.Enemies = append(g.Enemies, NewRandomEnemy(ScreenWidth, ScreenHeight, 6))
}

func main() {

	showDebug = false
	InvincibleMode = false

	LoadFont()

	LoadImages()

	game := &Game{
		Enemies:     make([]*Enemy, 0),
		Bullets:     make([]*Bullet, 0),
		Player:      Player{X: ScreenWidth/2 - 20, Y: ScreenHeight - 40, SizeMult: 1},
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
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Scramble Ghosts 👻")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
