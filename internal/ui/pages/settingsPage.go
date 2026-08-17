package pages

import (
	"github.com/go-gui-org/go-gui/gui"
)

func SettingsView() gui.View {
	return gui.Column(gui.ContainerCfg{
		Content: []gui.View{
			gui.Text(gui.TextCfg{
				Text: "Settings",
			}),
		},
	})
}
