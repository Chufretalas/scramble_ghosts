package main

import (
	"log"
	"os"

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
	var tIError error

	titleImage, _, tIError = ebitenutil.NewImageFromFile("./assets/title_screen_4k_16-9.png")

	if tIError != nil {
		log.Fatal("Title image did not load " + tIError.Error())
	}

	var goIError error

	gameoverImage, _, goIError = ebitenutil.NewImageFromFile("./assets/gamover_screen_1080p.png")

	if goIError != nil {
		log.Fatal("Gameover image did not load " + goIError.Error())
	}

	var buIError error

	bulletImage, _, buIError = ebitenutil.NewImageFromFile("./assets/bullet_30x30.png")

	if buIError != nil {
		log.Fatal("Bullet image did not load " + buIError.Error())
	}

	var pIError error

	playerImage, _, pIError = ebitenutil.NewImageFromFile("./assets/player_40x40.png")

	if pIError != nil {
		log.Fatal("Player sprite image did not load " + pIError.Error())
	}

}
