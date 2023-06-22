package main

import (
	"log"
	"os"

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
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}
