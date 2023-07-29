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
			if g.DWL.IsSpawning {
				g.TimerSystem.After(time.Millisecond*400, func() {
					g.ShowDWWL = true
					if g.DWL.IsSpawning {
						g.TimerSystem.After(time.Millisecond*800, func() {
							g.ShowDWWL = false
							if g.DWL.IsSpawning {
								g.DWL.IsSpawning = false
								g.DWL.Active = true
							}
						})
					} else {
						g.ShowDWWL = false
					}
				})
			} else {
				g.ShowDWWL = false
			}
		})
	case "right":
		g.ShowDWWR = true
		g.DWR.IsSpawning = true
		if g.DWR.IsSpawning {
			g.TimerSystem.After(time.Millisecond*800, func() {
				g.ShowDWWR = false
				if g.DWR.IsSpawning {
					g.TimerSystem.After(time.Millisecond*400, func() {
						g.ShowDWWR = true
						g.TimerSystem.After(time.Millisecond*800, func() {
							g.ShowDWWR = false
							if g.DWR.IsSpawning {
								g.DWR.IsSpawning = false
								g.DWR.Active = true
							}
						})
					})
				} else {
					g.ShowDWWR = false
				}
			})
		} else {
			g.ShowDWWR = false
		}
	}
}
