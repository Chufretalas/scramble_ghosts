package main

import (
	"image"
	_ "image/png"
	"io"
	"io/fs"
	"log"
	"reflect"

	u "github.com/Chufretalas/scramble_ghosts/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Sprite struct {
	Path string
	Img  *ebiten.Image
}

type GameSprites struct {
	Title          Sprite
	LDButton       Sprite
	LDButtonActive Sprite
	Gameover       Sprite
	GameoverHS     Sprite
	PlayerBullet   Sprite
	EnemyBullet30  Sprite
	EnemyBullet50  Sprite
	Player         Sprite
	CurveL         Sprite
	CurveR         Sprite
	Linear         Sprite
	DWL            Sprite
	DWR            Sprite
	DWWL           Sprite // death wall warning
	DWWR           Sprite // death wall warning
	Arcshot        Sprite
}

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

func LoadAssets(assetsFS fs.FS) {

	Sprites = GameSprites{
		Title:          Sprite{Path: "title_screen_4k_16-9.png"},
		Gameover:       Sprite{Path: "gamover_screen_1080p.png"},
		GameoverHS:     Sprite{Path: "gamover_screen_hs.png"},
		PlayerBullet:   Sprite{Path: "player_bullet.png"},
		EnemyBullet30:  Sprite{Path: "enemy_bullet_30.png"},
		EnemyBullet50:  Sprite{Path: "enemy_bullet_50.png"},
		Player:         Sprite{Path: "player_spritesheet.png"},
		Linear:         Sprite{Path: "linear_50x50.png"},
		CurveL:         Sprite{Path: "curve_spritesheet_l.png"},
		CurveR:         Sprite{Path: "curve_spritesheet_r.png"},
		DWL:            Sprite{Path: "deathwall_l.png"},
		DWR:            Sprite{Path: "deathwall_r.png"},
		DWWL:           Sprite{Path: "dw_warning_l.png"},
		DWWR:           Sprite{Path: "dw_warning_r.png"},
		LDButton:       Sprite{Path: "ldbutton.png"},
		LDButtonActive: Sprite{Path: "ldbutton_active.png"},
		Arcshot:        Sprite{Path: "arcshot_v2_sheet.png"},
	}

	ptrToSprites := &Sprites
	r := reflect.ValueOf(ptrToSprites).Elem()

	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)

		currentSprite, _ := field.Interface().(Sprite)

		spriteImg, _, imageError := ebitenutil.NewImageFromFileSystem(assetsFS, currentSprite.Path)

		if imageError != nil {
			u.ErrorAndDie(currentSprite.Path + " could not load: " + imageError.Error())
		}

		updatedSprite := Sprite{
			Path: currentSprite.Path,
			Img:  spriteImg,
		}
		field.Set(reflect.ValueOf(updatedSprite))
	}

	loadFont(assetsFS)
	loadIcon(assetsFS)
}
