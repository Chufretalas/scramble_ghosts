package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const dwV = math.Pi / 600

type DW struct {
	Image      *ebiten.Image
	Active     bool
	IsSpawning bool
	X          float64
	Rad        float64
	Side       string
}

func (dw *DW) Move() {
	switch dw.Side {
	case "left":
		dw.Rad += dwV
		if dw.Rad > math.Pi {
			dw.Reset()
			dw.Active = false
			return
		}
		dw.X = math.Sin(dw.Rad)*DWWidth - DWWidth
	case "right":
		dw.Rad += dwV
		if dw.Rad > math.Pi {
			dw.Reset()
			dw.Active = false
			return
		}
		dw.X = ScreenWidth - math.Sin(dw.Rad)*DWWidth
	}
}

func (dw *DW) Reset() {
	dw.Active = false
	dw.IsSpawning = false
	dw.Rad = 0
	switch dw.Side {
	case "left":
		dw.X = -DWWidth
	case "right":
		dw.X = ScreenWidth
	}
}
