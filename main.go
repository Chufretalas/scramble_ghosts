package main

import (
	"embed"
	"errors"
	"io/fs"
	"log"
	"time"

	u "github.com/Chufretalas/scramble_ghosts/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/ebitick"
	"golang.design/x/clipboard"
	"golang.org/x/image/font"

	_ "github.com/silbinarywolf/preferdiscretegpu"
)

//go:embed all:assets
var assetsFS embed.FS

const (
	SCREENWIDTH         = 1920
	SCREENHEIGHT        = 1080
	DIFF_INCREASE_DELAY = time.Second * 8
	PlayerBulletSize    = 26
	PlayerBaseSize      = 40
	EnemyW              = 50
	EnemyH              = 50
	EnemySpawnTime      = time.Millisecond * 50
	StoppingMult        = 4
	DWWidth             = 800
	DWSafeZone          = 80
	VERSION             = "0.4.0"
)

var (
	ShotDelay           time.Duration
	CanShoot            bool
	MyEpicGamerFont     font.Face
	showDebug           bool
	TitleImage          *ebiten.Image
	LDButtonImage       *ebiten.Image
	LDButtonActiveImage *ebiten.Image
	GameoverImage       *ebiten.Image
	GameoverImageHS     *ebiten.Image
	PlayerBulletImage   *ebiten.Image
	EnemyBullet30Image  *ebiten.Image
	EnemyBullet50Image  *ebiten.Image
	PlayerSheet         *ebiten.Image
	CurveLSheet         *ebiten.Image
	CurveRSheet         *ebiten.Image
	LinearImage         *ebiten.Image
	DWLImage            *ebiten.Image
	DWRImage            *ebiten.Image
	DWWLImage           *ebiten.Image // death wall warning
	DWWRImage           *ebiten.Image // death wall warning
	ArcshotSheet        *ebiten.Image
	InvincibleMode      bool
	UInfo               UserInfo
	LDConnection        string // anything that is not "ok" should not be trusted
	GotHighscore        bool
	editSelected        string
	canUseClipboard     bool
)

type Game struct {
	Enemies          []*Enemy
	PBullets         []*PBullet
	EHBullets        []*EHommingBullet
	Arcshot          Arcshot
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
	ShouldSpawnEnemy bool // this has to ne true on the start of the game
	Diff             Difficulty
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		go ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

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
	return SCREENWIDTH, SCREENHEIGHT
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

	assetsFS, err := fs.Sub(assetsFS, "assets")

	if err != nil {
		u.ErrorAndDie("Could not load assets: " + err.Error())
	}

	LoadAssets(assetsFS)

	game := &Game{
		Enemies:          make([]*Enemy, 0),
		PBullets:         make([]*PBullet, 0),
		EHBullets:        make([]*EHommingBullet, 0),
		Player:           Player{X: SCREENWIDTH/2 - 20, Y: SCREENHEIGHT - 40},
		Arcshot:          Arcshot{X: -150, Y: 0, ShotsFired: 0, State: "idle", Active: false},
		TimerSystem:      ebitick.NewTimerSystem(),
		Score:            0,
		Mode:             "title",
		DWL:              DW{Image: DWLImage, Active: false, IsSpawning: false, Rad: 0, X: -DWWidth, Side: "left"},
		DWR:              DW{Image: DWRImage, Active: false, IsSpawning: false, Rad: 0, X: SCREENWIDTH, Side: "right"},
		ShowDWWL:         false,
		ShowDWWR:         false,
		EditText:         "",
		Diff:             NewDefaultDifficulty(),
		ShouldSpawnEnemy: true,
	}
	ShotDelay = time.Millisecond * 200
	CanShoot = true
	ebiten.SetFullscreen(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Scramble Ghosts 👻")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
