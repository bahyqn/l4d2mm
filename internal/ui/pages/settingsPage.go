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
	items := []schema.SettingsGroupConfig{
		{
			Title: "Game & Paths",
			Rows: []schema.SettingsRowConfig{
				{
					LeftRowTitle:    "l4d2 Directory",
					LeftRowSubTitle: internal.GLOBALAPP.AppConfig.GamePath,
					RightView:       components.ButtonLaunchGameFolder(),
				},
				{
					LeftRowTitle:    "Game Version",
					LeftRowSubTitle: "Select the target game build or branch",
					RightView: components.Select(schema.TemplateSelect{
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
			},
		},
		{
			Title: "Game & Paths",
			Rows: []schema.SettingsRowConfig{
				{
					LeftRowTitle:    "l4d2mm",
					LeftRowSubTitle: "v0.1.0",
					RightView:       components.ButtonCheckForUpdates(),
				},
			},
		},
	}

	comps := []gui.View{
		components.VerticalGap(1, gui.FillFixed),
	}

	for _, el := range items {
		comps = append(comps, components.SettingsGroup(&el))
	}

	return gui.Column(gui.ContainerCfg{
		ID:      "settings-page",
		HAlign:  gui.HAlignCenter,
		Sizing:  gui.FillFill,
		Spacing: gui.SpacingPx(30),
		Width:   internal.GLOBALAPP.AppConfig.Width - internal.GLOBALAPP.AppConfig.AsideWidth,
		Content: comps,
	})
}
