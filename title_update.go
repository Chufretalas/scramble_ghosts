package main

import (
	"os/exec"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) TitleUpdate() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.Mode = "game"
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		g.Mode = "edit"
		return
	}
	if x, y := ebiten.CursorPosition(); x <= 350 && y <= 200 && inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) { // thanks to: https://gist.github.com/sevkin/9798d67b2cb9d07cb05f89f14ba682f8
		var cmd string
		var args []string

		switch runtime.GOOS {
		case "windows":
			cmd = "cmd"
			args = []string{"/c", "start"}
		case "darwin":
			cmd = "open"
		default: // "linux", "freebsd", "openbsd", "netbsd"
			cmd = "xdg-open"
		}
		args = append(args, UInfo.LD_URL)
		exec.Command(cmd, args...).Start()
	}
}
