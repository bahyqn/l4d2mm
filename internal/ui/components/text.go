package components

import (
	"l4d2mm/internal"
	"strconv"

	"github.com/go-gui-org/go-gui/gui"
)

func TextModsAcount() gui.View {
	amount := len(internal.GLOBALAPP.DI.Task.Selected)

	return gui.Column(gui.ContainerCfg{
		ID:      "column-mods-amount",
		Sizing:  gui.FillFit,
		Padding: gui.NewPadding(8, 8, 8, 30),
		Content: []gui.View{
			gui.Text(gui.TextCfg{
				ID:   "mods-amount",
				Text: strconv.Itoa(amount) + " mods",
			}),
		},
	})
}
