package main

import (
	"log"
	"os"

	u "github.com/Chufretalas/scramble_ghosts/utils"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func LoadFont() {
	f, err := os.ReadFile("./assets/PressStart2P-Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	tt, err := opentype.Parse(f)
	if err != nil {
		log.Fatal(err)
	}
	MyEpicGamerFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func LoadImages() {

	var imageError error

	titleImage, _, imageError = ebitenutil.NewImageFromFile("./assets/title_screen_4k_16-9.png")

	if imageError != nil {
		u.ErrorAndDie("Title image did not load " + imageError.Error())
	}

	imageError = nil

	gameoverImage, _, imageError = ebitenutil.NewImageFromFile("./assets/gamover_screen_1080p.png")

	if imageError != nil {
		u.ErrorAndDie("Gameover image did not load " + imageError.Error())
	}

	imageError = nil

	bulletImage, _, imageError = ebitenutil.NewImageFromFile("./assets/bullet_30x30.png")

	if imageError != nil {
		u.ErrorAndDie("Bullet image did not load " + imageError.Error())
	}

	imageError = nil

	playerImage, _, imageError = ebitenutil.NewImageFromFile("./assets/player_40x40.png")

	if imageError != nil {
		u.ErrorAndDie("Player sprite image did not load " + imageError.Error())
	}

	imageError = nil

	LinearImage, _, imageError = ebitenutil.NewImageFromFile("./assets/linear_50x50.png")

	if imageError != nil {
		u.ErrorAndDie("Linear enemy sprite image did not load " + imageError.Error())
	}

	imageError = nil

	CurveLImage, _, imageError = ebitenutil.NewImageFromFile("./assets/curve_l_50x50.png")

	if imageError != nil {
		u.ErrorAndDie("CurveL enemy sprite image did not load " + imageError.Error())
	}

	imageError = nil

	CurveRImage, _, imageError = ebitenutil.NewImageFromFile("./assets/curve_r_50x50.png")

	if imageError != nil {
		u.ErrorAndDie("CurveR enemy sprite image did not load " + imageError.Error())
	}

	imageError = nil

	DWLImage, _, imageError = ebitenutil.NewImageFromFile("./assets/deathwall_l.png")

	if imageError != nil {
		u.ErrorAndDie("DWL sprite image did not load " + imageError.Error())
	}

	imageError = nil

	DWRImage, _, imageError = ebitenutil.NewImageFromFile("./assets/deathwall_r.png")

	if imageError != nil {
		u.ErrorAndDie("DWR sprite image did not load: " + imageError.Error())
	}

	imageError = nil
}
