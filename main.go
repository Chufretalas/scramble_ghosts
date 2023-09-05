package main

import (
	"errors"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/ebitick"
	"golang.design/x/clipboard"
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
	DWSafeZone     = 80
	VERSION        = "0.2.2"
)

var (
	bulletsToRemove     []int
	ShotDelay           time.Duration
	CanShoot            bool
	MyEpicGamerFont     font.Face
	showDebug           bool
	titleImage          *ebiten.Image
	LDButtonImage       *ebiten.Image
	LDButtonActiveImage *ebiten.Image
	gameoverImage       *ebiten.Image
	gameoverImageHS     *ebiten.Image
	bulletImage         *ebiten.Image
	playerSheet         *ebiten.Image
	CurveLImage         *ebiten.Image
	CurveRImage         *ebiten.Image
	LinearImage         *ebiten.Image
	DWLImage            *ebiten.Image
	DWRImage            *ebiten.Image
	DWWLImage           *ebiten.Image // death wall warning
	DWWRImage           *ebiten.Image // death wall warning
	InvincibleMode      bool
	UInfo               UserInfo
	LDConnection        string
	GotHighscore        bool
	editSelected        string
	canUseClipboard     bool
)

type Bullet struct {
	X, Y     float32
	SizeMult float32
}

type Game struct {
	Enemies          []*Enemy
	Bullets          []*Bullet
	Player           Player
	TimerSystem      *ebitick.TimerSystem
	Score            int
	FinalScore       int
	Mode             string
	DWL              DW
	DWR              DW
	ShowDWWL         bool
	ShowDWWR         bool
	EditText         string
	EditRunes        []rune
	StartedTheTimers bool // such as the enemy spawner
}

func (g *Game) Update() error {

	switch g.Mode {
	case "title":
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			return errors.New("ahahaha")
		}
		g.TitleUpdate()
	case "gameover":
		g.GameoverModeUpdate()
	case "game":
		g.GameModeUpdate()
	case "edit":
		g.EditModeUpdate()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.Mode {
	case "title":
		g.TitleDraw(screen)
	case "gameover":
		g.GameoverModeDraw(screen)
	case "game":
		g.GameModeDraw(screen)
	case "edit":
		g.EditModeDraw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {

	err := clipboard.Init()
	if err != nil {
		canUseClipboard = false
	} else {
		canUseClipboard = true
	}

	showDebug = false
	InvincibleMode = false
	GotHighscore = false

	LoadUserInfo()

	LDConnection = "waiting..."
	go CheckLDConnection()

	LoadFont()

	LoadImages()

	game := &Game{
		Enemies:          make([]*Enemy, 0),
		Bullets:          make([]*Bullet, 0),
		Player:           Player{X: ScreenWidth/2 - 20, Y: ScreenHeight - 40, SizeMult: 1},
		TimerSystem:      ebitick.NewTimerSystem(),
		Score:            0,
		Mode:             "title",
		DWL:              DW{Image: DWLImage, Active: false, IsSpawning: false, Rad: 0, X: -DWWidth, Side: "left"},
		DWR:              DW{Image: DWRImage, Active: false, IsSpawning: false, Rad: 0, X: ScreenWidth, Side: "right"},
		ShowDWWL:         false,
		ShowDWWR:         false,
		EditText:         "",
		StartedTheTimers: false,
	}
	ShotDelay = time.Millisecond * 200
	CanShoot = true
	bulletsToRemove = make([]int, 0)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Scramble Ghosts 👻")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
