package main

import (
	"fmt"
	"time"
)

func (g *Game) SpawnDeathWall(side string) {
	fmt.Println(side)
	switch side {
	case "left":
		g.ShowDWWL = true
		g.DWL.IsSpawning = true
		g.TimerSystem.After(time.Millisecond*800, func() {
			g.ShowDWWL = false
			g.TimerSystem.After(time.Millisecond*400, func() {
				g.ShowDWWL = true
				g.TimerSystem.After(time.Millisecond*800, func() {
					g.ShowDWWL = false
					g.DWL.IsSpawning = false
					g.DWL.Active = true
				})
			})
		})
	case "right":
		g.ShowDWWR = true
		g.DWR.IsSpawning = true
		g.TimerSystem.After(time.Millisecond*800, func() {
			g.ShowDWWR = false
			g.TimerSystem.After(time.Millisecond*400, func() {
				g.ShowDWWR = true
				g.TimerSystem.After(time.Millisecond*800, func() {
					g.ShowDWWR = false
					g.DWR.IsSpawning = false
					g.DWR.Active = true
				})
			})
		})
	}
}
