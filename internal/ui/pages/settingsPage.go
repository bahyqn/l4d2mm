package pages

import (
	"l4d2mm/internal"
	"l4d2mm/internal/schema"
	"l4d2mm/internal/ui/components"

	"github.com/go-gui-org/go-gui/gui"
)

var branches = []gui.SelectOption{
	{
		Label: "l4d2",
		Value: "l4d2",
	},
	{
		Label: "outBreak",
		Value: "OutBreak",
	},
}

func SettingsView() gui.View {
	return gui.Column(gui.ContainerCfg{
		Content: []gui.View{
			gui.Text(gui.TextCfg{
				Text: "Settings",
			}),

			components.Select(schema.TemplateSelect{
				ID:       "game-branch",
				MaxWidth: 140,
				Selected: internal.GLOBALAPP.AppConfig.GameName,
				Options:  branches,
				OnSelectFunc: func(s []string, ec gui.EventCtx) {
					switch s[0] {
					case "l4d2":
						internal.GLOBALAPP.AppConfig.GameName[0] = s[0]
						internal.GLOBALAPP.AppConfig.Branch = 0
					case "outBreak":
					}
				},
			}),
		},
	})
}
