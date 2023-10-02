package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	u "github.com/Chufretalas/scramble_ghosts/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func LoadIcon() {
	iconFile, err := os.Open("./assets/icon.png")
	if err != nil {
		return
	}
	defer iconFile.Close()

	icon, _, err := image.Decode(iconFile)

	if err != nil {
		return
	}

	ebiten.SetWindowIcon([]image.Image{icon})
}

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

	TitleImage, _, imageError = ebitenutil.NewImageFromFile("./assets/title_screen_4k_16-9.png")

	if imageError != nil {
		u.ErrorAndDie("Title image did not load " + imageError.Error())
	}

	imageError = nil

	GameoverImage, _, imageError = ebitenutil.NewImageFromFile("./assets/gamover_screen_1080p.png")

	if imageError != nil {
		u.ErrorAndDie("Gameover image did not load " + imageError.Error())
	}

	imageError = nil

	GameoverImageHS, _, imageError = ebitenutil.NewImageFromFile("./assets/gamover_screen_hs.png")

	if imageError != nil {
		u.ErrorAndDie("Highscore Gameover image did not load " + imageError.Error())
	}

	imageError = nil

	PlayerBulletImage, _, imageError = ebitenutil.NewImageFromFile("./assets/player_bullet.png")

	if imageError != nil {
		u.ErrorAndDie("Player bullet image did not load " + imageError.Error())
	}

	imageError = nil

	EnemyBulletImage, _, imageError = ebitenutil.NewImageFromFile("./assets/enemy_bullet.png")

	if imageError != nil {
		u.ErrorAndDie("Enemy bullet image did not load " + imageError.Error())
	}

	imageError = nil

	PlayerSheet, _, imageError = ebitenutil.NewImageFromFile("./assets/player_spritesheet.png")

	if imageError != nil {
		u.ErrorAndDie("Player sprite sheet did not load " + imageError.Error())
	}

	imageError = nil

	LinearImage, _, imageError = ebitenutil.NewImageFromFile("./assets/linear_50x50.png")

	if imageError != nil {
		u.ErrorAndDie("Linear enemy sprite image did not load " + imageError.Error())
	}

	imageError = nil

	CurveLSheet, _, imageError = ebitenutil.NewImageFromFile("./assets/curve_spritesheet_l.png")

	if imageError != nil {
		u.ErrorAndDie("CurveL enemy spritesheet image did not load " + imageError.Error())
	}

	imageError = nil

	CurveRSheet, _, imageError = ebitenutil.NewImageFromFile("./assets/curve_spritesheet_r.png")

	if imageError != nil {
		u.ErrorAndDie("CurveR enemy spritesheet image did not load " + imageError.Error())
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

	DWWLImage, _, imageError = ebitenutil.NewImageFromFile("./assets/dw_warning_l.png")

	if imageError != nil {
		u.ErrorAndDie("DWWL sprite image did not load: " + imageError.Error())
	}

	imageError = nil

	DWWRImage, _, imageError = ebitenutil.NewImageFromFile("./assets/dw_warning_r.png")

	if imageError != nil {
		u.ErrorAndDie("DWWR sprite image did not load: " + imageError.Error())
	}

	imageError = nil

	LDButtonImage, _, imageError = ebitenutil.NewImageFromFile("./assets/ldbutton.png")

	if imageError != nil {
		u.ErrorAndDie("ldbutton image image did not load: " + imageError.Error())
	}

	imageError = nil

	LDButtonActiveImage, _, imageError = ebitenutil.NewImageFromFile("./assets/ldbutton_active.png")

	if imageError != nil {
		u.ErrorAndDie("ldbutton image image did not load: " + imageError.Error())
	}

	imageError = nil

	ArcshotSheet, _, imageError = ebitenutil.NewImageFromFile("./assets/arcshot_v2_sheet.png")

	if imageError != nil {
		u.ErrorAndDie("arcshot image image did not load: " + imageError.Error())
	}

	imageError = nil
}
