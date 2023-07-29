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
	DWWidth        = 800
	VERSION        = "0.2.0"
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
	CurveLImage     *ebiten.Image
	CurveRImage     *ebiten.Image
	LinearImage     *ebiten.Image
	DWLImage        *ebiten.Image
	DWRImage        *ebiten.Image
	InvincibleMode  bool
	UserName        string
	ApiPass         string
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
	DWL         DW
	DWR         DW
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

		versionOp := &ebiten.DrawImageOptions{}
		versionOp.GeoM.Scale(0.5, 0.5)
		versionOp.GeoM.Translate(10, ScreenHeight-20)
		text.DrawWithOptions(screen, fmt.Sprintf("Version: %v", VERSION), MyEpicGamerFont, versionOp)

		userNameOp := &ebiten.DrawImageOptions{}
		userNameOp.GeoM.Scale(0.5, 0.5)
		userNameOp.GeoM.Translate(10, ScreenHeight-45)
		text.DrawWithOptions(screen, fmt.Sprintf("User Name: %v", UserName), MyEpicGamerFont, userNameOp)
	} else if g.Mode == "gameover" {
		screen.DrawImage(gameoverImage, &ebiten.DrawImageOptions{})
		textSize := text.BoundString(MyEpicGamerFont, fmt.Sprintf("Score: %v", g.Score))
		text.Draw(screen, fmt.Sprintf("Score: %v", g.Score), MyEpicGamerFont, ScreenWidth/2-textSize.Size().X/2, ScreenHeight/2-30, color.White)
	} else {
		if showDebug {
			ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v\nBullets: %v\nEnemies: %v\nDWL X:%v", ebiten.ActualFPS(), len(g.Bullets), len(g.Enemies), g.DWL.X))
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

		// draw enemies
		enemyOp := &ebiten.DrawImageOptions{}
		for _, enemy := range g.Enemies {
			if enemy.Hit {
				enemyOp.ColorScale.SetB(255)
				enemyOp.ColorScale.SetG(100)
			}
			enemyOp.GeoM.Translate(float64(enemy.X), float64(enemy.Y))
			switch enemy.Type {
			case Linear:
				screen.DrawImage(LinearImage, enemyOp)
			case CurveL:
				screen.DrawImage(CurveLImage, enemyOp)
			case CurveR:
				screen.DrawImage(CurveRImage, enemyOp)
			}
			enemyOp.ColorScale.Reset()
			enemyOp.GeoM.Reset()
		}

		// Death Walls ☠️
		if g.DWL.Active {
			dwlOp := &ebiten.DrawImageOptions{}
			dwlOp.GeoM.Translate(g.DWL.X, 0)
			screen.DrawImage(g.DWL.Image, dwlOp)
		}

		if g.DWR.Active {
			dwrOp := &ebiten.DrawImageOptions{}
			dwrOp.GeoM.Translate(g.DWR.X, 0)
			screen.DrawImage(g.DWR.Image, dwrOp)
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

	LoadUserInfo()

	LoadFont()

	LoadImages()

	game := &Game{
		Enemies:     make([]*Enemy, 0),
		Bullets:     make([]*Bullet, 0),
		Player:      Player{X: ScreenWidth/2 - 20, Y: ScreenHeight - 40, SizeMult: 1},
		TimerSystem: ebitick.NewTimerSystem(),
		Score:       0,
		Mode:        "title",
		DWL:         DW{Image: DWLImage, Active: false, Rad: 0, X: -DWWidth, Side: "left"},
		DWR:         DW{Image: DWRImage, Active: false, Rad: 0, X: ScreenWidth, Side: "right"},
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
