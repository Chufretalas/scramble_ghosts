package main

import (
	"image"
	_ "image/png"
	"io"
	"io/fs"
	"log"

	u "github.com/Chufretalas/scramble_ghosts/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func loadIcon(assetsFS fs.FS) {
	iconFile, err := assetsFS.Open("icon.png")
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

func loadFont(assetsFS fs.FS) {

	f, err := assetsFS.Open("PressStart2P-Regular.ttf")
	if err != nil {
		log.Fatal(err)
		defer f.Close()
	}

	contents, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	tt, err := opentype.Parse(contents)
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

func loadImages(assetsFS fs.FS) {

	var imageError error

	TitleImage, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "title_screen_4k_16-9.png")

	if imageError != nil {
		u.ErrorAndDie("Title image did not load " + imageError.Error())
	}

	GameoverImage, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "gamover_screen_1080p.png")

	if imageError != nil {
		u.ErrorAndDie("Gameover image did not load " + imageError.Error())
	}

	GameoverImageHS, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "gamover_screen_hs.png")

	if imageError != nil {
		u.ErrorAndDie("Highscore Gameover image did not load " + imageError.Error())
	}

	PlayerBulletImage, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "player_bullet.png")

	if imageError != nil {
		u.ErrorAndDie("Player bullet image did not load " + imageError.Error())
	}

	EnemyBullet30Image, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "enemy_bullet_30.png")

	if imageError != nil {
		u.ErrorAndDie("Enemy bullet 30 image did not load " + imageError.Error())
	}

	EnemyBullet50Image, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "enemy_bullet_50.png")

	if imageError != nil {
		u.ErrorAndDie("Enemy bullet 50 image did not load " + imageError.Error())
	}

	PlayerSheet, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "player_spritesheet.png")

	if imageError != nil {
		u.ErrorAndDie("Player sprite sheet did not load " + imageError.Error())
	}

	LinearImage, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "linear_50x50.png")

	if imageError != nil {
		u.ErrorAndDie("Linear enemy sprite image did not load " + imageError.Error())
	}

	CurveLSheet, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "curve_spritesheet_l.png")

	if imageError != nil {
		u.ErrorAndDie("CurveL enemy spritesheet image did not load " + imageError.Error())
	}

	CurveRSheet, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "curve_spritesheet_r.png")

	if imageError != nil {
		u.ErrorAndDie("CurveR enemy spritesheet image did not load " + imageError.Error())
	}

	DWLImage, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "deathwall_l.png")

	if imageError != nil {
		u.ErrorAndDie("DWL sprite image did not load " + imageError.Error())
	}

	DWRImage, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "deathwall_r.png")

	if imageError != nil {
		u.ErrorAndDie("DWR sprite image did not load: " + imageError.Error())
	}

	DWWLImage, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "dw_warning_l.png")

	if imageError != nil {
		u.ErrorAndDie("DWWL sprite image did not load: " + imageError.Error())
	}

	DWWRImage, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "dw_warning_r.png")

	if imageError != nil {
		u.ErrorAndDie("DWWR sprite image did not load: " + imageError.Error())
	}

	LDButtonImage, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "ldbutton.png")

	if imageError != nil {
		u.ErrorAndDie("ldbutton image image did not load: " + imageError.Error())
	}

	LDButtonActiveImage, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "ldbutton_active.png")

	if imageError != nil {
		u.ErrorAndDie("ldbutton image image did not load: " + imageError.Error())
	}

	ArcshotSheet, _, imageError = ebitenutil.NewImageFromFileSystem(assetsFS, "arcshot_v2_sheet.png")

	if imageError != nil {
		u.ErrorAndDie("arcshot image image did not load: " + imageError.Error())
	}

}

func LoadAssets(assetsFS fs.FS) {
	loadFont(assetsFS)
	loadImages(assetsFS)
	loadIcon(assetsFS)
}
