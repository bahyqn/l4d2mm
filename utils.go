package main

import (
	"l4d2mm/internal"

	"github.com/go-gui-org/go-gui/gui"
)

func ButtonColorChoice(app *internal.App, cidx int, color1 gui.Color, color2 gui.Color) gui.Color {
	if app.ComponentStatus.AsideIdx == cidx {
		app.ComponentStatus.AsideIdx = cidx
		return color1
	}
	return color2
}
